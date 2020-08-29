
go:
	$(ECHO) go get -u google.golang.org/grpc
	$(ECHO) go get -u github.com/golang/protobuf/protoc-gen-go
	$(ECHO) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(ECHO) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

build_dir:
	$(ECHO) mkdir -p $(BUILD_SRC_DIR)
	$(ECHO) cp -R $(SRC) $(BUILD_SRC_DIR)

clean_build_dir:
	$(ECHO) rm -rf $(BUILD_DIR)
	$(ECHO) mkdir -p $(BUILD_SRC_DIR)
	$(ECHO) cp -R $(SRC) $(BUILD_SRC_DIR)

# build: clean_build_dir clean_bin
# 	$(ECHO) export GOPATH=$(BUILD_DIR)/go \
# 	&& cd $(BUILD_SRC_DIR) \
# 	&& go list ./... > GO_PKGS \
# 	&& go vet $$(cat GO_PKGS) \
# 	&& CGO_ENABLED=1 go build -v -ldflags="-X main.version=$(SEMVER) -X main.commit=$(CI_COMMIT_SHA_SHORT) -X main.commit_short=$(CI_COMMIT_SHA)" \
# 	&& mkdir -p $(TOP_DIR)/$(OUT_DIR) \
# 	&& cp -v $(PACKAGE) $(TOP_DIR)/$(OUT_DIR)


build:
    go build 

#  T3sgrPQ3Ra_9xfkFpVgo



