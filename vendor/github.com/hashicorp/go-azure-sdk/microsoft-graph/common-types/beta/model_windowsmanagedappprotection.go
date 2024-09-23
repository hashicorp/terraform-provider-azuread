package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedAppPolicy = WindowsManagedAppProtection{}

type WindowsManagedAppProtection struct {
	// Data can be transferred from/to these classes of apps
	AllowedInboundDataTransferSources *WindowsManagedAppDataTransferLevel `json:"allowedInboundDataTransferSources,omitempty"`

	// Represents the level to which the device's clipboard may be shared between apps
	AllowedOutboundClipboardSharingLevel *WindowsManagedAppClipboardSharingLevel `json:"allowedOutboundClipboardSharingLevel,omitempty"`

	// Data can be transferred from/to these classes of apps
	AllowedOutboundDataTransferDestinations *WindowsManagedAppDataTransferLevel `json:"allowedOutboundDataTransferDestinations,omitempty"`

	// If set, it will specify what action to take in the case where the user is unable to checkin because their
	// authentication token is invalid. This happens when the user is deleted or disabled in AAD. Some possible values are
	// block or wipe. If this property is not set, no action will be taken. Possible values are: block, wipe, warn.
	AppActionIfUnableToAuthenticateUser *ManagedAppRemediationAction `json:"appActionIfUnableToAuthenticateUser,omitempty"`

	// List of apps to which the policy is deployed.
	Apps *[]ManagedMobileApp `json:"apps,omitempty"`

	// Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
	Assignments *[]TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`

	// Indicates the total number of applications for which the current policy is deployed.
	DeployedAppCount *int64 `json:"deployedAppCount,omitempty"`

	// Navigation property to deployment summary of the configuration.
	DeploymentSummary *ManagedAppPolicyDeploymentSummary `json:"deploymentSummary,omitempty"`

	// When TRUE, indicates that the policy is deployed to some inclusion groups. When FALSE, indicates that the policy is
	// not deployed to any inclusion groups. Default value is FALSE.
	IsAssigned *bool `json:"isAssigned,omitempty"`

	// The maxium threat level allowed for an app to be compliant.
	MaximumAllowedDeviceThreatLevel *ManagedAppDeviceThreatLevel `json:"maximumAllowedDeviceThreatLevel,omitempty"`

	// Versions bigger than the specified version will block the managed app from accessing company data. For example:
	// '8.1.0' or '13.1.1'.
	MaximumRequiredOsVersion nullable.Type[string] `json:"maximumRequiredOsVersion,omitempty"`

	// Versions bigger than the specified version will result in warning message on the managed app from accessing company
	// data. For example: '8.1.0' or '13.1.1'.
	MaximumWarningOsVersion nullable.Type[string] `json:"maximumWarningOsVersion,omitempty"`

	// Versions bigger than the specified version will wipe the managed app and the associated company data. For example:
	// '8.1.0' or '13.1.1'.
	MaximumWipeOsVersion nullable.Type[string] `json:"maximumWipeOsVersion,omitempty"`

	// Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0'
	// or '13.1.1'.
	MinimumRequiredAppVersion nullable.Type[string] `json:"minimumRequiredAppVersion,omitempty"`

	// Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0'
	// or '13.1.1'.
	MinimumRequiredOsVersion nullable.Type[string] `json:"minimumRequiredOsVersion,omitempty"`

	// Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0'
	// or '13.1.1'.
	MinimumRequiredSdkVersion nullable.Type[string] `json:"minimumRequiredSdkVersion,omitempty"`

	// Versions less than the specified version will result in warning message on the managed app from accessing company
	// data. For example: '8.1.0' or '13.1.1'.
	MinimumWarningAppVersion nullable.Type[string] `json:"minimumWarningAppVersion,omitempty"`

	// Versions less than the specified version will result in warning message on the managed app from accessing company
	// data. For example: '8.1.0' or '13.1.1'.
	MinimumWarningOsVersion nullable.Type[string] `json:"minimumWarningOsVersion,omitempty"`

	// Versions less than the specified version will wipe the managed app and the associated company data. For example:
	// '8.1.0' or '13.1.1'.
	MinimumWipeAppVersion nullable.Type[string] `json:"minimumWipeAppVersion,omitempty"`

	// Versions less than the specified version will wipe the managed app and the associated company data. For example:
	// '8.1.0' or '13.1.1'.
	MinimumWipeOsVersion nullable.Type[string] `json:"minimumWipeOsVersion,omitempty"`

	// Versions less than the specified version will wipe the managed app and the associated company data. For example:
	// '8.1.0' or '13.1.1'.
	MinimumWipeSdkVersion nullable.Type[string] `json:"minimumWipeSdkVersion,omitempty"`

	// An admin initiated action to be applied on a managed app.
	MobileThreatDefenseRemediationAction *ManagedAppRemediationAction `json:"mobileThreatDefenseRemediationAction,omitempty"`

	// The period after which access is checked when the device is not connected to the internet. For example, PT5M
	// indicates that the interval is 5 minutes in duration. A timespan value of PT0S indicates that access will be blocked
	// immediately when the device is not connected to the internet.
	PeriodOfflineBeforeAccessCheck *string `json:"periodOfflineBeforeAccessCheck,omitempty"`

	// The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
	// For example, P5D indicates that the interval is 5 days in duration. A timespan value of PT0S indicates that managed
	// data will never be wiped when the device is not connected to the internet.
	PeriodOfflineBeforeWipeIsEnforced *string `json:"periodOfflineBeforeWipeIsEnforced,omitempty"`

	// When TRUE, indicates that printing is blocked from managed apps. When FALSE, indicates that printing is allowed from
	// managed apps. Default value is FALSE.
	PrintBlocked *bool `json:"printBlocked,omitempty"`

	// Fields inherited from ManagedAppPolicy

	// The date and time the policy was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The policy's description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Policy display name.
	DisplayName *string `json:"displayName,omitempty"`

	// Last time the policy was modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Version of the entity.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s WindowsManagedAppProtection) ManagedAppPolicy() BaseManagedAppPolicyImpl {
	return BaseManagedAppPolicyImpl{
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		RoleScopeTagIds:      s.RoleScopeTagIds,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s WindowsManagedAppProtection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsManagedAppProtection{}

func (s WindowsManagedAppProtection) MarshalJSON() ([]byte, error) {
	type wrapper WindowsManagedAppProtection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsManagedAppProtection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsManagedAppProtection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsManagedAppProtection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsManagedAppProtection: %+v", err)
	}

	return encoded, nil
}
