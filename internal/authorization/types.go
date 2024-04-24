/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package authorization

type User struct {
	email     string
	password  string
	challenge string
}

type Login struct {
	email    string
	authCode string
}
