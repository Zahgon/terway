/*
Copyright 2021 Terway Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controlplane

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	cfg         *Config
	viperConfig *viper.Viper
)

func GetConfig() *Config { _ = "STUB: not implemented"; return nil }

func SetConfig(c *Config) { _ = "STUB: not implemented"; return }

func GetViper() *viper.Viper { _ = "STUB: not implemented"; return nil }

func ParseAndValidateCredential(file string) (*Credential, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseAndValidate ready config and verify it
func ParseAndValidate(configFilePath, credentialFilePath string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InitViper initial viper
// only partial config is loaded
func InitViper(configFilePath string, onConfigChange func(e fsnotify.Event)) error {
	_ = "STUB: not implemented"
	return nil
}

// Read config first before starting watch to avoid race condition

// Start watching after initial read is complete

// IsControllerEnabled check if a specified controller enabled or not.
func IsControllerEnabled(name string, enable bool, controllers []string) bool {
	_ = "STUB: not implemented"
	return false
}
