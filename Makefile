

all: clean build

build: 
	go build -o mgob github.com/rskumar/mgobench/cmd/mgobench

clean:
	rm -f mgob
