package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ColumnDefinition{}

type ColumnDefinition struct {
	// This column stores Boolean values.
	Boolean *BooleanColumn `json:"boolean,omitempty"`

	// This column's data is calculated based on other columns.
	Calculated *CalculatedColumn `json:"calculated,omitempty"`

	// This column stores data from a list of choices.
	Choice *ChoiceColumn `json:"choice,omitempty"`

	// For site columns, the name of the group this column belongs to. Helps organize related columns.
	ColumnGroup nullable.Type[string] `json:"columnGroup,omitempty"`

	// This column stores content approval status.
	ContentApprovalStatus *ContentApprovalStatusColumn `json:"contentApprovalStatus,omitempty"`

	// This column stores currency values.
	Currency *CurrencyColumn `json:"currency,omitempty"`

	// This column stores DateTime values.
	DateTime *DateTimeColumn `json:"dateTime,omitempty"`

	// The default value for this column.
	DefaultValue *DefaultColumnValue `json:"defaultValue,omitempty"`

	// The user-facing description of the column.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The user-facing name of the column.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// If true, no two list items may have the same value for this column.
	EnforceUniqueValues nullable.Type[bool] `json:"enforceUniqueValues,omitempty"`

	// This column stores a geolocation.
	Geolocation *GeolocationColumn `json:"geolocation,omitempty"`

	// Specifies whether the column is displayed in the user interface.
	Hidden nullable.Type[bool] `json:"hidden,omitempty"`

	// This column stores hyperlink or picture values.
	HyperlinkOrPicture *HyperlinkOrPictureColumn `json:"hyperlinkOrPicture,omitempty"`

	// Specifies whether the column values can used for sorting and searching.
	Indexed nullable.Type[bool] `json:"indexed,omitempty"`

	// Indicates whether this column can be deleted.
	IsDeletable nullable.Type[bool] `json:"isDeletable,omitempty"`

	// Indicates whether values in the column can be reordered. Read-only.
	IsReorderable nullable.Type[bool] `json:"isReorderable,omitempty"`

	// Specifies whether the column can be changed.
	IsSealed nullable.Type[bool] `json:"isSealed,omitempty"`

	// This column's data is looked up from another source in the site.
	Lookup *LookupColumn `json:"lookup,omitempty"`

	// The API-facing name of the column as it appears in the fields on a listItem. For the user-facing name, see
	// displayName.
	Name nullable.Type[string] `json:"name,omitempty"`

	// This column stores number values.
	Number *NumberColumn `json:"number,omitempty"`

	// This column stores Person or Group values.
	PersonOrGroup *PersonOrGroupColumn `json:"personOrGroup,omitempty"`

	// If true, changes to this column will be propagated to lists that implement the column.
	PropagateChanges nullable.Type[bool] `json:"propagateChanges,omitempty"`

	// Specifies whether the column values can be modified.
	ReadOnly nullable.Type[bool] `json:"readOnly,omitempty"`

	// Specifies whether the column value isn't optional.
	Required nullable.Type[bool] `json:"required,omitempty"`

	// The source column for content type column.
	SourceColumn *ColumnDefinition `json:"sourceColumn,omitempty"`

	// ContentType from which this column is inherited from. Used only to fetch contentTypes columns.
	SourceContentType *ContentTypeInfo `json:"sourceContentType,omitempty"`

	// This column stores taxonomy terms.
	Term *TermColumn `json:"term,omitempty"`

	// This column stores text values.
	Text *TextColumn `json:"text,omitempty"`

	// This column stores thumbnail values.
	Thumbnail *ThumbnailColumn `json:"thumbnail,omitempty"`

	// For site columns, the type of column. Read-only.
	Type *ColumnTypes `json:"type,omitempty"`

	// This column stores validation formula and message for the column.
	Validation *ColumnValidation `json:"validation,omitempty"`

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

func (s ColumnDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ColumnDefinition{}

func (s ColumnDefinition) MarshalJSON() ([]byte, error) {
	type wrapper ColumnDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ColumnDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ColumnDefinition: %+v", err)
	}

	delete(decoded, "isReorderable")
	delete(decoded, "type")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.columnDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ColumnDefinition: %+v", err)
	}

	return encoded, nil
}
