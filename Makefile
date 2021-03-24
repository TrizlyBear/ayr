build:
	go build -o bin/ayr cmd/ayr/main.go

run:
	go run cmd/ayr/main.go

db:
	sudo mongod -f config/mongod.conf

stopdb:
	sudo mongod --shutdown