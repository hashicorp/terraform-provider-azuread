package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsUpdatesKnownIssue{}

type WindowsUpdatesKnownIssue struct {
	// The description of the particular known issue.
	Description *string `json:"description,omitempty"`

	KnownIssueHistories *[]WindowsUpdatesKnownIssueHistoryItem `json:"knownIssueHistories,omitempty"`

	// The date and time when the known issue was last updated. The timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	// Read-only.
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	// Knowledge base article associated with the release when the known issue was first reported.
	OriginatingKnowledgeBaseArticle *WindowsUpdatesKnowledgeBaseArticle `json:"originatingKnowledgeBaseArticle,omitempty"`

	// The date and time when the known issue was resolved or mitigated. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	ResolvedDateTime nullable.Type[string] `json:"resolvedDateTime,omitempty"`

	// Knowledge base article associated with the release when the known issue was resolved or mitigated.
	ResolvingKnowledgeBaseArticle *WindowsUpdatesKnowledgeBaseArticle `json:"resolvingKnowledgeBaseArticle,omitempty"`

	SafeguardHoldIds *[]int64 `json:"safeguardHoldIds,omitempty"`

	// The date and time when the known issue was first reported. The timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	StartDateTime *string `json:"startDateTime,omitempty"`

	Status *WindowsUpdatesWindowsReleaseHealthStatus `json:"status,omitempty"`

	// The title of the known issue.
	Title *string `json:"title,omitempty"`

	// The URL to the known issue in the Windows Release Health dashboard on Microsoft 365 admin center.
	WebViewUrl *string `json:"webViewUrl,omitempty"`

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

func (s WindowsUpdatesKnownIssue) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesKnownIssue{}

func (s WindowsUpdatesKnownIssue) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesKnownIssue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesKnownIssue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesKnownIssue: %+v", err)
	}

	delete(decoded, "lastUpdatedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.knownIssue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesKnownIssue: %+v", err)
	}

	return encoded, nil
}
