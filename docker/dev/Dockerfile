FROM golang:1.19.1-alpine3.16

WORKDIR /app

RUN apk update && \
    apk add protobuf=3.18.1-r2 make && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install github.com/favadi/protoc-go-inject-tag@v1.4.0

COPY go.mod go.sum ./
RUN go mod download

COPY ./pkg/proto ./pkg/proto
COPY Makefile ./
RUN make proto

COPY ./ent ./ent
RUN make ent

CMD ./air-install.sh && \
    ./bin/air
