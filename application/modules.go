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

package application

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lukemakhanu/access_control/config"
	"github.com/lukemakhanu/access_control/domain"
	"github.com/lukemakhanu/access_control/infrastructure/persistence"
)

// GetProject return a single module
func GetModule(id int) (*domain.Modules, error) {
	conn, err := config.MySQLConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewModulesRepositoryWithRDB(conn)
	return repo.Get(id)
}

// GetAllModules return all modules
func GetAllModules() ([]*domain.Modules, error) {
	conn, err := config.MySQLConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewModulesRepositoryWithRDB(conn)
	return repo.GetAll()
}

// AddModule saves new module
func AddModule(p domain.Modules) error {
	conn, err := config.MySQLConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewModulesRepositoryWithRDB(conn)
	return repo.Save(&p)
}

// ModuleStatus updates module status
func ModuleStatus(moduleID int, status int) error {
	conn, err := config.MySQLConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewModulesRepositoryWithRDB(conn)
	return repo.UpdateStatus(moduleID, status)
}

// UpdateModule updates a module
func UpdateModule(p domain.Modules) error {
	conn, err := config.MySQLConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewModulesRepositoryWithRDB(conn)
	return repo.Update(&p)
}
