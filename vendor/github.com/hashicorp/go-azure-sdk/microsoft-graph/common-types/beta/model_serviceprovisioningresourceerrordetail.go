package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceProvisioningResourceErrorDetail interface {
	ServiceProvisioningResourceErrorDetail() BaseServiceProvisioningResourceErrorDetailImpl
}

var _ ServiceProvisioningResourceErrorDetail = BaseServiceProvisioningResourceErrorDetailImpl{}

type BaseServiceProvisioningResourceErrorDetailImpl struct {
	Code    nullable.Type[string] `json:"code,omitempty"`
	Details nullable.Type[string] `json:"details,omitempty"`
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseServiceProvisioningResourceErrorDetailImpl) ServiceProvisioningResourceErrorDetail() BaseServiceProvisioningResourceErrorDetailImpl {
	return s
}

var _ ServiceProvisioningResourceErrorDetail = RawServiceProvisioningResourceErrorDetailImpl{}

// RawServiceProvisioningResourceErrorDetailImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawServiceProvisioningResourceErrorDetailImpl struct {
	serviceProvisioningResourceErrorDetail BaseServiceProvisioningResourceErrorDetailImpl
	Type                                   string
	Values                                 map[string]interface{}
}

func (s RawServiceProvisioningResourceErrorDetailImpl) ServiceProvisioningResourceErrorDetail() BaseServiceProvisioningResourceErrorDetailImpl {
	return s.serviceProvisioningResourceErrorDetail
}

func UnmarshalServiceProvisioningResourceErrorDetailImplementation(input []byte) (ServiceProvisioningResourceErrorDetail, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceProvisioningResourceErrorDetail into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceProvisioningLinkedResourceErrorDetail") {
		var out ServiceProvisioningLinkedResourceErrorDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceProvisioningLinkedResourceErrorDetail: %+v", err)
		}
		return out, nil
	}

	var parent BaseServiceProvisioningResourceErrorDetailImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseServiceProvisioningResourceErrorDetailImpl: %+v", err)
	}

	return RawServiceProvisioningResourceErrorDetailImpl{
		serviceProvisioningResourceErrorDetail: parent,
		Type:                                   value,
		Values:                                 temp,
	}, nil

}
