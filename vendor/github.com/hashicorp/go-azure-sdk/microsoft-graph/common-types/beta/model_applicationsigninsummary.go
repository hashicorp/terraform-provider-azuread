package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ApplicationSignInSummary{}

type ApplicationSignInSummary struct {
	// Name of the application that the user signed into.
	AppDisplayName *string `json:"appDisplayName,omitempty"`

	// Count of failed sign-ins made by the application.
	FailedSignInCount nullable.Type[int64] `json:"failedSignInCount,omitempty"`

	// Count of successful sign-ins made by the application.
	SuccessfulSignInCount nullable.Type[int64] `json:"successfulSignInCount,omitempty"`

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

func (s ApplicationSignInSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ApplicationSignInSummary{}

func (s ApplicationSignInSummary) MarshalJSON() ([]byte, error) {
	type wrapper ApplicationSignInSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApplicationSignInSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApplicationSignInSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.applicationSignInSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApplicationSignInSummary: %+v", err)
	}

	return encoded, nil
}
