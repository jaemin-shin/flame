/*
 * Fledge REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package controller

import (
	"context"
	"fmt"
	"net/http"

	"wwwin-github.cisco.com/eti/fledge/cmd/controller/app/database"
	"wwwin-github.cisco.com/eti/fledge/pkg/openapi"
)

// DesignSchemasApiService is a service that implents the logic for the DesignSchemasApiServicer
// This service should implement the business logic for every endpoint for the DesignSchemasApi API.
// Include any external packages or services that will be required by this service.
type DesignSchemasApiService struct {
}

// NewDesignSchemasApiService creates a default api service
func NewDesignSchemasApiService() openapi.DesignSchemasApiServicer {
	return &DesignSchemasApiService{}
}

// CreateDesignSchema - Update a design schema
func (s *DesignSchemasApiService) CreateDesignSchema(ctx context.Context, user string, designId string,
	designSchema openapi.DesignSchema) (openapi.ImplResponse, error) {
	err := database.CreateDesignSchema(user, designId, designSchema)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("insert design schema details request failed")
	}

	return openapi.Response(http.StatusOK, nil), nil
}

// GetDesignSchema - Get a design schema owned by user
func (s *DesignSchemasApiService) GetDesignSchema(ctx context.Context, user string, designId string,
	version string) (openapi.ImplResponse, error) {
	info, err := database.GetDesignSchema(user, designId, version)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("get design schema details request failed")
	}
	return openapi.Response(http.StatusOK, info), nil
}

// GetDesignSchemas - Get all design schemas in a design
func (s *DesignSchemasApiService) GetDesignSchemas(ctx context.Context, user string, designId string) (openapi.ImplResponse, error) {
	info, err := database.GetDesignSchemas(user, designId)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("get design schema details request failed")
	}
	return openapi.Response(http.StatusOK, info), nil
}

// UpdateDesignSchema - Update a schema for a given design
func (s *DesignSchemasApiService) UpdateDesignSchema(ctx context.Context, user string, designId string, version string,
	designSchema openapi.DesignSchema) (openapi.ImplResponse, error) {
	err := database.UpdateDesignSchema(user, designId, version, designSchema)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("schema update request failed")
	}

	return openapi.Response(http.StatusOK, nil), nil
}
