package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HostSecurityProfile{}

type HostSecurityProfile struct {
	AzureSubscriptionId       nullable.Type[string]      `json:"azureSubscriptionId,omitempty"`
	AzureTenantId             *string                    `json:"azureTenantId,omitempty"`
	FirstSeenDateTime         nullable.Type[string]      `json:"firstSeenDateTime,omitempty"`
	Fqdn                      nullable.Type[string]      `json:"fqdn,omitempty"`
	IsAzureAdJoined           nullable.Type[bool]        `json:"isAzureAdJoined,omitempty"`
	IsAzureAdRegistered       nullable.Type[bool]        `json:"isAzureAdRegistered,omitempty"`
	IsHybridAzureDomainJoined nullable.Type[bool]        `json:"isHybridAzureDomainJoined,omitempty"`
	LastSeenDateTime          nullable.Type[string]      `json:"lastSeenDateTime,omitempty"`
	LogonUsers                *[]LogonUser               `json:"logonUsers,omitempty"`
	NetBiosName               nullable.Type[string]      `json:"netBiosName,omitempty"`
	NetworkInterfaces         *[]NetworkInterface        `json:"networkInterfaces,omitempty"`
	Os                        nullable.Type[string]      `json:"os,omitempty"`
	OsVersion                 nullable.Type[string]      `json:"osVersion,omitempty"`
	ParentHost                nullable.Type[string]      `json:"parentHost,omitempty"`
	RelatedHostIds            *[]string                  `json:"relatedHostIds,omitempty"`
	RiskScore                 nullable.Type[string]      `json:"riskScore,omitempty"`
	Tags                      *[]string                  `json:"tags,omitempty"`
	VendorInformation         *SecurityVendorInformation `json:"vendorInformation,omitempty"`

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

func (s HostSecurityProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HostSecurityProfile{}

func (s HostSecurityProfile) MarshalJSON() ([]byte, error) {
	type wrapper HostSecurityProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HostSecurityProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HostSecurityProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.hostSecurityProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HostSecurityProfile: %+v", err)
	}

	return encoded, nil
}
