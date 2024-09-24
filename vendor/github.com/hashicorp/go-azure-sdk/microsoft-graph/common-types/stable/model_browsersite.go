package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = BrowserSite{}

type BrowserSite struct {
	// Controls the behavior of redirected sites. If true, indicates that the site will open in Internet Explorer 11 or
	// Microsoft Edge even if the site is navigated to as part of a HTTP or meta refresh redirection chain.
	AllowRedirect *bool `json:"allowRedirect,omitempty"`

	// The comment for the site.
	Comment *string `json:"comment,omitempty"`

	CompatibilityMode *BrowserSiteCompatibilityMode `json:"compatibilityMode,omitempty"`

	// The date and time when the site was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The date and time when the site was deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

	// The history of modifications applied to the site.
	History *[]BrowserSiteHistory `json:"history,omitempty"`

	// The user who last modified the site.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The date and time when the site was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	MergeType         *BrowserSiteMergeType         `json:"mergeType,omitempty"`
	Status            *BrowserSiteStatus            `json:"status,omitempty"`
	TargetEnvironment *BrowserSiteTargetEnvironment `json:"targetEnvironment,omitempty"`

	// The URL of the site.
	WebUrl *string `json:"webUrl,omitempty"`

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

func (s BrowserSite) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BrowserSite{}

func (s BrowserSite) MarshalJSON() ([]byte, error) {
	type wrapper BrowserSite
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BrowserSite: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BrowserSite: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.browserSite"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BrowserSite: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BrowserSite{}

func (s *BrowserSite) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowRedirect        *bool                         `json:"allowRedirect,omitempty"`
		Comment              *string                       `json:"comment,omitempty"`
		CompatibilityMode    *BrowserSiteCompatibilityMode `json:"compatibilityMode,omitempty"`
		CreatedDateTime      *string                       `json:"createdDateTime,omitempty"`
		DeletedDateTime      nullable.Type[string]         `json:"deletedDateTime,omitempty"`
		History              *[]BrowserSiteHistory         `json:"history,omitempty"`
		LastModifiedDateTime *string                       `json:"lastModifiedDateTime,omitempty"`
		MergeType            *BrowserSiteMergeType         `json:"mergeType,omitempty"`
		Status               *BrowserSiteStatus            `json:"status,omitempty"`
		TargetEnvironment    *BrowserSiteTargetEnvironment `json:"targetEnvironment,omitempty"`
		WebUrl               *string                       `json:"webUrl,omitempty"`
		Id                   *string                       `json:"id,omitempty"`
		ODataId              *string                       `json:"@odata.id,omitempty"`
		ODataType            *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowRedirect = decoded.AllowRedirect
	s.Comment = decoded.Comment
	s.CompatibilityMode = decoded.CompatibilityMode
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DeletedDateTime = decoded.DeletedDateTime
	s.History = decoded.History
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.MergeType = decoded.MergeType
	s.Status = decoded.Status
	s.TargetEnvironment = decoded.TargetEnvironment
	s.WebUrl = decoded.WebUrl
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BrowserSite into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BrowserSite': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
