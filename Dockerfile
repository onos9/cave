ARG  BUILDER_IMAGE=golang:alpine
ARG BUILD_SERVICE

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
    ca-certificates \
    git \
    && update-ca-certificates

# Fetch dependencies.
COPY . /app
COPY dbinit.js /docker-entrypoint-initdb.d/
RUN go mod tidy -compat=1.17 \
    && go mod download \
    && go get github.com/githubnemo/CompileDaemon

ENTRYPOINT export STORAGE_HOST=db && CompileDaemon --build="go build cmd/api/main.go" --command="./main"


###############################
# STEP 2 Build services image
###############################
FROM dev AS builder
WORKDIR /app

# Fetch dependencies.
COPY . /app
RUN go mod tidy \
    && go mod download \
    && go mod verify

# Buid for production
RUN go build -gcflags "all=-N -l" -o ./api


################################
# STEP 3 build a small image
################################
FROM scratch AS prod
COPY --from=builder ./api ./api
ENTRYPOINT ["./api"]
