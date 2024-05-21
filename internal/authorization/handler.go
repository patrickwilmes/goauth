/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package authorization

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/patrickwilmes/goauth/internal/db"
	"html/template"
	"net/http"
)

type authorizationContext struct {
	backend *db.DatabaseBackend
	service service
}

func InitializeAuthorizationHandler(router *mux.Router, backend *db.DatabaseBackend) {
	context := authorizationContext{
		backend: backend,
		service: newInstance(newRepository(*backend)),
	}
	router.Methods("GET").Path("/token").Name("generateToken").HandlerFunc(context.GenerateToken)
	router.Methods("POST").Path("/login").Name("registerUser").HandlerFunc(context.Login)
	router.Methods("GET").Path("/login").Name("loginPage").HandlerFunc(context.LoginPage)
	router.Methods("POST").Path("/token/verify").Name("verifyToken").HandlerFunc(context.VerifyToken)
}

type loginContext struct {
	RedirectUrl string
	Challenge   string
}

type tokenVerification struct {
	Token string
}

func (context authorizationContext) VerifyToken(w http.ResponseWriter, r *http.Request) {
	var tokenVerification tokenVerification
	err := json.NewDecoder(r.Body).Decode(&tokenVerification)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	isValid, err := context.service.IsValidToken(tokenVerification.Token)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !isValid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	return
}

func (context authorizationContext) GenerateToken(w http.ResponseWriter, r *http.Request) {
	authCode := r.FormValue("authCode")
	challenge := r.FormValue("challenge")
	token, err := context.service.GenerateToken(authCode, challenge)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(token))
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (context authorizationContext) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Can't parse form", http.StatusInternalServerError)
		return
	}
	redirectUrl := r.FormValue("redirect_url")
	user := User{
		email:     r.FormValue("email"),
		password:  r.FormValue("password"),
		challenge: r.FormValue("challenge"),
	}
	login, err := context.service.Login(user)
	if err != nil {
		http.Error(w, "login failed", http.StatusInternalServerError)
		return
	}
	if login == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	redirectUrl = redirectUrl + "?authCode=" + login.authCode
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
}

func (context authorizationContext) LoginPage(w http.ResponseWriter, r *http.Request) {
	ctx := loginContext{
		RedirectUrl: r.URL.Query().Get("redirect_url"),
		Challenge:   r.URL.Query().Get("challenge"),
	}
	tpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = tpl.Execute(w, ctx); err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
