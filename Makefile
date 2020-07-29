all:
	go run app.go


db-exec:
	docker exec -it app_postgres psql -d app -U postgres 