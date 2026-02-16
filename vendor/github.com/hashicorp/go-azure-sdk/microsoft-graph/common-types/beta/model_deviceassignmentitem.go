package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAssignmentItem struct {
	// A list of possible assignment item action intent values on the application or configuration when executing this
	// action on the managed device. For example, if the application or configuration is intended to be removed on the
	// managed device, then the intent value is remove, and if the application or configuration already under removal
	// through previous actions and is now intended to be restored on the managed device, then the intent value is restore
	AssignmentItemActionIntent *DeviceAssignmentItemIntent `json:"assignmentItemActionIntent,omitempty"`

	// A list of possible assignment item action status values for the application or configuration regarding their executed
	// action on the managed device. For example, a configuration included in the deviceAssignmentItems list has just been
	// executed the action. Its status starts with inProgress until it's successfully removed to reflect as removed status
	// or failed to be removed to reflect as error status on the managed device. Similar status change happens for
	// restoration process
	AssignmentItemActionStatus *DeviceAssignmentItemStatus `json:"assignmentItemActionStatus,omitempty"`

	// The error code for the application or configuration regarding the failed executed action on the managed device.
	// Read-Only. Returned in the action result. 0 is default value and indicates no failure. Valid values
	// -9.22337203685478E+18 to 9.22337203685478E+18. This property is read-only.
	ErrorCode *int64 `json:"errorCode,omitempty"`

	// The intent action message for the application or configuration regarding the executed action on the managed device.
	// When the action is on error, this property provides message on the reason of failure. When the action is in progress,
	// this property provides message on what's being processed on the device. Read-Only. Returned in the action result. Can
	// be null. Max length is 1500. This property is read-only.
	IntentActionMessage nullable.Type[string] `json:"intentActionMessage,omitempty"`

	// The item displayName name for the application or configuration. Read-Only. Returned in the action result. Default
	// value is null. The property value cannot be modified and is automatically populated with the action result. Max
	// length is 200. This property is read-only.
	ItemDisplayName nullable.Type[string] `json:"itemDisplayName,omitempty"`

	// The unique identifier for the application or configuration. ItemId is required property which needs to be set in the
	// action POST request parameter for the DeviceAssignmentItem intended to remove. Max length is 40
	ItemId *string `json:"itemId,omitempty"`

	// Indicates the specific type for the application or configuration. For example, unknown, application,
	// appConfiguration, exploitProtection, bitLocker, deviceControl, microsoftEdgeBaseline,
	// attackSurfaceReductionRulesConfigMgr, endpointDetectionandResponse, windowsUpdateforBusiness,
	// microsoftDefenderFirewallRules, applicationControl, microsoftDefenderAntivirusexclusions, microsoftDefenderAntivirus,
	// wiredNetwork, derivedPersonalIdentityVerificationCredential, windowsHealthMonitoring, extensions, mxProfileZebraOnly,
	// deviceFirmwareConfigurationInterface, deliveryOptimization, identityProtection, kiosk, overrideGroupPolicy,
	// domainJoinPreview, pkcsImportedCertificate, networkBoundary, endpointProtection,
	// microsoftDefenderAtpWindows10Desktop, sharedMultiUserDevice, deviceFeatures, secureAssessmentEducation, wiFiImport,
	// editionUpgradeAndModeSwitch, vpn, custom, softwareUpdates, deviceRestrictionsWindows10Team, email,
	// trustedCertificate, scepCertificate, emailSamsungKnoxOnly, pkcsCertificate, deviceRestrictions, wiFi,
	// settingsCatalog. Read-Only. Returned in the action result. Default value is null. The property value cannot be
	// modified and is automatically populated with the action result. Max length is 200. This property is read-only.
	ItemSubTypeDisplayName nullable.Type[string] `json:"itemSubTypeDisplayName,omitempty"`

	// A list of possible device assignment item types to execute this action on the managed device. Device assignment item
	// represents existing assigned Intune resource such as application or configuration. Currently supported device
	// assignment item types are Application, DeviceConfiguration, DeviceManagementConfigurationPolicy and
	// MobileAppConfiguration
	ItemType *DeviceAssignmentItemType `json:"itemType,omitempty"`

	// The date and time when the application or configuration was initiated an action execution. Read-Only. Returned in the
	// action result. The property value cannot be modified and is automatically populated when the action is initiated. The
	// Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2025 would look like this: '2025-01-01T00:00:00Z'. This property is read-only.
	LastActionDateTime *string `json:"lastActionDateTime,omitempty"`

	// The date and time when the application or configuration was last modified because of either action execution or
	// status change. Read-Only. Returned in the action result. The property value cannot be modified and is automatically
	// populated when the action is initiated or the device has a status change. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2025 would look like
	// this: '2025-01-01T00:00:00Z'. This property is read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = DeviceAssignmentItem{}

func (s DeviceAssignmentItem) MarshalJSON() ([]byte, error) {
	type wrapper DeviceAssignmentItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceAssignmentItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAssignmentItem: %+v", err)
	}

	delete(decoded, "errorCode")
	delete(decoded, "intentActionMessage")
	delete(decoded, "itemDisplayName")
	delete(decoded, "itemSubTypeDisplayName")
	delete(decoded, "lastActionDateTime")
	delete(decoded, "lastModifiedDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceAssignmentItem: %+v", err)
	}

	return encoded, nil
}
