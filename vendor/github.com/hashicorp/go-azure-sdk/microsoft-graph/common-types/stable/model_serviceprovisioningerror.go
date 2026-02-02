package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceProvisioningError interface {
	ServiceProvisioningError() BaseServiceProvisioningErrorImpl
}

var _ ServiceProvisioningError = BaseServiceProvisioningErrorImpl{}

type BaseServiceProvisioningErrorImpl struct {
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

func (s BaseServiceProvisioningErrorImpl) ServiceProvisioningError() BaseServiceProvisioningErrorImpl {
	return s
}

var _ ServiceProvisioningError = RawServiceProvisioningErrorImpl{}

// RawServiceProvisioningErrorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawServiceProvisioningErrorImpl struct {
	serviceProvisioningError BaseServiceProvisioningErrorImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawServiceProvisioningErrorImpl) ServiceProvisioningError() BaseServiceProvisioningErrorImpl {
	return s.serviceProvisioningError
}

func UnmarshalServiceProvisioningErrorImplementation(input []byte) (ServiceProvisioningError, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceProvisioningError into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceProvisioningXmlError") {
		var out ServiceProvisioningXmlError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceProvisioningXmlError: %+v", err)
		}
		return out, nil
	}

	var parent BaseServiceProvisioningErrorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseServiceProvisioningErrorImpl: %+v", err)
	}

	return RawServiceProvisioningErrorImpl{
		serviceProvisioningError: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
