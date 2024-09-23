package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceAnnouncementBase = ServiceHealthIssue{}

type ServiceHealthIssue struct {
	Classification *ServiceHealthClassificationType `json:"classification,omitempty"`

	// The feature name of the service issue.
	Feature nullable.Type[string] `json:"feature,omitempty"`

	// The feature group name of the service issue.
	FeatureGroup nullable.Type[string] `json:"featureGroup,omitempty"`

	// The description of the service issue impact.
	ImpactDescription *string `json:"impactDescription,omitempty"`

	// Indicates whether the issue is resolved.
	IsResolved *bool `json:"isResolved,omitempty"`

	Origin *ServiceHealthOrigin `json:"origin,omitempty"`

	// Collection of historical posts for the service issue.
	Posts *[]ServiceHealthIssuePost `json:"posts,omitempty"`

	// Indicates the service affected by the issue.
	Service *string `json:"service,omitempty"`

	Status *ServiceHealthStatus `json:"status,omitempty"`

	// Fields inherited from ServiceAnnouncementBase

	// More details about service event. This property doesn't support filters.
	Details *[]KeyValuePair `json:"details,omitempty"`

	// The end time of the service event.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The last modified time of the service event.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The start time of the service event.
	StartDateTime *string `json:"startDateTime,omitempty"`

	// The title of the service event.
	Title *string `json:"title,omitempty"`

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

func (s ServiceHealthIssue) ServiceAnnouncementBase() BaseServiceAnnouncementBaseImpl {
	return BaseServiceAnnouncementBaseImpl{
		Details:              s.Details,
		EndDateTime:          s.EndDateTime,
		LastModifiedDateTime: s.LastModifiedDateTime,
		StartDateTime:        s.StartDateTime,
		Title:                s.Title,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s ServiceHealthIssue) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServiceHealthIssue{}

func (s ServiceHealthIssue) MarshalJSON() ([]byte, error) {
	type wrapper ServiceHealthIssue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceHealthIssue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceHealthIssue: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.serviceHealthIssue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceHealthIssue: %+v", err)
	}

	return encoded, nil
}
