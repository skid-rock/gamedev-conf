# Собрать бинарник key-detect в этом каталоге
BIN := key-detect

.PHONY: build clean install

build:
	go build -o $(BIN) .

# В GOPATH/bin или GOBIN; каталог должен быть в PATH
install:
	go install .

clean:
	rm -f $(BIN)
