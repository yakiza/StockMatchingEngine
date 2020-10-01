include .env
check_if_installed:
	which  swagger || GO111MODULE=off  go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_if_installed
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

run:
	go run main.go