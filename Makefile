ALL:clean
	go build -ldflags="-w " -o server ./cmd
	go build -ldflags="-w " -o test ./testonce
clean:
	rm -f server
	rm -f test
