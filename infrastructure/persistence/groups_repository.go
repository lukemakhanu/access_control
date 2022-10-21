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

package persistence

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lukemakhanu/access_control/domain"
	"github.com/lukemakhanu/access_control/domain/repository"
)

// GroupsRepositoryImpl Implements repository.TopicRepository
type GroupsRepositoryImpl struct {
	Conn *sql.DB
}

// NewGroupsRepositoryWithRDB returns initialized GroupsRepositoryImpl
func NewGroupsRepositoryWithRDB(conn *sql.DB) repository.GroupsRepository {
	return &GroupsRepositoryImpl{Conn: conn}
}

// Get Group by id return domain.Groups
func (r *GroupsRepositoryImpl) Get(id int) (*domain.Groups, error) {
	var gc domain.Groups
	statement := fmt.Sprintf("SELECT group_id,group_name,description,project_id,active,created_by, \n"+
		"updated_by,date_created,date_modified FROM access_control.groups where group_id =  '%d' ", id)
	log.Printf("statement : %s", statement)
	rows := r.Conn.QueryRow(statement)
	log.Printf("rows : %v", rows)
	err := rows.Scan(&gc.GroupID, &gc.GroupName, &gc.Description, &gc.ProjectID, &gc.Active, &gc.CreatedBy,
		&gc.UpdatedBy, &gc.DateCreated, &gc.DateModified)
	log.Printf("Err : %v", err)
	switch err {
	case sql.ErrNoRows:
		return &gc, err
	case nil:
		return &gc, nil
	default:
		return &gc, err
	}
}

// GetAll returns all groups
func (r *GroupsRepositoryImpl) GetAll() ([]*domain.Groups, error) {
	qry := fmt.Sprintf("SELECT group_id,group_name,description,project_id,active,created_by, \n" +
		"updated_by,date_created,date_modified FROM access_control.groups ")
	gc := make([]*domain.Groups, 0)
	results, err := r.Conn.Query(qry)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		u := &domain.Groups{}
		err = results.Scan(&u.GroupID, &u.GroupName, &u.Description, &u.ProjectID, &u.Active, &u.CreatedBy,
			&u.UpdatedBy, &u.DateCreated, &u.DateModified)
		if err != nil {
			return nil, err
		}
		gc = append(gc, u)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	results.Close()
	return gc, nil
}

// Save to add group
func (r *GroupsRepositoryImpl) Save(p *domain.Groups) error {
	tx, err := r.Conn.Begin()
	qry := "insert into access_control.groups(group_name,description,project_id,active,created_by, \n" +
		"updated_by,date_created,date_modified) values (?,?,?,?,?,?,?,?)"
	log.Printf("Query : %s | GroupID : %d", qry, p.ProjectID)
	response, err := tx.Exec(qry, p.GroupName, p.Description, p.ProjectID, p.Active, p.CreatedBy,
		p.UpdatedBy, p.DateCreated, p.DateModified)
	if err != nil {
		tx.Rollback()
		return err
	}
	betID, err := response.LastInsertId()
	log.Printf("last groupID : %d", betID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// Update to update group
func (r *GroupsRepositoryImpl) Update(p *domain.Groups) error {
	tx, err := r.Conn.Begin()
	statement := "update access_control.groups set group_name=?,description=?,project_id=?,\n" +
		"active=?,created_by=?,updated_by=?,date_created=?,date_modified=? where group_id=? "
	log.Printf("Query : %s | GroupID : %d", statement, p.ProjectID)
	response, err := tx.Exec(statement, p.GroupName, p.Description, p.ProjectID,
		p.Active, p.CreatedBy, p.UpdatedBy, p.DateCreated, p.DateModified, p.GroupID)
	if err != nil {
		tx.Rollback()
		return err
	}
	groupID, err := response.RowsAffected()
	log.Printf("RowsAffected : %d", groupID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// Update to update Group status
func (r *GroupsRepositoryImpl) UpdateStatus(groupID int, status int) error {
	tx, err := r.Conn.Begin()
	statement := "update access_control.groups set active=?,date_modified=now() where group_id=? limit 1"
	log.Printf("Query : %s | GroupID : %d", statement, groupID)
	response, err := tx.Exec(statement, status, groupID)
	if err != nil {
		tx.Rollback()
		return err
	}
	betID, err := response.RowsAffected()
	log.Printf("RowsAffected : %d", betID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
