package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomCalloutExtension = AccessPackageAssignmentRequestWorkflowExtension{}

type AccessPackageAssignmentRequestWorkflowExtension struct {
	// The callback configuration for a custom extension.
	CallbackConfiguration CustomExtensionCallbackConfiguration `json:"callbackConfiguration"`

	// The userPrincipalName of the user or identity of the subject that created this resource. Read-only.
	CreatedBy nullable.Type[string] `json:"createdBy,omitempty"`

	// When the object was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The userPrincipalName of the identity that last modified the object.
	LastModifiedBy nullable.Type[string] `json:"lastModifiedBy,omitempty"`

	// When the object was last modified.
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

func (s AccessPackageAssignmentRequestWorkflowExtension) CustomCalloutExtension() BaseCustomCalloutExtensionImpl {
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

func (s AccessPackageAssignmentRequestWorkflowExtension) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageAssignmentRequestWorkflowExtension{}

func (s AccessPackageAssignmentRequestWorkflowExtension) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageAssignmentRequestWorkflowExtension
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageAssignmentRequestWorkflowExtension: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageAssignmentRequestWorkflowExtension: %+v", err)
	}

	delete(decoded, "createdBy")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageAssignmentRequestWorkflowExtension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageAssignmentRequestWorkflowExtension: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessPackageAssignmentRequestWorkflowExtension{}

func (s *AccessPackageAssignmentRequestWorkflowExtension) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedBy            nullable.Type[string]               `json:"createdBy,omitempty"`
		CreatedDateTime      nullable.Type[string]               `json:"createdDateTime,omitempty"`
		LastModifiedBy       nullable.Type[string]               `json:"lastModifiedBy,omitempty"`
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

	s.CreatedBy = decoded.CreatedBy
	s.CreatedDateTime = decoded.CreatedDateTime
	s.LastModifiedBy = decoded.LastModifiedBy
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ClientConfiguration = decoded.ClientConfiguration
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageAssignmentRequestWorkflowExtension into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authenticationConfiguration"]; ok {
		impl, err := UnmarshalCustomExtensionAuthenticationConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthenticationConfiguration' for 'AccessPackageAssignmentRequestWorkflowExtension': %+v", err)
		}
		s.AuthenticationConfiguration = impl
	}

	if v, ok := temp["callbackConfiguration"]; ok {
		impl, err := UnmarshalCustomExtensionCallbackConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CallbackConfiguration' for 'AccessPackageAssignmentRequestWorkflowExtension': %+v", err)
		}
		s.CallbackConfiguration = impl
	}

	if v, ok := temp["endpointConfiguration"]; ok {
		impl, err := UnmarshalCustomExtensionEndpointConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EndpointConfiguration' for 'AccessPackageAssignmentRequestWorkflowExtension': %+v", err)
		}
		s.EndpointConfiguration = impl
	}

	return nil
}
