VERSION=`git rev-parse HEAD`

.PHONY: help
help: ## - Show help message
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build:	## - Build the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Building the cave image \n\033[0m"
	@docker-compose -p cave -f docker/docker-compose.yml build

.PHONY: config
config:	## - Build the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Building the cave image \n\033[0m"
	@docker-compose -f docker/docker-compose.yml config
	
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
	@docker-compose -p cave -f docker/docker-compose.yml up  -d

.PHONY: stop
stop:	## - Stop the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Stopping cave docker image  \n\033[0m"
	@docker-compose -p cave -f docker/docker-compose.yml down

.PHONY: push-to-azure
push-to-azure:	## - Push docker image to azurecr.io Container Registry
	@az acr login --name chemidy

.PHONY: scan
scan:	## - Scan for known vulnerabilities the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Scaning for known vulnerabilities   \n\033[0m"
	@docker scan -f Dockerfile 

.PHONY: deploy
deploy:stop	## - Scan for known vulnerabilities the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Deploying to VPS \n\033[0m"
	@docker rmi -f cave_api 
	@docker-compose -p cave -f docker/docker-compose.yml up -d
	
.PHONY: rm
rm:	## - Scan for known vulnerabilities the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Stopping and removing servers \n\033[0m"
	@docker-compose -p cave -f docker/docker-compose.yml stop api
	@docker-compose -p cave -f docker/docker-compose.yml rm -f

.PHONY: prune
prune:	## - Scan for known vulnerabilities the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Deploying to VPS \n\033[0m"
	@docker system prune -a -f