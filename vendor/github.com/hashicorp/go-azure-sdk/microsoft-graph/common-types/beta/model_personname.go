package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ItemFacet = PersonName{}

type PersonName struct {
	// Provides an ordered rendering of firstName and lastName depending on the locale of the user or their device.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// First name of the user.
	First nullable.Type[string] `json:"first,omitempty"`

	// Initials of the user.
	Initials nullable.Type[string] `json:"initials,omitempty"`

	// Contains the name for the language (en-US, no-NB, en-AU) following IETF BCP47 format.
	LanguageTag nullable.Type[string] `json:"languageTag,omitempty"`

	// Last name of the user.
	Last nullable.Type[string] `json:"last,omitempty"`

	// Maiden name of the user.
	Maiden nullable.Type[string] `json:"maiden,omitempty"`

	// Middle name of the user.
	Middle nullable.Type[string] `json:"middle,omitempty"`

	// Nickname of the user.
	Nickname nullable.Type[string] `json:"nickname,omitempty"`

	// Guidance on how to pronounce the users name.
	Pronunciation *PersonNamePronounciation `json:"pronunciation,omitempty"`

	// Designators used after the users name (eg: PhD.)
	Suffix nullable.Type[string] `json:"suffix,omitempty"`

	// Honorifics used to prefix a users name (eg: Dr, Sir, Madam, Mrs.)
	Title nullable.Type[string] `json:"title,omitempty"`

	// Fields inherited from ItemFacet

	// The audiences that are able to see the values contained within the associated entity. Possible values are: me,
	// family, contacts, groupMembers, organization, federatedOrganizations, everyone, unknownFutureValue.
	AllowedAudiences *AllowedAudiences `json:"allowedAudiences,omitempty"`

	CreatedBy IdentitySet `json:"createdBy"`

	// Provides the dateTimeOffset for when the entity was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Contains inference detail if the entity is inferred by the creating or modifying application.
	Inference *InferenceData `json:"inference,omitempty"`

	IsSearchable   nullable.Type[bool] `json:"isSearchable,omitempty"`
	LastModifiedBy IdentitySet         `json:"lastModifiedBy"`

	// Provides the dateTimeOffset for when the entity was created.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Where the values within an entity originated if synced from another service.
	Source *PersonDataSources `json:"source,omitempty"`

	// Where the values within an entity originated if synced from another source.
	Sources *[]ProfileSourceAnnotation `json:"sources,omitempty"`

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

func (s PersonName) ItemFacet() BaseItemFacetImpl {
	return BaseItemFacetImpl{
		AllowedAudiences:     s.AllowedAudiences,
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		Inference:            s.Inference,
		IsSearchable:         s.IsSearchable,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Source:               s.Source,
		Sources:              s.Sources,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s PersonName) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PersonName{}

func (s PersonName) MarshalJSON() ([]byte, error) {
	type wrapper PersonName
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PersonName: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PersonName: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.personName"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PersonName: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PersonName{}

func (s *PersonName) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName          nullable.Type[string]      `json:"displayName,omitempty"`
		First                nullable.Type[string]      `json:"first,omitempty"`
		Initials             nullable.Type[string]      `json:"initials,omitempty"`
		LanguageTag          nullable.Type[string]      `json:"languageTag,omitempty"`
		Last                 nullable.Type[string]      `json:"last,omitempty"`
		Maiden               nullable.Type[string]      `json:"maiden,omitempty"`
		Middle               nullable.Type[string]      `json:"middle,omitempty"`
		Nickname             nullable.Type[string]      `json:"nickname,omitempty"`
		Pronunciation        *PersonNamePronounciation  `json:"pronunciation,omitempty"`
		Suffix               nullable.Type[string]      `json:"suffix,omitempty"`
		Title                nullable.Type[string]      `json:"title,omitempty"`
		AllowedAudiences     *AllowedAudiences          `json:"allowedAudiences,omitempty"`
		CreatedDateTime      *string                    `json:"createdDateTime,omitempty"`
		Inference            *InferenceData             `json:"inference,omitempty"`
		IsSearchable         nullable.Type[bool]        `json:"isSearchable,omitempty"`
		LastModifiedDateTime *string                    `json:"lastModifiedDateTime,omitempty"`
		Source               *PersonDataSources         `json:"source,omitempty"`
		Sources              *[]ProfileSourceAnnotation `json:"sources,omitempty"`
		Id                   *string                    `json:"id,omitempty"`
		ODataId              *string                    `json:"@odata.id,omitempty"`
		ODataType            *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.First = decoded.First
	s.Initials = decoded.Initials
	s.LanguageTag = decoded.LanguageTag
	s.Last = decoded.Last
	s.Maiden = decoded.Maiden
	s.Middle = decoded.Middle
	s.Nickname = decoded.Nickname
	s.Pronunciation = decoded.Pronunciation
	s.Suffix = decoded.Suffix
	s.Title = decoded.Title
	s.AllowedAudiences = decoded.AllowedAudiences
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.Inference = decoded.Inference
	s.IsSearchable = decoded.IsSearchable
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Source = decoded.Source
	s.Sources = decoded.Sources

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PersonName into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'PersonName': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'PersonName': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
