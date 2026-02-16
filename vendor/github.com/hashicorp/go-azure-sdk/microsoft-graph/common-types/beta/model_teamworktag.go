package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamworkTag{}

type TeamworkTag struct {
	// Tag description as it appears to the user in Microsoft Teams. A teamworkTag can't have more than 200
	// teamworkTagMembers.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Tag name as it appears to the user in Microsoft Teams.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The number of users assigned to the tag.
	MemberCount nullable.Type[int64] `json:"memberCount,omitempty"`

	// Users assigned to the tag.
	Members *[]TeamworkTagMember `json:"members,omitempty"`

	// The type of tag. Default is standard.
	TagType *TeamworkTagType `json:"tagType,omitempty"`

	// ID of the team in which the tag is defined.
	TeamId nullable.Type[string] `json:"teamId,omitempty"`

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

func (s TeamworkTag) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamworkTag{}

func (s TeamworkTag) MarshalJSON() ([]byte, error) {
	type wrapper TeamworkTag
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamworkTag: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamworkTag: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamworkTag"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamworkTag: %+v", err)
	}

	return encoded, nil
}
