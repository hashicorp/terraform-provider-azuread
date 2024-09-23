package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCFrontLineServicePlan{}

type CloudPCFrontLineServicePlan struct {
	AllotmentLicensesCount nullable.Type[int64] `json:"allotmentLicensesCount,omitempty"`

	// The display name of the front-line service plan. For example, 2vCPU/8GB/128GB Front-line or 4vCPU/16GB/256GB
	// Front-line.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The total number of front-line service plans purchased by the customer.
	TotalCount nullable.Type[int64] `json:"totalCount,omitempty"`

	// The number of service plans that have been used for the account.
	UsedCount nullable.Type[int64] `json:"usedCount,omitempty"`

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

func (s CloudPCFrontLineServicePlan) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCFrontLineServicePlan{}

func (s CloudPCFrontLineServicePlan) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCFrontLineServicePlan
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCFrontLineServicePlan: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCFrontLineServicePlan: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcFrontLineServicePlan"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCFrontLineServicePlan: %+v", err)
	}

	return encoded, nil
}
