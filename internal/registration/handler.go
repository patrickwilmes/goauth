/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package registration

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/patrickwilmes/goauth/internal/db"
	"net/http"
)

type registrationContext struct {
	backend *db.DatabaseBackend
	service service
}

func InitializeRegistrationHandlers(router *mux.Router, backend *db.DatabaseBackend) {
	context := registrationContext{
		backend: backend,
		service: newInstance(newRepository(backend)),
	}
	router.Methods("POST").Path("/register").Name("registerUser").HandlerFunc(context.Register)
	router.Methods("GET").Path("/register/activate").Name("activateAccount").HandlerFunc(context.ActivateAccount)
}

type registrationDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (context registrationContext) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	email := r.URL.Query().Get("email")
	if err := context.service.Activate(token, email); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (context registrationContext) Register(w http.ResponseWriter, r *http.Request) {
	var registrationDto registrationDto
	err := json.NewDecoder(r.Body).Decode(&registrationDto)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = context.service.Register(User{
		Email:    registrationDto.Email,
		Password: registrationDto.Password,
	})
	// todo - proper error handling
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
