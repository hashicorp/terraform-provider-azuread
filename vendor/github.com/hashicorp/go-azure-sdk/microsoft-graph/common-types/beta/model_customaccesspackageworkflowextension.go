package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomCalloutExtension = CustomAccessPackageWorkflowExtension{}

type CustomAccessPackageWorkflowExtension struct {
	// Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Fields inherited from CustomCalloutExtension

	// Configuration for securing the API call to the logic app. For example, using OAuth client credentials flow.
	AuthenticationConfiguration CustomExtensionAuthenticationConfiguration `json:"authenticationConfiguration"`

	// HTTP connection settings that define how long Microsoft Entra ID can wait for a connection to a logic app, how many
	// times you can retry a timed-out connection and the exception scenarios when retries are allowed.
	ClientConfiguration *CustomExtensionClientConfiguration `json:"clientConfiguration,omitempty"`

	// Description for the customCalloutExtension object.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name for the customCalloutExtension object.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The type and details for configuring the endpoint to call the logic app's workflow.
	EndpointConfiguration CustomExtensionEndpointConfiguration `json:"endpointConfiguration"`

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

func (s CustomAccessPackageWorkflowExtension) CustomCalloutExtension() BaseCustomCalloutExtensionImpl {
	return BaseCustomCalloutExtensionImpl{
		AuthenticationConfiguration: s.AuthenticationConfiguration,
		ClientConfiguration:         s.ClientConfiguration,
		Description:                 s.Description,
		DisplayName:                 s.DisplayName,
		EndpointConfiguration:       s.EndpointConfiguration,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s CustomAccessPackageWorkflowExtension) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CustomAccessPackageWorkflowExtension{}

func (s CustomAccessPackageWorkflowExtension) MarshalJSON() ([]byte, error) {
	type wrapper CustomAccessPackageWorkflowExtension
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomAccessPackageWorkflowExtension: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomAccessPackageWorkflowExtension: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customAccessPackageWorkflowExtension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomAccessPackageWorkflowExtension: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CustomAccessPackageWorkflowExtension{}

func (s *CustomAccessPackageWorkflowExtension) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string]               `json:"createdDateTime,omitempty"`
		LastModifiedDateTime nullable.Type[string]               `json:"lastModifiedDateTime,omitempty"`
		ClientConfiguration  *CustomExtensionClientConfiguration `json:"clientConfiguration,omitempty"`
		Description          nullable.Type[string]               `json:"description,omitempty"`
		DisplayName          nullable.Type[string]               `json:"displayName,omitempty"`
		Id                   *string                             `json:"id,omitempty"`
		ODataId              *string                             `json:"@odata.id,omitempty"`
		ODataType            *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ClientConfiguration = decoded.ClientConfiguration
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CustomAccessPackageWorkflowExtension into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authenticationConfiguration"]; ok {
		impl, err := UnmarshalCustomExtensionAuthenticationConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthenticationConfiguration' for 'CustomAccessPackageWorkflowExtension': %+v", err)
		}
		s.AuthenticationConfiguration = impl
	}

	if v, ok := temp["endpointConfiguration"]; ok {
		impl, err := UnmarshalCustomExtensionEndpointConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EndpointConfiguration' for 'CustomAccessPackageWorkflowExtension': %+v", err)
		}
		s.EndpointConfiguration = impl
	}

	return nil
}
