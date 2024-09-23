package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = BrowserSharedCookie{}

type BrowserSharedCookie struct {
	// The comment for the shared cookie.
	Comment *string `json:"comment,omitempty"`

	// The date and time when the shared cookie was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The date and time when the shared cookie was deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

	// The name of the cookie.
	DisplayName *string `json:"displayName,omitempty"`

	// The history of modifications applied to the cookie.
	History *[]BrowserSharedCookieHistory `json:"history,omitempty"`

	// Controls whether a cookie is a host-only or domain cookie.
	HostOnly *bool `json:"hostOnly,omitempty"`

	// The URL of the cookie.
	HostOrDomain *string `json:"hostOrDomain,omitempty"`

	// The user who last modified the cookie.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The date and time when the cookie was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The path of the cookie.
	Path *string `json:"path,omitempty"`

	SourceEnvironment *BrowserSharedCookieSourceEnvironment `json:"sourceEnvironment,omitempty"`
	Status            *BrowserSharedCookieStatus            `json:"status,omitempty"`

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

func (s BrowserSharedCookie) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BrowserSharedCookie{}

func (s BrowserSharedCookie) MarshalJSON() ([]byte, error) {
	type wrapper BrowserSharedCookie
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BrowserSharedCookie: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BrowserSharedCookie: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.browserSharedCookie"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BrowserSharedCookie: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BrowserSharedCookie{}

func (s *BrowserSharedCookie) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Comment              *string                               `json:"comment,omitempty"`
		CreatedDateTime      *string                               `json:"createdDateTime,omitempty"`
		DeletedDateTime      nullable.Type[string]                 `json:"deletedDateTime,omitempty"`
		DisplayName          *string                               `json:"displayName,omitempty"`
		History              *[]BrowserSharedCookieHistory         `json:"history,omitempty"`
		HostOnly             *bool                                 `json:"hostOnly,omitempty"`
		HostOrDomain         *string                               `json:"hostOrDomain,omitempty"`
		LastModifiedDateTime *string                               `json:"lastModifiedDateTime,omitempty"`
		Path                 *string                               `json:"path,omitempty"`
		SourceEnvironment    *BrowserSharedCookieSourceEnvironment `json:"sourceEnvironment,omitempty"`
		Status               *BrowserSharedCookieStatus            `json:"status,omitempty"`
		Id                   *string                               `json:"id,omitempty"`
		ODataId              *string                               `json:"@odata.id,omitempty"`
		ODataType            *string                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Comment = decoded.Comment
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DeletedDateTime = decoded.DeletedDateTime
	s.DisplayName = decoded.DisplayName
	s.History = decoded.History
	s.HostOnly = decoded.HostOnly
	s.HostOrDomain = decoded.HostOrDomain
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Path = decoded.Path
	s.SourceEnvironment = decoded.SourceEnvironment
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BrowserSharedCookie into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BrowserSharedCookie': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
