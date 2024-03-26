#!/bin/bash


docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:latest
echo "Waiting for postgres to start..."
sleep 3

docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE productapp"
sleep 3
echo "Database created successfully"

docker exec -it postgres-test psql -U postgres -d productapp -c "
create table if not exists products 
(
    id bigserial not null primary key, 
    name varchar(255) not null, 
    price double precision not null,
    discount double precision,
    store varchar(255) not null
);
"
sleep 3
echo "Table created successfully"