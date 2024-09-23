package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MicrosoftTunnelServer{}

type MicrosoftTunnelServer struct {
	// The digest of the current agent image running on this server. Supports: $filter, $select, $top, $skip, $orderby.
	// $search is not supported. Read-only.
	AgentImageDigest nullable.Type[string] `json:"agentImageDigest,omitempty"`

	// Microsoft Tunnel server deployment mode. The value is set when the server is registered. Possible values are
	// standaloneRootful, standaloneRootless, podRootful, podRootless. Default value: standaloneRootful. Supports: $filter,
	// $select, $top, $skip, $orderby. $search is not supported. Read-only. Possible values are: standaloneRootful,
	// standaloneRootless, podRootful, podRootless, unknownFutureValue.
	DeploymentMode *MicrosoftTunnelDeploymentMode `json:"deploymentMode,omitempty"`

	// The display name of the server. It is the same as the host name during registration and can be changed later.
	// Supports: $filter, $select, $top, $skip, $orderby. $search is not supported. Max allowed length is 200 chars.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates when the server last checked in. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this:
	// '2014-01-01T00:00:00Z'. Supports: $filter, $select, $top, $skip, $orderby. $search is not supported Read-only.
	LastCheckinDateTime *string `json:"lastCheckinDateTime,omitempty"`

	// The digest of the current server image running on this server. Supports: $filter, $select, $top, $skip, $orderby.
	// $search is not supported. Read-only.
	ServerImageDigest nullable.Type[string] `json:"serverImageDigest,omitempty"`

	// Enum of possible MicrosoftTunnelServer health status types
	TunnelServerHealthStatus *MicrosoftTunnelServerHealthStatus `json:"tunnelServerHealthStatus,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MicrosoftTunnelServer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MicrosoftTunnelServer{}

func (s MicrosoftTunnelServer) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftTunnelServer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftTunnelServer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftTunnelServer: %+v", err)
	}

	delete(decoded, "agentImageDigest")
	delete(decoded, "deploymentMode")
	delete(decoded, "lastCheckinDateTime")
	delete(decoded, "serverImageDigest")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftTunnelServer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftTunnelServer: %+v", err)
	}

	return encoded, nil
}
