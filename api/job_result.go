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

type JobResult struct {
	Id          int64       `db:"id" json:"id"`
	Job         int64       `db:"job" json:"job"`
	Agent       int64       `db:"agent" json:"agent"`
	Result      string      `db:"result" json:"result"`
	Description null.String `db:"description" json:"description"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleJobResult(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getJobResult(id)
	} else if method == "POST" {
		return postJobResult(payload)
	} else if method == "PUT" {
		return putJobResult(id, payload)
	} else if method == "DELETE" {
		return delJobResult(id)
	}
	return nil, nil
}

func getJobResult(id int) (interface{}, error) {
	if id >= 0 {
		return getJobResultById(id)
	} else {
		return getJobResults()
	}
}

// @Title getJobResultById
// @Description retrieves the job_result information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    JobResult
// @Resource /api/2.0
// @Router /api/2.0/job_result/{id} [get]
func getJobResultById(id int) (interface{}, error) {
	ret := []JobResult{}
	arg := JobResult{Id: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from job_result where id=:id`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getJobResults
// @Description retrieves the job_result information for a certain id
// @Accept  application/json
// @Success 200 {array}    JobResult
// @Resource /api/2.0
// @Router /api/2.0/job_result [get]
func getJobResults() (interface{}, error) {
	ret := []JobResult{}
	queryStr := "select * from job_result"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postJobResult
// @Description enter a new job_result
// @Accept  application/json
// @Param                  Job json      int64   false "job description"
// @Param                Agent json      int64   false "agent description"
// @Param               Result json     string   false "result description"
// @Param          Description json     string    true "description description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/job_result [post]
func postJobResult(payload []byte) (interface{}, error) {
	var v JobResult
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO job_result("
	sqlString += "job"
	sqlString += ",agent"
	sqlString += ",result"
	sqlString += ",description"
	sqlString += ") VALUES ("
	sqlString += ":job"
	sqlString += ",:agent"
	sqlString += ",:result"
	sqlString += ",:description"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putJobResult
// @Description modify an existing job_resultentry
// @Accept  application/json
// @Param                  Job json      int64   false "job description"
// @Param                Agent json      int64   false "agent description"
// @Param               Result json     string   false "result description"
// @Param          Description json null.String    true "description description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/job_result [put]
func putJobResult(id int, payload []byte) (interface{}, error) {
	var v JobResult
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE job_result SET "
	sqlString += "job = :job"
	sqlString += ",agent = :agent"
	sqlString += ",result = :result"
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

// @Title delJobResultById
// @Description deletes job_result information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    JobResult
// @Resource /api/2.0
// @Router /api/2.0/job_result/{id} [delete]
func delJobResult(id int) (interface{}, error) {
	arg := JobResult{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM job_result WHERE id=:id", arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
