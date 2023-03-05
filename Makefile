test:
	go test -v -count=1 ./...

cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html coverage.out -o cover.html
	rm coverage.out
