build:
	go build -o ./bin/blockchain
run:
	./bin/blockchain

test:
	go test -v ./...