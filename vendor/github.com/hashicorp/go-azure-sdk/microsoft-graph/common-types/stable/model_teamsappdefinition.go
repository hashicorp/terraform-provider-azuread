package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamsAppDefinition{}

type TeamsAppDefinition struct {
	// Authorization requirements specified in the Teams app manifest.
	Authorization *TeamsAppAuthorization `json:"authorization,omitempty"`

	// The details of the bot specified in the Teams app manifest.
	Bot *TeamworkBot `json:"bot,omitempty"`

	CreatedBy IdentitySet `json:"createdBy"`

	// Verbose description of the application.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of the app provided by the app developer.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The published status of a specific version of a Teams app. Possible values are:submitted—The specific version of
	// the Teams app was submitted and is under review.published—The request to publish the specific version of the Teams
	// app was approved by the admin and the app is published.rejected—The admin rejected the request to publish the
	// specific version of the Teams app.
	PublishingState *TeamsAppPublishingState `json:"publishingState,omitempty"`

	// Short description of the application.
	ShortDescription nullable.Type[string] `json:"shortDescription,omitempty"`

	// The ID from the Teams app manifest.
	TeamsAppId nullable.Type[string] `json:"teamsAppId,omitempty"`

	// The version number of the application.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s TeamsAppDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsAppDefinition{}

func (s TeamsAppDefinition) MarshalJSON() ([]byte, error) {
	type wrapper TeamsAppDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsAppDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAppDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamsAppDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsAppDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamsAppDefinition{}

func (s *TeamsAppDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Authorization        *TeamsAppAuthorization   `json:"authorization,omitempty"`
		Bot                  *TeamworkBot             `json:"bot,omitempty"`
		Description          nullable.Type[string]    `json:"description,omitempty"`
		DisplayName          nullable.Type[string]    `json:"displayName,omitempty"`
		LastModifiedDateTime nullable.Type[string]    `json:"lastModifiedDateTime,omitempty"`
		PublishingState      *TeamsAppPublishingState `json:"publishingState,omitempty"`
		ShortDescription     nullable.Type[string]    `json:"shortDescription,omitempty"`
		TeamsAppId           nullable.Type[string]    `json:"teamsAppId,omitempty"`
		Version              nullable.Type[string]    `json:"version,omitempty"`
		Id                   *string                  `json:"id,omitempty"`
		ODataId              *string                  `json:"@odata.id,omitempty"`
		ODataType            *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Authorization = decoded.Authorization
	s.Bot = decoded.Bot
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.PublishingState = decoded.PublishingState
	s.ShortDescription = decoded.ShortDescription
	s.TeamsAppId = decoded.TeamsAppId
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamsAppDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'TeamsAppDefinition': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
