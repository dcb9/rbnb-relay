VERSION := $(shell git describe --tags)
COMMIT  := $(shell git log -1 --format='%H')

all: build

LD_FLAGS = -X rbnb-relay/cmd.Version=$(VERSION) \
	-X rbnb-relay/cmd.Commit=$(COMMIT) \

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'

get:
	@echo "  >  \033[32mDownloading & Installing all the modules...\033[0m "
	go mod tidy && go mod download

build:
	@echo " > \033[32mBuilding rbnb-relay...\033[0m "
	go build -mod readonly $(BUILD_FLAGS) -o build/rbnb-relay main.go

build-linux:
	@GOOS=linux GOARCH=amd64 go build --mod readonly $(BUILD_FLAGS) -o ./build/rbnb-relay main.go

install:
	@echo " > \033[32mInstalling rbnb-relay...\033[0m "
	go install -mod readonly $(BUILD_FLAGS) ./...

abi:
	@echo " > \033[32mGenabi...\033[0m "
	abigen --abi ./bindings/StakeManager/StakeManager_abi.json --pkg stake_manager --type StakeManager --out ./bindings/StakeManager/StakeManager.go
	abigen --abi ./bindings/StakePool/StakePool_abi.json --pkg stake_pool --type StakePool --out ./bindings/StakePool/StakePool.go


clean:
	@echo " > \033[32mCleanning build files ...\033[0m "
	rm -rf build
fmt :
	@echo " > \033[32mFormatting go files ...\033[0m "
	go fmt ./...

get-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest

lint:
	golangci-lint run ./... --skip-files ".+_test.go"

.PHONY: all lint test race msan tools clean build
