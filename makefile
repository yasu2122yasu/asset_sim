migrate:
	migrate -path db/migrations -database "mysql://root:password@tcp(localhost:3306)/app?multiStatements=true" up 1
