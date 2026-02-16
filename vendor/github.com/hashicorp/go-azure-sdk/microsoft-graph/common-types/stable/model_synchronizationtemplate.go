package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SynchronizationTemplate{}

type SynchronizationTemplate struct {
	// Identifier of the application this template belongs to.
	ApplicationId *string `json:"applicationId,omitempty"`

	// true if this template is recommended to be the default for the application.
	Default *bool `json:"default,omitempty"`

	// Description of the template.
	Description nullable.Type[string] `json:"description,omitempty"`

	// true if this template should appear in the collection of templates available for the application instance (service
	// principal).
	Discoverable *bool `json:"discoverable,omitempty"`

	// One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization
	// engine which implementation to use when processing jobs based on this template.
	FactoryTag nullable.Type[string] `json:"factoryTag,omitempty"`

	// Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
	Metadata *[]SynchronizationMetadataEntry `json:"metadata,omitempty"`

	// Default synchronization schema for the jobs based on this template.
	Schema *SynchronizationSchema `json:"schema,omitempty"`

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

func (s SynchronizationTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SynchronizationTemplate{}

func (s SynchronizationTemplate) MarshalJSON() ([]byte, error) {
	type wrapper SynchronizationTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SynchronizationTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SynchronizationTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.synchronizationTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SynchronizationTemplate: %+v", err)
	}

	return encoded, nil
}
