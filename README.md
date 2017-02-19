# Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL [![Build Status](https://semaphoreci.com/api/v1/nad2000/gorilla-mux-pg-rest-api/branches/master/badge.svg)](https://semaphoreci.com/nad2000/gorilla-mux-pg-rest-api)

This primer is loosly based on https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

**TODO**: try to implement without Gorilla MUX or replace with something lighter.

1. Run PostgreSQL: `docker run --name PG nad2000/postgres-uint`
1. Connect to DB and create the table: `psql -h 172.17.0.2 -U postgres`:

    ```sql
    CREATE TABLE products
    (
        id SERIAL,
        name TEXT NOT NULL,
        price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
        CONSTRAINT products_pkey PRIMARY KEY (id)
    )
    ```
1. Fetch dependencies: `go get github.com/gorilla/mux github.com/lib/pq`
1. Fetch better support for testing: `go get github.com/smartystreets/goconvey` ([http://goconvey.co])


## Run Tests

1. `export PGHOST=172.17.0.2`
1. `export PGSSLMODE=disable` (in order to disable SSL since docker PG wasn't built with SSL support)
1. `export TEST_DB_USERNAME=postgres` or **PGDATABASE**
1. `export TEST_DB_NAME=postgres` or **PGUSER**
```
export PGHOST=172.17.0.2
export PGSSLMODE=disable
export PGDATABASE=test
export PGUSER=postgres
```
Run either manually: `go test -v` or monitor the project directory with **goconvey**: `goconvey`

## Goconvey Screenshot
![Goconvey](/.screenshots/test.png?raw=true "Goconvey Console")
