package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchHit struct {
	// The name of the content source that the externalItem is part of.
	ContentSource nullable.Type[string] `json:"contentSource,omitempty"`

	// The internal identifier for the item. The format of the identifier varies based on the entity type. For details, see
	// hitId format.
	HitId nullable.Type[string] `json:"hitId,omitempty"`

	// Indicates whether the current result is collapses when the collapseProperties property in the searchRequest is used.
	IsCollapsed nullable.Type[bool] `json:"isCollapsed,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The rank or the order of the result.
	Rank nullable.Type[int64] `json:"rank,omitempty"`

	Resource Entity `json:"resource"`

	// ID of the result template for rendering the search result. This ID must map to a display layout in the
	// resultTemplates dictionary, included in the searchresponse as well.
	ResultTemplateId nullable.Type[string] `json:"resultTemplateId,omitempty"`

	// A summary of the result, if a summary is available.
	Summary nullable.Type[string] `json:"summary,omitempty"`

	_Id      nullable.Type[string] `json:"_id,omitempty"`
	_Score   nullable.Type[int64]  `json:"_score,omitempty"`
	_Source  Entity                `json:"_source"`
	_Summary nullable.Type[string] `json:"_summary,omitempty"`
}

var _ json.Unmarshaler = &SearchHit{}

func (s *SearchHit) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ContentSource    nullable.Type[string] `json:"contentSource,omitempty"`
		HitId            nullable.Type[string] `json:"hitId,omitempty"`
		IsCollapsed      nullable.Type[bool]   `json:"isCollapsed,omitempty"`
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
		Rank             nullable.Type[int64]  `json:"rank,omitempty"`
		ResultTemplateId nullable.Type[string] `json:"resultTemplateId,omitempty"`
		Summary          nullable.Type[string] `json:"summary,omitempty"`
		_Id              nullable.Type[string] `json:"_id,omitempty"`
		_Score           nullable.Type[int64]  `json:"_score,omitempty"`
		_Summary         nullable.Type[string] `json:"_summary,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ContentSource = decoded.ContentSource
	s.HitId = decoded.HitId
	s.IsCollapsed = decoded.IsCollapsed
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Rank = decoded.Rank
	s.ResultTemplateId = decoded.ResultTemplateId
	s.Summary = decoded.Summary
	s._Id = decoded._Id
	s._Score = decoded._Score
	s._Summary = decoded._Summary

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SearchHit into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["resource"]; ok {
		impl, err := UnmarshalEntityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Resource' for 'SearchHit': %+v", err)
		}
		s.Resource = impl
	}

	if v, ok := temp["_source"]; ok {
		impl, err := UnmarshalEntityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field '_Source' for 'SearchHit': %+v", err)
		}
		s._Source = impl
	}

	return nil
}
