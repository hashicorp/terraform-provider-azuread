package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationResource = EducationTeamsAppResource{}

type EducationTeamsAppResource struct {
	// URL that points to the icon of the app.
	AppIconWebUrl nullable.Type[string] `json:"appIconWebUrl,omitempty"`

	// Teams app ID of the application.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// URL for the app resource that will be opened by Teams.
	TeamsEmbeddedContentUrl nullable.Type[string] `json:"teamsEmbeddedContentUrl,omitempty"`

	// URL for the app resource that can be opened in the browser.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

	// Fields inherited from EducationResource

	// Who created the resource?
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Display name of resource.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Who was the last user to modify the resource?
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Moment in time when the resource was last modified. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationTeamsAppResource) EducationResource() BaseEducationResourceImpl {
	return BaseEducationResourceImpl{
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		DisplayName:          s.DisplayName,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

var _ json.Marshaler = EducationTeamsAppResource{}

func (s EducationTeamsAppResource) MarshalJSON() ([]byte, error) {
	type wrapper EducationTeamsAppResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationTeamsAppResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationTeamsAppResource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationTeamsAppResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationTeamsAppResource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationTeamsAppResource{}

func (s *EducationTeamsAppResource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AppIconWebUrl           nullable.Type[string] `json:"appIconWebUrl,omitempty"`
		AppId                   nullable.Type[string] `json:"appId,omitempty"`
		TeamsEmbeddedContentUrl nullable.Type[string] `json:"teamsEmbeddedContentUrl,omitempty"`
		WebUrl                  nullable.Type[string] `json:"webUrl,omitempty"`
		CreatedDateTime         nullable.Type[string] `json:"createdDateTime,omitempty"`
		DisplayName             nullable.Type[string] `json:"displayName,omitempty"`
		LastModifiedDateTime    nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		ODataId                 *string               `json:"@odata.id,omitempty"`
		ODataType               *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppIconWebUrl = decoded.AppIconWebUrl
	s.AppId = decoded.AppId
	s.TeamsEmbeddedContentUrl = decoded.TeamsEmbeddedContentUrl
	s.WebUrl = decoded.WebUrl
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationTeamsAppResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EducationTeamsAppResource': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'EducationTeamsAppResource': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
