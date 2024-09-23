package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ServicePrincipalCreationConditionSet{}

type ServicePrincipalCreationConditionSet struct {
	ApplicationIds                        *[]string           `json:"applicationIds,omitempty"`
	ApplicationPublisherIds               *[]string           `json:"applicationPublisherIds,omitempty"`
	ApplicationTenantIds                  *[]string           `json:"applicationTenantIds,omitempty"`
	ApplicationsFromVerifiedPublisherOnly nullable.Type[bool] `json:"applicationsFromVerifiedPublisherOnly,omitempty"`
	CertifiedApplicationsOnly             nullable.Type[bool] `json:"certifiedApplicationsOnly,omitempty"`

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

func (s ServicePrincipalCreationConditionSet) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServicePrincipalCreationConditionSet{}

func (s ServicePrincipalCreationConditionSet) MarshalJSON() ([]byte, error) {
	type wrapper ServicePrincipalCreationConditionSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServicePrincipalCreationConditionSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServicePrincipalCreationConditionSet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.servicePrincipalCreationConditionSet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServicePrincipalCreationConditionSet: %+v", err)
	}

	return encoded, nil
}
