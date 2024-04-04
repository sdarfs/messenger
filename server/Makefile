postgresinit:
	docker run --name postgres14 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14-alpine

postgres:
	docker exec -it postgres14 psql

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root chatGO

dropdb:
	docker exec -it postgres14 dropdb chatGO

.PHONY:postgresinit