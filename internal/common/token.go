/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package common

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

var secretKey = []byte(GetStringValue("security.secretKey"))

func CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Minute * time.Duration(GetIntValue("security.jwt.expirationOffsetMins"))).Unix(),
		})
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
