package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ StatusBase = StatusDetails{}

type StatusDetails struct {
	// Additional details if there is an error.
	AdditionalDetails nullable.Type[string] `json:"additionalDetails,omitempty"`

	// Categorizes the error code. Possible values are Failure, NonServiceFailure, Success.
	ErrorCategory *ProvisioningStatusErrorCategory `json:"errorCategory,omitempty"`

	// Unique error code if any occurred. Learn more
	ErrorCode nullable.Type[string] `json:"errorCode,omitempty"`

	// Summarizes the status and describes why the status happened.
	Reason nullable.Type[string] `json:"reason,omitempty"`

	// Provides the resolution for the corresponding error.
	RecommendedAction nullable.Type[string] `json:"recommendedAction,omitempty"`

	// Fields inherited from StatusBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Possible values are: success, warning, failure, skipped, unknownFutureValue. Supports $filter (eq, contains).
	Status *ProvisioningResult `json:"status,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s StatusDetails) StatusBase() BaseStatusBaseImpl {
	return BaseStatusBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Status:    s.Status,
	}
}

var _ json.Marshaler = StatusDetails{}

func (s StatusDetails) MarshalJSON() ([]byte, error) {
	type wrapper StatusDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StatusDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StatusDetails: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.statusDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StatusDetails: %+v", err)
	}

	return encoded, nil
}
