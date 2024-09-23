package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CustomSecurityAttributeAudit{}

type CustomSecurityAttributeAudit struct {
	ActivityDateTime    *string                 `json:"activityDateTime,omitempty"`
	ActivityDisplayName *string                 `json:"activityDisplayName,omitempty"`
	AdditionalDetails   *[]KeyValue             `json:"additionalDetails,omitempty"`
	Category            *string                 `json:"category,omitempty"`
	CorrelationId       nullable.Type[string]   `json:"correlationId,omitempty"`
	InitiatedBy         *AuditActivityInitiator `json:"initiatedBy,omitempty"`
	LoggedByService     nullable.Type[string]   `json:"loggedByService,omitempty"`
	OperationType       nullable.Type[string]   `json:"operationType,omitempty"`
	Result              *OperationResult        `json:"result,omitempty"`
	ResultReason        nullable.Type[string]   `json:"resultReason,omitempty"`
	TargetResources     *[]TargetResource       `json:"targetResources,omitempty"`
	UserAgent           nullable.Type[string]   `json:"userAgent,omitempty"`

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

func (s CustomSecurityAttributeAudit) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CustomSecurityAttributeAudit{}

func (s CustomSecurityAttributeAudit) MarshalJSON() ([]byte, error) {
	type wrapper CustomSecurityAttributeAudit
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomSecurityAttributeAudit: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomSecurityAttributeAudit: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customSecurityAttributeAudit"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomSecurityAttributeAudit: %+v", err)
	}

	return encoded, nil
}
