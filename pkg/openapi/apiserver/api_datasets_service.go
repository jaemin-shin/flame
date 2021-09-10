/*
 * Fledge REST API
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
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"wwwin-github.cisco.com/eti/fledge/pkg/openapi"
	"wwwin-github.cisco.com/eti/fledge/pkg/restapi"
)

// DatasetsApiService is a service that implents the logic for the DatasetsApiServicer
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
		"user": user,
	}
	url := restapi.CreateURL(Host, Port, restapi.CreateDatasetEndPoint, uriMap)

	// send post request
	code, _, err := restapi.HTTPPost(url, datasetInfo, "application/json")

	// response to the user
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("create new dataset request failed")
	}

	if err = restapi.CheckStatusCode(code); err != nil {
		return openapi.Response(code, nil), err
	}

	return openapi.Response(http.StatusCreated, nil), nil
}

// GetAllDatasets - Get the meta info on all the datasets
func (s *DatasetsApiService) GetAllDatasets(ctx context.Context, limit int32) (openapi.ImplResponse, error) {
	// TODO - update GetAllDatasets with the required logic for this service method.
	// Add api_datasets_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []DatasetInfo{}) or use other options such as http.Ok ...
	//return Response(200, []DatasetInfo{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetAllDatasets method not implemented")
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
	// TODO - update GetDatasets with the required logic for this service method.
	// Add api_datasets_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []DatasetInfo{}) or use other options such as http.Ok ...
	//return Response(200, []DatasetInfo{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetDatasets method not implemented")
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
