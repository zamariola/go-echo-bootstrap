OUTPUT=build

all: clean test build

clean:
	rm -rf ./$(OUTPUT)

build: clean
	mkdir -p ./$(OUTPUT) && go build -o $(OUTPUT)/server

test:
	go test -v .../.
