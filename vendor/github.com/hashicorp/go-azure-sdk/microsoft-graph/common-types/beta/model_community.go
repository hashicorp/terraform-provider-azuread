package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Community{}

type Community struct {
	// The description of the community. The maximum length is 1,024 characters.
	Description *string `json:"description,omitempty"`

	// The name of the community. The maximum length is 255 characters.
	DisplayName *string `json:"displayName,omitempty"`

	// The Microsoft 365 group that manages the membership of this community.
	Group *Group `json:"group,omitempty"`

	// The ID of the Microsoft 365 group that manages the membership of this community.
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// The admins of the community. Limited to 100 users. If this property isn't specified when you create the community,
	// the calling user is automatically assigned as the community owner.
	Owners *[]User `json:"owners,omitempty"`

	// Types of communityPrivacy.
	Privacy *CommunityPrivacy `json:"privacy,omitempty"`

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

func (s Community) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Community{}

func (s Community) MarshalJSON() ([]byte, error) {
	type wrapper Community
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Community: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Community: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.community"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Community: %+v", err)
	}

	return encoded, nil
}
