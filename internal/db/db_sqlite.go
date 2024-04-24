/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	file       string = "auth.db"
	driverName string = "sqlite3"
)

// table definitions
const (
	// Stores all users that should have access to the backend system.
	// On creation / registration the user account will be set to inactive (0)
	// and will only be set to active ones the email verification is completed
	// successfully.
	accountTable string = `
		CREATE TABLE IF NOT EXISTS accounts (
			email 		TEXT NOT NULL PRIMARY KEY,
			password 	TEXT NOT NULL,
			is_active 	INT DEFAULT 0
		);
	`
	// used to store the token when the user first registers and needs
	// to verify the email address. The token sent via mail, is stored
	// in this table, and when the user clicks on the corresponding
	// button in the mail, the token can be checked against this table.
	// Afterward the account will be set to active.
	verifierTable string = `
		CREATE TABLE IF NOT EXISTS verifier (
			email TEXT NOT NULL PRIMARY KEY,
			token TEXT NOT NULL
		);
	`
	// table to store the active user logins, so every user who
	// successfully logged in and received a valid JWT. To keep
	// it simple for now, the JWT expiration time is written directly
	// into the table column INVAL_DATE, to simplify check for now.
	logins = `
		CREATE TABLE IF NOT EXISTS logins (
			email 		TEXT NOT NULL PRIMARY KEY,
			challenge 	TEXT NOT NULL,
			auth_code 	TEXT NOT NULL,
		    jwt			TEXT,
			inval_date 	TEXT
		);
	`
)

type SqliteBackend struct {
	DB *sql.DB
}

func New() (DatabaseBackend, error) {
	db, err := sql.Open(driverName, file)
	if err != nil {
		return nil, err
	}
	backend := &SqliteBackend{
		DB: db,
	}
	if err := backend.Exec(accountTable); err != nil {
		return nil, err
	}
	if err := backend.Exec(verifierTable); err != nil {
		return nil, err
	}
	if err := backend.Exec(logins); err != nil {
		return nil, err
	}
	return backend, nil
}

func (backend *SqliteBackend) Database() *sql.DB {
	return backend.DB
}
func (backend *SqliteBackend) Close() error {
	return backend.DB.Close()
}

func (backend *SqliteBackend) Exec(sql string) error {
	if _, err := backend.DB.Exec(sql); err != nil {
		return err
	}
	return nil
}
