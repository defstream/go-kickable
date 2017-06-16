clean:
	@echo "😁  cleaning"
	@echo "⚙️  cleaning bin/*"
	@rm -rf bin/*
	@echo "⚙️  cleaning vendor/*"
	@rm -rf vendor/*
	@echo "😁  clean"
.PHONY: clean

all:
	@echo "😁  make all "
	@echo "⚙️  initializing: building cani"
	@make build
	@echo "⚙️  initializing: built cani"

	@echo "☁️  service: building http"
	@make http
	@echo "☁️  service: built http"

	@echo "☁️  service: building rpc-gob"
	@make rpc-gob
	@echo "☁️  service: built rpc-gob"

	@echo "☁️  service: building rpc-json"
	@make rpc-json
	@echo "☁️  service: built rpc-json"

	@echo "☁️  service: building kite"
	@make kite
	@echo "☁️  service: built kite"

	@echo "☁️  service: building grpc"
	@make grpc
	@echo "☁️  service: built grpc"

	@echo "☁️  service: building micro"
	@make micro
	@echo "☁️  service: built micro"

	@echo "☁️  service: building graphql"
	@make graphql
	@echo "☁️  service: built graphql"

	@echo "😁  make all complete."

.PHONY: all
	
install:
	@echo "😁  make install"
	@echo "⚙️  installing dependencies..."
	@./scripts/make-install.sh
	@echo "⚙️  installed dependencies."
	@echo "⚙️  building cani..."
	@make build
	@echo "⚙️  built cani."
	@echo "⚙️  installing cani..."
	@cp bin/cani ${GOPATH}/bin
	@echo "⚙️  installed cani."
	@echo "😁  make install complete."
.PHONY: install

# service
build:
	@go build -o bin/cani ./cmd
.PHONY: build

http:
	@go build -o bin/cani.http.server ./service/http/cmd/server
	@go build -o bin/cani.http.client ./service/http/cmd/client
.PHONY: http

rpc-gob: 	
	@go build -o bin/cani.gob.rpc.server ./service/rpc/gob/cmd/server
	@go build -o bin/cani.gob.rpc.client ./service/rpc/gob/cmd/client
.PHONY: rpc-gob

rpc-json: 	
	@go build -o bin/cani.json.rpc.server ./service/rpc/json/cmd/server
	@go build -o bin/cani.json.rpc.client ./service/rpc/json/cmd/client
.PHONY: rpc-json

kite:
	@go build -o bin/cani.kite.server ./service/kite/cmd/server
	@go build -o bin/cani.kite.client ./service/kite/cmd/client
.PHONY: kite

grpc: 	
	@protoc -I${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/defstream/go-kickable/service/grpc/proto/cani.proto 
	@go build -o bin/cani.grpc.server ./service/grpc/cmd/server
	@go build -o bin/cani.grpc.client ./service/grpc/cmd/client
.PHONY: grpc

micro:
	@protoc -I${GOPATH}/src --go_out=plugins=micro:${GOPATH}/src ${GOPATH}/src/github.com/defstream/go-kickable/service/micro/proto/cani.proto
	@go build -o bin/cani.micro.server ./service/micro/cmd/server
	@go build -o bin/cani.micro.client ./service/micro/cmd/client

.PHONY: micro

graphql: 
	@go build -o bin/cani.graphql.server ./service/graphql/cmd/server
.PHONY: graphql


# DEPENDENCIES
start consul:  # micro
	@docker run -p 8500:8500 -p 8300:8300 consul &
