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

// ProjectsRepositoryImpl Implements repository.TopicRepository
type ProjectsRepositoryImpl struct {
	Conn *sql.DB
}

// NewTopicRepositoryWithRDB returns initialized ProjectsRepositoryImpl
func NewProjectsRepositoryWithRDB(conn *sql.DB) repository.ProjectsRepository {
	return &ProjectsRepositoryImpl{Conn: conn}
}

// Get Project by id return domain.news
func (r *ProjectsRepositoryImpl) Get(id int) (*domain.Projects, error) {
	var gc domain.Projects
	statement := fmt.Sprintf("SELECT project_id,country_id,name,email,phone,active,created_by, \n"+
		"updated_by,date_created,date_modified FROM projects where project_id =  '%d' ", id)
	log.Printf("statement : %s", statement)
	rows := r.Conn.QueryRow(statement)
	log.Printf("rows : %v", rows)
	err := rows.Scan(&gc.ProjectID, &gc.CountryID, &gc.Name, &gc.Email, &gc.Phone, &gc.Active, &gc.CreatedBy,
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
func (r *ProjectsRepositoryImpl) GetAll() ([]*domain.Projects, error) {
	qry := fmt.Sprintf("SELECT project_id,country_id,name,email,phone,active,created_by, \n" +
		"updated_by,date_created,date_modified FROM projects ")
	gc := make([]*domain.Projects, 0)
	results, err := r.Conn.Query(qry)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		u := &domain.Projects{}
		err = results.Scan(&u.ProjectID, &u.CountryID, &u.Name, &u.Email, &u.Phone, &u.Active, &u.CreatedBy,
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
func (r *ProjectsRepositoryImpl) Save(p *domain.Projects) error {
	tx, err := r.Conn.Begin()
	qry := "insert into projects(country_id,name,email,phone,active,created_by,\n" +
		"updated_by,date_created,date_modified) values (?,?,?,?,?,?,?,?,?)"
	log.Printf("Query : %s | ProjectID : %d", qry, p.ProjectID)
	response, err := tx.Exec(qry, p.CountryID, p.Name, p.Email, p.Phone, p.Active, p.CreatedBy,
		p.UpdatedBy, p.DateCreated, p.DateModified)
	if err != nil {
		tx.Rollback()
		return err
	}
	betID, err := response.LastInsertId()
	log.Printf("last projectID : %d", betID)
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
func (r *ProjectsRepositoryImpl) Update(p *domain.Projects) error {
	tx, err := r.Conn.Begin()
	statement := "update projects set country_id=?,name=?,email=?,phone=?,\n" +
		"active=?,created_by=?,updated_by=?,date_created=?,date_modified=? where project_id=? "
	log.Printf("Query : %s | ProjectID : %d", statement, p.ProjectID)
	response, err := tx.Exec(statement, p.CountryID, p.Name, p.Email, p.Phone,
		p.Active, p.CreatedBy, p.UpdatedBy, p.DateCreated, p.DateModified, p.ProjectID)
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
func (r *ProjectsRepositoryImpl) UpdateStatus(projectID int, status int) error {
	tx, err := r.Conn.Begin()
	statement := "update projects set active=?,date_modified=now() where project_id=? limit 1"
	log.Printf("Query : %s | ProjectID : %d", statement, projectID)
	response, err := tx.Exec(statement, status, projectID)
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
