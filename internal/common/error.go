/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package common

import "github.com/rs/zerolog/log"

func PanicOnError(err error) {
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		panic(err)
	}
}
