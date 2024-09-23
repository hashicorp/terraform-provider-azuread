package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = VppToken{}

type VppToken struct {
	// The apple Id associated with the given Apple Volume Purchase Program Token.
	AppleId nullable.Type[string] `json:"appleId,omitempty"`

	// Whether or not apps for the VPP token will be automatically updated.
	AutomaticallyUpdateApps *bool `json:"automaticallyUpdateApps,omitempty"`

	// Admin consent to allow claiming token management from external MDM.
	ClaimTokenManagementFromExternalMdm *bool `json:"claimTokenManagementFromExternalMdm,omitempty"`

	// Whether or not apps for the VPP token will be automatically updated.
	CountryOrRegion nullable.Type[string] `json:"countryOrRegion,omitempty"`

	// Consent granted for data sharing with the Apple Volume Purchase Program.
	DataSharingConsentGranted *bool `json:"dataSharingConsentGranted,omitempty"`

	// An admin specified token friendly name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The expiration date time of the Apple Volume Purchase Program Token.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Last modification date time associated with the Apple Volume Purchase Program Token.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The last time when an application sync was done with the Apple volume purchase program service using the the Apple
	// Volume Purchase Program Token.
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// Possible sync statuses associated with an Apple Volume Purchase Program token.
	LastSyncStatus *VppTokenSyncStatus `json:"lastSyncStatus,omitempty"`

	// Token location returned from Apple VPP.
	LocationName nullable.Type[string] `json:"locationName,omitempty"`

	// The organization associated with the Apple Volume Purchase Program Token
	OrganizationName nullable.Type[string] `json:"organizationName,omitempty"`

	// Role Scope Tags IDs assigned to this entity.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Possible states associated with an Apple Volume Purchase Program token.
	State *VppTokenState `json:"state,omitempty"`

	// The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
	Token nullable.Type[string] `json:"token,omitempty"`

	// The collection of statuses of the actions performed on the Apple Volume Purchase Program Token.
	TokenActionResults *[]VppTokenActionResult `json:"tokenActionResults,omitempty"`

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

var _ json.Unmarshaler = &VppToken{}

func (s *VppToken) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AppleId                             nullable.Type[string] `json:"appleId,omitempty"`
		AutomaticallyUpdateApps             *bool                 `json:"automaticallyUpdateApps,omitempty"`
		ClaimTokenManagementFromExternalMdm *bool                 `json:"claimTokenManagementFromExternalMdm,omitempty"`
		CountryOrRegion                     nullable.Type[string] `json:"countryOrRegion,omitempty"`
		DataSharingConsentGranted           *bool                 `json:"dataSharingConsentGranted,omitempty"`
		DisplayName                         nullable.Type[string] `json:"displayName,omitempty"`
		ExpirationDateTime                  *string               `json:"expirationDateTime,omitempty"`
		LastModifiedDateTime                *string               `json:"lastModifiedDateTime,omitempty"`
		LastSyncDateTime                    *string               `json:"lastSyncDateTime,omitempty"`
		LastSyncStatus                      *VppTokenSyncStatus   `json:"lastSyncStatus,omitempty"`
		LocationName                        nullable.Type[string] `json:"locationName,omitempty"`
		OrganizationName                    nullable.Type[string] `json:"organizationName,omitempty"`
		RoleScopeTagIds                     *[]string             `json:"roleScopeTagIds,omitempty"`
		State                               *VppTokenState        `json:"state,omitempty"`
		Token                               nullable.Type[string] `json:"token,omitempty"`
		VppTokenAccountType                 *VppTokenAccountType  `json:"vppTokenAccountType,omitempty"`
		Id                                  *string               `json:"id,omitempty"`
		ODataId                             *string               `json:"@odata.id,omitempty"`
		ODataType                           *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppleId = decoded.AppleId
	s.AutomaticallyUpdateApps = decoded.AutomaticallyUpdateApps
	s.ClaimTokenManagementFromExternalMdm = decoded.ClaimTokenManagementFromExternalMdm
	s.CountryOrRegion = decoded.CountryOrRegion
	s.DataSharingConsentGranted = decoded.DataSharingConsentGranted
	s.DisplayName = decoded.DisplayName
	s.ExpirationDateTime = decoded.ExpirationDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.LastSyncDateTime = decoded.LastSyncDateTime
	s.LastSyncStatus = decoded.LastSyncStatus
	s.LocationName = decoded.LocationName
	s.OrganizationName = decoded.OrganizationName
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.State = decoded.State
	s.Token = decoded.Token
	s.VppTokenAccountType = decoded.VppTokenAccountType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling VppToken into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["tokenActionResults"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling TokenActionResults into list []json.RawMessage: %+v", err)
		}

		output := make([]VppTokenActionResult, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalVppTokenActionResultImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'TokenActionResults' for 'VppToken': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TokenActionResults = &output
	}

	return nil
}
