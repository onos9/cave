VERSION=`git rev-parse HEAD`
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"
DOCKER_IMAGE=smallest-secured-golang-docker-image

# AWS related variables, eu-west-3 is Paris region
AWS_REGION=eu-west-3
AWS_ACCOUNT_NUMBER=123412341234

#GCP related variables
GCP_PROJECT_ID='chemidy'

.PHONY: help
help: ## - Show help message
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build:	## - Build the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Building the cave image \n\033[0m"
	

.PHONY: build-no-cache
build-no-cache:docker-pull	## - Build the cave docker image without cache
	@printf "\033[32m\xE2\x9c\x93 Build the image with cache disabled \n\033[0m"

.PHONY: ls
ls: ## - List 'cave' docker images
	@printf "\033[32m\xE2\x9c\x93 Look at the size dude !\n\033[0m"
	@docker image ls cave

.PHONY: run
run:	## - Run the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Running cave docker image  \n\033[0m"
	@docker-compose up -d

.PHONY: stop
stop:	## - Stop the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Stopping cave docker image  \n\033[0m"
	@docker-compose down

.PHONY: push-to-azure
push-to-azure:	## - Push docker image to azurecr.io Container Registry
	@az acr login --name chemidy

.PHONY: scan
scan:	## - Scan for known vulnerabilities the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Scan for known vulnerabilities the smallest and secured golang docker image  \n\033[0m"
	@docker scan -f Dockerfile smallest-secured-golang