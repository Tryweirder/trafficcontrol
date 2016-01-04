// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	"fmt"
	"github.com/Comcast/traffic_control/traffic_ops/goto2/db"
	_ "github.com/Comcast/traffic_control/traffic_ops/goto2/output_format" // needed for swagger
	null "gopkg.in/guregu/null.v3"
	"time"
)

type Status struct {
	Id          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	Description null.String `db:"description" json:"description"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleStatus(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getStatus(id)
	} else if method == "POST" {
		return postStatus(payload)
	} else if method == "PUT" {
		return putStatus(id, payload)
	} else if method == "DELETE" {
		return delStatus(id)
	}
	return nil, nil
}

func getStatus(id int) (interface{}, error) {
	if id >= 0 {
		return getStatusById(id)
	} else {
		return getStatuss()
	}
}

// @Title getStatusById
// @Description retrieves the status information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Status
// @Resource /api/2.0
// @Router /api/2.0/status/{id} [get]
func getStatusById(id int) (interface{}, error) {
	ret := []Status{}
	arg := Status{Id: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from status where id=:id`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getStatuss
// @Description retrieves the status information for a certain id
// @Accept  application/json
// @Success 200 {array}    Status
// @Resource /api/2.0
// @Router /api/2.0/status [get]
func getStatuss() (interface{}, error) {
	ret := []Status{}
	queryStr := "select * from status"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postStatus
// @Description enter a new status
// @Accept  application/json
// @Param                 Name json     string   false "name description"
// @Param          Description json     string    true "description description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/status [post]
func postStatus(payload []byte) (interface{}, error) {
	var v Status
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO status("
	sqlString += "name"
	sqlString += ",description"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:description"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putStatus
// @Description modify an existing statusentry
// @Accept  application/json
// @Param                 Name json     string   false "name description"
// @Param          Description json null.String    true "description description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/status [put]
func putStatus(id int, payload []byte) (interface{}, error) {
	var v Status
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE status SET "
	sqlString += "name = :name"
	sqlString += ",description = :description"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delStatusById
// @Description deletes status information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Status
// @Resource /api/2.0
// @Router /api/2.0/status/{id} [delete]
func delStatus(id int) (interface{}, error) {
	arg := Status{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM status WHERE id=:id", arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
