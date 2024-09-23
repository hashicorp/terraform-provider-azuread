package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OmaSetting = OmaSettingInteger{}

type OmaSettingInteger struct {
	// By setting to true, the CSP (configuration service provider) specified in the OMA-URI will perform a get, instead of
	// set
	IsReadOnly *bool `json:"isReadOnly,omitempty"`

	// Value.
	Value *int64 `json:"value,omitempty"`

	// Fields inherited from OmaSetting

	// Description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display Name.
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates whether the value field is encrypted. This property is read-only.
	IsEncrypted *bool `json:"isEncrypted,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// OMA.
	OmaUri *string `json:"omaUri,omitempty"`

	// ReferenceId for looking up secret for decryption. This property is read-only.
	SecretReferenceValueId nullable.Type[string] `json:"secretReferenceValueId,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OmaSettingInteger) OmaSetting() BaseOmaSettingImpl {
	return BaseOmaSettingImpl{
		Description:            s.Description,
		DisplayName:            s.DisplayName,
		IsEncrypted:            s.IsEncrypted,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
		OmaUri:                 s.OmaUri,
		SecretReferenceValueId: s.SecretReferenceValueId,
	}
}

var _ json.Marshaler = OmaSettingInteger{}

func (s OmaSettingInteger) MarshalJSON() ([]byte, error) {
	type wrapper OmaSettingInteger
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OmaSettingInteger: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OmaSettingInteger: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.omaSettingInteger"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OmaSettingInteger: %+v", err)
	}

	return encoded, nil
}
