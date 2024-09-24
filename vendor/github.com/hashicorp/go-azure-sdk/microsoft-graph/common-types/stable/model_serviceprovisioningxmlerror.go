package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceProvisioningError = ServiceProvisioningXmlError{}

type ServiceProvisioningXmlError struct {
	// Error Information published by the Federated Service as an xml string.
	ErrorDetail nullable.Type[string] `json:"errorDetail,omitempty"`

	// Fields inherited from ServiceProvisioningError

	// The date and time at which the error occurred.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Indicates whether the error has been attended to.
	IsResolved nullable.Type[bool] `json:"isResolved,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Qualified service instance (for example, 'SharePoint/Dublin') that published the service error information.
	ServiceInstance nullable.Type[string] `json:"serviceInstance,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ServiceProvisioningXmlError) ServiceProvisioningError() BaseServiceProvisioningErrorImpl {
	return BaseServiceProvisioningErrorImpl{
		CreatedDateTime: s.CreatedDateTime,
		IsResolved:      s.IsResolved,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
		ServiceInstance: s.ServiceInstance,
	}
}

var _ json.Marshaler = ServiceProvisioningXmlError{}

func (s ServiceProvisioningXmlError) MarshalJSON() ([]byte, error) {
	type wrapper ServiceProvisioningXmlError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceProvisioningXmlError: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceProvisioningXmlError: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.serviceProvisioningXmlError"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceProvisioningXmlError: %+v", err)
	}

	return encoded, nil
}
