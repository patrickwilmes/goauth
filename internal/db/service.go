/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package db

import "database/sql"

type DatabaseBackend interface {
	Exec(sql string) error
	Database() *sql.DB
	Close() error
}
