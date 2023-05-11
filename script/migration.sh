#!/bin/bash

migrate create -ext sql -dir migrations/postgres create_table_user

migrate -database "postgres://postgres:postgres@localhost:5432/simple_attendance?sslmode=disable" -path migrations/postgres up

migrate -database "postgres://postgres:postgres@localhost:5432/simple_attendance?sslmode=disable" -path migrations/postgres down

# force dirty state to clear from version
migrate -database "postgres://postgres:postgres@localhost:5432/simple_attendance?sslmode=disable" -path migrations/postgres force VER

# check version 
migrate -database "postgres://postgres:postgres@localhost:5432/simple_attendance?sslmode=disable" -path migrations/postgres version