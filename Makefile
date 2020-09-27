check_if_installed:
	which  swagger || GO111MODULE=off  go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_if_installed
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

run:
	export APP_DB_USERNAME=postgres
	export APP_DB_PASSWORD=1q2w3e4r
	export APP_DB_NAME=stockmatching