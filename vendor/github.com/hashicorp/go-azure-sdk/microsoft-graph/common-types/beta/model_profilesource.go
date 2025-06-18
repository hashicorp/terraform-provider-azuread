package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ProfileSource{}

type ProfileSource struct {
	// Name of the profile source intended to inform users about the profile source name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Type of the profile source.
	Kind nullable.Type[string] `json:"kind,omitempty"`

	// Alternative localized labels specified by an administrator.
	Localizations *[]ProfileSourceLocalization `json:"localizations,omitempty"`

	// Profile source identifier used as an alternate key.
	SourceId nullable.Type[string] `json:"sourceId,omitempty"`

	// Web URL of the profile source that directs users to the page view of profile data.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

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

func (s ProfileSource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ProfileSource{}

func (s ProfileSource) MarshalJSON() ([]byte, error) {
	type wrapper ProfileSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProfileSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProfileSource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.profileSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProfileSource: %+v", err)
	}

	return encoded, nil
}
