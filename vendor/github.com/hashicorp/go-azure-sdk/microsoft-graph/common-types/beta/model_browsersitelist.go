package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = BrowserSiteList{}

type BrowserSiteList struct {
	// The description of the site list.
	Description *string `json:"description,omitempty"`

	// The name of the site list.
	DisplayName *string `json:"displayName,omitempty"`

	// The user who last modified the site list.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The date and time when the site list was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The user who published the site list.
	PublishedBy IdentitySet `json:"publishedBy"`

	// The date and time when the site list was published.
	PublishedDateTime nullable.Type[string] `json:"publishedDateTime,omitempty"`

	// The current revision of the site list.
	Revision *string `json:"revision,omitempty"`

	// A collection of shared cookies defined for the site list.
	SharedCookies *[]BrowserSharedCookie `json:"sharedCookies,omitempty"`

	// A collection of sites defined for the site list.
	Sites *[]BrowserSite `json:"sites,omitempty"`

	Status *BrowserSiteListStatus `json:"status,omitempty"`

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

func (s BrowserSiteList) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BrowserSiteList{}

func (s BrowserSiteList) MarshalJSON() ([]byte, error) {
	type wrapper BrowserSiteList
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BrowserSiteList: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BrowserSiteList: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.browserSiteList"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BrowserSiteList: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BrowserSiteList{}

func (s *BrowserSiteList) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description          *string                `json:"description,omitempty"`
		DisplayName          *string                `json:"displayName,omitempty"`
		LastModifiedDateTime *string                `json:"lastModifiedDateTime,omitempty"`
		PublishedDateTime    nullable.Type[string]  `json:"publishedDateTime,omitempty"`
		Revision             *string                `json:"revision,omitempty"`
		SharedCookies        *[]BrowserSharedCookie `json:"sharedCookies,omitempty"`
		Sites                *[]BrowserSite         `json:"sites,omitempty"`
		Status               *BrowserSiteListStatus `json:"status,omitempty"`
		Id                   *string                `json:"id,omitempty"`
		ODataId              *string                `json:"@odata.id,omitempty"`
		ODataType            *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.PublishedDateTime = decoded.PublishedDateTime
	s.Revision = decoded.Revision
	s.SharedCookies = decoded.SharedCookies
	s.Sites = decoded.Sites
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BrowserSiteList into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BrowserSiteList': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	if v, ok := temp["publishedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PublishedBy' for 'BrowserSiteList': %+v", err)
		}
		s.PublishedBy = impl
	}

	return nil
}
