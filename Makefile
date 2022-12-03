all: build/bin/app

build/bin/app: main.go
	GOOS=linux GOARCH=amd64 go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o build/bin/app .

.PHONY: init plan apply destroy

init: build main.tf
	terraform init

plan: build main.tf
	terraform plan

apply: build main.tf
	terraform apply --auto-approve

destroy: build main.tf
	terraform destroy --auto-approve
