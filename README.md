1. Run PostgreSQL: docker run --name PG nad2000/postgres-uint
1. Connect to DB and create the table: psql -h 172.17.0.2 -U postgres
CREATE TABLE products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)
1. Fetch dependencies: `go get github.com/gorilla/mux github.com/lib/pq`

## Run Tests

1. `export PGHOST=172.17.0.2`
1. `export PGSSLMODE=disable` (in order to disable SSL since docker PG wasn't built with SSL support)
1. `export TEST_DB_USERNAME=postgres` or **PGDATABASE**
1. `export TEST_DB_NAME=postgres` or **PGUSER**

