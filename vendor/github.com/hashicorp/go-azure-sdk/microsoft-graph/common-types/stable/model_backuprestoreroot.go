package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = BackupRestoreRoot{}

type BackupRestoreRoot struct {
	// The list of drive inclusion rules applied to the tenant.
	DriveInclusionRules *[]DriveProtectionRule `json:"driveInclusionRules,omitempty"`

	// The list of drive protection units in the tenant.
	DriveProtectionUnits *[]DriveProtectionUnit `json:"driveProtectionUnits,omitempty"`

	DriveProtectionUnitsBulkAdditionJobs *[]DriveProtectionUnitsBulkAdditionJob `json:"driveProtectionUnitsBulkAdditionJobs,omitempty"`

	// The list of Exchange protection policies in the tenant.
	ExchangeProtectionPolicies *[]ExchangeProtectionPolicy `json:"exchangeProtectionPolicies,omitempty"`

	// The list of Exchange restore sessions available in the tenant.
	ExchangeRestoreSessions *[]ExchangeRestoreSession `json:"exchangeRestoreSessions,omitempty"`

	// The list of mailbox inclusion rules applied to the tenant.
	MailboxInclusionRules *[]MailboxProtectionRule `json:"mailboxInclusionRules,omitempty"`

	// The list of mailbox protection units in the tenant.
	MailboxProtectionUnits *[]MailboxProtectionUnit `json:"mailboxProtectionUnits,omitempty"`

	MailboxProtectionUnitsBulkAdditionJobs *[]MailboxProtectionUnitsBulkAdditionJob `json:"mailboxProtectionUnitsBulkAdditionJobs,omitempty"`

	// The list of OneDrive for Business protection policies in the tenant.
	OneDriveForBusinessProtectionPolicies *[]OneDriveForBusinessProtectionPolicy `json:"oneDriveForBusinessProtectionPolicies,omitempty"`

	// The list of OneDrive for Business restore sessions available in the tenant.
	OneDriveForBusinessRestoreSessions *[]OneDriveForBusinessRestoreSession `json:"oneDriveForBusinessRestoreSessions,omitempty"`

	// List of protection policies in the tenant.
	ProtectionPolicies *[]ProtectionPolicyBase `json:"protectionPolicies,omitempty"`

	// List of protection units in the tenant.
	ProtectionUnits *[]ProtectionUnitBase `json:"protectionUnits,omitempty"`

	// List of restore points in the tenant.
	RestorePoints *[]RestorePoint `json:"restorePoints,omitempty"`

	// List of restore sessions in the tenant.
	RestoreSessions *[]RestoreSessionBase `json:"restoreSessions,omitempty"`

	// List of Backup Storage apps in the tenant.
	ServiceApps *[]ServiceApp `json:"serviceApps,omitempty"`

	// Represents the tenant-level status of the Backup Storage service.
	ServiceStatus *ServiceStatus `json:"serviceStatus,omitempty"`

	// The list of SharePoint protection policies in the tenant.
	SharePointProtectionPolicies *[]SharePointProtectionPolicy `json:"sharePointProtectionPolicies,omitempty"`

	// The list of SharePoint restore sessions available in the tenant.
	SharePointRestoreSessions *[]SharePointRestoreSession `json:"sharePointRestoreSessions,omitempty"`

	// The list of site inclusion rules applied to the tenant.
	SiteInclusionRules *[]SiteProtectionRule `json:"siteInclusionRules,omitempty"`

	// The list of site protection units in the tenant.
	SiteProtectionUnits *[]SiteProtectionUnit `json:"siteProtectionUnits,omitempty"`

	SiteProtectionUnitsBulkAdditionJobs *[]SiteProtectionUnitsBulkAdditionJob `json:"siteProtectionUnitsBulkAdditionJobs,omitempty"`

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

func (s BackupRestoreRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BackupRestoreRoot{}

func (s BackupRestoreRoot) MarshalJSON() ([]byte, error) {
	type wrapper BackupRestoreRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BackupRestoreRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BackupRestoreRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.backupRestoreRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BackupRestoreRoot: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BackupRestoreRoot{}

func (s *BackupRestoreRoot) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DriveInclusionRules                    *[]DriveProtectionRule                   `json:"driveInclusionRules,omitempty"`
		DriveProtectionUnits                   *[]DriveProtectionUnit                   `json:"driveProtectionUnits,omitempty"`
		DriveProtectionUnitsBulkAdditionJobs   *[]DriveProtectionUnitsBulkAdditionJob   `json:"driveProtectionUnitsBulkAdditionJobs,omitempty"`
		ExchangeProtectionPolicies             *[]ExchangeProtectionPolicy              `json:"exchangeProtectionPolicies,omitempty"`
		ExchangeRestoreSessions                *[]ExchangeRestoreSession                `json:"exchangeRestoreSessions,omitempty"`
		MailboxInclusionRules                  *[]MailboxProtectionRule                 `json:"mailboxInclusionRules,omitempty"`
		MailboxProtectionUnits                 *[]MailboxProtectionUnit                 `json:"mailboxProtectionUnits,omitempty"`
		MailboxProtectionUnitsBulkAdditionJobs *[]MailboxProtectionUnitsBulkAdditionJob `json:"mailboxProtectionUnitsBulkAdditionJobs,omitempty"`
		OneDriveForBusinessProtectionPolicies  *[]OneDriveForBusinessProtectionPolicy   `json:"oneDriveForBusinessProtectionPolicies,omitempty"`
		OneDriveForBusinessRestoreSessions     *[]OneDriveForBusinessRestoreSession     `json:"oneDriveForBusinessRestoreSessions,omitempty"`
		RestorePoints                          *[]RestorePoint                          `json:"restorePoints,omitempty"`
		ServiceApps                            *[]ServiceApp                            `json:"serviceApps,omitempty"`
		ServiceStatus                          *ServiceStatus                           `json:"serviceStatus,omitempty"`
		SharePointProtectionPolicies           *[]SharePointProtectionPolicy            `json:"sharePointProtectionPolicies,omitempty"`
		SharePointRestoreSessions              *[]SharePointRestoreSession              `json:"sharePointRestoreSessions,omitempty"`
		SiteInclusionRules                     *[]SiteProtectionRule                    `json:"siteInclusionRules,omitempty"`
		SiteProtectionUnits                    *[]SiteProtectionUnit                    `json:"siteProtectionUnits,omitempty"`
		SiteProtectionUnitsBulkAdditionJobs    *[]SiteProtectionUnitsBulkAdditionJob    `json:"siteProtectionUnitsBulkAdditionJobs,omitempty"`
		Id                                     *string                                  `json:"id,omitempty"`
		ODataId                                *string                                  `json:"@odata.id,omitempty"`
		ODataType                              *string                                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DriveInclusionRules = decoded.DriveInclusionRules
	s.DriveProtectionUnits = decoded.DriveProtectionUnits
	s.DriveProtectionUnitsBulkAdditionJobs = decoded.DriveProtectionUnitsBulkAdditionJobs
	s.ExchangeProtectionPolicies = decoded.ExchangeProtectionPolicies
	s.ExchangeRestoreSessions = decoded.ExchangeRestoreSessions
	s.MailboxInclusionRules = decoded.MailboxInclusionRules
	s.MailboxProtectionUnits = decoded.MailboxProtectionUnits
	s.MailboxProtectionUnitsBulkAdditionJobs = decoded.MailboxProtectionUnitsBulkAdditionJobs
	s.OneDriveForBusinessProtectionPolicies = decoded.OneDriveForBusinessProtectionPolicies
	s.OneDriveForBusinessRestoreSessions = decoded.OneDriveForBusinessRestoreSessions
	s.RestorePoints = decoded.RestorePoints
	s.ServiceApps = decoded.ServiceApps
	s.ServiceStatus = decoded.ServiceStatus
	s.SharePointProtectionPolicies = decoded.SharePointProtectionPolicies
	s.SharePointRestoreSessions = decoded.SharePointRestoreSessions
	s.SiteInclusionRules = decoded.SiteInclusionRules
	s.SiteProtectionUnits = decoded.SiteProtectionUnits
	s.SiteProtectionUnitsBulkAdditionJobs = decoded.SiteProtectionUnitsBulkAdditionJobs
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BackupRestoreRoot into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["protectionPolicies"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ProtectionPolicies into list []json.RawMessage: %+v", err)
		}

		output := make([]ProtectionPolicyBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalProtectionPolicyBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ProtectionPolicies' for 'BackupRestoreRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ProtectionPolicies = &output
	}

	if v, ok := temp["protectionUnits"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ProtectionUnits into list []json.RawMessage: %+v", err)
		}

		output := make([]ProtectionUnitBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalProtectionUnitBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ProtectionUnits' for 'BackupRestoreRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ProtectionUnits = &output
	}

	if v, ok := temp["restoreSessions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RestoreSessions into list []json.RawMessage: %+v", err)
		}

		output := make([]RestoreSessionBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRestoreSessionBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RestoreSessions' for 'BackupRestoreRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RestoreSessions = &output
	}

	return nil
}
