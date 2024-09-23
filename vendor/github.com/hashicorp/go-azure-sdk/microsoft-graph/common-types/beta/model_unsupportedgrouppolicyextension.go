package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnsupportedGroupPolicyExtension{}

type UnsupportedGroupPolicyExtension struct {
	// ExtensionType of the unsupported extension.
	ExtensionType nullable.Type[string] `json:"extensionType,omitempty"`

	// Namespace Url of the unsupported extension.
	NamespaceUrl nullable.Type[string] `json:"namespaceUrl,omitempty"`

	// Node name of the unsupported extension.
	NodeName nullable.Type[string] `json:"nodeName,omitempty"`

	// Scope of the group policy setting.
	SettingScope *GroupPolicySettingScope `json:"settingScope,omitempty"`

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

func (s UnsupportedGroupPolicyExtension) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnsupportedGroupPolicyExtension{}

func (s UnsupportedGroupPolicyExtension) MarshalJSON() ([]byte, error) {
	type wrapper UnsupportedGroupPolicyExtension
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnsupportedGroupPolicyExtension: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnsupportedGroupPolicyExtension: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unsupportedGroupPolicyExtension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnsupportedGroupPolicyExtension: %+v", err)
	}

	return encoded, nil
}
