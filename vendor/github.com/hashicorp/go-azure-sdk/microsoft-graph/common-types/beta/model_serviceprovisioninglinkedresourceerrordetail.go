package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceProvisioningResourceErrorDetail = ServiceProvisioningLinkedResourceErrorDetail{}

type ServiceProvisioningLinkedResourceErrorDetail struct {
	PropertyName nullable.Type[string] `json:"propertyName,omitempty"`
	Target       nullable.Type[string] `json:"target,omitempty"`

	// Fields inherited from ServiceProvisioningResourceErrorDetail

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

func (s ServiceProvisioningLinkedResourceErrorDetail) ServiceProvisioningResourceErrorDetail() BaseServiceProvisioningResourceErrorDetailImpl {
	return BaseServiceProvisioningResourceErrorDetailImpl{
		Code:      s.Code,
		Details:   s.Details,
		Message:   s.Message,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServiceProvisioningLinkedResourceErrorDetail{}

func (s ServiceProvisioningLinkedResourceErrorDetail) MarshalJSON() ([]byte, error) {
	type wrapper ServiceProvisioningLinkedResourceErrorDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceProvisioningLinkedResourceErrorDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceProvisioningLinkedResourceErrorDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.serviceProvisioningLinkedResourceErrorDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceProvisioningLinkedResourceErrorDetail: %+v", err)
	}

	return encoded, nil
}
