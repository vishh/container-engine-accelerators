// usage: go run predict_client.go --server_addr 127.0.0.1:9000 --model_name dense --model_version 1
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/context"

	framework "tensorflow/core/framework"
	pb "tensorflow_serving"

	"github.com/gin-gonic/gin"
	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
	imageupload "github.com/olahol/go-imageupload"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

var (
	serverAddr         = flag.String("server_addr", "127.0.0.1:9000", "The server address in the format of host:port")
	modelName          = flag.String("model_name", "cancer", "TensorFlow model name")
	modelVersion       = flag.Int64("model_version", 1, "TensorFlow model version")
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "testdata/ca.pem", "The file containning the CA root cert file")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	imageFile          = flag.String("image_file", "", "Sample image file for prediction")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		var sn string
		if *serverHostOverride != "" {
			sn = *serverHostOverride
		}
		var creds credentials.TransportCredentials
		if *caFile != "" {
			var err error
			creds, err = credentials.NewClientTLSFromFile(*caFile, sn)
			if err != nil {
				grpclog.Fatalf("Failed to create TLS credentials %v", err)
			}
		} else {
			creds = credentials.NewClientTLSFromCert(nil, sn)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPredictionServiceClient(conn)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.GET("/smart", func(c *gin.Context) {
		c.File("index-smart.html")
	})

	r.POST("/upload", func(c *gin.Context) {
		img, err := imageupload.Process(c.Request, "file")

		if err != nil {
			panic(err)
		}
		if !strings.Contains(strings.ToLower(img.Filename), "hotdog") {
			c.Writer.WriteString("<font size=\"6\">Not a hotdog</font>")
		} else {
			c.Writer.WriteString("<font size=\"6\">It's a hotdog!</font>")
		}
	})

	r.POST("/uploadSmart", func(c *gin.Context) {
		img, err := imageupload.Process(c.Request, "file")
		if err != nil {
			panic(err)
		}
		pr := newInceptionPredictRequest(img.Data)
		resp, err := client.Predict(context.Background(), pr)
		if err != nil {
			fmt.Println(err)
			return
		}

		if val, ok := resp.Outputs["classes"]; ok {
			log.Println(string(val.StringVal[0]))
			if strings.Contains(string(val.StringVal[0]), "hotdog") {
				c.Writer.WriteString("<font size=\"6\">It's a hotdog!</font>")
			} else {
				c.Writer.WriteString("<font size=\"6\">Not a hotdog</font>")
			}
		}
	})

	if *imageFile == "" {
		r.Run(":5000")
	} else {
		var pr *pb.PredictRequest
		imageBytes, err := ioutil.ReadFile(*imageFile)
		if err != nil {
			log.Fatalln(err)
		}
		pr = newInceptionPredictRequest(imageBytes)
		resp, err := client.Predict(context.Background(), pr)
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Printf("%q\n", string(resp.Outputs["classes"].StringVal[0]))
	}
}

func newInceptionPredictRequest(image []byte) *pb.PredictRequest {
	tensor, err := tf.NewTensor(string(image))
	if err != nil {
		log.Fatalln("Cannot read image file")
	}

	tensorString, ok := tensor.Value().(string)
	if !ok {
		log.Fatalln("Cannot type assert tensor value to string")
	}

	request := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          "inception",
			SignatureName: "predict_images",
			Version: &google_protobuf.Int64Value{
				Value: int64(1),
			},
		},
		Inputs: map[string]*framework.TensorProto{
			"images": &framework.TensorProto{
				Dtype: framework.DataType_DT_STRING,
				TensorShape: &framework.TensorShapeProto{
					Dim: []*framework.TensorShapeProto_Dim{
						&framework.TensorShapeProto_Dim{
							Size: int64(1),
						},
					},
				},
				StringVal: [][]byte{[]byte(tensorString)},
			},
		},
	}
	return request

}
