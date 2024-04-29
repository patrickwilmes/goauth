/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package main

import (
	"github.com/gorilla/mux"
	"github.com/patrickwilmes/goauth/internal/authorization"
	"github.com/patrickwilmes/goauth/internal/common"
	"github.com/patrickwilmes/goauth/internal/db"
	"github.com/patrickwilmes/goauth/internal/registration"
	"net/http"
)

const (
	configServerAddress = "server.address"
)

func configureRouting(backend *db.DatabaseBackend) *mux.Router {
	router := mux.NewRouter()
	registration.InitializeRegistrationHandlers(router, backend)
	authorization.InitializeAuthorizationHandler(router, backend)
	return router
}

func main() {
	common.InitializeConfiguration()
	dbBackend, err := db.CreateDatabaseBackend()
	common.PanicOnError(err)

	defer func() {
		common.PanicOnError(dbBackend.Close())
	}()

	router := configureRouting(&dbBackend)
	serverAddress := common.GetStringValue(configServerAddress)
	err = http.ListenAndServe(serverAddress, router)
	common.PanicOnError(err)
}
