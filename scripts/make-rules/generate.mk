# ==============================================================================
# 用来进行代码生成的 Makefile
#

.PHONY: gen.add-copyright
gen.add-copyright: ## dd license header
	@addlicense -v -f $(ROOT_DIR)/scripts/boilerplate.txt $(ROOT_DIR) --skip-dirs=third_party,vendor,$(OUTPUT_DIR)

.PHONY: gen.ca
gen.ca: ## generate CA file
	@mkdir -p $(OUTPUT_DIR)/cert
	# 1. generate root certificate private key
	@openssl genrsa -out $(OUTPUT_DIR)/cert/ca.key 4096
 	# 2. generate request file
	@openssl req -new -key $(OUTPUT_DIR)/cert/ca.key -out $(OUTPUT_DIR)/cert/ca.csr \
  	-subj "/C=CN/ST=Shannxi/L=Xi'an/O=devops/OU=XiyouLUG/CN=127.0.0.1/emailAddress=daz-3ux@proton.me"
  	# 3. generate root certificate
	@openssl x509 -req -in $(OUTPUT_DIR)/cert/ca.csr -signkey $(OUTPUT_DIR)/cert/ca.key -out $(OUTPUT_DIR)/cert/ca.crt
	# 4. generate server private key
	@openssl genrsa -out $(OUTPUT_DIR)/cert/server.key 4096
	# 5. generate server public key
	@openssl rsa -in $(OUTPUT_DIR)/cert/server.key -pubout -out $(OUTPUT_DIR)/cert/server.pem
	# 6. generate CSR for server to request signing from CA
	@openssl req -new -key $(OUTPUT_DIR)/cert/server.key -out $(OUTPUT_DIR)/cert/server.csr \
  	-subj "/C=CN/ST=Guangdong/L=Shenzhen/O=serverdevops/OU=serverit/CN=127.0.0.1/emailAddress=daz-3ux@proton.me"
  	# 7.  generate server certificate signed by CA
	@openssl x509 -req -CA $(OUTPUT_DIR)/cert/ca.crt -CAkey $(OUTPUT_DIR)/cert/ca.key \
  	-CAcreateserial -in $(OUTPUT_DIR)/cert/server.csr -out $(OUTPUT_DIR)/cert/server.crt

.PHONY: gen.protoc
gen.protoc: ## compile protobuf file
	@echo "===========> Generate protobuf files"
	@protoc                                            \
		--proto_path=$(APIROOT)                          \
		--proto_path=$(ROOT_DIR)/third_party             \
		--go_out=paths=source_relative:$(APIROOT)        \
		--go-grpc_out=paths=source_relative:$(APIROOT)   \
		$(shell find $(APIROOT) -name *.proto)

.PHONY: gen.deps
gen.deps:
	@go generate $(ROOT_DIR)/...
