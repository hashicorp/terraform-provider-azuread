package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupPolicyUploadedPresentation = GroupPolicyPresentationListBox{}

type GroupPolicyPresentationListBox struct {
	// If this option is specified true the user must specify the registry subkey value and the registry subkey name. The
	// list box shows two columns, one for the name and one for the data. The default value is false.
	ExplicitValue *bool `json:"explicitValue,omitempty"`

	ValuePrefix nullable.Type[string] `json:"valuePrefix,omitempty"`

	// Fields inherited from GroupPolicyPresentation

	// The group policy definition associated with the presentation.
	Definition *GroupPolicyDefinition `json:"definition,omitempty"`

	// Localized text label for any presentation entity. The default value is empty.
	Label nullable.Type[string] `json:"label,omitempty"`

	// The date and time the entity was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

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

func (s GroupPolicyPresentationListBox) GroupPolicyUploadedPresentation() BaseGroupPolicyUploadedPresentationImpl {
	return BaseGroupPolicyUploadedPresentationImpl{
		Definition:           s.Definition,
		Label:                s.Label,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s GroupPolicyPresentationListBox) GroupPolicyPresentation() BaseGroupPolicyPresentationImpl {
	return BaseGroupPolicyPresentationImpl{
		Definition:           s.Definition,
		Label:                s.Label,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s GroupPolicyPresentationListBox) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyPresentationListBox{}

func (s GroupPolicyPresentationListBox) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyPresentationListBox
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyPresentationListBox: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyPresentationListBox: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyPresentationListBox"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyPresentationListBox: %+v", err)
	}

	return encoded, nil
}
