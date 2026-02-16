package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerTaskCreation = PlannerTeamsPublicationInfo{}

type PlannerTeamsPublicationInfo struct {
	// The date and time when this task was last modified by the publication process. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The identifier of the publication. Read-only.
	PublicationId nullable.Type[string] `json:"publicationId,omitempty"`

	// The name of the published task list. Read-only.
	PublicationName nullable.Type[string] `json:"publicationName,omitempty"`

	// The identifier of the plannerPlan this task was originally placed in. Read-only.
	PublishedToPlanId nullable.Type[string] `json:"publishedToPlanId,omitempty"`

	// The identifier of the team that initiated the publication process. Read-only.
	PublishingTeamId nullable.Type[string] `json:"publishingTeamId,omitempty"`

	// The display name of the team that initiated the publication process. This display name is for reference only, and
	// might not represent the most up-to-date name of the team. Read-only.
	PublishingTeamName nullable.Type[string] `json:"publishingTeamName,omitempty"`

	// Fields inherited from PlannerTaskCreation

	// Specifies what kind of creation source the task is created with. The possible values are: external, publication and
	// unknownFutureValue.
	CreationSourceKind *PlannerCreationSourceKind `json:"creationSourceKind,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Information about the publication process that created this task. This field is deprecated and clients should move to
	// using the new inheritance model.
	TeamsPublicationInfo *PlannerTeamsPublicationInfo `json:"teamsPublicationInfo,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PlannerTeamsPublicationInfo) PlannerTaskCreation() BasePlannerTaskCreationImpl {
	return BasePlannerTaskCreationImpl{
		CreationSourceKind:   s.CreationSourceKind,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
		TeamsPublicationInfo: s.TeamsPublicationInfo,
	}
}

var _ json.Marshaler = PlannerTeamsPublicationInfo{}

func (s PlannerTeamsPublicationInfo) MarshalJSON() ([]byte, error) {
	type wrapper PlannerTeamsPublicationInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerTeamsPublicationInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerTeamsPublicationInfo: %+v", err)
	}

	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "publicationId")
	delete(decoded, "publicationName")
	delete(decoded, "publishedToPlanId")
	delete(decoded, "publishingTeamId")
	delete(decoded, "publishingTeamName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerTeamsPublicationInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerTeamsPublicationInfo: %+v", err)
	}

	return encoded, nil
}
