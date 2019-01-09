SRC_DIR=.
GO_DST_DIR=./go/cfagproto
DART_DST_DIR=./dart/cfagproto

build:
	docker build -t protoc .

compile:
	docker run --rm -v $(PWD):$(PWD) -w $(PWD) protoc --dart_out=${DART_DST_DIR} --go_out=${GO_DST_DIR} -I. cfag.proto
