package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OmaSetting = OmaSettingStringXml{}

type OmaSettingStringXml struct {
	// File name associated with the Value property (.xml).
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// Value. (UTF8 encoded byte array)
	Value *string `json:"value,omitempty"`

	// Fields inherited from OmaSetting

	// Description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display Name.
	DisplayName *string `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// OMA.
	OmaUri *string `json:"omaUri,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OmaSettingStringXml) OmaSetting() BaseOmaSettingImpl {
	return BaseOmaSettingImpl{
		Description: s.Description,
		DisplayName: s.DisplayName,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		OmaUri:      s.OmaUri,
	}
}

var _ json.Marshaler = OmaSettingStringXml{}

func (s OmaSettingStringXml) MarshalJSON() ([]byte, error) {
	type wrapper OmaSettingStringXml
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OmaSettingStringXml: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OmaSettingStringXml: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.omaSettingStringXml"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OmaSettingStringXml: %+v", err)
	}

	return encoded, nil
}
