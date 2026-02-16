package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ContentType{}

type ContentType struct {
	// List of canonical URLs for hub sites with which this content type is associated to. This will contain all hub sites
	// where this content type is queued to be enforced or is already enforced. Enforcing a content type means that the
	// content type is applied to the lists in the enforced sites.
	AssociatedHubsUrls *[]string `json:"associatedHubsUrls,omitempty"`

	// Parent contentType from which this content type is derived.
	Base *ContentType `json:"base,omitempty"`

	// The collection of content types that are ancestors of this content type.
	BaseTypes *[]ContentType `json:"baseTypes,omitempty"`

	// The collection of columns that are required by this content type.
	ColumnLinks *[]ColumnLink `json:"columnLinks,omitempty"`

	// Column order information in a content type.
	ColumnPositions *[]ColumnDefinition `json:"columnPositions,omitempty"`

	// The collection of column definitions for this content type.
	Columns *[]ColumnDefinition `json:"columns,omitempty"`

	// The descriptive text for the item.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Document Set metadata.
	DocumentSet *DocumentSet `json:"documentSet,omitempty"`

	// Document template metadata. To make sure that documents have consistent content across a site and its subsites, you
	// can associate a Word, Excel, or PowerPoint template with a site content type.
	DocumentTemplate *DocumentSetContent `json:"documentTemplate,omitempty"`

	// The name of the group this content type belongs to. Helps organize related content types.
	Group nullable.Type[string] `json:"group,omitempty"`

	// Indicates whether the content type is hidden in the list's 'New' menu.
	Hidden nullable.Type[bool] `json:"hidden,omitempty"`

	// If this content type is inherited from another scope (like a site), provides a reference to the item where the
	// content type is defined.
	InheritedFrom *ItemReference `json:"inheritedFrom,omitempty"`

	// Specifies if a content type is a built-in content type.
	IsBuiltIn nullable.Type[bool] `json:"isBuiltIn,omitempty"`

	// The name of the content type.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Specifies the order in which the content type appears in the selection UI.
	Order *ContentTypeOrder `json:"order,omitempty"`

	// The unique identifier of the content type.
	ParentId nullable.Type[string] `json:"parentId,omitempty"`

	// If true, any changes made to the content type are pushed to inherited content types and lists that implement the
	// content type.
	PropagateChanges nullable.Type[bool] `json:"propagateChanges,omitempty"`

	// If true, the content type can't be modified unless this value is first set to false.
	ReadOnly nullable.Type[bool] `json:"readOnly,omitempty"`

	// If true, the content type can't be modified by users or through push-down operations. Only site collection
	// administrators can seal or unseal content types.
	Sealed nullable.Type[bool] `json:"sealed,omitempty"`

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

func (s ContentType) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ContentType{}

func (s ContentType) MarshalJSON() ([]byte, error) {
	type wrapper ContentType
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ContentType: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ContentType: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.contentType"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ContentType: %+v", err)
	}

	return encoded, nil
}
