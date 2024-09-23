package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityInformationProtectionPolicySetting{}

type SecurityInformationProtectionPolicySetting struct {
	DefaultLabelId nullable.Type[string] `json:"defaultLabelId,omitempty"`

	// Exposes whether justification input is required on label downgrade.
	IsDowngradeJustificationRequired *bool `json:"isDowngradeJustificationRequired,omitempty"`

	// Exposes whether mandatory labeling is enabled.
	IsMandatory *bool `json:"isMandatory,omitempty"`

	// Exposes the more information URL that can be configured by the administrator.
	MoreInfoUrl nullable.Type[string] `json:"moreInfoUrl,omitempty"`

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

func (s SecurityInformationProtectionPolicySetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityInformationProtectionPolicySetting{}

func (s SecurityInformationProtectionPolicySetting) MarshalJSON() ([]byte, error) {
	type wrapper SecurityInformationProtectionPolicySetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityInformationProtectionPolicySetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityInformationProtectionPolicySetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.informationProtectionPolicySetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityInformationProtectionPolicySetting: %+v", err)
	}

	return encoded, nil
}
