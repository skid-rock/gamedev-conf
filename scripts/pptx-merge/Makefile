# Собрать бинарник pptx-merge в этом каталоге
BIN := pptx-merge

.PHONY: build clean install

build:
	go build -o $(BIN) .

# В GOPATH/bin или GOBIN; каталог должен быть в PATH
install:
	go install .

clean:
	rm -f $(BIN)
