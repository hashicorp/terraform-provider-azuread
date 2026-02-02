package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PartnerSecurityPartnerSecurity{}

type PartnerSecurityPartnerSecurity struct {
	// The security alerts or a vulnerability of a Cloud Solution Provider (CSP) partner's customer that the partner must be
	// made aware of for further action.
	SecurityAlerts *[]PartnerSecurityPartnerSecurityAlert `json:"securityAlerts,omitempty"`

	// The security score calculated for the CSP partner and their customers.
	SecurityScore *PartnerSecurityPartnerSecurityScore `json:"securityScore,omitempty"`

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

func (s PartnerSecurityPartnerSecurity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PartnerSecurityPartnerSecurity{}

func (s PartnerSecurityPartnerSecurity) MarshalJSON() ([]byte, error) {
	type wrapper PartnerSecurityPartnerSecurity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PartnerSecurityPartnerSecurity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnerSecurityPartnerSecurity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.partner.security.partnerSecurity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PartnerSecurityPartnerSecurity: %+v", err)
	}

	return encoded, nil
}
