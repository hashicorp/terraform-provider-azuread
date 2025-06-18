package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EncryptContent = EncryptWithUserDefinedRights{}

type EncryptWithUserDefinedRights struct {
	AllowAdHocPermissions                nullable.Type[bool]   `json:"allowAdHocPermissions,omitempty"`
	AllowMailForwarding                  nullable.Type[bool]   `json:"allowMailForwarding,omitempty"`
	DecryptionRightsManagementTemplateId nullable.Type[string] `json:"decryptionRightsManagementTemplateId,omitempty"`

	// Fields inherited from EncryptContent

	EncryptWith *EncryptWith `json:"encryptWith,omitempty"`

	// Fields inherited from LabelActionBase

	// The name of the action (for example, 'Encrypt', 'AddHeader').
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EncryptWithUserDefinedRights) EncryptContent() BaseEncryptContentImpl {
	return BaseEncryptContentImpl{
		EncryptWith: s.EncryptWith,
		Name:        s.Name,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s EncryptWithUserDefinedRights) LabelActionBase() BaseLabelActionBaseImpl {
	return BaseLabelActionBaseImpl{
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EncryptWithUserDefinedRights{}

func (s EncryptWithUserDefinedRights) MarshalJSON() ([]byte, error) {
	type wrapper EncryptWithUserDefinedRights
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EncryptWithUserDefinedRights: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EncryptWithUserDefinedRights: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.encryptWithUserDefinedRights"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EncryptWithUserDefinedRights: %+v", err)
	}

	return encoded, nil
}
