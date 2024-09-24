package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AuditEvent{}

type AuditEvent struct {
	// Friendly name of the activity.
	Activity nullable.Type[string] `json:"activity,omitempty"`

	// The date time in UTC when the activity was performed.
	ActivityDateTime *string `json:"activityDateTime,omitempty"`

	// The HTTP operation type of the activity.
	ActivityOperationType nullable.Type[string] `json:"activityOperationType,omitempty"`

	// The result of the activity.
	ActivityResult nullable.Type[string] `json:"activityResult,omitempty"`

	// The type of activity that was being performed.
	ActivityType nullable.Type[string] `json:"activityType,omitempty"`

	// AAD user and application that are associated with the audit event.
	Actor *AuditActor `json:"actor,omitempty"`

	// Audit category.
	Category nullable.Type[string] `json:"category,omitempty"`

	// Component name.
	ComponentName nullable.Type[string] `json:"componentName,omitempty"`

	// The client request Id that is used to correlate activity within the system.
	CorrelationId *string `json:"correlationId,omitempty"`

	// Event display name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Resources being modified.
	Resources *[]AuditResource `json:"resources,omitempty"`

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

func (s AuditEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AuditEvent{}

func (s AuditEvent) MarshalJSON() ([]byte, error) {
	type wrapper AuditEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuditEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuditEvent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.auditEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuditEvent: %+v", err)
	}

	return encoded, nil
}
