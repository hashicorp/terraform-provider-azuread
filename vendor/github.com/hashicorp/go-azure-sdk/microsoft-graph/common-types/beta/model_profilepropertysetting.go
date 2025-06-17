package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ProfilePropertySetting{}

type ProfilePropertySetting struct {
	// A privacy setting that reflects the allowed audience for the configured property. The possible values are: me,
	// organization, federatedOrganizations, everyone, unknownFutureValue.
	AllowedAudiences *OrganizationAllowedAudiences `json:"allowedAudiences,omitempty"`

	// Defines whether a user is allowed to override the tenant admin privacy setting.
	IsUserOverrideForAudienceEnabled nullable.Type[bool] `json:"isUserOverrideForAudienceEnabled,omitempty"`

	// Name of the property-level setting.
	Name nullable.Type[string] `json:"name,omitempty"`

	// A collection of prioritized profile source URLs ordered by data precedence within an organization.
	PrioritizedSourceUrls *[]string `json:"prioritizedSourceUrls,omitempty"`

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

func (s ProfilePropertySetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ProfilePropertySetting{}

func (s ProfilePropertySetting) MarshalJSON() ([]byte, error) {
	type wrapper ProfilePropertySetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProfilePropertySetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProfilePropertySetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.profilePropertySetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProfilePropertySetting: %+v", err)
	}

	return encoded, nil
}
