package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityHealthIssue{}

type SecurityHealthIssue struct {
	AdditionalInformation     *[]string                    `json:"additionalInformation,omitempty"`
	CreatedDateTime           *string                      `json:"createdDateTime,omitempty"`
	Description               *string                      `json:"description,omitempty"`
	DisplayName               nullable.Type[string]        `json:"displayName,omitempty"`
	DomainNames               *[]string                    `json:"domainNames,omitempty"`
	HealthIssueType           *SecurityHealthIssueType     `json:"healthIssueType,omitempty"`
	IssueTypeId               nullable.Type[string]        `json:"issueTypeId,omitempty"`
	LastModifiedDateTime      *string                      `json:"lastModifiedDateTime,omitempty"`
	Recommendations           *[]string                    `json:"recommendations,omitempty"`
	RecommendedActionCommands *[]string                    `json:"recommendedActionCommands,omitempty"`
	SensorDNSNames            *[]string                    `json:"sensorDNSNames,omitempty"`
	Severity                  *SecurityHealthIssueSeverity `json:"severity,omitempty"`
	Status                    *SecurityHealthIssueStatus   `json:"status,omitempty"`

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

func (s SecurityHealthIssue) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityHealthIssue{}

func (s SecurityHealthIssue) MarshalJSON() ([]byte, error) {
	type wrapper SecurityHealthIssue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityHealthIssue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityHealthIssue: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.healthIssue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityHealthIssue: %+v", err)
	}

	return encoded, nil
}
