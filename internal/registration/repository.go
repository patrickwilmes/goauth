/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package registration

import "github.com/patrickwilmes/goauth/internal/db"

type registrationRepository struct {
	backend db.DatabaseBackend
}

func newRepository(backend *db.DatabaseBackend) *registrationRepository {
	return &registrationRepository{backend: *backend}
}

func (repo *registrationRepository) IsValidToken(token string) (bool, error) {
	database := repo.backend.Database()
	var exists bool
	err := database.QueryRow("SELECT EXISTS(SELECT 1 FROM verifier WHERE token=?);", token).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (repo *registrationRepository) Activate(email string) error {
	_, err := repo.backend.Database().Exec("UPDATE accounts SET is_active=1 WHERE email=?;", email)
	return err
}

func (repo *registrationRepository) Save(user User) error {
	database := repo.backend.Database()
	// todo - the password must be stored encrypted not in plain text
	_, err := database.Exec("INSERT INTO accounts VALUES(?,?, 0);", user.Email, user.Password)
	return err
}
func (repo *registrationRepository) RegisterToken(email, token string) error {
	database := repo.backend.Database()
	_, err := database.Exec("INSERT INTO verifier VALUES (?,?);", email, token)
	return err
}
