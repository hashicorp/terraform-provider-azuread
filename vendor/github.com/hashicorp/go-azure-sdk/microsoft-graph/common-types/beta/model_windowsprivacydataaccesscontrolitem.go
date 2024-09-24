package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsPrivacyDataAccessControlItem{}

type WindowsPrivacyDataAccessControlItem struct {
	// Determine the access level to specific Windows privacy data category.
	AccessLevel *WindowsPrivacyDataAccessLevel `json:"accessLevel,omitempty"`

	// The Package Family Name of a Windows app. When set, the access level applies to the specified application.
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// The Package Family Name of a Windows app. When set, the access level applies to the specified application.
	AppPackageFamilyName nullable.Type[string] `json:"appPackageFamilyName,omitempty"`

	// Windows privacy data category specifier for privacy data access.
	DataCategory *WindowsPrivacyDataCategory `json:"dataCategory,omitempty"`

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

func (s WindowsPrivacyDataAccessControlItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsPrivacyDataAccessControlItem{}

func (s WindowsPrivacyDataAccessControlItem) MarshalJSON() ([]byte, error) {
	type wrapper WindowsPrivacyDataAccessControlItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsPrivacyDataAccessControlItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsPrivacyDataAccessControlItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsPrivacyDataAccessControlItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsPrivacyDataAccessControlItem: %+v", err)
	}

	return encoded, nil
}
