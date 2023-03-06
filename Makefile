DIR="$(shell echo $(NAME) | sed 's/ /_/g')"

kata:
	mkdir -p $(DIR)
	touch $(DIR)/kata.go
	touch $(DIR)/kata_test.go
	touch $(DIR)/README.md
	echo "# $(NAME)" > $(DIR)/README.md
	echo "[link]($(LINK))" >> $(DIR)/README.md

tests:
	go test -v -count=1 ./...

cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html coverage.out -o cover.html
	rm coverage.out
