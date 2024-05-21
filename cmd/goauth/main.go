/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patrickwilmes/goauth/internal/authorization"
	"github.com/patrickwilmes/goauth/internal/common"
	"github.com/patrickwilmes/goauth/internal/db"
	"github.com/patrickwilmes/goauth/internal/registration"
)

const (
	configServerAddress = "server.address"
)

func init() {
	common.InitializeConfiguration()
}

func configureRouting(backend *db.DatabaseBackend) *mux.Router {
	router := mux.NewRouter()
	registration.InitializeRegistrationHandlers(router, backend)
	authorization.InitializeAuthorizationHandler(router, backend)
	return router
}

func main() {
	defer common.CloseLogging()

	logger := common.GetLogger()
	logger.Info().Msg("Booting server ...")
	dbBackend, err := db.NewDBBackend()
	common.PanicOnError(err)

	defer func() {
		common.PanicOnError(dbBackend.Close())
	}()

	router := configureRouting(&dbBackend)
	serverAddress := common.GetStringValue(configServerAddress)
	err = http.ListenAndServe(serverAddress, router)
	common.PanicOnError(err)
}
