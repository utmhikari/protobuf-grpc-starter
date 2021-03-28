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
	--go_opt=Mcache.proto=$(PROJECT_PATH)/$(PB_CODE_ROOT)/cache \
    --go-grpc_opt=Mcache.proto=$(PROJECT_PATH)/$(PB_CODE_ROOT)/cache \
	base.proto \
	cache.proto

clean:
	@echo "clean all builds..."

clean: clean_proto

clean_proto:
	@echo "clean all generated proto codes..."
	rm -rf $(PB_CODE_ROOT)
