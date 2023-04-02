OUT_BINARY=brokerApp
PB_DIR = ./pkg/adapter/grpc/pb/
PB_FILES_MASK = ${PB_DIR}*.proto

build: pb_gen_all
	env GOOS=linux CGO_ENABLED=0 go build -o ${OUT_BINARY} ./cmd/api

pb_gen_all: pb_gen_user
	@echo "[DONE] Generation"

pb_gen_user: ${PB_FILES_MASK}
	@echo "Generating code from user.proto..."
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	${PB_FILES_MASK}