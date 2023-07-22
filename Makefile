createdb:
	# add command later
	echo "db created"

dropdb:
	# add command later
	echo "db dropped"

migrateup:
	migrate -path db/migrations -database "postgres://deltech:secret@localhost:54320/devemm?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://deltech:secret@localhost:54320/devemm?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

dev:
	go run main.go web serve

pre_build:
	rm -rf bin
	mkdir bin

build_linux: pre_build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/web .
	
build_windows: pre_build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/web.exe .

build_darwin: pre_build
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/web .
	
.PHONY: createdb migrateup migratedown sqlc dev build_linux build_windows build_darwin