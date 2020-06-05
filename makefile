all: api

api:
	go build -o api main.go

clean:
	rm api