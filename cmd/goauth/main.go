/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/patrickwilmes/goauth/internal/authorization"
	"github.com/patrickwilmes/goauth/internal/common"
	"github.com/patrickwilmes/goauth/internal/db"
	"github.com/patrickwilmes/goauth/internal/registration"
	"github.com/rs/zerolog"
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
	logFile, _ := os.OpenFile(
		"application.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)

	defer logFile.Close()

	multi := zerolog.MultiLevelWriter(os.Stdout, logFile)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	logger.Info().Msg("Booting server ...")
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
