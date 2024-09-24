package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamTemplateDefinition{}

type TeamTemplateDefinition struct {
	// Describes the audience the team template is available to. The possible values are: organization, user, public,
	// unknownFutureValue.
	Audience *TeamTemplateAudience `json:"audience,omitempty"`

	// The assigned categories for the team template.
	Categories *[]string `json:"categories,omitempty"`

	// A brief description of the team template as it will appear to the users in Microsoft Teams.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The user defined name of the team template.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The icon url for the team template.
	IconUrl nullable.Type[string] `json:"iconUrl,omitempty"`

	// Language the template is available in.
	LanguageTag nullable.Type[string] `json:"languageTag,omitempty"`

	// The identity of the user who last modified the team template.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The date time of when the team template was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The templateId for the team template
	ParentTemplateId nullable.Type[string] `json:"parentTemplateId,omitempty"`

	// The organization which published the team template.
	PublisherName nullable.Type[string] `json:"publisherName,omitempty"`

	// A short-description of the team template as it will appear to the users in Microsoft Teams.
	ShortDescription nullable.Type[string] `json:"shortDescription,omitempty"`

	// Collection of channel objects. A channel represents a topic, and therefore a logical isolation of discussion, within
	// a team.
	TeamDefinition *Team `json:"teamDefinition,omitempty"`

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

func (s TeamTemplateDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamTemplateDefinition{}

func (s TeamTemplateDefinition) MarshalJSON() ([]byte, error) {
	type wrapper TeamTemplateDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamTemplateDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamTemplateDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamTemplateDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamTemplateDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamTemplateDefinition{}

func (s *TeamTemplateDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Audience             *TeamTemplateAudience `json:"audience,omitempty"`
		Categories           *[]string             `json:"categories,omitempty"`
		Description          nullable.Type[string] `json:"description,omitempty"`
		DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
		IconUrl              nullable.Type[string] `json:"iconUrl,omitempty"`
		LanguageTag          nullable.Type[string] `json:"languageTag,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		ParentTemplateId     nullable.Type[string] `json:"parentTemplateId,omitempty"`
		PublisherName        nullable.Type[string] `json:"publisherName,omitempty"`
		ShortDescription     nullable.Type[string] `json:"shortDescription,omitempty"`
		TeamDefinition       *Team                 `json:"teamDefinition,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Audience = decoded.Audience
	s.Categories = decoded.Categories
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.IconUrl = decoded.IconUrl
	s.LanguageTag = decoded.LanguageTag
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ParentTemplateId = decoded.ParentTemplateId
	s.PublisherName = decoded.PublisherName
	s.ShortDescription = decoded.ShortDescription
	s.TeamDefinition = decoded.TeamDefinition
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamTemplateDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'TeamTemplateDefinition': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
