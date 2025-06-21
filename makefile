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