GOC = go build -ldflags "-s -w"
GOV = go vet
GOL = golint

BIN = $(shell basename $(CURDIR))
SRC = $(BIN).go

build:
	$(GOC) $(SRC)

check:
	$(GOV) $(SRC)
	$(GOL) $(SRC)

clean:
	go clean
	rm -f $(BIN)

run:
	go run ./$(SRC)

.PHONY: build check clean run
