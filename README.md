# GO-CLI Application
## Development command
### Build and Run
```bash
# Build the application
$ make build APP_NAME="myapp"

# Run serve
$ make serve PORT="8081"

# Clean build
$ make clean

# Run database migrations
$ make migrate_create MIGRATE_NAME="create_user_tbl"
$ make migrate_up
$ make migrate_down