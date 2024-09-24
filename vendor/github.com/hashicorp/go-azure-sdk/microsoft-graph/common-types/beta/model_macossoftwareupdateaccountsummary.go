package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MacOSSoftwareUpdateAccountSummary{}

type MacOSSoftwareUpdateAccountSummary struct {
	// Summary of the updates by category.
	CategorySummaries *[]MacOSSoftwareUpdateCategorySummary `json:"categorySummaries,omitempty"`

	// The device ID.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The device name.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The name of the report
	DisplayName *string `json:"displayName,omitempty"`

	// Number of failed updates on the device.
	FailedUpdateCount *int64 `json:"failedUpdateCount,omitempty"`

	// Last date time the report for this device was updated.
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	// The OS version.
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// Number of successful updates on the device.
	SuccessfulUpdateCount *int64 `json:"successfulUpdateCount,omitempty"`

	// Number of total updates on the device.
	TotalUpdateCount *int64 `json:"totalUpdateCount,omitempty"`

	// The user ID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The user principal name
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s MacOSSoftwareUpdateAccountSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOSSoftwareUpdateAccountSummary{}

func (s MacOSSoftwareUpdateAccountSummary) MarshalJSON() ([]byte, error) {
	type wrapper MacOSSoftwareUpdateAccountSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOSSoftwareUpdateAccountSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSSoftwareUpdateAccountSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSSoftwareUpdateAccountSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOSSoftwareUpdateAccountSummary: %+v", err)
	}

	return encoded, nil
}
