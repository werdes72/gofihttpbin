.PHONY: build
build:
	docker build -t werdes72/gofihttpbin:latest .

.PHONY: run
run:
	go run ./cmd/gofihttpbin/gofihttpbin.go

.PHONY: test
test:
	go test -v ./...

.PHONY: unitcoverage
unitcoverage:
	go test -coverprofile=coverage.out ./... ; go tool cover -html=coverage.out
