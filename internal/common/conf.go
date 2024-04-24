/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package common

// We don't want to spill viper dependencies all over the place. Therefor, this
// module is used to hide the viper configuration part, and only expose certain
// function for accessing the configuration as interface.

import "github.com/spf13/viper"

const (
	configType              string = "yaml"      // we are only capable of using yaml config
	primaryConfigLocation   string = "./config/" // we are expecting our yaml file to be in a config folder right aside the binary
	secondaryConfigLocation string = "."         // if not, the config file might directly be aside the binary
)

// InitializeConfiguration we initially need to tell viper how to access the configuration
// which type the configuration is of, and where it's located.
func InitializeConfiguration() {
	viper.SetConfigType(configType)
	viper.AddConfigPath(primaryConfigLocation)
	viper.AddConfigPath(secondaryConfigLocation)
	PanicOnError(viper.ReadInConfig())
}

// GetStringValue accepts the config key "root_key.sub_key", and returns
// the associated value.
func GetStringValue(key string) string {
	return viper.GetString(key)
}

func GetIntValue(key string) int {
	return viper.GetInt(key)
}
