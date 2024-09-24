package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingLink struct {
	// The app the link is associated with.
	Application Identity `json:"application"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If true then the user can only use this link to view the item on the web, and cannot use it to download the contents
	// of the item. Only for OneDrive for Business and SharePoint.
	PreventsDownload nullable.Type[bool] `json:"preventsDownload,omitempty"`

	// The scope of the link represented by this permission. Value anonymous indicates the link is usable by anyone,
	// organization indicates the link is only usable for users signed into the same tenant.
	Scope nullable.Type[string] `json:"scope,omitempty"`

	// The type of the link created.
	Type nullable.Type[string] `json:"type,omitempty"`

	// For embed links, this property contains the HTML code for an <iframe> element that will embed the item in a webpage.
	WebHtml nullable.Type[string] `json:"webHtml,omitempty"`

	// A URL that opens the item in the browser on the OneDrive website.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`
}

var _ json.Unmarshaler = &SharingLink{}

func (s *SharingLink) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
		PreventsDownload nullable.Type[bool]   `json:"preventsDownload,omitempty"`
		Scope            nullable.Type[string] `json:"scope,omitempty"`
		Type             nullable.Type[string] `json:"type,omitempty"`
		WebHtml          nullable.Type[string] `json:"webHtml,omitempty"`
		WebUrl           nullable.Type[string] `json:"webUrl,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PreventsDownload = decoded.PreventsDownload
	s.Scope = decoded.Scope
	s.Type = decoded.Type
	s.WebHtml = decoded.WebHtml
	s.WebUrl = decoded.WebUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SharingLink into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["application"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Application' for 'SharingLink': %+v", err)
		}
		s.Application = impl
	}

	return nil
}
