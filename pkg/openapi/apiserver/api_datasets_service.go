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

/*
 * Flame REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/cisco-open/flame/pkg/openapi"
	"github.com/cisco-open/flame/pkg/openapi/constants"
	"github.com/cisco-open/flame/pkg/restapi"
	"github.com/cisco-open/flame/pkg/util"
)

// DatasetsApiService is a service that implements the logic for the DatasetsApiServicer
// This service should implement the business logic for every endpoint for the DatasetsApi API.
// Include any external packages or services that will be required by this service.
type DatasetsApiService struct {
}

// NewDatasetsApiService creates a default api service
func NewDatasetsApiService() openapi.DatasetsApiServicer {
	return &DatasetsApiService{}
}

// CreateDataset - Create meta info for a new dataset.
func (s *DatasetsApiService) CreateDataset(ctx context.Context, user string,
	datasetInfo openapi.DatasetInfo) (openapi.ImplResponse, error) {
	//TODO input validation
	zap.S().Debugf("New dataset request received for user: %s | datasetInfo: %v", user, datasetInfo)

	// create controller request
	uriMap := map[string]string{
		constants.ParamUser: user,
	}
	url := restapi.CreateURL(HostEndpoint, restapi.CreateDatasetEndPoint, uriMap)

	// send post request
	code, body, err := restapi.HTTPPost(url, datasetInfo, "application/json")
	errResp, retErr := errorResponse(code, body, err)
	if retErr != nil {
		return errResp, retErr
	}

	return openapi.Response(http.StatusCreated, string(body)), nil
}

// GetAllDatasets - Get the meta info on all the datasets
func (s *DatasetsApiService) GetAllDatasets(ctx context.Context, limit int32) (openapi.ImplResponse, error) {
	zap.S().Debugf("get list of open datasets for limit: %d", limit)

	uriMap := map[string]string{
		constants.ParamLimit: strconv.Itoa(int(limit)),
	}
	url := restapi.CreateURL(HostEndpoint, restapi.GetAllDatasetsEndPoint, uriMap)

	code, body, err := restapi.HTTPGet(url)
	errResp, retErr := errorResponse(code, body, err)
	if retErr != nil {
		return errResp, retErr
	}

	var datasetInfoList []openapi.DatasetInfo
	err = util.ByteToStruct(body, &datasetInfoList)

	return openapi.Response(http.StatusOK, datasetInfoList), err
}

// GetDataset - Get dataset meta information
func (s *DatasetsApiService) GetDataset(ctx context.Context, user string, datasetId string) (openapi.ImplResponse, error) {
	// TODO - update GetDataset with the required logic for this service method.
	// Add api_datasets_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, DatasetInfo{}) or use other options such as http.Ok ...
	//return Response(200, DatasetInfo{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetDataset method not implemented")
}

// GetDatasets - Get the meta info on all the datasets owned by user
func (s *DatasetsApiService) GetDatasets(ctx context.Context, user string, limit int32) (openapi.ImplResponse, error) {
	zap.S().Debugf("get list of datasets for user: %s | limit: %d", user, limit)

	//create controller request
	//construct URL
	uriMap := map[string]string{
		constants.ParamUser:  user,
		constants.ParamLimit: strconv.Itoa(int(limit)),
	}
	url := restapi.CreateURL(HostEndpoint, restapi.GetDatasetsEndPoint, uriMap)

	//send get request
	code, body, err := restapi.HTTPGet(url)
	errResp, retErr := errorResponse(code, body, err)
	if retErr != nil {
		return errResp, retErr
	}

	var datasetInfoList []openapi.DatasetInfo
	err = util.ByteToStruct(body, &datasetInfoList)

	return openapi.Response(http.StatusOK, datasetInfoList), err
}

// UpdateDataset - Update meta info for a given dataset
func (s *DatasetsApiService) UpdateDataset(ctx context.Context, user string, datasetId string,
	datasetInfo openapi.DatasetInfo) (openapi.ImplResponse, error) {
	// TODO - update UpdateDataset with the required logic for this service method.
	// Add api_datasets_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("UpdateDataset method not implemented")
}
