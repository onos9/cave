FROM golang:1.17

WORKDIR /caveapi

COPY go.mod ./

RUN go mod tidy

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT export STORAGE_HOST=db && CompileDaemon --build="go build cmd/api/main.go" --command="./main"
