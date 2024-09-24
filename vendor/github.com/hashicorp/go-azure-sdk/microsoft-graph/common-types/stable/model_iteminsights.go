package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OfficeGraphInsights = ItemInsights{}

type ItemInsights struct {

	// Fields inherited from OfficeGraphInsights

	// Calculated relationship that identifies documents shared with or by the user. This includes URLs, file attachments,
	// and reference attachments to OneDrive for work or school and SharePoint files found in Outlook messages and meetings.
	// This also includes URLs and reference attachments to Teams conversations. Ordered by recency of share.
	Shared *[]SharedInsight `json:"shared,omitempty"`

	// Calculated relationship that identifies documents trending around a user. Trending documents are calculated based on
	// activity of the user's closest network of people and include files stored in OneDrive for work or school and
	// SharePoint. Trending insights help the user to discover potentially useful content that the user has access to, but
	// has never viewed before.
	Trending *[]Trending `json:"trending,omitempty"`

	// Calculated relationship that identifies the latest documents viewed or modified by a user, including OneDrive for
	// work or school and SharePoint documents, ranked by recency of use.
	Used *[]UsedInsight `json:"used,omitempty"`

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

func (s ItemInsights) OfficeGraphInsights() BaseOfficeGraphInsightsImpl {
	return BaseOfficeGraphInsightsImpl{
		Shared:    s.Shared,
		Trending:  s.Trending,
		Used:      s.Used,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s ItemInsights) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ItemInsights{}

func (s ItemInsights) MarshalJSON() ([]byte, error) {
	type wrapper ItemInsights
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ItemInsights: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ItemInsights: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.itemInsights"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ItemInsights: %+v", err)
	}

	return encoded, nil
}
