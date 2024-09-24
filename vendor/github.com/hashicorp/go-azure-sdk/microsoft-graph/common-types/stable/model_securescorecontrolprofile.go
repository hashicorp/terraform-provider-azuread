package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecureScoreControlProfile{}

type SecureScoreControlProfile struct {
	// Control action type (Config, Review, Behavior).
	ActionType nullable.Type[string] `json:"actionType,omitempty"`

	// URL to where the control can be actioned.
	ActionUrl nullable.Type[string] `json:"actionUrl,omitempty"`

	// GUID string for tenant ID.
	AzureTenantId *string `json:"azureTenantId,omitempty"`

	// The collection of compliance information associated with secure score control
	ComplianceInformation *[]ComplianceInformation `json:"complianceInformation,omitempty"`

	// Control action category (Identity, Data, Device, Apps, Infrastructure).
	ControlCategory nullable.Type[string] `json:"controlCategory,omitempty"`

	// Flag to indicate where the tenant has marked a control (ignored, thirdParty, reviewed) (supports update).
	ControlStateUpdates *[]SecureScoreControlStateUpdate `json:"controlStateUpdates,omitempty"`

	// Flag to indicate if a control is depreciated.
	Deprecated nullable.Type[bool] `json:"deprecated,omitempty"`

	// Resource cost of implemmentating control (low, moderate, high).
	ImplementationCost nullable.Type[string] `json:"implementationCost,omitempty"`

	// Time at which the control profile entity was last modified. The Timestamp type represents date and time
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Microsoft's stack ranking of control.
	Rank nullable.Type[int64] `json:"rank,omitempty"`

	// Description of what the control will help remediate.
	Remediation nullable.Type[string] `json:"remediation,omitempty"`

	// Description of the impact on users of the remediation.
	RemediationImpact nullable.Type[string] `json:"remediationImpact,omitempty"`

	// Service that owns the control (Exchange, Sharepoint, Microsoft Entra ID).
	Service nullable.Type[string] `json:"service,omitempty"`

	// List of threats the control mitigates (accountBreach, dataDeletion, dataExfiltration, dataSpillage,
	// elevationOfPrivilege, maliciousInsider, passwordCracking, phishingOrWhaling, spoofing).
	Threats *[]string `json:"threats,omitempty"`

	// Control tier (Core, Defense in Depth, Advanced.)
	Tier nullable.Type[string] `json:"tier,omitempty"`

	// Title of the control.
	Title nullable.Type[string] `json:"title,omitempty"`

	// User impact of implementing control (low, moderate, high).
	UserImpact nullable.Type[string] `json:"userImpact,omitempty"`

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

func (s SecureScoreControlProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecureScoreControlProfile{}

func (s SecureScoreControlProfile) MarshalJSON() ([]byte, error) {
	type wrapper SecureScoreControlProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecureScoreControlProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecureScoreControlProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.secureScoreControlProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecureScoreControlProfile: %+v", err)
	}

	return encoded, nil
}
