package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstanceDecisionItemResource interface {
	AccessReviewInstanceDecisionItemResource() BaseAccessReviewInstanceDecisionItemResourceImpl
}

var _ AccessReviewInstanceDecisionItemResource = BaseAccessReviewInstanceDecisionItemResourceImpl{}

type BaseAccessReviewInstanceDecisionItemResourceImpl struct {
	// Display name of the resource
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Resource ID
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of resource. Types include: Group, ServicePrincipal, DirectoryRole, AzureRole, AccessPackageAssignmentPolicy.
	Type nullable.Type[string] `json:"type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAccessReviewInstanceDecisionItemResourceImpl) AccessReviewInstanceDecisionItemResource() BaseAccessReviewInstanceDecisionItemResourceImpl {
	return s
}

var _ AccessReviewInstanceDecisionItemResource = RawAccessReviewInstanceDecisionItemResourceImpl{}

// RawAccessReviewInstanceDecisionItemResourceImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAccessReviewInstanceDecisionItemResourceImpl struct {
	accessReviewInstanceDecisionItemResource BaseAccessReviewInstanceDecisionItemResourceImpl
	Type                                     string
	Values                                   map[string]interface{}
}

func (s RawAccessReviewInstanceDecisionItemResourceImpl) AccessReviewInstanceDecisionItemResource() BaseAccessReviewInstanceDecisionItemResourceImpl {
	return s.accessReviewInstanceDecisionItemResource
}

func (s RawAccessReviewInstanceDecisionItemResourceImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalAccessReviewInstanceDecisionItemResourceImplementation(input []byte) (AccessReviewInstanceDecisionItemResource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewInstanceDecisionItemResource into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource") {
		var out AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewInstanceDecisionItemAzureRoleResource") {
		var out AccessReviewInstanceDecisionItemAzureRoleResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewInstanceDecisionItemAzureRoleResource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewInstanceDecisionItemServicePrincipalResource") {
		var out AccessReviewInstanceDecisionItemServicePrincipalResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewInstanceDecisionItemServicePrincipalResource: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessReviewInstanceDecisionItemResourceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessReviewInstanceDecisionItemResourceImpl: %+v", err)
	}

	return RawAccessReviewInstanceDecisionItemResourceImpl{
		accessReviewInstanceDecisionItemResource: parent,
		Type:                                     value,
		Values:                                   temp,
	}, nil

}
