GO_MODULE := github.com/Ilhamkawe/commerce-proto

.PHONY: all init
all: init

init:
	@mkdir -p .github/workflows
	@mkdir -p proto
	@mkdir -p protogen

new-proto: 
ifndef name
	@printf "\033[31mVariable 'name' is required. Use: make new-proto name=protoname\033[0m\n";
	@exit 1;
endif
	@mkdir -p proto/$(name)/type
	@touch proto/$(name)/type/$(name).proto
	@touch proto/$(name)/service.proto
	@echo "proto/$(name) Created succesfuly"

del-proto: 
ifndef name
	@printf "\033[31mVariable 'name' is required. Use: make del-proto name=protoname\033[0m\n"; 
	@exit 1;
endif
	@rm -rf proto/$(name)
	@echo "proto/$(name) Deleted succesfuly"

clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen
	mkdir protogen\go
else
	rm -fR ./protogen
	mkdir -p ./protogen/go
	ls
endif

PROTO_FILES := $(shell find proto -name "*.proto")
protoc-go:
	@protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=require_unimplemented_servers=false:. \
	$(PROTO_FILES)
	@printf "\033[36mProtogen succesfuly generated\033[0m\n"

.PHONY: build
build: clean protoc-go

.PHONY: pipeline-init
pipeline-init:
	sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


.PHONY: pipeline-build
pipeline-build: pipeline-init build

## gateway ##

.PHONY: clean-gateway
clean-gateway:
ifeq ($(OS), Windows_NT)
	if exist "protogen\gateway" rd /s /q protogen\gateway
	mkdir protogen\gateway\go
	mkdir protogen\gateway\openapiv2
else
	rm -fR ./protogen/gateway
	mkdir -p ./protogen/gateway/go
	mkdir -p ./protogen/gateway/openapiv2
endif


.PHONY: protoc-go-gateway
protoc-go-gateway:
	protoc -I .\ ./proto/hello/*.proto


.PHONY: protoc-openapiv2-gateway
protoc-openapiv2-gateway:
	protoc -I .\ ./proto/hello/*.proto



.PHONY: build-gateway
build-gateway: clean-gateway protoc-go-gateway


.PHONY: pipeline-init-gateway
pipeline-init-gateway:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest


.PHONY: pipeline-build-gateway
pipeline-build-gateway: pipeline-init-gateway build-gateway protoc-openapiv2-gateway