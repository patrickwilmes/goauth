/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package authorization

import (
	"errors"
	"github.com/patrickwilmes/goauth/internal/common"
)

type service interface {
	Login(user User) (*Login, error)
	GenerateToken(authCode, challenge string) (string, error)
}

type repository interface {
	GetUserByMail(mail string) (User, error)
	RegisterActiveLoginFlow(user User, authCode string) error
	CheckAuthCodeAndChallenge(authCode, challenge string) (bool, error)
	UpdateLoginFlow(authCode, token, challenge string) error
}

type authService struct {
	repo repository
}

func newInstance(repo repository) service {
	return &authService{repo: repo}
}

func (srv authService) GenerateToken(authCode, challenge string) (string, error) {
	// todo - store the expiration date in the database too
	exists, err := srv.repo.CheckAuthCodeAndChallenge(authCode, challenge)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errors.New("no matching active authentication flow")
	}
	token, err := common.CreateToken("") // todo - include email
	if err != nil {
		return "", err
	}
	if err = srv.repo.UpdateLoginFlow(authCode, token, challenge); err != nil {
		return "", err
	}
	return token, nil
}

func (srv authService) Login(user User) (*Login, error) {
	existingUser, err := srv.repo.GetUserByMail(user.email)
	if err != nil {
		return nil, err
	}
	if user.password != existingUser.password {
		return nil, errors.New("invalid password")
	}
	authCode := common.GenerateRandomString(128)
	err = srv.repo.RegisterActiveLoginFlow(user, authCode)
	if err != nil {
		return nil, err
	}
	return &Login{
		email:    user.email,
		authCode: authCode,
	}, nil
}
