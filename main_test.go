package main_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/jackc/pgx/v4/stdlib" // force psql driver import
)

/*
	TEST DB (postgres):
	create table test (id int)
*/

func TestMain(m *testing.M) {
	txdb.Register("txdb", "pgx", "postgres://root@localhost:5432/test?sslmode=disable")
	os.Exit(m.Run())
}

func openDB(id string) (*sql.DB, error) {
	// Create a single transaction database based on a unique id
	db, err := sql.Open("txdb", id)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestInsertRecord(t *testing.T) {
	t.Parallel()
	db, err := openDB(t.Name())
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO product (id) VALUES (1)")
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsert2Record(t *testing.T) {
	t.Parallel()
	db, err := openDB(t.Name())
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO product (id) VALUES (1)")
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsertDeleteRecord(t *testing.T) {
	t.Parallel()
	db, err := openDB(t.Name())
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO product (id) VALUES (1)")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("DELETE FROM product where id = 1")
	if err != nil {
		t.Fatal(err)
	}
}

