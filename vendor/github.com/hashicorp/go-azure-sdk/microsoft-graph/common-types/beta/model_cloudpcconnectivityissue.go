package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCConnectivityIssue{}

type CloudPCConnectivityIssue struct {
	// The Intune DeviceId of the device the connection is associated with.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The error code of the connectivity issue.
	ErrorCode nullable.Type[string] `json:"errorCode,omitempty"`

	// The time that the connection initiated. The time is shown in ISO 8601 format and Coordinated Universal Time (UTC)
	// time.
	ErrorDateTime *string `json:"errorDateTime,omitempty"`

	// The detailed description of what went wrong.
	ErrorDescription nullable.Type[string] `json:"errorDescription,omitempty"`

	// The recommended action to fix the corresponding error.
	RecommendedAction nullable.Type[string] `json:"recommendedAction,omitempty"`

	// The unique id of user who initialize the connection.
	UserId nullable.Type[string] `json:"userId,omitempty"`

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

func (s CloudPCConnectivityIssue) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCConnectivityIssue{}

func (s CloudPCConnectivityIssue) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCConnectivityIssue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCConnectivityIssue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCConnectivityIssue: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPCConnectivityIssue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCConnectivityIssue: %+v", err)
	}

	return encoded, nil
}
