package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecureScore{}

type SecureScore struct {
	// Active user count of the given tenant.
	ActiveUserCount nullable.Type[int64] `json:"activeUserCount,omitempty"`

	// Average score by different scopes (for example, average by industry, average by seating) and control category
	// (Identity, Data, Device, Apps, Infrastructure) within the scope.
	AverageComparativeScores *[]AverageComparativeScore `json:"averageComparativeScores,omitempty"`

	// GUID string for tenant ID.
	AzureTenantId *string `json:"azureTenantId,omitempty"`

	// Contains tenant scores for a set of controls.
	ControlScores *[]ControlScore `json:"controlScores,omitempty"`

	// When the report was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Microsoft-provided services for the tenant (for example, Exchange online, Skype, Sharepoint).
	EnabledServices *[]string `json:"enabledServices,omitempty"`

	// Licensed user count of the given tenant.
	LicensedUserCount nullable.Type[int64] `json:"licensedUserCount,omitempty"`

	// Complex type containing details about the security product/service vendor, provider, and subprovider (for example,
	// vendor=Microsoft; provider=SecureScore). Required.
	VendorInformation SecurityVendorInformation `json:"vendorInformation"`

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

func (s SecureScore) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecureScore{}

func (s SecureScore) MarshalJSON() ([]byte, error) {
	type wrapper SecureScore
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecureScore: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecureScore: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.secureScore"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecureScore: %+v", err)
	}

	return encoded, nil
}
