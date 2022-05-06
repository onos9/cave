VERSION=`git rev-parse HEAD`
API_SERVICE=api
DOCKERFILE=docker/docker-compose.yml
DOCKERFILE_PROD=docker/docker-compose.prod.yml

.PHONY: help
help: ## - Show help message
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: stop
stop:  ## - Stop the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Stopping cave docker image  \n\033[0m"
	@docker-compose -p cave -f $(DOCKERFILE_PROD) down
	
.PHONY: build
build:	## - Build the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Building the cave image \n\033[0m"
	@docker-compose -p cave -f $(DOCKERFILE) build

.PHONY: cache
cache:	## - Build the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Building the cave image \n\033[0m"
	@docker-compose -p cave -f $(DOCKERFILE) build --no-cache

.PHONY: config
config:	## - Build the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Building the cave image \n\033[0m"
	@docker-compose -f $(DOCKERFILE) config
	
.PHONY: run
run:	## - Run the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Running cave docker image  \n\033[0m"
	@docker-compose -p cave -f $(DOCKERFILE) up  -d

.PHONY: deploy
deploy:	## - Scan for known vulnerabilities the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Deploying to VPS \n\033[0m"
	@docker-compose -p cave -f $(DOCKERFILE_PROD) up  -d

.PHONY: deploy1
deploy1:	## - Scan for known vulnerabilities the cave docker image  
	@printf "\033[32m\xE2\x9c\x93 Deploying to VPS \n\033[0m"
	$(eval OLD_CONTAINER_ID=$(shell docker ps -f name=$(SERVICE_NAME) -q | tail -n1))

	@docker-compose -p cave -f $(DOCKERFILE_PROD) up -d --no-deps --scale $(API_SERVICE)=2 --no-recreate $(API_SERVICE)

	$(eval NEW_CONTAINER_IP=$(shell docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(NEW_CONTAINER_ID)))
	$(eval NEW_CONTAINER_ID=$(shell docker ps -f name=$(SERVICE_NAME) -q | head -n1))

	@curl --silent --include --retry-connrefused --retry 30 --retry-delay 1 --fail http://$(NEW_CONTAINER_IP):3000/ || exit 1

	@docker exec nginx /usr/sbin/nginx -s reload 
	@docker stop $(OLD_CONTAINER_ID)
	@docker rm $(OLD_CONTAINER_ID)
	@docker-compose -p cave -f $(DOCKERFILE_PROD) up -d --no-deps --scale $(API_SERVICE)=1 --no-recreate $(API_SERVICE)

	@docker exec nginx /usr/sbin/nginx -s reload
	@echo $(NEW_CONTAINER_IP)
