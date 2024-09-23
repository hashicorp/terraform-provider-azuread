package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DataSharingConsent{}

type DataSharingConsent struct {
	// The time consent was granted for this account
	GrantDateTime *string `json:"grantDateTime,omitempty"`

	// The granted state for the data sharing consent
	Granted *bool `json:"granted,omitempty"`

	// The Upn of the user that granted consent for this account
	GrantedByUpn nullable.Type[string] `json:"grantedByUpn,omitempty"`

	// The UserId of the user that granted consent for this account
	GrantedByUserId nullable.Type[string] `json:"grantedByUserId,omitempty"`

	// The display name of the service work flow
	ServiceDisplayName nullable.Type[string] `json:"serviceDisplayName,omitempty"`

	// The TermsUrl for the data sharing consent
	TermsUrl nullable.Type[string] `json:"termsUrl,omitempty"`

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

func (s DataSharingConsent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DataSharingConsent{}

func (s DataSharingConsent) MarshalJSON() ([]byte, error) {
	type wrapper DataSharingConsent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataSharingConsent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataSharingConsent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.dataSharingConsent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataSharingConsent: %+v", err)
	}

	return encoded, nil
}
