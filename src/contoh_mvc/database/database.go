package database

import (
  "database/sql" // datase
  _ "github.com/go-sql-driver/mysql" // import package db
)

func Connect() (*sql.DB, error) {
  con, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/golang") // terhubung ke database
  if err != nil { // cek error
      return nil, err
  }
  return con, nil // kembalikan instansi database
}
