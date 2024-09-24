package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GroupPolicyObjectFile{}

type GroupPolicyObjectFile struct {
	// The Group Policy Object file content.
	Content *string `json:"content,omitempty"`

	// The date and time at which the GroupPolicy was first uploaded.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The Group Policy Object GUID from GPO Xml content
	GroupPolicyObjectId *string `json:"groupPolicyObjectId,omitempty"`

	// The date and time at which the GroupPolicyObjectFile was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The distinguished name of the OU.
	OuDistinguishedName *string `json:"ouDistinguishedName,omitempty"`

	// The list of scope tags for the configuration.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

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

func (s GroupPolicyObjectFile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyObjectFile{}

func (s GroupPolicyObjectFile) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyObjectFile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyObjectFile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyObjectFile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyObjectFile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyObjectFile: %+v", err)
	}

	return encoded, nil
}
