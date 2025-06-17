package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EnhancedPersonalizationSetting{}

type EnhancedPersonalizationSetting struct {
	// The ID of a Microsoft Entra group to which the value is used to disable the control for populated users. The default
	// value is null. This parameter is optional.
	DisabledForGroup nullable.Type[string] `json:"disabledForGroup,omitempty"`

	// If true, enables the enhanced personalization control and therefore related features as defined in control enhanced
	// personalization privacy
	IsEnabledInOrganization *bool `json:"isEnabledInOrganization,omitempty"`

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

func (s EnhancedPersonalizationSetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EnhancedPersonalizationSetting{}

func (s EnhancedPersonalizationSetting) MarshalJSON() ([]byte, error) {
	type wrapper EnhancedPersonalizationSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EnhancedPersonalizationSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EnhancedPersonalizationSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.enhancedPersonalizationSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EnhancedPersonalizationSetting: %+v", err)
	}

	return encoded, nil
}
