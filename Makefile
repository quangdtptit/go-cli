APP_NAME=myapp
PORT=8080

.PHONY: build serve clean
build:
	go build -o bin/$(APP_NAME)
serve:
	go run main.go serve --port $(PORT)
clean:
	rm -Rf bin/

MIGRATE_NAME=migrate_name
.PHONY: migrate_create migrate_up migrate_down
migrate_create:
	go run main.go migrate create $(MIGRATE_NAME)
migrate_up:
	go run main.go migrate up
migrate_down:
	go run main.go migrate down