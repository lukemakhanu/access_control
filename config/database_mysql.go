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
	"database/sql"
	"log"

	"github.com/BurntSushi/toml"
)

var mysqlconfig = MysqlConfig{}

// MysqlConfig db seting
type MysqlConfig struct {
	User        string
	Password    string
	Host        string
	Port        string
	Dbname      string
	MysqlDNS    string
	MysqlDriver string
}

// MySQLConnectDB returns initialized conn.DB for mysql
func MySQLConnectDB() (*sql.DB, error) {
	mysqlconfig.MysqlRead()

	log.Printf("mysqlconfig.MysqlDNS : %s\n", mysqlconfig.MysqlDNS)
	db, err := sql.Open(mysqlconfig.MysqlDriver, mysqlconfig.MysqlDNS)
	if err != nil {
		log.Println("Error |", err)
	}
	//defer db.Close()
	return db, nil
}

// MysqlRead and parse the configuration file
func (c *MysqlConfig) MysqlRead() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
