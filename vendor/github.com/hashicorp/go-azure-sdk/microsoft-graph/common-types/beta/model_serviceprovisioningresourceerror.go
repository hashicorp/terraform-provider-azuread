package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceProvisioningError = ServiceProvisioningResourceError{}

type ServiceProvisioningResourceError struct {
	Errors *[]ServiceProvisioningResourceErrorDetail `json:"errors,omitempty"`

	// Fields inherited from ServiceProvisioningError

	// The date and time at which the error occurred.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Indicates whether the Error has been attended to.
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

func (s ServiceProvisioningResourceError) ServiceProvisioningError() BaseServiceProvisioningErrorImpl {
	return BaseServiceProvisioningErrorImpl{
		CreatedDateTime: s.CreatedDateTime,
		IsResolved:      s.IsResolved,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
		ServiceInstance: s.ServiceInstance,
	}
}

var _ json.Marshaler = ServiceProvisioningResourceError{}

func (s ServiceProvisioningResourceError) MarshalJSON() ([]byte, error) {
	type wrapper ServiceProvisioningResourceError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceProvisioningResourceError: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceProvisioningResourceError: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.serviceProvisioningResourceError"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceProvisioningResourceError: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ServiceProvisioningResourceError{}

func (s *ServiceProvisioningResourceError) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`
		IsResolved      nullable.Type[bool]   `json:"isResolved,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
		ServiceInstance nullable.Type[string] `json:"serviceInstance,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.IsResolved = decoded.IsResolved
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ServiceInstance = decoded.ServiceInstance

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ServiceProvisioningResourceError into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["errors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Errors into list []json.RawMessage: %+v", err)
		}

		output := make([]ServiceProvisioningResourceErrorDetail, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalServiceProvisioningResourceErrorDetailImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Errors' for 'ServiceProvisioningResourceError': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Errors = &output
	}

	return nil
}
