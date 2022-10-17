// Copyright 2022 lukemakhanu
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
)

var config = Config{}

// ConfigDB db seting
type Config struct {
	User        string
	Password    string
	Host        string
	Port        string
	Dbname      string
	MysqlConn   string
	MysqlDriver string
}

// ConnectDB returns initialized gorm.DB
func ConnectDB() (*gorm.DB, error) {
	config.Read()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.Dbname)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
