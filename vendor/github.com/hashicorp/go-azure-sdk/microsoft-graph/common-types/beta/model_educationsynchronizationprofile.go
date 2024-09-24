package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationSynchronizationProfile{}

type EducationSynchronizationProfile struct {
	DataProvider EducationSynchronizationDataProvider `json:"dataProvider"`

	// Name of the configuration profile for syncing identities.
	DisplayName *string `json:"displayName,omitempty"`

	// All errors associated with this synchronization profile.
	Errors *[]EducationSynchronizationError `json:"errors,omitempty"`

	// The date the profile should be considered expired and cease syncing. Provide the date in YYYY-MM-DD format, following
	// ISO 8601. Maximum value is 18 months from profile creation. (optional)
	ExpirationDate nullable.Type[string] `json:"expirationDate,omitempty"`

	// Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
	HandleSpecialCharacterConstraint nullable.Type[bool] `json:"handleSpecialCharacterConstraint,omitempty"`

	IdentitySynchronizationConfiguration EducationIdentitySynchronizationConfiguration `json:"identitySynchronizationConfiguration"`

	// License setup configuration.
	LicensesToAssign *[]EducationSynchronizationLicenseAssignment `json:"licensesToAssign,omitempty"`

	// The synchronization status.
	ProfileStatus *EducationSynchronizationProfileStatus `json:"profileStatus,omitempty"`

	// The state of the profile. Possible values are: provisioning, provisioned, provisioningFailed, deleting,
	// deletionFailed.
	State *EducationSynchronizationProfileState `json:"state,omitempty"`

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

func (s EducationSynchronizationProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationSynchronizationProfile{}

func (s EducationSynchronizationProfile) MarshalJSON() ([]byte, error) {
	type wrapper EducationSynchronizationProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationSynchronizationProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSynchronizationProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationSynchronizationProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationSynchronizationProfile: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationSynchronizationProfile{}

func (s *EducationSynchronizationProfile) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName                      *string                                      `json:"displayName,omitempty"`
		Errors                           *[]EducationSynchronizationError             `json:"errors,omitempty"`
		ExpirationDate                   nullable.Type[string]                        `json:"expirationDate,omitempty"`
		HandleSpecialCharacterConstraint nullable.Type[bool]                          `json:"handleSpecialCharacterConstraint,omitempty"`
		LicensesToAssign                 *[]EducationSynchronizationLicenseAssignment `json:"licensesToAssign,omitempty"`
		ProfileStatus                    *EducationSynchronizationProfileStatus       `json:"profileStatus,omitempty"`
		State                            *EducationSynchronizationProfileState        `json:"state,omitempty"`
		Id                               *string                                      `json:"id,omitempty"`
		ODataId                          *string                                      `json:"@odata.id,omitempty"`
		ODataType                        *string                                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.Errors = decoded.Errors
	s.ExpirationDate = decoded.ExpirationDate
	s.HandleSpecialCharacterConstraint = decoded.HandleSpecialCharacterConstraint
	s.LicensesToAssign = decoded.LicensesToAssign
	s.ProfileStatus = decoded.ProfileStatus
	s.State = decoded.State
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationSynchronizationProfile into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dataProvider"]; ok {
		impl, err := UnmarshalEducationSynchronizationDataProviderImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DataProvider' for 'EducationSynchronizationProfile': %+v", err)
		}
		s.DataProvider = impl
	}

	if v, ok := temp["identitySynchronizationConfiguration"]; ok {
		impl, err := UnmarshalEducationIdentitySynchronizationConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentitySynchronizationConfiguration' for 'EducationSynchronizationProfile': %+v", err)
		}
		s.IdentitySynchronizationConfiguration = impl
	}

	return nil
}
