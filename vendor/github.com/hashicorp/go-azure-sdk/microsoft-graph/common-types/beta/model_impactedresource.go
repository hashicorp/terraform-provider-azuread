package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ImpactedResource{}

type ImpactedResource struct {
	// The date and time when the impactedResource object was initially associated with the recommendation.
	AddedDateTime *string `json:"addedDateTime,omitempty"`

	// Additional information unique to the impactedResource to help contextualize the recommendation.
	AdditionalDetails *[]KeyValue `json:"additionalDetails,omitempty"`

	// The URL link to the corresponding Microsoft Entra resource.
	ApiUrl nullable.Type[string] `json:"apiUrl,omitempty"`

	// Friendly name of the Microsoft Entra resource.
	DisplayName *string `json:"displayName,omitempty"`

	// Name of the user or service that last updated the status.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The date and time when the status was last updated.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The user responsible for maintaining the resource.
	Owner nullable.Type[string] `json:"owner,omitempty"`

	// The URL link to the corresponding Microsoft Entra admin center page of the resource.
	PortalUrl nullable.Type[string] `json:"portalUrl,omitempty"`

	// The future date and time when the status of a postponed impactedResource will be active again.
	PostponeUntilDateTime nullable.Type[string] `json:"postponeUntilDateTime,omitempty"`

	// Indicates the importance of the resource. A resource with a rank equal to 1 is of the highest importance.
	Rank nullable.Type[int64] `json:"rank,omitempty"`

	// The unique identifier of the recommendation that the resource is associated with.
	RecommendationId *string `json:"recommendationId,omitempty"`

	// Indicates the type of Microsoft Entra resource. Examples include user, application.
	ResourceType *string `json:"resourceType,omitempty"`

	Status *RecommendationStatus `json:"status,omitempty"`

	// The related unique identifier, depending on the resourceType. For example, this property is set to the applicationId
	// if the resourceType is an application.
	SubjectId *string `json:"subjectId,omitempty"`

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

func (s ImpactedResource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ImpactedResource{}

func (s ImpactedResource) MarshalJSON() ([]byte, error) {
	type wrapper ImpactedResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ImpactedResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ImpactedResource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.impactedResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ImpactedResource: %+v", err)
	}

	return encoded, nil
}
