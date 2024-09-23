package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MacOSSoftwareUpdateCategorySummary{}

type MacOSSoftwareUpdateCategorySummary struct {
	// The device ID.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The name of the report
	DisplayName *string `json:"displayName,omitempty"`

	// Number of failed updates on the device
	FailedUpdateCount *int64 `json:"failedUpdateCount,omitempty"`

	// Last date time the report for this device was updated.
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	// Number of successful updates on the device
	SuccessfulUpdateCount *int64 `json:"successfulUpdateCount,omitempty"`

	// Number of total updates on the device
	TotalUpdateCount *int64 `json:"totalUpdateCount,omitempty"`

	// MacOS Software Update Category
	UpdateCategory *MacOSSoftwareUpdateCategory `json:"updateCategory,omitempty"`

	// Summary of the update states.
	UpdateStateSummaries *[]MacOSSoftwareUpdateStateSummary `json:"updateStateSummaries,omitempty"`

	// The user ID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

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

func (s MacOSSoftwareUpdateCategorySummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOSSoftwareUpdateCategorySummary{}

func (s MacOSSoftwareUpdateCategorySummary) MarshalJSON() ([]byte, error) {
	type wrapper MacOSSoftwareUpdateCategorySummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOSSoftwareUpdateCategorySummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSSoftwareUpdateCategorySummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSSoftwareUpdateCategorySummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOSSoftwareUpdateCategorySummary: %+v", err)
	}

	return encoded, nil
}
