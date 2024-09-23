package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TargetedManagedAppProtection = IosManagedAppProtection{}

type IosManagedAppProtection struct {
	// Indicates if content sync for widgets is allowed for iOS on App Protection Policies
	AllowWidgetContentSync *bool `json:"allowWidgetContentSync,omitempty"`

	// Semicolon seperated list of device models allowed, as a string, for the managed app to work.
	AllowedIosDeviceModels nullable.Type[string] `json:"allowedIosDeviceModels,omitempty"`

	// Defines a managed app behavior, either block or warn, if the user is clocked out (non-working time). Possible values
	// are: block, wipe, warn.
	AppActionIfAccountIsClockedOut *ManagedAppRemediationAction `json:"appActionIfAccountIsClockedOut,omitempty"`

	// An admin initiated action to be applied on a managed app.
	AppActionIfIosDeviceModelNotAllowed *ManagedAppRemediationAction `json:"appActionIfIosDeviceModelNotAllowed,omitempty"`

	// Represents the level to which app data is encrypted for managed apps
	AppDataEncryptionType *ManagedAppDataEncryptionType `json:"appDataEncryptionType,omitempty"`

	// List of apps to which the policy is deployed.
	Apps *[]ManagedMobileApp `json:"apps,omitempty"`

	// A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired
	// should be true.
	CustomBrowserProtocol nullable.Type[string] `json:"customBrowserProtocol,omitempty"`

	// Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
	CustomDialerAppProtocol nullable.Type[string] `json:"customDialerAppProtocol,omitempty"`

	// Count of apps to which the current policy is deployed.
	DeployedAppCount *int64 `json:"deployedAppCount,omitempty"`

	// Navigation property to deployment summary of the configuration.
	DeploymentSummary *ManagedAppPolicyDeploymentSummary `json:"deploymentSummary,omitempty"`

	// Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be
	// True when AllowedOutboundDataTransferDestinations is set to ManagedApps.
	DisableProtectionOfManagedOutboundOpenInData *bool `json:"disableProtectionOfManagedOutboundOpenInData,omitempty"`

	// Apps in this list will be exempt from the policy and will be able to receive data from managed apps.
	ExemptedAppProtocols *[]KeyValuePair `json:"exemptedAppProtocols,omitempty"`

	// A list of custom urls that are allowed to invocate an unmanaged app
	ExemptedUniversalLinks *[]string `json:"exemptedUniversalLinks,omitempty"`

	// Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
	FaceIdBlocked *bool `json:"faceIdBlocked,omitempty"`

	// Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting
	// only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and
	// DisableProtectionOfManagedOutboundOpenInData is set to False.
	FilterOpenInToOnlyManagedApps *bool `json:"filterOpenInToOnlyManagedApps,omitempty"`

	// A list of custom urls that are allowed to invocate a managed app
	ManagedUniversalLinks *[]string `json:"managedUniversalLinks,omitempty"`

	// When a specific app redirection is enforced by protectedMessagingRedirectAppType in an App Protection Policy, this
	// value defines the app url redirect schemes which are allowed to be used.
	MessagingRedirectAppUrlScheme nullable.Type[string] `json:"messagingRedirectAppUrlScheme,omitempty"`

	// Versions less than the specified version will block the managed app from accessing company data.
	MinimumRequiredSdkVersion nullable.Type[string] `json:"minimumRequiredSdkVersion,omitempty"`

	// Versions less than the specified version will result in warning message on the managed app from accessing company
	// data.
	MinimumWarningSdkVersion nullable.Type[string] `json:"minimumWarningSdkVersion,omitempty"`

	// Versions less than the specified version will block the managed app from accessing company data.
	MinimumWipeSdkVersion nullable.Type[string] `json:"minimumWipeSdkVersion,omitempty"`

	// Protect incoming data from unknown source. This setting is only allowed to be True when
	// AllowedInboundDataTransferSources is set to AllApps.
	ProtectInboundDataFromUnknownSources *bool `json:"protectInboundDataFromUnknownSources,omitempty"`

	// Defines if third party keyboards are allowed while accessing a managed app
	ThirdPartyKeyboardsBlocked *bool `json:"thirdPartyKeyboardsBlocked,omitempty"`

	// Fields inherited from TargetedManagedAppProtection

	// Indicates a collection of apps to target which can be one of several pre-defined lists of apps or a manually selected
	// list of apps
	AppGroupType *TargetedManagedAppGroupType `json:"appGroupType,omitempty"`

	// Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
	Assignments *[]TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`

	// Indicates if the policy is deployed to any inclusion groups or not.
	IsAssigned *bool `json:"isAssigned,omitempty"`

	// Management levels for apps
	TargetedAppManagementLevels *AppManagementLevel `json:"targetedAppManagementLevels,omitempty"`

	// Fields inherited from ManagedAppProtection

	// Data storage locations where a user may store managed data.
	AllowedDataIngestionLocations *[]ManagedAppDataIngestionLocation `json:"allowedDataIngestionLocations,omitempty"`

	// Data storage locations where a user may store managed data.
	AllowedDataStorageLocations *[]ManagedAppDataStorageLocation `json:"allowedDataStorageLocations,omitempty"`

	// Data can be transferred from/to these classes of apps
	AllowedInboundDataTransferSources *ManagedAppDataTransferLevel `json:"allowedInboundDataTransferSources,omitempty"`

	// Specify the number of characters that may be cut or copied from Org data and accounts to any application. This
	// setting overrides the AllowedOutboundClipboardSharingLevel restriction. Default value of '0' means no exception is
	// allowed.
	AllowedOutboundClipboardSharingExceptionLength *int64 `json:"allowedOutboundClipboardSharingExceptionLength,omitempty"`

	// Represents the level to which the device's clipboard may be shared between apps
	AllowedOutboundClipboardSharingLevel *ManagedAppClipboardSharingLevel `json:"allowedOutboundClipboardSharingLevel,omitempty"`

	// Data can be transferred from/to these classes of apps
	AllowedOutboundDataTransferDestinations *ManagedAppDataTransferLevel `json:"allowedOutboundDataTransferDestinations,omitempty"`

	// An admin initiated action to be applied on a managed app.
	AppActionIfDeviceComplianceRequired *ManagedAppRemediationAction `json:"appActionIfDeviceComplianceRequired,omitempty"`

	// An admin initiated action to be applied on a managed app.
	AppActionIfMaximumPinRetriesExceeded *ManagedAppRemediationAction `json:"appActionIfMaximumPinRetriesExceeded,omitempty"`

	// If set, it will specify what action to take in the case where the user is unable to checkin because their
	// authentication token is invalid. This happens when the user is deleted or disabled in AAD. Possible values are:
	// block, wipe, warn.
	AppActionIfUnableToAuthenticateUser *ManagedAppRemediationAction `json:"appActionIfUnableToAuthenticateUser,omitempty"`

	// Indicates whether a user can bring data into org documents.
	BlockDataIngestionIntoOrganizationDocuments *bool `json:"blockDataIngestionIntoOrganizationDocuments,omitempty"`

	// Indicates whether contacts can be synced to the user's device.
	ContactSyncBlocked *bool `json:"contactSyncBlocked,omitempty"`

	// Indicates whether the backup of a managed app's data is blocked.
	DataBackupBlocked *bool `json:"dataBackupBlocked,omitempty"`

	// Indicates whether device compliance is required.
	DeviceComplianceRequired *bool `json:"deviceComplianceRequired,omitempty"`

	// The classes of apps that are allowed to click-to-open a phone number, for making phone calls or sending text
	// messages.
	DialerRestrictionLevel *ManagedAppPhoneNumberRedirectLevel `json:"dialerRestrictionLevel,omitempty"`

	// Indicates whether use of the app pin is required if the device pin is set.
	DisableAppPinIfDevicePinIsSet *bool `json:"disableAppPinIfDevicePinIsSet,omitempty"`

	// Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
	FingerprintBlocked *bool `json:"fingerprintBlocked,omitempty"`

	// A grace period before blocking app access during off clock hours.
	GracePeriodToBlockAppsDuringOffClockHours nullable.Type[string] `json:"gracePeriodToBlockAppsDuringOffClockHours,omitempty"`

	// Type of managed browser
	ManagedBrowser *ManagedBrowserType `json:"managedBrowser,omitempty"`

	// Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by
	// CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
	ManagedBrowserToOpenLinksRequired *bool `json:"managedBrowserToOpenLinksRequired,omitempty"`

	// The maxium threat level allowed for an app to be compliant.
	MaximumAllowedDeviceThreatLevel *ManagedAppDeviceThreatLevel `json:"maximumAllowedDeviceThreatLevel,omitempty"`

	// Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
	MaximumPinRetries *int64 `json:"maximumPinRetries,omitempty"`

	// Versions bigger than the specified version will block the managed app from accessing company data.
	MaximumRequiredOsVersion nullable.Type[string] `json:"maximumRequiredOsVersion,omitempty"`

	// Versions bigger than the specified version will block the managed app from accessing company data.
	MaximumWarningOsVersion nullable.Type[string] `json:"maximumWarningOsVersion,omitempty"`

	// Versions bigger than the specified version will block the managed app from accessing company data.
	MaximumWipeOsVersion nullable.Type[string] `json:"maximumWipeOsVersion,omitempty"`

	// Minimum pin length required for an app-level pin if PinRequired is set to True
	MinimumPinLength *int64 `json:"minimumPinLength,omitempty"`

	// Versions less than the specified version will block the managed app from accessing company data.
	MinimumRequiredAppVersion nullable.Type[string] `json:"minimumRequiredAppVersion,omitempty"`

	// Versions less than the specified version will block the managed app from accessing company data.
	MinimumRequiredOsVersion nullable.Type[string] `json:"minimumRequiredOsVersion,omitempty"`

	// Versions less than the specified version will result in warning message on the managed app.
	MinimumWarningAppVersion nullable.Type[string] `json:"minimumWarningAppVersion,omitempty"`

	// Versions less than the specified version will result in warning message on the managed app from accessing company
	// data.
	MinimumWarningOsVersion nullable.Type[string] `json:"minimumWarningOsVersion,omitempty"`

	// Versions less than or equal to the specified version will wipe the managed app and the associated company data.
	MinimumWipeAppVersion nullable.Type[string] `json:"minimumWipeAppVersion,omitempty"`

	// Versions less than or equal to the specified version will wipe the managed app and the associated company data.
	MinimumWipeOsVersion nullable.Type[string] `json:"minimumWipeOsVersion,omitempty"`

	// Indicates how to prioritize which Mobile Threat Defense (MTD) partner is enabled for a given platform, when more than
	// one is enabled. An app can only be actively using a single Mobile Threat Defense partner. When NULL, Microsoft
	// Defender will be given preference. Otherwise setting the value to defenderOverThirdPartyPartner or
	// thirdPartyPartnerOverDefender will make explicit which partner to prioritize. Possible values are: null,
	// defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender and unknownFutureValue. Default value is null. Possible
	// values are: defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender, unknownFutureValue.
	MobileThreatDefensePartnerPriority *MobileThreatDefensePartnerPriority `json:"mobileThreatDefensePartnerPriority,omitempty"`

	// An admin initiated action to be applied on a managed app.
	MobileThreatDefenseRemediationAction *ManagedAppRemediationAction `json:"mobileThreatDefenseRemediationAction,omitempty"`

	// Restrict managed app notification
	NotificationRestriction *ManagedAppNotificationRestriction `json:"notificationRestriction,omitempty"`

	// Indicates whether organizational credentials are required for app use.
	OrganizationalCredentialsRequired *bool `json:"organizationalCredentialsRequired,omitempty"`

	// TimePeriod before the all-level pin must be reset if PinRequired is set to True.
	PeriodBeforePinReset *string `json:"periodBeforePinReset,omitempty"`

	// The period after which access is checked when the device is not connected to the internet.
	PeriodOfflineBeforeAccessCheck *string `json:"periodOfflineBeforeAccessCheck,omitempty"`

	// The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
	PeriodOfflineBeforeWipeIsEnforced *string `json:"periodOfflineBeforeWipeIsEnforced,omitempty"`

	// The period after which access is checked when the device is connected to the internet.
	PeriodOnlineBeforeAccessCheck *string `json:"periodOnlineBeforeAccessCheck,omitempty"`

	// Character set which is to be used for a user's app PIN
	PinCharacterSet *ManagedAppPinCharacterSet `json:"pinCharacterSet,omitempty"`

	// Indicates whether an app-level pin is required.
	PinRequired *bool `json:"pinRequired,omitempty"`

	// Timeout in minutes for an app pin instead of non biometrics passcode
	PinRequiredInsteadOfBiometricTimeout nullable.Type[string] `json:"pinRequiredInsteadOfBiometricTimeout,omitempty"`

	// Requires a pin to be unique from the number specified in this property.
	PreviousPinBlockCount *int64 `json:"previousPinBlockCount,omitempty"`

	// Indicates whether printing is allowed from managed apps.
	PrintBlocked *bool `json:"printBlocked,omitempty"`

	// Defines how app messaging redirection is protected by an App Protection Policy. Default is anyApp.
	ProtectedMessagingRedirectAppType *MessagingRedirectAppType `json:"protectedMessagingRedirectAppType,omitempty"`

	// Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
	SaveAsBlocked *bool `json:"saveAsBlocked,omitempty"`

	// Indicates whether simplePin is blocked.
	SimplePinBlocked *bool `json:"simplePinBlocked,omitempty"`

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

func (s IosManagedAppProtection) TargetedManagedAppProtection() BaseTargetedManagedAppProtectionImpl {
	return BaseTargetedManagedAppProtectionImpl{
		AppGroupType:                                   s.AppGroupType,
		Assignments:                                    s.Assignments,
		IsAssigned:                                     s.IsAssigned,
		TargetedAppManagementLevels:                    s.TargetedAppManagementLevels,
		AllowedDataIngestionLocations:                  s.AllowedDataIngestionLocations,
		AllowedDataStorageLocations:                    s.AllowedDataStorageLocations,
		AllowedInboundDataTransferSources:              s.AllowedInboundDataTransferSources,
		AllowedOutboundClipboardSharingExceptionLength: s.AllowedOutboundClipboardSharingExceptionLength,
		AllowedOutboundClipboardSharingLevel:           s.AllowedOutboundClipboardSharingLevel,
		AllowedOutboundDataTransferDestinations:        s.AllowedOutboundDataTransferDestinations,
		AppActionIfDeviceComplianceRequired:            s.AppActionIfDeviceComplianceRequired,
		AppActionIfMaximumPinRetriesExceeded:           s.AppActionIfMaximumPinRetriesExceeded,
		AppActionIfUnableToAuthenticateUser:            s.AppActionIfUnableToAuthenticateUser,
		BlockDataIngestionIntoOrganizationDocuments:    s.BlockDataIngestionIntoOrganizationDocuments,
		ContactSyncBlocked:                             s.ContactSyncBlocked,
		DataBackupBlocked:                              s.DataBackupBlocked,
		DeviceComplianceRequired:                       s.DeviceComplianceRequired,
		DialerRestrictionLevel:                         s.DialerRestrictionLevel,
		DisableAppPinIfDevicePinIsSet:                  s.DisableAppPinIfDevicePinIsSet,
		FingerprintBlocked:                             s.FingerprintBlocked,
		GracePeriodToBlockAppsDuringOffClockHours:      s.GracePeriodToBlockAppsDuringOffClockHours,
		ManagedBrowser:                                 s.ManagedBrowser,
		ManagedBrowserToOpenLinksRequired:              s.ManagedBrowserToOpenLinksRequired,
		MaximumAllowedDeviceThreatLevel:                s.MaximumAllowedDeviceThreatLevel,
		MaximumPinRetries:                              s.MaximumPinRetries,
		MaximumRequiredOsVersion:                       s.MaximumRequiredOsVersion,
		MaximumWarningOsVersion:                        s.MaximumWarningOsVersion,
		MaximumWipeOsVersion:                           s.MaximumWipeOsVersion,
		MinimumPinLength:                               s.MinimumPinLength,
		MinimumRequiredAppVersion:                      s.MinimumRequiredAppVersion,
		MinimumRequiredOsVersion:                       s.MinimumRequiredOsVersion,
		MinimumWarningAppVersion:                       s.MinimumWarningAppVersion,
		MinimumWarningOsVersion:                        s.MinimumWarningOsVersion,
		MinimumWipeAppVersion:                          s.MinimumWipeAppVersion,
		MinimumWipeOsVersion:                           s.MinimumWipeOsVersion,
		MobileThreatDefensePartnerPriority:             s.MobileThreatDefensePartnerPriority,
		MobileThreatDefenseRemediationAction:           s.MobileThreatDefenseRemediationAction,
		NotificationRestriction:                        s.NotificationRestriction,
		OrganizationalCredentialsRequired:              s.OrganizationalCredentialsRequired,
		PeriodBeforePinReset:                           s.PeriodBeforePinReset,
		PeriodOfflineBeforeAccessCheck:                 s.PeriodOfflineBeforeAccessCheck,
		PeriodOfflineBeforeWipeIsEnforced:              s.PeriodOfflineBeforeWipeIsEnforced,
		PeriodOnlineBeforeAccessCheck:                  s.PeriodOnlineBeforeAccessCheck,
		PinCharacterSet:                                s.PinCharacterSet,
		PinRequired:                                    s.PinRequired,
		PinRequiredInsteadOfBiometricTimeout:           s.PinRequiredInsteadOfBiometricTimeout,
		PreviousPinBlockCount:                          s.PreviousPinBlockCount,
		PrintBlocked:                                   s.PrintBlocked,
		ProtectedMessagingRedirectAppType:              s.ProtectedMessagingRedirectAppType,
		SaveAsBlocked:                                  s.SaveAsBlocked,
		SimplePinBlocked:                               s.SimplePinBlocked,
		CreatedDateTime:                                s.CreatedDateTime,
		Description:                                    s.Description,
		DisplayName:                                    s.DisplayName,
		LastModifiedDateTime:                           s.LastModifiedDateTime,
		RoleScopeTagIds:                                s.RoleScopeTagIds,
		Version:                                        s.Version,
		Id:                                             s.Id,
		ODataId:                                        s.ODataId,
		ODataType:                                      s.ODataType,
	}
}

func (s IosManagedAppProtection) ManagedAppProtection() BaseManagedAppProtectionImpl {
	return BaseManagedAppProtectionImpl{
		AllowedDataIngestionLocations:                  s.AllowedDataIngestionLocations,
		AllowedDataStorageLocations:                    s.AllowedDataStorageLocations,
		AllowedInboundDataTransferSources:              s.AllowedInboundDataTransferSources,
		AllowedOutboundClipboardSharingExceptionLength: s.AllowedOutboundClipboardSharingExceptionLength,
		AllowedOutboundClipboardSharingLevel:           s.AllowedOutboundClipboardSharingLevel,
		AllowedOutboundDataTransferDestinations:        s.AllowedOutboundDataTransferDestinations,
		AppActionIfDeviceComplianceRequired:            s.AppActionIfDeviceComplianceRequired,
		AppActionIfMaximumPinRetriesExceeded:           s.AppActionIfMaximumPinRetriesExceeded,
		AppActionIfUnableToAuthenticateUser:            s.AppActionIfUnableToAuthenticateUser,
		BlockDataIngestionIntoOrganizationDocuments:    s.BlockDataIngestionIntoOrganizationDocuments,
		ContactSyncBlocked:                             s.ContactSyncBlocked,
		DataBackupBlocked:                              s.DataBackupBlocked,
		DeviceComplianceRequired:                       s.DeviceComplianceRequired,
		DialerRestrictionLevel:                         s.DialerRestrictionLevel,
		DisableAppPinIfDevicePinIsSet:                  s.DisableAppPinIfDevicePinIsSet,
		FingerprintBlocked:                             s.FingerprintBlocked,
		GracePeriodToBlockAppsDuringOffClockHours:      s.GracePeriodToBlockAppsDuringOffClockHours,
		ManagedBrowser:                                 s.ManagedBrowser,
		ManagedBrowserToOpenLinksRequired:              s.ManagedBrowserToOpenLinksRequired,
		MaximumAllowedDeviceThreatLevel:                s.MaximumAllowedDeviceThreatLevel,
		MaximumPinRetries:                              s.MaximumPinRetries,
		MaximumRequiredOsVersion:                       s.MaximumRequiredOsVersion,
		MaximumWarningOsVersion:                        s.MaximumWarningOsVersion,
		MaximumWipeOsVersion:                           s.MaximumWipeOsVersion,
		MinimumPinLength:                               s.MinimumPinLength,
		MinimumRequiredAppVersion:                      s.MinimumRequiredAppVersion,
		MinimumRequiredOsVersion:                       s.MinimumRequiredOsVersion,
		MinimumWarningAppVersion:                       s.MinimumWarningAppVersion,
		MinimumWarningOsVersion:                        s.MinimumWarningOsVersion,
		MinimumWipeAppVersion:                          s.MinimumWipeAppVersion,
		MinimumWipeOsVersion:                           s.MinimumWipeOsVersion,
		MobileThreatDefensePartnerPriority:             s.MobileThreatDefensePartnerPriority,
		MobileThreatDefenseRemediationAction:           s.MobileThreatDefenseRemediationAction,
		NotificationRestriction:                        s.NotificationRestriction,
		OrganizationalCredentialsRequired:              s.OrganizationalCredentialsRequired,
		PeriodBeforePinReset:                           s.PeriodBeforePinReset,
		PeriodOfflineBeforeAccessCheck:                 s.PeriodOfflineBeforeAccessCheck,
		PeriodOfflineBeforeWipeIsEnforced:              s.PeriodOfflineBeforeWipeIsEnforced,
		PeriodOnlineBeforeAccessCheck:                  s.PeriodOnlineBeforeAccessCheck,
		PinCharacterSet:                                s.PinCharacterSet,
		PinRequired:                                    s.PinRequired,
		PinRequiredInsteadOfBiometricTimeout:           s.PinRequiredInsteadOfBiometricTimeout,
		PreviousPinBlockCount:                          s.PreviousPinBlockCount,
		PrintBlocked:                                   s.PrintBlocked,
		ProtectedMessagingRedirectAppType:              s.ProtectedMessagingRedirectAppType,
		SaveAsBlocked:                                  s.SaveAsBlocked,
		SimplePinBlocked:                               s.SimplePinBlocked,
		CreatedDateTime:                                s.CreatedDateTime,
		Description:                                    s.Description,
		DisplayName:                                    s.DisplayName,
		LastModifiedDateTime:                           s.LastModifiedDateTime,
		RoleScopeTagIds:                                s.RoleScopeTagIds,
		Version:                                        s.Version,
		Id:                                             s.Id,
		ODataId:                                        s.ODataId,
		ODataType:                                      s.ODataType,
	}
}

func (s IosManagedAppProtection) ManagedAppPolicy() BaseManagedAppPolicyImpl {
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

func (s IosManagedAppProtection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosManagedAppProtection{}

func (s IosManagedAppProtection) MarshalJSON() ([]byte, error) {
	type wrapper IosManagedAppProtection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosManagedAppProtection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosManagedAppProtection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosManagedAppProtection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosManagedAppProtection: %+v", err)
	}

	return encoded, nil
}
