/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package authorization

import "github.com/patrickwilmes/goauth/internal/db"

type authRepository struct {
	backend db.DatabaseBackend
}

func newRepository(backend db.DatabaseBackend) *authRepository {
	return &authRepository{backend: backend}
}

func (repo authRepository) TokenExists(token string) (bool, error) {
	var exists bool
	err := repo.backend.Database().QueryRow("SELECT EXISTS(SELECT 1 FROM logins WHERE jwt=?", token).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (repo authRepository) CheckAuthCodeAndChallenge(authCode, challenge string) (bool, error) {
	var exists bool
	err := repo.backend.Database().QueryRow("SELECT EXISTS(SELECT 1 FROM logins WHERE challenge=? AND auth_code=?);", challenge, authCode).Scan(&exists)
	return exists, err
}

func (repo authRepository) UpdateLoginFlow(authCode, token, challenge string) error {
	_, err := repo.backend.Database().Exec("UPDATE logins SET jwt=? WHERE auth_code=? AND challenge=?;", token, authCode, challenge)
	return err
}

func (repo authRepository) GetUserByMail(mail string) (User, error) {
	var user User
	err := repo.backend.Database().QueryRow("SELECT * FROM accounts WHERE email = ?", mail).Scan(&user.email, &user.password, &user.challenge)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (repo authRepository) RegisterActiveLoginFlow(user User, authCode string) error {
	_, err := repo.backend.Database().Exec("INSERT INTO logins (email, auth_code, challenge) VALUES (?, ?, ?)", user.email, authCode, user.challenge)
	if err != nil {
		return err
	}
	return nil
}
