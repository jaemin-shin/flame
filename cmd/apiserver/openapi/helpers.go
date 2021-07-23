/*
 * Fledge REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import "wwwin-github.cisco.com/eti/fledge/pkg/util"

//Response return a ImplResponse struct filled
func Response(code int, body interface{}) ImplResponse {
	return ImplResponse{
		Code: code,
		Body: body,
	}
}

func CreateURI(endPoint string, uriMap map[string]string) string {
	//TODO - implement value retrieval from config/environment
	return util.CreateURI("localhost", util.ControllerRestApiPort, endPoint, uriMap)
}
