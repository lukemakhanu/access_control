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

// ModulesRepositoryImpl Implements repository.TopicRepository
type ModulesRepositoryImpl struct {
	Conn *sql.DB
}

// NewModulesRepositoryWithRDB returns initialized ModulesRepositoryImpl
func NewModulesRepositoryWithRDB(conn *sql.DB) repository.ModulesRepository {
	return &ModulesRepositoryImpl{Conn: conn}
}

// Get Project by id return domain.news
func (r *ModulesRepositoryImpl) Get(id int) (*domain.Modules, error) {
	var gc domain.Modules
	statement := fmt.Sprintf("SELECT module_id,project_id,active,created_by, \n"+
		"updated_by,date_created,date_modified FROM modules where project_id =  '%d' ", id)
	log.Printf("statement : %s", statement)
	rows := r.Conn.QueryRow(statement)
	log.Printf("rows : %v", rows)
	err := rows.Scan(&gc.ModuleID, &gc.ProjectID, &gc.Active, &gc.CreatedBy,
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

// Get topic by id return domain.topic
func (r *ModulesRepositoryImpl) GetAll() ([]*domain.Modules, error) {
	qry := fmt.Sprintf("SELECT module_id,project_id,active,created_by, \n" +
		"updated_by,date_created,date_modified FROM modules ")
	gc := make([]*domain.Modules, 0)
	results, err := r.Conn.Query(qry)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		u := &domain.Modules{}
		err = results.Scan(&u.ModuleID, &u.ProjectID, &u.Active, &u.CreatedBy,
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

// Save to add project
func (r *ModulesRepositoryImpl) Save(p *domain.Modules) error {
	tx, err := r.Conn.Begin()
	qry := "insert into modules(project_id,active,created_by,\n" +
		"updated_by,date_created,date_modified) values (?,?,?,?,?,?)"
	log.Printf("Query : %s | ModuleID : %d", qry, p.ModuleID)
	response, err := tx.Exec(qry, p.ProjectID, p.Active, p.CreatedBy,
		p.UpdatedBy, p.DateCreated, p.DateModified)
	if err != nil {
		tx.Rollback()
		return err
	}
	betID, err := response.LastInsertId()
	log.Printf("last moduleID : %d", betID)
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

// Update to update project
func (r *ModulesRepositoryImpl) Update(p *domain.Modules) error {
	tx, err := r.Conn.Begin()
	statement := "update modules set project_id=?,\n" +
		"active=?,created_by=?,updated_by=?,date_created=?,date_modified=? where module_id=? "
	log.Printf("Query : %s | ProjectID : %d", statement, p.ProjectID)
	response, err := tx.Exec(statement, p.ProjectID,
		p.Active, p.CreatedBy, p.UpdatedBy, p.DateCreated, p.DateModified, p.ModuleID)
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

// Update to update project
func (r *ModulesRepositoryImpl) UpdateStatus(moduleID int, status int) error {
	tx, err := r.Conn.Begin()
	statement := "update modules set active=?,date_modified=now() where module_id=? limit 1"
	log.Printf("Query : %s | ModuleID : %d", statement, moduleID)
	response, err := tx.Exec(statement, status, moduleID)
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
