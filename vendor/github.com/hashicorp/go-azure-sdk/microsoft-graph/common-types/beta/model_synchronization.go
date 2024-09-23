package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Synchronization{}

type Synchronization struct {
	// Performs synchronization by periodically running in the background, polling for changes in one directory, and pushing
	// them to another directory.
	Jobs *[]SynchronizationJob `json:"jobs,omitempty"`

	// Represents a collection of credentials to access provisioned cloud applications.
	Secrets *[]SynchronizationSecretKeyStringValuePair `json:"secrets,omitempty"`

	// Pre-configured synchronization settings for a particular application.
	Templates *[]SynchronizationTemplate `json:"templates,omitempty"`

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

func (s Synchronization) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Synchronization{}

func (s Synchronization) MarshalJSON() ([]byte, error) {
	type wrapper Synchronization
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Synchronization: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Synchronization: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.synchronization"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Synchronization: %+v", err)
	}

	return encoded, nil
}
