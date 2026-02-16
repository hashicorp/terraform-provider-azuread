package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsUpdatesProduct{}

type WindowsUpdatesProduct struct {
	// Represents an edition of a particular Windows product.
	Editions *[]WindowsUpdatesEdition `json:"editions,omitempty"`

	// The friendly names of the product. For example, Version 22H2 (OS build 22621). Read-only.
	FriendlyNames *[]string `json:"friendlyNames,omitempty"`

	// The name of the product group. For example, Windows 11. Read-only.
	GroupName nullable.Type[string] `json:"groupName,omitempty"`

	// Represents a known issue related to a Windows product.
	KnownIssues *[]WindowsUpdatesKnownIssue `json:"knownIssues,omitempty"`

	// The name of the product. For example, Windows 11, version 22H2. Read-only.
	Name *string `json:"name,omitempty"`

	// Represents a product revision.
	Revisions *[]WindowsUpdatesProductRevision `json:"revisions,omitempty"`

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

func (s WindowsUpdatesProduct) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesProduct{}

func (s WindowsUpdatesProduct) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesProduct
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesProduct: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesProduct: %+v", err)
	}

	delete(decoded, "friendlyNames")
	delete(decoded, "groupName")
	delete(decoded, "name")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.product"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesProduct: %+v", err)
	}

	return encoded, nil
}
