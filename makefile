build: fmt vet
	go build -o bin/GoMergeExcel ./main.go

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run -E golint,goimports ./...

clean:
	rm -rf bin/