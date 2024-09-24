package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityHealthIssue{}

type SecurityHealthIssue struct {
	// Contains additional information about the issue, such as a list of items to fix.
	AdditionalInformation *[]string `json:"additionalInformation,omitempty"`

	// The date and time of when the health issue was generated.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Contains more detailed information about the health issue.
	Description *string `json:"description,omitempty"`

	// The display name of the health issue.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// A list of the fully qualified domain names of the domains or the sensors the health issue is related to.
	DomainNames *[]string `json:"domainNames,omitempty"`

	// The type of the health issue. The possible values are: sensor, global, unknownFutureValue. For a list of all health
	// issues and their identifiers, see Microsoft Defender for Identity health issues.
	HealthIssueType *SecurityHealthIssueType `json:"healthIssueType,omitempty"`

	// The type identifier of the health issue. For a list of all health issues and their identifiers, see Microsoft
	// Defender for Identity health issues.
	IssueTypeId nullable.Type[string] `json:"issueTypeId,omitempty"`

	// The date and time of when the health issue was last updated.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// This field contains a list of recommended actions that can be taken to resolve the issue effectively and efficiently.
	// These actions might include how to investigate the issue further. Not limited to prewritten responses.
	Recommendations *[]string `json:"recommendations,omitempty"`

	// Contains a list of commands from the product's PowerShell module that can be used to resolve the issue, if available.
	// If there aren't any commands that can be used to solve the issue, this field is empty. The commands, if present,
	// provide a quick and efficient way to address the issue. The commands run in order for the single recommended fix.
	RecommendedActionCommands *[]string `json:"recommendedActionCommands,omitempty"`

	// A list of the dns names of the sensors the health issue is related to.
	SensorDNSNames *[]string `json:"sensorDNSNames,omitempty"`

	// The severity of the health issue. The possible values are: low, medium, high, unknownFutureValue.
	Severity *SecurityHealthIssueSeverity `json:"severity,omitempty"`

	// The status of the health issue. The possible values are: open, closed, suppressed, unknownFutureValue.
	Status *SecurityHealthIssueStatus `json:"status,omitempty"`

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
