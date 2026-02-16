package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSharedCookieHistory struct {
	// The comment for the shared cookie.
	Comment nullable.Type[string] `json:"comment,omitempty"`

	// The name of the cookie.
	DisplayName *string `json:"displayName,omitempty"`

	// Controls whether a cookie is a host-only or domain cookie.
	HostOnly *bool `json:"hostOnly,omitempty"`

	// The URL of the cookie.
	HostOrDomain nullable.Type[string] `json:"hostOrDomain,omitempty"`

	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The path of the cookie.
	Path nullable.Type[string] `json:"path,omitempty"`

	// The date and time when the cookie was last published.
	PublishedDateTime *string `json:"publishedDateTime,omitempty"`

	// Specifies how the cookies are shared between Microsoft Edge and Internet Explorer. The possible values are:
	// microsoftEdge, internetExplorer11, both, unknownFutureValue.
	SourceEnvironment *BrowserSharedCookieSourceEnvironment `json:"sourceEnvironment,omitempty"`
}

var _ json.Unmarshaler = &BrowserSharedCookieHistory{}

func (s *BrowserSharedCookieHistory) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Comment           nullable.Type[string]                 `json:"comment,omitempty"`
		DisplayName       *string                               `json:"displayName,omitempty"`
		HostOnly          *bool                                 `json:"hostOnly,omitempty"`
		HostOrDomain      nullable.Type[string]                 `json:"hostOrDomain,omitempty"`
		ODataId           *string                               `json:"@odata.id,omitempty"`
		ODataType         *string                               `json:"@odata.type,omitempty"`
		Path              nullable.Type[string]                 `json:"path,omitempty"`
		PublishedDateTime *string                               `json:"publishedDateTime,omitempty"`
		SourceEnvironment *BrowserSharedCookieSourceEnvironment `json:"sourceEnvironment,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Comment = decoded.Comment
	s.DisplayName = decoded.DisplayName
	s.HostOnly = decoded.HostOnly
	s.HostOrDomain = decoded.HostOrDomain
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Path = decoded.Path
	s.PublishedDateTime = decoded.PublishedDateTime
	s.SourceEnvironment = decoded.SourceEnvironment

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BrowserSharedCookieHistory into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BrowserSharedCookieHistory': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
