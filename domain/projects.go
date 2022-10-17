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

package domain

// Projects represent entity of the project
type Projects struct {
	ProjectID    int    `json:"project_id"`
	CountryID    int    `json:"country_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Active       int    `json:"active"`
	CreatedBy    int    `json:"created_by"`
	UpdatedBy    int    `json:"updated_by"`
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
}
