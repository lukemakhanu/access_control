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

// GetProject return a single project
func GetProject(id int) (*domain.Projects, error) {
	conn, err := config.MySQLConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewProjectsRepositoryWithRDB(conn)
	return repo.Get(id)
}

// GetAllProjects return all topics
func GetAllProjects() ([]*domain.Projects, error) {
	conn, err := config.MySQLConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewProjectsRepositoryWithRDB(conn)
	return repo.GetAll()
}

// AddProject saves new project
func AddProject(p domain.Projects) error {
	conn, err := config.MySQLConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewProjectsRepositoryWithRDB(conn)
	return repo.Save(&p)
}

// UpdateProject updates a project
func ProjectStatus(projectID int, status int) error {
	conn, err := config.MySQLConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewProjectsRepositoryWithRDB(conn)
	return repo.UpdateStatus(projectID, status)
}

// UpdateProject updates a project
func UpdateProject(p domain.Projects) error {
	conn, err := config.MySQLConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewProjectsRepositoryWithRDB(conn)
	return repo.Update(&p)
}
