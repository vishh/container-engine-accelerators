FROM ubuntu
ADD libtensorflow.so /usr/local/lib/
ENV LIBRARY_PATH /usr/local/lib
ENV LD_LIBRARY_PATH /usr/local/lib
ADD src /hotdog_server
ADD index.html /
ADD index-smart.html /
RUN chmod a+x /hotdog_server
ENTRYPOINT ["/hotdog_server"]