run: go.mod
	godog

go.mod:
	go mod init
	GO111MODULE=on go get github.com/cucumber/godog/cmd/godog@v0.9.0
	GO111MODULE=on go get github.com/cucumber/godog@v0.9.0
