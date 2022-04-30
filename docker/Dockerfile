###############################
# STEP 1 create dev image
###############################
FROM golang:1.17-alpine AS dev
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0
ENV GOARCH=amd64

# System dependencies
RUN apk update && apk upgrade \
    && apk add --no-cache \
    git

# Fetch dependencies.
COPY  . /app

# COPY dbinit.js /docker-entrypoint-initdb.d/
RUN go mod tidy -compat=1.17 \
    && go mod download \
    && go mod verify \
    && go install github.com/go-delve/delve/cmd/dlv@latest

RUN go build -gcflags "all=-N -l"  cmd/api/main.go

ENTRYPOINT /go/bin/dlv --listen=:2345 --headless=true --log=true --log-output=debugger,debuglineerr,gdbwire,lldbout,rpc --accept-multiclient --api-version=2 exec ./main


###############################
# STEP 2 Build services image
###############################
FROM golang:1.17-alpine AS builder

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0
ENV GOARCH=amd64

WORKDIR /app

# System dependencies
RUN apk update && apk upgrade \
    && apk add --no-cache \
    git 

COPY  . /app

# Fetch dependencies.
RUN go mod tidy -compat=1.17 \
    && go mod download \
    && go mod verify 

# Buid for production
RUN cd cmd/api \
    && go build -gcflags "all=-N -l" -o ../../api

################################
# STEP 3 build a small image
################################
FROM scratch AS server
COPY --from=builder /app/api ./api
ENTRYPOINT ["./api"]

################################
# STEP 3 build a small image
################################
FROM scratch AS web
COPY --from=builder /app/web/build ./web
ENTRYPOINT ["./api"]
