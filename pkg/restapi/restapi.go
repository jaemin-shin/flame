// Copyright 2022 Cisco Systems, Inc. and its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"text/template"

	"go.uber.org/zap"
)

const (
	// Keys for dataset endpoints
	CreateDatasetEndPoint  = "CREATE_DATASET"
	GetDatasetEndPoint     = "GET_DATASET"
	GetDatasetsEndPoint    = "GET_DATASETS"
	GetAllDatasetsEndPoint = "GET_ALL_DATASETS"

	// Keys for design endpoints
	CreateDesignEndPoint = "CREATE_DESIGN"
	GetDesignsEndPoint   = "GET_DESIGNS"
	GetDesignEndPoint    = "GET_DESIGN"
	DeleteDesignEndPoint = "DELETE_DESIGN"

	// Keys for design schema endpoints
	CreateDesignSchemaEndPoint = "CREATE_DESIGN_SCHEMA"
	GetDesignSchemaEndPoint    = "GET_DESIGN_SCHEMA"
	GetDesignSchemasEndPoint   = "GET_DESIGN_SCHEMAS"
	UpdateDesignSchemaEndPoint = "UPDATE_DESIGN_SCHEMA"
	DeleteDesignSchemaEndPoint = "DELETE_DESIGN_SCHEMA"

	// Keys for design code endpoints
	CreateDesignCodeEndPoint = "CREATE_DESIGN_CODE"
	GetDesignCodeEndPoint    = "GET_DESIGN_CODE"
	UpdateDesignCodeEndPoint = "UPDATE_DESIGN_CODE"
	DeleteDesignCodeEndPoint = "DELETE_DESIGN_CODE"

	// Keys for job endpoints
	CreateJobEndpoint        = "CREATE_JOB"
	GetJobEndPoint           = "GET_JOB"
	GetJobsEndPoint          = "GET_JOBS"
	GetJobsByComputeEndPoint = "GET_JOBS_BY_COMPUTE"
	GetJobStatusEndPoint     = "GET_JOB_STATUS"
	GetTasksInfoEndpoint     = "GET_TASKS_INFO"
	GetTaskInfoEndpoint      = "GET_TASK_INFO"
	DeleteJobEndPoint        = "DELETE_JOB"
	UpdateJobEndPoint        = "UPDATE_JOB"
	ChangeJobSchemaEndPoint  = "CHANGE_SCHEMA_JOB"
	UpdateJobStatusEndPoint  = "UPDATE_JOB_STATUS"

	// Keys for task
	GetTaskEndpoint          = "GET_TASK"
	UpdateTaskStatusEndPoint = "UPDATE_TASK_STATUS"

	// Keys for computes endpoints
	RegisterComputeEndpoint     = "REGISTER_COMPUTE"
	GetComputeStatusEndpoint    = "GET_COMPUTE_STATUS"
	UpdateComputeEndpoint       = "UPDATE_COMPUTE"
	DeleteComputeEndpoint       = "DELETE_COMPUTE"
	GetComputeConfigEndpoint    = "GET_COMPUTE_CONFIG"
	GetDeploymentsEndpoint      = "GET_DEPLOYMENTS"
	GetDeploymentConfigEndpoint = "GET_DEPLOYMENT_CONFIG"
	PutDeploymentStatusEndpoint = "PUT_DEPLOYMENT_STATUS"
	GetDeploymentStatusEndpoint = "GET_DEPLOYMENT_STATUS"
)

var URI = map[string]string{
	// Dataset
	CreateDatasetEndPoint:  "/users/{{.user}}/datasets",
	GetDatasetEndPoint:     "/users/{{.user}}/datasets/{{.datasetId}}",
	GetDatasetsEndPoint:    "/users/{{.user}}/datasets/?limit={{.limit}}",
	GetAllDatasetsEndPoint: "/datasets/?limit={{.limit}}",

	// Design
	CreateDesignEndPoint: "/users/{{.user}}/designs",
	GetDesignEndPoint:    "/users/{{.user}}/designs/{{.designId}}",
	GetDesignsEndPoint:   "/users/{{.user}}/designs/?limit={{.limit}}",
	DeleteDesignEndPoint: "/users/{{.user}}/designs/{{.designId}}",

	// Design schema
	CreateDesignSchemaEndPoint: "/users/{{.user}}/designs/{{.designId}}/schemas",
	GetDesignSchemaEndPoint:    "/users/{{.user}}/designs/{{.designId}}/schemas/{{.version}}",
	GetDesignSchemasEndPoint:   "/users/{{.user}}/designs/{{.designId}}/schemas",
	UpdateDesignSchemaEndPoint: "/users/{{.user}}/designs/{{.designId}}/schemas/{{.version}}",
	DeleteDesignSchemaEndPoint: "/users/{{.user}}/designs/{{.designId}}/schemas/{{.version}}",

	// Design Code
	CreateDesignCodeEndPoint: "/users/{{.user}}/designs/{{.designId}}/codes",
	GetDesignCodeEndPoint:    "/users/{{.user}}/designs/{{.designId}}/codes/{{.version}}",
	UpdateDesignCodeEndPoint: "/users/{{.user}}/designs/{{.designId}}/codes/{{.version}}",
	DeleteDesignCodeEndPoint: "/users/{{.user}}/designs/{{.designId}}/codes/{{.version}}",

	// Job
	GetJobsByComputeEndPoint: "/jobs/{{.computeId}}",
	CreateJobEndpoint:        "/users/{{.user}}/jobs",
	GetJobEndPoint:           "/users/{{.user}}/jobs/{{.jobId}}",
	GetJobsEndPoint:          "/users/{{.user}}/jobs/?limit={{.limit}}",
	GetJobStatusEndPoint:     "/users/{{.user}}/jobs/{{.jobId}}/status",
	GetTasksInfoEndpoint:     "/users/{{.user}}/jobs/{{.jobId}}/tasks/?limit={{.limit}}&generic={{.generic}}",
	GetTaskInfoEndpoint:      "/users/{{.user}}/jobs/{{.jobId}}/tasks/{{.taskId}}",
	UpdateJobEndPoint:        "/users/{{.user}}/jobs/{{.jobId}}",
	DeleteJobEndPoint:        "/users/{{.user}}/jobs/{{.jobId}}",
	ChangeJobSchemaEndPoint:  "/users/{{.user}}/jobs/{{.jobId}}/schema/{{.schemaId}}/design/{{.designId}}",
	UpdateJobStatusEndPoint:  "/users/{{.user}}/jobs/{{.jobId}}/status",

	// Task
	GetTaskEndpoint:          "/jobs/{{.jobId}}/{{.taskId}}/task/?key={{.key}}",
	UpdateTaskStatusEndPoint: "/jobs/{{.jobId}}/{{.taskId}}/task/status",

	// Computes
	RegisterComputeEndpoint:     "/computes",
	GetComputeStatusEndpoint:    "/computes/{{.computeId}}",
	UpdateComputeEndpoint:       "/computes/{{.computeId}}",
	DeleteComputeEndpoint:       "/computes/{{.computeId}}",
	GetComputeConfigEndpoint:    "/computes/{{.computeId}}/config",
	GetDeploymentsEndpoint:      "/computes/{{.computeId}}/deployments",
	GetDeploymentConfigEndpoint: "/computes/{{.computeId}}/deployments/{{.jobId}}/config",
	PutDeploymentStatusEndpoint: "/computes/{{.computeId}}/deployments/{{.jobId}}/status",
	GetDeploymentStatusEndpoint: "/computes/{{.computeId}}/deployments/{{.jobId}}/status",
}

func FromTemplate(skeleton string, inputMap map[string]string) (string, error) {
	//https://stackoverflow.com/questions/29071212/implementing-dynamic-strings-in-golang
	var t = template.Must(template.New("").Parse(skeleton))
	buf := bytes.Buffer{}
	err := t.Execute(&buf, inputMap)
	if err != nil {
		zap.S().Errorf("error creating a text from skeleton. %v", err)
		return "", err
	}
	return buf.String(), nil
}

func CreateURL(hostEndpoint string, endPoint string, inputMap map[string]string) string {
	msg, err := FromTemplate(URI[endPoint], inputMap)
	if err != nil {
		zap.S().Errorf("error creating a uri. End point: %s", endPoint)
		return ""
	}

	return hostEndpoint + msg
}

func HTTPPost(url string, msg interface{}, contentType string) (int, []byte, error) {
	postBody, err := json.Marshal(msg)
	if err != nil {
		zap.S().Errorf("error encoding the payload")
		return -1, nil, err
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(postBody))
	if ErrorNilCheck(GetFunctionName(HTTPPost), err) != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := io.ReadAll(resp.Body)
	if ErrorNilCheck(GetFunctionName(HTTPPost), err) != nil {
		return -1, nil, err
	}

	return resp.StatusCode, body, nil
}

func HTTPPut(url string, msg interface{}, contentType string) (int, []byte, error) {
	putBody, err := json.Marshal(msg)
	if err != nil {
		zap.S().Errorf("error encoding the payload")
		return -1, nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(putBody))
	if ErrorNilCheck(GetFunctionName(HTTPPut), err) != nil {
		return -1, nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)

	if ErrorNilCheck(GetFunctionName(HTTPPut), err) != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := io.ReadAll(resp.Body)
	if ErrorNilCheck(GetFunctionName(HTTPPut), err) != nil {
		return -1, nil, err
	}

	return resp.StatusCode, body, nil
}

func HTTPDelete(url string, msg interface{}, contentType string) (int, []byte, error) {
	deleteBody, err := json.Marshal(msg)
	if err != nil {
		zap.S().Errorf("error encoding the payload for delete")
		return -1, nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(deleteBody))
	if ErrorNilCheck(GetFunctionName(HTTPDelete), err) != nil {
		return -1, nil, err
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)

	if ErrorNilCheck(GetFunctionName(HTTPDelete), err) != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := io.ReadAll(resp.Body)
	if ErrorNilCheck(GetFunctionName(HTTPDelete), err) != nil {
		return -1, nil, err
	}

	return resp.StatusCode, body, nil
}

func HTTPGet(url string) (int, []byte, error) {
	resp, err := http.Get(url)

	if ErrorNilCheck(GetFunctionName(HTTPGet), err) != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if ErrorNilCheck(GetFunctionName(HTTPGet), err) != nil {
		return -1, nil, err
	}

	return resp.StatusCode, body, nil
}

func HTTPGetMultipart(url string) (int, map[string][]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "multipart/form-data; charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return resp.StatusCode, nil, err
	}

	result := make(map[string][]byte)
	mr := multipart.NewReader(resp.Body, params["boundary"])
	for part, err := mr.NextPart(); err == nil; part, err = mr.NextPart() {
		data, err := io.ReadAll(part)
		if err != nil {
			return -1, nil, err
		}

		result[part.FormName()] = data
	}

	return resp.StatusCode, result, nil
}

func CreateMultipartFormData(kv map[string]io.Reader) (*bytes.Buffer, *multipart.Writer, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	for key, reader := range kv {
		var fw io.Writer
		var err error
		if x, ok := reader.(io.Closer); ok {
			defer x.Close()
		}

		if x, ok := reader.(*os.File); ok {
			// Add a file
			if fw, err = writer.CreateFormFile(key, x.Name()); err != nil {
				return nil, nil, err
			}
		} else {
			// Add other fields
			if fw, err = writer.CreateFormField(key); err != nil {
				return nil, nil, err
			}
		}

		if _, err = io.Copy(fw, reader); err != nil {
			return nil, nil, err
		}
	}
	writer.Close()

	return &buf, writer, nil
}

// ErrorNilCheck logger function to avoid re-writing the checks
func ErrorNilCheck(method string, err error) error {
	if err != nil {
		zap.S().Errorf("[%s] an error occurred %v", method, err)
	}

	return err
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func CheckStatusCode(code int) error {
	if code >= 400 && code <= 599 {
		return fmt.Errorf("status code: %d", code)
	}

	return nil
}
