package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ServiceNowConnection{}

type ServiceNowConnection struct {
	// Indicates the method used by Intune to authenticate with ServiceNow. Currently supports only web authentication with
	// ServiceNow using the specified app id.
	AuthenticationMethod ServiceNowAuthenticationMethod `json:"authenticationMethod"`

	// Date Time when connection properties were created. The value cannot be modified and is automatically populated when
	// the connection properties were entered.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Indicates the ServiceNow incident API URL that Intune will use the fetch incidents. Saved in the format of
	// /api/now/table/incident
	IncidentApiUrl nullable.Type[string] `json:"incidentApiUrl,omitempty"`

	// Indicates the ServiceNow instance URL that Intune will connect to. Saved in the format of
	// https://<instance>.service-now.com
	InstanceUrl nullable.Type[string] `json:"instanceUrl,omitempty"`

	// Date Time when connection properties were last updated. The value cannot be modified and is automatically populated
	// when the connection properties were updated.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Date Time when incidents from ServiceNow were last queried
	LastQueriedDateTime nullable.Type[string] `json:"lastQueriedDateTime,omitempty"`

	// Status of ServiceNow Connection
	ServiceNowConnectionStatus *ServiceNowConnectionStatus `json:"serviceNowConnectionStatus,omitempty"`

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

func (s ServiceNowConnection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServiceNowConnection{}

func (s ServiceNowConnection) MarshalJSON() ([]byte, error) {
	type wrapper ServiceNowConnection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceNowConnection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceNowConnection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.serviceNowConnection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceNowConnection: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ServiceNowConnection{}

func (s *ServiceNowConnection) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime            nullable.Type[string]       `json:"createdDateTime,omitempty"`
		IncidentApiUrl             nullable.Type[string]       `json:"incidentApiUrl,omitempty"`
		InstanceUrl                nullable.Type[string]       `json:"instanceUrl,omitempty"`
		LastModifiedDateTime       nullable.Type[string]       `json:"lastModifiedDateTime,omitempty"`
		LastQueriedDateTime        nullable.Type[string]       `json:"lastQueriedDateTime,omitempty"`
		ServiceNowConnectionStatus *ServiceNowConnectionStatus `json:"serviceNowConnectionStatus,omitempty"`
		Id                         *string                     `json:"id,omitempty"`
		ODataId                    *string                     `json:"@odata.id,omitempty"`
		ODataType                  *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.IncidentApiUrl = decoded.IncidentApiUrl
	s.InstanceUrl = decoded.InstanceUrl
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.LastQueriedDateTime = decoded.LastQueriedDateTime
	s.ServiceNowConnectionStatus = decoded.ServiceNowConnectionStatus
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ServiceNowConnection into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authenticationMethod"]; ok {
		impl, err := UnmarshalServiceNowAuthenticationMethodImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthenticationMethod' for 'ServiceNowConnection': %+v", err)
		}
		s.AuthenticationMethod = impl
	}

	return nil
}
