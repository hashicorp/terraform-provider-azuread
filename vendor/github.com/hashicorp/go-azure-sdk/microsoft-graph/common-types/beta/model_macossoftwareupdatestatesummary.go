package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MacOSSoftwareUpdateStateSummary{}

type MacOSSoftwareUpdateStateSummary struct {
	// Human readable name of the software update
	DisplayName *string `json:"displayName,omitempty"`

	// Last date time the report for this device and product key was updated.
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	// Product key of the software update.
	ProductKey nullable.Type[string] `json:"productKey,omitempty"`

	// MacOS Software Update State
	State *MacOSSoftwareUpdateState `json:"state,omitempty"`

	// MacOS Software Update Category
	UpdateCategory *MacOSSoftwareUpdateCategory `json:"updateCategory,omitempty"`

	// Version of the software update
	UpdateVersion nullable.Type[string] `json:"updateVersion,omitempty"`

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

func (s MacOSSoftwareUpdateStateSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOSSoftwareUpdateStateSummary{}

func (s MacOSSoftwareUpdateStateSummary) MarshalJSON() ([]byte, error) {
	type wrapper MacOSSoftwareUpdateStateSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOSSoftwareUpdateStateSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSSoftwareUpdateStateSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSSoftwareUpdateStateSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOSSoftwareUpdateStateSummary: %+v", err)
	}

	return encoded, nil
}
