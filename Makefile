BIN := bin

all:
	go build -o $(BIN)/ ./...

test:
	go test ./...
	$(BIN)/codex convert impl/convert/convert.go out 34:9 71:11
	
setup:
	mkdir $(BIN)

clean:
	rm -rf $(BIN)