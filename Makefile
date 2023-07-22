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

web:
	go run main.go web serve -production=false
	
.PHONY: createdb migrateup migratedown sqlc web