/*
 * Job REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"wwwin-github.cisco.com/eti/fledge/cmd/controller/app/database"
	grpcctlr "wwwin-github.cisco.com/eti/fledge/cmd/controller/app/grpc"
	"wwwin-github.cisco.com/eti/fledge/pkg/objects"
	"wwwin-github.cisco.com/eti/fledge/pkg/openapi"
	pbNotification "wwwin-github.cisco.com/eti/fledge/pkg/proto/go/notification"
	"wwwin-github.cisco.com/eti/fledge/pkg/util"
)

// JobApiService is a service that implents the logic for the JobApiServicer
// This service should implement the business logic for every endpoint for the JobApi API.
// Include any external packages or services that will be required by this service.
type JobApiService struct {
}

// NewJobApiService creates a default api service
func NewJobApiService() openapi.JobApiServicer {
	return &JobApiService{}
}

// DeleteJob - Delete job by id.
func (s *JobApiService) DeleteJob(ctx context.Context, user string, jobId string) (openapi.ImplResponse, error) {
	// TODO - update DeleteJob with the required logic for this service method.
	// Add api_job_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation
	// when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("DeleteJob method not implemented")
}

// GetJob - Get job detail.
func (s *JobApiService) GetJob(ctx context.Context, user string, jobId string) (openapi.ImplResponse, error) {
	jInfo, err := database.GetJob(user, jobId)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("get job details request failed")
	}
	return openapi.Response(http.StatusOK, jInfo), nil
}

// SubmitJob - Submit a new job.
func (s *JobApiService) SubmitJob(ctx context.Context, user string, jobInfo openapi.JobInfo) (openapi.ImplResponse, error) {
	//insert in database. If failed, abort
	jId, err := database.SubmitJob(user, jobInfo)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("submit new job request failed")
	}

	// get design detail that is passed to the nodes as part of notification
	schemaInfo, err := database.GetDesignSchema(util.InternalUser, jobInfo.DesignId, "FIXME")
	if err != nil {
		err = fmt.Errorf("submit new job request failed: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	//Notify the agents about new job
	//Step 1 - get nodes for the job
	var agentInfo = getNodes(jobInfo.DesignId)

	//Step 2 - add nodes details into database
	err = database.UpdateJobDetails(jId, util.AddJobNodes, agentInfo)
	if err != nil {
		return openapi.Response(http.StatusMultiStatus, nil), errors.New("job request created but failed to initialized")
	}

	//Step 3 - update cache
	Cache.jobAgents[jId] = agentInfo
	Cache.jobSchema[jId] = schemaInfo

	//Step 4 - Notifying the agents of new job. Sending a init request allows to re-use the fledgelet nodes in the future, if needed.
	jobInfo.Id = jId
	jobMsg := objects.JobNotification{
		Agents:           agentInfo,
		Job:              jobInfo,
		SchemaInfo:       schemaInfo,
		NotificationType: util.InitState,
	}
	zap.S().Debugf("Sending notification to all the agents (count: %d) for new job id: %s. Info: %v", len(agentInfo), jId, jobMsg)
	resp, err := grpcctlr.ControllerGRPC.SendNotification(grpcctlr.JobNotification, jobMsg)
	if err != nil {
		zap.S().Errorf("failed to notify the agents. %v", err)
		return openapi.Response(http.StatusCreated, jId), err
	}

	//Check for partial error
	if resp.GetStatus() == pbNotification.Response_SUCCESS_WITH_ERROR {
		zap.S().Errorf("error while sending out new job notification for jobId: %s. Only partial clients notified.", jId)
		msResponse := map[string]interface{}{
			util.ID:     jId,
			util.Errors: resp.GetDetails(),
		}
		return openapi.Response(http.StatusMultiStatus, msResponse), nil
	}
	return openapi.Response(http.StatusCreated, map[string]string{util.ID: jId}), nil
}

// TODO: Code related to calling cluster manager to get nodes to be added here.
// For development purpose we are assuming that information is present in-memory
func getNodes(designId string) []openapi.ServerInfo {
	zap.S().Debugf("Getting nodes for designId: %s", designId)
	var agentInfo []openapi.ServerInfo
	for _, node := range JobNodesInMem[designId].Nodes {
		node.State = util.InitState

		agentInfo = append(agentInfo, node)
	}
	return agentInfo
}

// UpdateJob - Update job by id.
func (s *JobApiService) UpdateJob(ctx context.Context, user string, jobId string, jobInfo openapi.JobInfo) (openapi.ImplResponse, error) {
	// TODO - update UpdateJobEndPoint with the required logic for this service method.
	// Add api_job_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation
	// when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("UpdateJobEndPoint method not implemented")
}

// ChangeJobSchema - Change the schema for the given job
func (s *JobApiService) ChangeJobSchema(ctx context.Context, user string, jobId string, newSchemaId string,
	designId string) (openapi.ImplResponse, error) {
	//step 1 - get old schema details
	oldSchema := Cache.jobSchema[jobId]

	//step 2 - get schema details for new schema id
	//get design detail that is passed to the nodes as part of notification
	res, err := database.GetDesignSchema(util.InternalUser, designId, "FIXME")
	if err != nil {
		err = fmt.Errorf("change job schema request failed - job: %s | new schema id: %s | %v", jobId, newSchemaId, err)
		zap.S().Error(err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	newSchema := res
	zap.S().Debugf("schema : %v | %v", oldSchema, newSchema)

	//step 3 - check the old and new schema to determine the -new nodes required and -determine changes in the existing nodes
	existingNodes, newNodes := getNodesToNotify(designId)

	//step 4 - update the job details in Database - 1) change schema id to new schema id and 2) add new nodes
	err = database.UpdateJobDetails(jobId, util.ChangeJobSchema, map[string]interface{}{
		util.DBFieldSchemaId: newSchemaId,
		util.DBFieldNodes:    newNodes,
	})
	if err != nil {
		zap.S().Errorf("error while updating the schema for existing job: %s | new schema id: %s. %v", jobId, newSchemaId, err)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("error updating the schema for the given job")
	}

	//Step 5 - Since database has been updated update or invalidate the cache
	cacheAgentList := Cache.jobAgents[jobId]
	Cache.jobAgents[jobId] = append(cacheAgentList, newNodes...)
	Cache.jobSchema[jobId] = newSchema

	//step 6 - Get updated job information
	jobInfo, err := database.GetJob(util.InternalUser, jobId)
	if err != nil {
		err = fmt.Errorf("getting updated job info failed - job: %s | new schema id: %s | %v", jobId, newSchemaId, err)
		zap.S().Error(err)
		return openapi.Response(http.StatusMultiStatus, nil), err
	}

	//step 7 - Send corresponding notifications to the nodes.
	sendNotification := func(agentList []openapi.ServerInfo, nsType string, nodeType string) bool {
		isError := false
		if len(agentList) > 0 {
			jobMsg := objects.JobNotification{
				Agents:           agentList,
				Job:              jobInfo,
				SchemaInfo:       newSchema,
				NotificationType: nsType,
			}
			zap.S().Debugf("Notifying %d nodes (type: %s) for job: %s; msg: %v", len(agentList), nodeType, jobId, jobMsg)
			resp, err := grpcctlr.ControllerGRPC.SendNotification(grpcctlr.JobNotification, jobMsg)

			if err != nil {
				// TODO: currently we are ignoring the failure. Add a flag to check
				//       the job configuration if setting is to stop or continue
				zap.S().Errorf("failed to notify the new nodes about the job: %v", err)
				isError = true
			}

			// Check for partial error
			if resp.GetStatus() == pbNotification.Response_SUCCESS_WITH_ERROR {
				msgTemplate := "error while sending out job notification to %s nodes for job: %s; clients notified partially"
				zap.S().Errorf(msgTemplate, nodeType, jobId)
				isError = true
				//TODO what steps should be taken if an error occurs for few of the new nodes.
				//msResponse := map[string]interface{}{
				//	util.ID:     jId,
				//	util.Errors: resp.GetDetails(),
				//}
				//return errors.New("failed to notify all the new nodes. Only partial nodes we notified")
				//return openapi.Response(http.StatusMultiStatus, msResponse), nil
			}
		}
		return isError
	}

	nsError := sendNotification(newNodes, util.InitState, "NEW")
	nsError = sendNotification(existingNodes, util.InitState, "EXISTING") || nsError
	if !nsError {
		err := fmt.Errorf("successfully changed the schema for the existing job; error while notifying the new/existing nodes")
		return openapi.Response(http.StatusMultiStatus, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}

// TODO: Based on the changes in the schema determine 1) the new worker nodes that are required to be added
// and 2) nodes that will get affected by schema update
func getNodesToNotify(designId string) ([]openapi.ServerInfo, []openapi.ServerInfo) {
	var existingAgentInfo []openapi.ServerInfo
	var newAgentInfo []openapi.ServerInfo

	for i, node := range JobNodesInMem[designId].Nodes {
		if !node.IsExistingNode {
			node.State = util.InitState
			newAgentInfo = append(newAgentInfo, node)
		} else if node.IsUpdated {
			existingAgentInfo = append(existingAgentInfo, node)
		}

		//since the node information is used time they will now be considered as existing nodes
		JobNodesInMem[designId].Nodes[i].IsExistingNode = true
		JobNodesInMem[designId].Nodes[i].IsUpdated = false
	}

	zap.S().Debugf("new nodes: %v for job: %s", newAgentInfo)
	zap.S().Debugf("existing nodes: %v for job: %s", existingAgentInfo)
	return existingAgentInfo, newAgentInfo
}
