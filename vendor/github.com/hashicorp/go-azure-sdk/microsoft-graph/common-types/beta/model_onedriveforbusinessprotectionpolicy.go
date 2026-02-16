package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionPolicyBase = OneDriveForBusinessProtectionPolicy{}

type OneDriveForBusinessProtectionPolicy struct {
	// Contains the details of the OneDrive for Work or School protection rule.
	DriveInclusionRules *[]DriveProtectionRule `json:"driveInclusionRules,omitempty"`

	// Contains the protection units associated with a OneDrive for Work or School protection policy.
	DriveProtectionUnits *[]DriveProtectionUnit `json:"driveProtectionUnits,omitempty"`

	DriveProtectionUnitsBulkAdditionJobs *[]DriveProtectionUnitsBulkAdditionJob `json:"driveProtectionUnitsBulkAdditionJobs,omitempty"`

	// Fields inherited from ProtectionPolicyBase

	// The identity of person who created the policy.
	CreatedBy IdentitySet `json:"createdBy"`

	// The time of creation of the policy.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The name of the policy to be created.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The identity of the person who last modified the policy.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The timestamp of the last modification of the policy.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Contains the retention setting details for the policy.
	RetentionSettings *[]RetentionSetting `json:"retentionSettings,omitempty"`

	// The aggregated status of the protection units associated with the policy. The possible values are: inactive,
	// activeWithErrors, updating, active, unknownFutureValue.
	Status *ProtectionPolicyStatus `json:"status,omitempty"`

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

func (s OneDriveForBusinessProtectionPolicy) ProtectionPolicyBase() BaseProtectionPolicyBaseImpl {
	return BaseProtectionPolicyBaseImpl{
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		DisplayName:          s.DisplayName,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		RetentionSettings:    s.RetentionSettings,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s OneDriveForBusinessProtectionPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OneDriveForBusinessProtectionPolicy{}

func (s OneDriveForBusinessProtectionPolicy) MarshalJSON() ([]byte, error) {
	type wrapper OneDriveForBusinessProtectionPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OneDriveForBusinessProtectionPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OneDriveForBusinessProtectionPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.oneDriveForBusinessProtectionPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OneDriveForBusinessProtectionPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OneDriveForBusinessProtectionPolicy{}

func (s *OneDriveForBusinessProtectionPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DriveInclusionRules                  *[]DriveProtectionRule                 `json:"driveInclusionRules,omitempty"`
		DriveProtectionUnits                 *[]DriveProtectionUnit                 `json:"driveProtectionUnits,omitempty"`
		DriveProtectionUnitsBulkAdditionJobs *[]DriveProtectionUnitsBulkAdditionJob `json:"driveProtectionUnitsBulkAdditionJobs,omitempty"`
		CreatedDateTime                      nullable.Type[string]                  `json:"createdDateTime,omitempty"`
		DisplayName                          nullable.Type[string]                  `json:"displayName,omitempty"`
		LastModifiedDateTime                 nullable.Type[string]                  `json:"lastModifiedDateTime,omitempty"`
		RetentionSettings                    *[]RetentionSetting                    `json:"retentionSettings,omitempty"`
		Status                               *ProtectionPolicyStatus                `json:"status,omitempty"`
		Id                                   *string                                `json:"id,omitempty"`
		ODataId                              *string                                `json:"@odata.id,omitempty"`
		ODataType                            *string                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DriveInclusionRules = decoded.DriveInclusionRules
	s.DriveProtectionUnits = decoded.DriveProtectionUnits
	s.DriveProtectionUnitsBulkAdditionJobs = decoded.DriveProtectionUnitsBulkAdditionJobs
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RetentionSettings = decoded.RetentionSettings
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OneDriveForBusinessProtectionPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'OneDriveForBusinessProtectionPolicy': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'OneDriveForBusinessProtectionPolicy': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
