proxy:
	@go env -w GO111MODULE="on"
	@go env -w GOPROXY="https://goproxy.cn,direct"

clean:
	@bash build/clean.sh

tidy:
	@go mod tidy -e -v

fmt:
	@goctl api format -dir ./gateway
	@find . -name '*.go' -not -path "./vendor/*" -not -name "*.pb.go" | xargs gofumpt -w -s -extra
	@find . -name '*.go' -not -path "./vendor/*" -not -name "*.pb.go" | xargs -n 1 -I {} -t goimports-reviser -file-path {} -local "gitlab.33.cn" project-name "gitlab.33.cn/proof/backend-micro/" -rm-unused
	@find . -name '*.sh' -not -path "./vendor/*" | xargs shfmt -w -s -i 2 -ci -bn -sr

lint:
	@golangci-lint run ./...

erc721:
	@bash build/single-client.sh erc721

transfer:
	@bash build/single-client.sh transfer

proof:
	@bash build/single-client.sh proof