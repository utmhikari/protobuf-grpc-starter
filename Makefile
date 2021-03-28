.PHONY: all clean

PROJECT_PATH = github.com/utmhikari/protobuf-grpc-starter

# for protos
PB_PROTO_ROOT = api/proto
PB_CODE_ROOT = api/pb

all: proto

proto:
	@echo "make proto -> pb & grpc..."
	mkdir -p $(PB_CODE_ROOT)
	protoc --proto_path=$(PB_PROTO_ROOT) \
	--go_out=$(GOPATH)/src \
	--go-grpc_out=$(GOPATH)/src \
	--go_opt=Mbase.proto=$(PROJECT_PATH)/$(PB_CODE_ROOT)/base \
	--go-grpc_opt=Mbase.proto=$(PROJECT_PATH)/$(PB_CODE_ROOT)/base \
	--go_opt=Mw2c.proto=$(PROJECT_PATH)/$(PB_CODE_ROOT)/w2c \
    --go-grpc_opt=Mw2c.proto=$(PROJECT_PATH)/$(PB_CODE_ROOT)/w2c \
	base.proto \
	w2c.proto

clean:
	@echo "clean all builds..."

clean: clean_proto

clean_proto:
	@echo "clean all generated proto codes..."
	rm -rf $(PB_CODE_ROOT)
