FROM google/dart

RUN apt-get update && \ 
	apt-get install -y curl unzip wget

########################################
### Install the dart protobuf plugin ###
########################################
RUN cd / && \
	git clone https://github.com/dart-lang/protobuf.git dart-pb && \
	cd /dart-pb && \
	git checkout 2f6fa9f97384ae4749e4f414d7d00d4b79f8fb01 && \
	cd /dart-pb/protoc_plugin && \
	pub install
	
ENV PATH="/dart-pb/protoc_plugin/bin:${PATH}"

##################################################################################
### Copied from https://github.com/znly/docker-protobuf/blob/master/Dockerfile ###
##################################################################################
RUN mkdir -p /protobuf/google/protobuf && \
    for f in any duration descriptor empty struct timestamp wrappers; do \
    curl -L -o /protobuf/google/protobuf/${f}.proto https://raw.githubusercontent.com/google/protobuf/master/src/google/protobuf/${f}.proto; \
    done && \
    mkdir -p /protobuf/google/api && \
    for f in annotations http; do \
    curl -L -o /protobuf/google/api/${f}.proto https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/third_party/googleapis/google/api/${f}.proto; \
    done && \
    mkdir -p /protobuf/github.com/gogo/protobuf/gogoproto && \
    curl -L -o /protobuf/github.com/gogo/protobuf/gogoproto/gogo.proto https://raw.githubusercontent.com/gogo/protobuf/master/gogoproto/gogo.proto && \
    mkdir -p /protobuf/github.com/mwitkow/go-proto-validators && \
    curl -L -o /protobuf/github.com/mwitkow/go-proto-validators/validator.proto https://raw.githubusercontent.com/mwitkow/go-proto-validators/master/validator.proto

##################################
### Install the protoc command ###
##################################
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip && \
	unzip protoc-3.2.0-linux-x86_64.zip -d protoc3 && \
	mv protoc3/bin/* /usr/bin/ && \
	mv protoc3/include/* /usr/local/include/

##################
### Install Go ###
##################
RUN wget https://storage.googleapis.com/golang/go1.11.4.linux-amd64.tar.gz && \
	tar -C /usr/local -xzf go1.11.4.linux-amd64.tar.gz && \
	mkdir /go

ENV GOPATH /root/
ENV PATH $PATH:/root/bin
ENV PATH $PATH:/usr/local/go/bin

######################################
### Install the Go protobuf plugin ###
######################################
RUN go get -u github.com/golang/protobuf/proto && \
	go get -u github.com/golang/protobuf/protoc-gen-go

#####################################################
### Setup the entrypoint and default library path ###
#####################################################
ENTRYPOINT ["/usr/bin/protoc", "-I/protobuf"]
