package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessReviewHistoryInstance{}

type AccessReviewHistoryInstance struct {
	// Uri that can be used to retrieve review history data. This URI will be active for 24 hours after being generated.
	// Required.
	DownloadUri nullable.Type[string] `json:"downloadUri,omitempty"`

	// Timestamp when this instance and associated data expires and the history is deleted. Required.
	ExpirationDateTime string `json:"expirationDateTime"`

	// Timestamp when all of the available data for this instance was collected. This will be set after this instance's
	// status is set to done. Required.
	FulfilledDateTime nullable.Type[string] `json:"fulfilledDateTime,omitempty"`

	// The date and time for which reviews ended before this date are included in the fetched history data.
	ReviewHistoryPeriodEndDateTime nullable.Type[string] `json:"reviewHistoryPeriodEndDateTime,omitempty"`

	// The date and time for which reviews started on or after this date are included in the fetched history data.
	ReviewHistoryPeriodStartDateTime nullable.Type[string] `json:"reviewHistoryPeriodStartDateTime,omitempty"`

	// The date and time when the instance's history data is scheduled to be generated.
	RunDateTime *string `json:"runDateTime,omitempty"`

	// Represents the status of the review history data collection. The possible values are: done, inProgress, error,
	// requested, unknownFutureValue. Once the status is marked as done, you can generate a link retrieve the instance's
	// data by calling generateDownloadUri method.
	Status *AccessReviewHistoryStatus `json:"status,omitempty"`

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

func (s AccessReviewHistoryInstance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewHistoryInstance{}

func (s AccessReviewHistoryInstance) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewHistoryInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewHistoryInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewHistoryInstance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewHistoryInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewHistoryInstance: %+v", err)
	}

	return encoded, nil
}
