default:
    echo 'Hello, world!'

hello param:
  echo "hello {{param}}"

clean:
  rm -f ./bin/*

build:
  - GOFLAGS=-mod=mod go build -o bin/aoc-2022 main.go 

run:
  - GOFLAGS=-mod=mod go run main.go

test-all:
  - GOFLAGS=-mod=mod go test ./...

test day:
  - GOFLAGS=-mod=mod go test ./internal/{{day}}...
