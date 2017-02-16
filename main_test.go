package main_test

import (
	"log"
	"os"
	"testing"

	"."
)

var app main.App

func TestMain(m *testing.M) {
	app = main.App{}
	app.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := app.DB.Exec(`
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL,
		name TEXT NOT NULL,
		price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
		CONSTRAINT products_pkey PRIMARY KEY (id))`); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("TRUNCATE TABLE products")
	app.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}
