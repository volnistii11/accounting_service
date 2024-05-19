# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN:=$(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

.PHONY: .vendor-rm
.vendor-rm:
	rm -rf vendor-proto

# Устанавливаем proto описания google/protobuf
vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
	cd vendor-proto/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p vendor-proto/google
	mv vendor-proto/protobuf/src/google/protobuf vendor-proto/google
	rm -rf vendor-proto/protobuf

# Устанавливаем proto описания validate
vendor-proto/validate:
	git clone -b main --single-branch --depth=2 --filter=tree:0 \
		https://github.com/bufbuild/protoc-gen-validate vendor-proto/tmp && \
		cd vendor-proto/tmp && \
		git sparse-checkout set --no-cone validate &&\
		git checkout
		mkdir -p vendor-proto/validate
		mv vendor-proto/tmp/validate vendor-proto/
		rm -rf vendor-proto/tmp

# Вендоринг внешних proto файлов
.vendor-proto: .vendor-rm  vendor-proto/google/protobuf vendor-proto/validate

.PHONY: .bin-deps
.bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
    GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 && \
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4

.create-pkg:
	mkdir -p pkg/api/transfer/v1

TRANSFER_PROTO_PATH:="api/transfer/v1"

PHONY: .protoc-generate
protoc-generate: .create-pkg .bin-deps .vendor-proto
	protoc \
  -I ${TRANSFER_PROTO_PATH} \
  -I vendor-proto \
  --plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
  --go_out pkg/${TRANSFER_PROTO_PATH} \
  --go_opt paths=source_relative \
  --plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
  --go-grpc_out pkg/${TRANSFER_PROTO_PATH} \
  --go-grpc_opt paths=source_relative \
  --plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate \
  --validate_out="lang=go,paths=source_relative:pkg/api/transfer/v1" \
  api/transfer/v1/transfer.proto
	go mod tidy
