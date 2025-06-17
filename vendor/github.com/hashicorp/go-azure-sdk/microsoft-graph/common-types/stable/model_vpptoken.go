package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = VppToken{}

type VppToken struct {
	// The Apple ID associated with the given Apple Volume Purchase Program Token.
	AppleId nullable.Type[string] `json:"appleId,omitempty"`

	// Whether or not apps for the VPP token will be automatically updated.
	AutomaticallyUpdateApps *bool `json:"automaticallyUpdateApps,omitempty"`

	// The country or region associated with the Apple Volume Purchase Program Token.
	CountryOrRegion nullable.Type[string] `json:"countryOrRegion,omitempty"`

	// The expiration date time of the Apple Volume Purchase Program Token.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Last modification date time associated with the Apple Volume Purchase Program Token.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The last time when an application sync was done with the Apple volume purchase program service using the the Apple
	// Volume Purchase Program Token.
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// Possible sync statuses associated with an Apple Volume Purchase Program token.
	LastSyncStatus *VppTokenSyncStatus `json:"lastSyncStatus,omitempty"`

	// The organization associated with the Apple Volume Purchase Program Token
	OrganizationName nullable.Type[string] `json:"organizationName,omitempty"`

	// Possible states associated with an Apple Volume Purchase Program token.
	State *VppTokenState `json:"state,omitempty"`

	// The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
	Token nullable.Type[string] `json:"token,omitempty"`

	// Possible types of an Apple Volume Purchase Program token.
	VppTokenAccountType *VppTokenAccountType `json:"vppTokenAccountType,omitempty"`

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

func (s VppToken) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VppToken{}

func (s VppToken) MarshalJSON() ([]byte, error) {
	type wrapper VppToken
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VppToken: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VppToken: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.vppToken"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VppToken: %+v", err)
	}

	return encoded, nil
}
