/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package registration

import "github.com/patrickwilmes/goauth/internal/common"

const (
	definedTokenLength = 8
)

type service interface {
	Register(user User) error
	Activate(token string, email string) error
}

type repository interface {
	Save(user User) error
	RegisterToken(email, token string) error
	IsValidToken(token string) (bool, error)
	Activate(email string) error
}

type registrationService struct {
	repo repository
}

func newInstance(repo repository) service {
	return &registrationService{repo: repo}
}

func (srv registrationService) Register(user User) error {
	err := srv.repo.Save(user)
	if err != nil {
		return err
	}
	token := common.GenerateRandomString(definedTokenLength)
	err = srv.repo.RegisterToken(user.Email, token)
	if err != nil {
		return err
	}
	msg := common.Message{
		Token:    token,
		Subject:  "Email Verification",
		Receiver: user.Email,
	}
	err = common.SentMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
func (srv registrationService) Activate(token string, email string) error {
	exists, err := srv.repo.IsValidToken(token)
	if err != nil {
		return err
	}
	if !exists {
		return nil // todo - custom error as the verification attempt failed to an invalid token
	}
	if err := srv.repo.Activate(email); err != nil {
		return err
	}
	return nil
}
