package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DomainSecurityProfile{}

type DomainSecurityProfile struct {
	ActivityGroupNames       *[]string                  `json:"activityGroupNames,omitempty"`
	AzureSubscriptionId      nullable.Type[string]      `json:"azureSubscriptionId,omitempty"`
	AzureTenantId            *string                    `json:"azureTenantId,omitempty"`
	CountHits                nullable.Type[int64]       `json:"countHits,omitempty"`
	CountInOrg               nullable.Type[int64]       `json:"countInOrg,omitempty"`
	DomainCategories         *[]ReputationCategory      `json:"domainCategories,omitempty"`
	DomainRegisteredDateTime nullable.Type[string]      `json:"domainRegisteredDateTime,omitempty"`
	FirstSeenDateTime        nullable.Type[string]      `json:"firstSeenDateTime,omitempty"`
	LastSeenDateTime         nullable.Type[string]      `json:"lastSeenDateTime,omitempty"`
	Name                     nullable.Type[string]      `json:"name,omitempty"`
	Registrant               *DomainRegistrant          `json:"registrant,omitempty"`
	RiskScore                nullable.Type[string]      `json:"riskScore,omitempty"`
	Tags                     *[]string                  `json:"tags,omitempty"`
	VendorInformation        *SecurityVendorInformation `json:"vendorInformation,omitempty"`

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

func (s DomainSecurityProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DomainSecurityProfile{}

func (s DomainSecurityProfile) MarshalJSON() ([]byte, error) {
	type wrapper DomainSecurityProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DomainSecurityProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DomainSecurityProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.domainSecurityProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DomainSecurityProfile: %+v", err)
	}

	return encoded, nil
}
