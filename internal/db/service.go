/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package db

import (
	"database/sql"
	"errors"
	"github.com/patrickwilmes/goauth/internal/common"
	"log"
)

const (
	dbDriverName string = "database.driverName"
	dbUrl        string = "database.url"
)

type DatabaseBackend interface {
	Exec(sql string) error
	Database() *sql.DB
	Close() error
}

func CreateDatabaseBackend() (DatabaseBackend, error) {
	driverName := common.GetStringValue(dbDriverName)
	url := common.GetStringValue(dbUrl)

	if driverName == "sqlite3" {
		return New(driverName, url)
	} else {
		log.Printf("an unsupported database driver %s is configured", driverName)
		return nil, errors.New("unsupported database driver configured")
	}
}
