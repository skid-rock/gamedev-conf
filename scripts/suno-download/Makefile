# Собрать бинарник suno-download в этом каталоге
BIN := suno-download

.PHONY: build clean install

build:
	go build -o $(BIN) .

# В GOPATH/bin или GOBIN; каталог должен быть в PATH
install:
	go install .

clean:
	rm -f $(BIN)
