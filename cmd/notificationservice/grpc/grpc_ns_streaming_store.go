package grpcnotify

import (
	"errors"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"
	"wwwin-github.cisco.com/eti/fledge/pkg/util"

	pbNotification "wwwin-github.cisco.com/eti/fledge/pkg/proto/go/notification"
)

// SetupAgentStream is called by the client to subscribe to the notification service.
// Adds the client to the server client map and stores the client stream.
func (s *notificationServer) SetupAgentStream(in *pbNotification.AgentInfo, stream pbNotification.NotificationStreamingStore_SetupAgentStreamServer) error {
	s.addNewClient(in, &stream)

	// the stream should not be killed so we do not return from this server
	// loop infinitely to keep stream alive else this stream will be closed
	for {
	}
	return nil
}

// addNewClient is responsible to add new client to the server map.
func (s *notificationServer) addNewClient(in *pbNotification.AgentInfo, stream *pbNotification.NotificationStreamingStore_SetupAgentStreamServer) {
	uuid := in.GetUuid()
	zap.S().Debugf("Adding new client to the collection | %v", in)
	s.clientStreams[uuid] = stream
	s.clients[uuid] = in
}

// pushNotification is called to send notification to the specific clients.
func (s *notificationServer) pushNotification(clientID string, notifyType pbNotification.StreamResponse_ResponseType, in interface{}) error {
	zap.S().Debugf("Sending notification to client: %v", clientID)

	//Step 1 - create notification object
	m, err := util.StructToMapInterface(in)
	if err != nil {
		zap.S().Errorf("error converting notification object into map interface. %v", err)
		return err
	}
	details, err := structpb.NewStruct(m)
	if err != nil {
		zap.S().Errorf("error creating proto struct. %v", err)
		return err
	}

	//Step 2 - send notification
	return s.send(clientID, notifyType, details)
}

//pushNotifications send notification to list of clients in on go
func (s *notificationServer) pushNotifications(clients []string, notifyType pbNotification.StreamResponse_ResponseType, in interface{}) map[string]interface{} {
	//create single notification object since same object is sent to all the clients
	m, err := util.StructToMapInterface(in)
	eList := map[string]interface{}{}
	if err != nil {
		zap.S().Errorf("error converting notification object into map interface. %v", err)
		eList[util.GenericError] = err
		return eList
	}
	details, err := structpb.NewStruct(m)
	if err != nil {
		zap.S().Errorf("error creating proto struct. %v", err)
		eList[util.GenericError] = err
		return eList
	}

	//iterate and send notification to all the clients
	for _, clientID := range clients {
		err = s.send(clientID, notifyType, details)
		if err != nil {
			eList[clientID] = err
		}
	}

	//if no error
	if len(eList) == 0 {
		eList = nil
	}
	return eList
}

// broadcastNotification is called to broadcast notification to all the connected clients.
func (s *notificationServer) broadcastNotification(notifyType pbNotification.StreamResponse_ResponseType, in interface{}) map[string]interface{} {
	var clients []string
	for clientID := range s.clients {
		clients = append(clients, clientID)
	}
	return s.pushNotifications(clients, notifyType, in)
}

//send implements helper method to push notification to the corresponding streams
func (s *notificationServer) send(clientID string, notifyType pbNotification.StreamResponse_ResponseType, msg *structpb.Struct) error {
	stream := s.clientStreams[clientID]

	if stream == nil {
		zap.S().Errorf("clientId %s is not registerred with the notifiaction service", clientID)
		return errors.New("client not registered with the service")
	}

	resp := pbNotification.StreamResponse{
		Type:      notifyType,
		Message:   msg,
		AgentUuid: clientID,
	}

	err := (*stream).Send(&resp)
	if err != nil {
		zap.S().Errorf("notification send error. clientId: %s | %v", clientID, err)
	}
	return err
}
