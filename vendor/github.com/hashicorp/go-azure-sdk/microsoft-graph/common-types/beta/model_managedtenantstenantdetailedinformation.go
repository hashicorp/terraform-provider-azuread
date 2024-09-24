package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsTenantDetailedInformation{}

type ManagedTenantsTenantDetailedInformation struct {
	// The city where the managed tenant is located. Optional. Read-only.
	City nullable.Type[string] `json:"city,omitempty"`

	// The code for the country where the managed tenant is located. Optional. Read-only.
	CountryCode nullable.Type[string] `json:"countryCode,omitempty"`

	// The name for the country where the managed tenant is located. Optional. Read-only.
	CountryName nullable.Type[string] `json:"countryName,omitempty"`

	// The default domain name for the managed tenant. Optional. Read-only.
	DefaultDomainName nullable.Type[string] `json:"defaultDomainName,omitempty"`

	// The display name for the managed tenant.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The business industry associated with the managed tenant. Optional. Read-only.
	IndustryName nullable.Type[string] `json:"industryName,omitempty"`

	// The region where the managed tenant is located. Optional. Read-only.
	Region nullable.Type[string] `json:"region,omitempty"`

	// The business segment associated with the managed tenant. Optional. Read-only.
	SegmentName nullable.Type[string] `json:"segmentName,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// The vertical associated with the managed tenant. Optional. Read-only.
	VerticalName nullable.Type[string] `json:"verticalName,omitempty"`

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

func (s ManagedTenantsTenantDetailedInformation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsTenantDetailedInformation{}

func (s ManagedTenantsTenantDetailedInformation) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsTenantDetailedInformation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsTenantDetailedInformation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsTenantDetailedInformation: %+v", err)
	}

	delete(decoded, "city")
	delete(decoded, "countryCode")
	delete(decoded, "countryName")
	delete(decoded, "defaultDomainName")
	delete(decoded, "industryName")
	delete(decoded, "region")
	delete(decoded, "segmentName")
	delete(decoded, "verticalName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.tenantDetailedInformation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsTenantDetailedInformation: %+v", err)
	}

	return encoded, nil
}
