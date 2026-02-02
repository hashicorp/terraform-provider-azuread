package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = FileSecurityProfile{}

type FileSecurityProfile struct {
	ActivityGroupNames    *[]string                  `json:"activityGroupNames,omitempty"`
	AzureSubscriptionId   nullable.Type[string]      `json:"azureSubscriptionId,omitempty"`
	AzureTenantId         *string                    `json:"azureTenantId,omitempty"`
	CertificateThumbprint nullable.Type[string]      `json:"certificateThumbprint,omitempty"`
	Extensions            *[]string                  `json:"extensions,omitempty"`
	FileType              nullable.Type[string]      `json:"fileType,omitempty"`
	FirstSeenDateTime     nullable.Type[string]      `json:"firstSeenDateTime,omitempty"`
	Hashes                *[]FileHash                `json:"hashes,omitempty"`
	LastSeenDateTime      nullable.Type[string]      `json:"lastSeenDateTime,omitempty"`
	MalwareStates         *[]MalwareState            `json:"malwareStates,omitempty"`
	Names                 *[]string                  `json:"names,omitempty"`
	RiskScore             nullable.Type[string]      `json:"riskScore,omitempty"`
	Size                  nullable.Type[int64]       `json:"size,omitempty"`
	Tags                  *[]string                  `json:"tags,omitempty"`
	VendorInformation     *SecurityVendorInformation `json:"vendorInformation,omitempty"`
	VulnerabilityStates   *[]VulnerabilityState      `json:"vulnerabilityStates,omitempty"`

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

func (s FileSecurityProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = FileSecurityProfile{}

func (s FileSecurityProfile) MarshalJSON() ([]byte, error) {
	type wrapper FileSecurityProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FileSecurityProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FileSecurityProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.fileSecurityProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FileSecurityProfile: %+v", err)
	}

	return encoded, nil
}
