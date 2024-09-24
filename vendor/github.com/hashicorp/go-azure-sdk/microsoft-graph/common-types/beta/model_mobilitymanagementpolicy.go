package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MobilityManagementPolicy{}

type MobilityManagementPolicy struct {
	// Indicates the user scope of the mobility management policy. Possible values are: none, all, selected.
	AppliesTo *PolicyScope `json:"appliesTo,omitempty"`

	// Compliance URL of the mobility management application.
	ComplianceUrl nullable.Type[string] `json:"complianceUrl,omitempty"`

	// Description of the mobility management application.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Discovery URL of the mobility management application.
	DiscoveryUrl nullable.Type[string] `json:"discoveryUrl,omitempty"`

	// Display name of the mobility management application.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Microsoft Entra groups under the scope of the mobility management application if appliesTo is selected
	IncludedGroups *[]Group `json:"includedGroups,omitempty"`

	// Whether policy is valid. Invalid policies may not be updated and should be deleted.
	IsValid nullable.Type[bool] `json:"isValid,omitempty"`

	// Terms of Use URL of the mobility management application.
	TermsOfUseUrl nullable.Type[string] `json:"termsOfUseUrl,omitempty"`

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

func (s MobilityManagementPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobilityManagementPolicy{}

func (s MobilityManagementPolicy) MarshalJSON() ([]byte, error) {
	type wrapper MobilityManagementPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobilityManagementPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobilityManagementPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobilityManagementPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobilityManagementPolicy: %+v", err)
	}

	return encoded, nil
}
