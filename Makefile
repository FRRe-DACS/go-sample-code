SERVER_OUT=bin/hospital_server
BANNER=banner.txt
BANNER_OUT=bin/banner.txt
PKG := "github.com/FRRe-DACS/hospital-server"
DOCKER_TAG=frredacs/hospital-server:1.0

.PHONY: docker

dep: ## Get the dependencies
	@go get -v -d ./...
copy_banner:

server: dep 
	@go build -i -v -o $(SERVER_OUT) && cp $(BANNER) $(BANNER_OUT)

docker:
	@docker build -t $(DOCKER_TAG) .

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(BANNER_OUT) 