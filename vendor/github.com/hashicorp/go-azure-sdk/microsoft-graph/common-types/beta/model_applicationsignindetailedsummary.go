package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ApplicationSignInDetailedSummary{}

type ApplicationSignInDetailedSummary struct {
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	AggregatedEventDateTime nullable.Type[string] `json:"aggregatedEventDateTime,omitempty"`

	// Name of the application that the user signed in to.
	AppDisplayName *string `json:"appDisplayName,omitempty"`

	// ID of the application that the user signed in to.
	AppId *string `json:"appId,omitempty"`

	// Count of sign-ins made by the application.
	SignInCount nullable.Type[int64] `json:"signInCount,omitempty"`

	// Details of the sign-in status.
	Status *SignInStatus `json:"status,omitempty"`

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

func (s ApplicationSignInDetailedSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ApplicationSignInDetailedSummary{}

func (s ApplicationSignInDetailedSummary) MarshalJSON() ([]byte, error) {
	type wrapper ApplicationSignInDetailedSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApplicationSignInDetailedSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApplicationSignInDetailedSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.applicationSignInDetailedSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApplicationSignInDetailedSummary: %+v", err)
	}

	return encoded, nil
}
