package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SensitivityPolicySettings{}

type SensitivityPolicySettings struct {
	ApplicableTo                              *SensitivityLabelTarget `json:"applicableTo,omitempty"`
	DowngradeSensitivityRequiresJustification nullable.Type[bool]     `json:"downgradeSensitivityRequiresJustification,omitempty"`
	HelpWebUrl                                nullable.Type[string]   `json:"helpWebUrl,omitempty"`
	IsMandatory                               nullable.Type[bool]     `json:"isMandatory,omitempty"`

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

func (s SensitivityPolicySettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SensitivityPolicySettings{}

func (s SensitivityPolicySettings) MarshalJSON() ([]byte, error) {
	type wrapper SensitivityPolicySettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SensitivityPolicySettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SensitivityPolicySettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sensitivityPolicySettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SensitivityPolicySettings: %+v", err)
	}

	return encoded, nil
}
