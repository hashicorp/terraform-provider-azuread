package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteHistory struct {
	// Controls the behavior of redirected sites. If true, indicates that the site will open in Internet Explorer 11 or
	// Microsoft Edge even if the site is navigated to as part of a HTTP or meta refresh redirection chain.
	AllowRedirect nullable.Type[bool] `json:"allowRedirect,omitempty"`

	// The comment for the site.
	Comment *string `json:"comment,omitempty"`

	// Controls what compatibility setting is used for specific sites or domains. The possible values are: default,
	// internetExplorer8Enterprise, internetExplorer7Enterprise, internetExplorer11, internetExplorer10, internetExplorer9,
	// internetExplorer8, internetExplorer7, internetExplorer5, unknownFutureValue.
	CompatibilityMode *BrowserSiteCompatibilityMode `json:"compatibilityMode,omitempty"`

	// The user who last modified the site.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The merge type of the site. The possible values are: noMerge, default, unknownFutureValue.
	MergeType *BrowserSiteMergeType `json:"mergeType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date and time when the site was last published.
	PublishedDateTime *string `json:"publishedDateTime,omitempty"`

	// The target environment that the site should open in. The possible values are: internetExplorerMode,
	// internetExplorer11, microsoftEdge, configurable, none, unknownFutureValue.Prior to June 15, 2022, the
	// internetExplorer11 option would allow opening a site in the Internet Explorer 11 (IE11) desktop application.
	// Following the retirement of IE11 on June 15, 2022, the internetExplorer11 option will no longer open an IE11 window
	// and will instead behave the same as the internetExplorerMode option.
	TargetEnvironment *BrowserSiteTargetEnvironment `json:"targetEnvironment,omitempty"`
}

var _ json.Unmarshaler = &BrowserSiteHistory{}

func (s *BrowserSiteHistory) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowRedirect     nullable.Type[bool]           `json:"allowRedirect,omitempty"`
		Comment           *string                       `json:"comment,omitempty"`
		CompatibilityMode *BrowserSiteCompatibilityMode `json:"compatibilityMode,omitempty"`
		MergeType         *BrowserSiteMergeType         `json:"mergeType,omitempty"`
		ODataId           *string                       `json:"@odata.id,omitempty"`
		ODataType         *string                       `json:"@odata.type,omitempty"`
		PublishedDateTime *string                       `json:"publishedDateTime,omitempty"`
		TargetEnvironment *BrowserSiteTargetEnvironment `json:"targetEnvironment,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowRedirect = decoded.AllowRedirect
	s.Comment = decoded.Comment
	s.CompatibilityMode = decoded.CompatibilityMode
	s.MergeType = decoded.MergeType
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PublishedDateTime = decoded.PublishedDateTime
	s.TargetEnvironment = decoded.TargetEnvironment

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BrowserSiteHistory into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BrowserSiteHistory': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
