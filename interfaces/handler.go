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

package interfaces

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lukemakhanu/access_control/application"
	"github.com/lukemakhanu/access_control/domain"
)

// Router : gin router
var Router *gin.Engine

func Run(port int) {

	Router = gin.Default()
	Router.Use(cors.Default())
	v1 := Router.Group("/v1")
	{
		// ======== PROJECTS
		v1.GET("/projects/all", GetProjects)
		v1.GET("/projects/single", GetProject)
		v1.POST("/projects/create", CreateProject)
		v1.POST("/projects/update", UpdateProject)
		v1.POST("/projects/status", ProjectStatus)
		// ======== GROUPS
		v1.GET("/groups/all", GetGroups)
		v1.GET("/groups/single", GetGroup)
		v1.POST("/groups/create", CreateGroup)
		v1.POST("/groups/update", UpdateGroup)
		v1.POST("/groups/status", GroupStatus)
	}
	portStr := fmt.Sprintf(":%d", port)
	log.Printf("Running on port : %s", portStr)
	Router.Run(":8080")
}

// ProjectStatus : update status whether its active or not active
func ProjectStatus(c *gin.Context) {
	type payload struct {
		ProjectID int `json:"project_id"`
		Status    int `json:"status"`
	}
	var p payload
	err := c.Bind(&p)
	if err != nil {
		Error(c.Writer, http.StatusBadRequest, err, err.Error())
		return
	}
	err = application.ProjectStatus(p.ProjectID, p.Status)
	if err != nil {
		Error(c.Writer, http.StatusNotFound, err, err.Error())
		return
	}
	JSON(c.Writer, http.StatusCreated, nil)
}

// UpdateProject : updates a project
func UpdateProject(c *gin.Context) {
	var p domain.Projects
	err := c.Bind(&p)
	if err != nil {
		Error(c.Writer, http.StatusBadRequest, err, err.Error())
		return
	}
	err = application.UpdateProject(p)
	if err != nil {
		Error(c.Writer, http.StatusNotFound, err, err.Error())
		return
	}
	JSON(c.Writer, http.StatusCreated, nil)
}

// CreateProject : creates a new project
func CreateProject(c *gin.Context) {
	var p domain.Projects
	err := c.Bind(&p)
	if err != nil {
		Error(c.Writer, http.StatusBadRequest, err, err.Error())
		return
	}
	err = application.AddProject(p)
	if err != nil {
		Error(c.Writer, http.StatusNotFound, err, err.Error())
		return
	}
	JSON(c.Writer, http.StatusCreated, nil)
}

// GetProjects : return projects
func GetProjects(c *gin.Context) {
	topics, err := application.GetAllProjects()
	if err != nil {
		log.Printf("Err is : %v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	JSON(c.Writer, http.StatusOK, topics)
}

// GetProject : return a single project
func GetProject(c *gin.Context) {
	projectID := c.DefaultQuery("project_id", "0")
	pID, _ := strconv.Atoi(projectID)
	topics, err := application.GetProject(pID)
	if err != nil {
		log.Printf("Err is : %v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	JSON(c.Writer, http.StatusOK, topics)
}

// GROUPS

// GroupStatus : update group status whether its active or not active
func GroupStatus(c *gin.Context) {
	type payload struct {
		GroupID int `json:"group_id"`
		Status  int `json:"status"`
	}
	var p payload
	err := c.Bind(&p)
	if err != nil {
		Error(c.Writer, http.StatusBadRequest, err, err.Error())
		return
	}
	err = application.GroupStatus(p.GroupID, p.Status)
	if err != nil {
		Error(c.Writer, http.StatusNotFound, err, err.Error())
		return
	}
	JSON(c.Writer, http.StatusCreated, nil)
}

// UpdateGroup : updates a group
func UpdateGroup(c *gin.Context) {
	var p domain.Groups
	err := c.Bind(&p)
	if err != nil {
		Error(c.Writer, http.StatusBadRequest, err, err.Error())
		return
	}
	err = application.UpdateGroup(p)
	if err != nil {
		Error(c.Writer, http.StatusNotFound, err, err.Error())
		return
	}
	JSON(c.Writer, http.StatusCreated, nil)
}

// CreateGroup : creates a new group
func CreateGroup(c *gin.Context) {
	var p domain.Groups
	err := c.Bind(&p)
	if err != nil {
		Error(c.Writer, http.StatusBadRequest, err, err.Error())
		return
	}
	err = application.AddGroup(p)
	if err != nil {
		Error(c.Writer, http.StatusNotFound, err, err.Error())
		return
	}
	JSON(c.Writer, http.StatusCreated, nil)
}

// GetGroups : return all group
func GetGroups(c *gin.Context) {
	groups, err := application.GetAllGroups()
	if err != nil {
		log.Printf("Err is : %v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	JSON(c.Writer, http.StatusOK, groups)
}

// GetGroup : return a single group
func GetGroup(c *gin.Context) {
	groupID := c.DefaultQuery("group_id", "0")
	pID, _ := strconv.Atoi(groupID)
	topics, err := application.GetProject(pID)
	if err != nil {
		log.Printf("Err is : %v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	JSON(c.Writer, http.StatusOK, topics)
}
