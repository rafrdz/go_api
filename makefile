all: api

api:
	go build -o go_api main.go

clean:
	rm api