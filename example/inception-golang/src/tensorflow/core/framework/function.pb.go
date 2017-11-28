// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/function.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A library is a set of named functions.
type FunctionDefLibrary struct {
	Function []*FunctionDef `protobuf:"bytes,1,rep,name=function" json:"function,omitempty"`
	Gradient []*GradientDef `protobuf:"bytes,2,rep,name=gradient" json:"gradient,omitempty"`
}

func (m *FunctionDefLibrary) Reset()                    { *m = FunctionDefLibrary{} }
func (m *FunctionDefLibrary) String() string            { return proto.CompactTextString(m) }
func (*FunctionDefLibrary) ProtoMessage()               {}
func (*FunctionDefLibrary) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *FunctionDefLibrary) GetFunction() []*FunctionDef {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *FunctionDefLibrary) GetGradient() []*GradientDef {
	if m != nil {
		return m.Gradient
	}
	return nil
}

// A function can be instantiated when the runtime can bind every attr
// with a value. When a GraphDef has a call to a function, it must
// have binding for every attr defined in the signature.
//
// TODO(zhifengc):
//   * device spec, etc.
type FunctionDef struct {
	// The definition of the function's name, arguments, return values,
	// attrs etc.
	Signature *OpDef `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// Attributes specific to this function definition.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The body of the function.
	Node []*FunctionDef_Node `protobuf:"bytes,2,rep,name=node" json:"node,omitempty"`
	// The body of the function.  Unlike the NodeDefs in a GraphDef, attrs
	// may have values of type `placeholder` and the `input` field uses
	// the "output" format above.
	NodeDef []*NodeDef `protobuf:"bytes,3,rep,name=node_def,json=nodeDef" json:"node_def,omitempty"`
	// A mapping from the output arg names from `signature` to the
	// outputs from `node_def` that should be returned by the function.
	Ret map[string]string `protobuf:"bytes,4,rep,name=ret" json:"ret,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FunctionDef) Reset()                    { *m = FunctionDef{} }
func (m *FunctionDef) String() string            { return proto.CompactTextString(m) }
func (*FunctionDef) ProtoMessage()               {}
func (*FunctionDef) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *FunctionDef) GetSignature() *OpDef {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FunctionDef) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *FunctionDef) GetNode() []*FunctionDef_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *FunctionDef) GetNodeDef() []*NodeDef {
	if m != nil {
		return m.NodeDef
	}
	return nil
}

func (m *FunctionDef) GetRet() map[string]string {
	if m != nil {
		return m.Ret
	}
	return nil
}

// A node is a multi-value assignment:
//   (ret[0], ret[1], ...) = func(arg[0], arg[1], ...)
//
// By convention, "func" is resolved by consulting with a user-defined
// library first. If not resolved, "func" is assumed to be a builtin op.
type FunctionDef_Node struct {
	// This node produces multiple outputs. They are named ret[0],
	// ret[1], ..., etc.
	//
	// REQUIRES: function.node.ret[*] are unique across all nodes.
	// REQUIRES: ret.size == func/op def's number of output args.
	Ret []string `protobuf:"bytes,1,rep,name=ret" json:"ret,omitempty"`
	// The op/function name.
	Op string `protobuf:"bytes,2,opt,name=op" json:"op,omitempty"`
	// Arguments passed to this func/op.
	//
	// arg[i] must be either one of
	// function.signature.input_args[*].name or one of
	// function.node[*].ret[*].
	//
	// REQUIRES: arg.size == func/op def's number of input args.
	Arg []string `protobuf:"bytes,3,rep,name=arg" json:"arg,omitempty"`
	// Control dependencies.
	//
	// dep[i] must be one of function.node[*].ret[*] or one of
	// function.signature.input_args[*].name.
	Dep []string `protobuf:"bytes,4,rep,name=dep" json:"dep,omitempty"`
	// Attrs.
	//
	// 'attr' maps names defined by 'func's attr defs to attr values.
	// attr values may have placeholders which are substituted
	// recursively by concrete values when this node is instantiated.
	// These placeholders must name an attr listed in the FunctionDef's
	// signature.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FunctionDef_Node) Reset()                    { *m = FunctionDef_Node{} }
func (m *FunctionDef_Node) String() string            { return proto.CompactTextString(m) }
func (*FunctionDef_Node) ProtoMessage()               {}
func (*FunctionDef_Node) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1, 1} }

func (m *FunctionDef_Node) GetRet() []string {
	if m != nil {
		return m.Ret
	}
	return nil
}

func (m *FunctionDef_Node) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *FunctionDef_Node) GetArg() []string {
	if m != nil {
		return m.Arg
	}
	return nil
}

func (m *FunctionDef_Node) GetDep() []string {
	if m != nil {
		return m.Dep
	}
	return nil
}

func (m *FunctionDef_Node) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

// GradientDef defines the gradient function of a function defined in
// a function library.
//
// A gradient function g (specified by gradient_func) for a function f
// (specified by function_name) must follow the following:
//
// The function 'f' must be a numerical function which takes N inputs
// and produces M outputs. Its gradient function 'g', which is a
// function taking N + M inputs and produces N outputs.
//
// I.e. if we have
//    (y1, y2, ..., y_M) = f(x1, x2, ..., x_N),
// then, g is
//    (dL/dx1, dL/dx2, ..., dL/dx_N) = g(x1, x2, ..., x_N,
//                                      dL/dy1, dL/dy2, ..., dL/dy_M),
// where L is a scalar-value function of (x1, x2, ..., xN) (e.g., the
// loss function). dL/dx_i is the partial derivative of L with respect
// to x_i.
type GradientDef struct {
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName" json:"function_name,omitempty"`
	GradientFunc string `protobuf:"bytes,2,opt,name=gradient_func,json=gradientFunc" json:"gradient_func,omitempty"`
}

func (m *GradientDef) Reset()                    { *m = GradientDef{} }
func (m *GradientDef) String() string            { return proto.CompactTextString(m) }
func (*GradientDef) ProtoMessage()               {}
func (*GradientDef) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *GradientDef) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *GradientDef) GetGradientFunc() string {
	if m != nil {
		return m.GradientFunc
	}
	return ""
}

func init() {
	proto.RegisterType((*FunctionDefLibrary)(nil), "tensorflow.FunctionDefLibrary")
	proto.RegisterType((*FunctionDef)(nil), "tensorflow.FunctionDef")
	proto.RegisterType((*FunctionDef_Node)(nil), "tensorflow.FunctionDef.Node")
	proto.RegisterType((*GradientDef)(nil), "tensorflow.GradientDef")
}

func init() { proto.RegisterFile("tensorflow/core/framework/function.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0x6e, 0xb6, 0x10, 0x4f, 0x4a, 0x05, 0x06, 0x84, 0x15, 0x71, 0x08, 0x41, 0xaa, 0x22,
	0x90, 0x76, 0x51, 0x2a, 0x10, 0xea, 0x8d, 0x8a, 0x8f, 0x0b, 0x0a, 0x95, 0x0f, 0x70, 0x8c, 0xdc,
	0xee, 0x6c, 0x14, 0xb5, 0xb1, 0x57, 0x8e, 0x43, 0x95, 0x0b, 0xbf, 0x93, 0x5f, 0xc1, 0x99, 0x23,
	0x1a, 0x67, 0x9d, 0x98, 0x8f, 0xcd, 0x89, 0xdb, 0xc8, 0x7e, 0xef, 0xcd, 0x9b, 0xe7, 0x31, 0x8c,
	0x1c, 0xea, 0xa5, 0xb1, 0xd5, 0xb5, 0xb9, 0x29, 0x2e, 0x8d, 0xc5, 0xa2, 0xb2, 0x6a, 0x81, 0x37,
	0xc6, 0x5e, 0x15, 0xd5, 0x4a, 0x5f, 0xba, 0xb9, 0xd1, 0x79, 0x6d, 0x8d, 0x33, 0x1c, 0x76, 0xc8,
	0xfe, 0xb3, 0x76, 0x96, 0x72, 0xce, 0x4e, 0xbf, 0xaa, 0xeb, 0x15, 0x6e, 0x78, 0xfd, 0x3d, 0x1d,
	0xb4, 0x29, 0x71, 0x5a, 0x62, 0xd5, 0x20, 0x8f, 0xdb, 0x91, 0xa6, 0xde, 0xe1, 0x86, 0xdf, 0x80,
	0xbf, 0x6f, 0xbc, 0xbd, 0xc5, 0xea, 0xe3, 0xfc, 0xc2, 0x2a, 0xbb, 0xe6, 0x27, 0xd0, 0x0d, 0x8e,
	0x45, 0x32, 0xe8, 0x8c, 0x7a, 0xe3, 0x47, 0xf9, 0x4e, 0x30, 0x8f, 0x18, 0x72, 0x0b, 0x24, 0xd2,
	0xcc, 0xaa, 0x72, 0x8e, 0xda, 0x89, 0xf4, 0x6f, 0xd2, 0x87, 0xe6, 0xce, 0x93, 0x02, 0x70, 0xf8,
	0x23, 0x83, 0x5e, 0x24, 0xc7, 0x0b, 0x60, 0xcb, 0xf9, 0x4c, 0x2b, 0xb7, 0xb2, 0x28, 0x92, 0x41,
	0x32, 0xea, 0x8d, 0xef, 0xc5, 0x2a, 0x9f, 0x6a, 0xe2, 0xef, 0x30, 0xfc, 0x25, 0x64, 0x14, 0x93,
	0x38, 0xf0, 0x1d, 0x9f, 0xb4, 0xd8, 0xcc, 0xdf, 0x38, 0x67, 0xdf, 0x69, 0x67, 0xd7, 0xd2, 0xc3,
	0xf9, 0x0b, 0xc8, 0x28, 0xb1, 0xc6, 0xe8, 0xe3, 0x36, 0xda, 0xc4, 0x94, 0x28, 0x3d, 0x92, 0xe7,
	0xd0, 0x0d, 0x19, 0x8b, 0x8e, 0x67, 0xdd, 0x8f, 0x59, 0x84, 0x24, 0x6b, 0xb7, 0xf5, 0xa6, 0xe0,
	0x63, 0xe8, 0x58, 0x74, 0x22, 0xf3, 0xd0, 0x41, 0x5b, 0x03, 0x89, 0x6e, 0x63, 0x8b, 0xc0, 0xfd,
	0x09, 0xb0, 0xad, 0x51, 0x7e, 0x17, 0x3a, 0x57, 0xb8, 0xf6, 0x21, 0x30, 0x49, 0x25, 0x7f, 0x0e,
	0x07, 0x7e, 0x1b, 0x44, 0xea, 0x83, 0x79, 0x18, 0x8b, 0x12, 0xef, 0x33, 0x5d, 0xca, 0x0d, 0xe6,
	0x34, 0x7d, 0x9d, 0xf4, 0xbf, 0x27, 0x90, 0x91, 0x31, 0xd2, 0x22, 0x33, 0xf4, 0x96, 0xcc, 0xb7,
	0xe2, 0x47, 0x90, 0x9a, 0xda, 0x0b, 0x31, 0x99, 0x9a, 0x9a, 0x10, 0xca, 0xce, 0xfc, 0x64, 0x4c,
	0x52, 0x49, 0x27, 0x25, 0xd6, 0x7e, 0x00, 0x26, 0xa9, 0xe4, 0xa7, 0xbf, 0x65, 0x7d, 0xbc, 0x2f,
	0xb4, 0x3f, 0x03, 0xff, 0xef, 0xa3, 0xbd, 0x82, 0x6e, 0xc8, 0xee, 0x1f, 0x72, 0x0f, 0x62, 0x39,
	0x16, 0xf1, 0x86, 0x5f, 0xa0, 0x17, 0x6d, 0x22, 0x7f, 0x0a, 0x77, 0xc2, 0x02, 0x4f, 0xb5, 0x5a,
	0x60, 0x23, 0x72, 0x18, 0x0e, 0x27, 0x6a, 0x81, 0x04, 0x0a, 0x0b, 0x3b, 0xa5, 0x8b, 0x46, 0xf5,
	0x30, 0x1c, 0xd2, 0xf0, 0x67, 0x05, 0x08, 0x63, 0x67, 0xb1, 0xef, 0xed, 0x97, 0x3b, 0x3b, 0x0a,
	0xf1, 0x9c, 0xd3, 0xa7, 0x5b, 0x9e, 0x27, 0x3f, 0x93, 0xe4, 0xe2, 0x96, 0xff, 0x81, 0x27, 0xbf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x44, 0xe4, 0x66, 0x37, 0x04, 0x00, 0x00,
}
