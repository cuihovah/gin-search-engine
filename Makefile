all:
	go build main.go
start:
	sudo ./main
test:
	go run main.go
clean:
	rm main
