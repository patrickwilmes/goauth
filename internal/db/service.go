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
	dbDriverName     string = "database.driverName"
	dbUrl            string = "database.url"
	sqliteDriverName string = "sqlite3"
	noDriver         string = "noDriver"
)

var (
	errorUnsupportedDBDriver = errors.New("unsupported database driver configured")
)

type DatabaseBackend interface {
	Exec(sql string) error
	Database() *sql.DB
	Close() error
}

func NewDBBackend() (DatabaseBackend, error) {
	driverName := common.GetStringValue(dbDriverName)
	url := common.GetStringValue(dbUrl)
	dbFactories := map[string]func() (DatabaseBackend, error){
		sqliteDriverName: func() (DatabaseBackend, error) {
			return New(driverName, url)
		},
		noDriver: func() (DatabaseBackend, error) {
			log.Printf("an unsupported database driver %s is configured", driverName)
			return nil, errorUnsupportedDBDriver
		},
	}
	if val, present := dbFactories[driverName]; present {
		return val()
	} else {
		return dbFactories[noDriver]()
	}
}
