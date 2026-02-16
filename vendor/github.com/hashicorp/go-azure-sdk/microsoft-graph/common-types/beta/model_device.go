package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = Device{}

type Device struct {
	// true if the account is enabled; otherwise, false. Default is true. Supports $filter (eq, ne, not, in). Only callers
	// with at least the Cloud Device Administrator role can set this property.
	AccountEnabled nullable.Type[bool] `json:"accountEnabled,omitempty"`

	// List of alternative names for the device.
	AlternativeNames *[]string `json:"alternativeNames,omitempty"`

	// For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
	AlternativeSecurityIds *[]AlternativeSecurityId `json:"alternativeSecurityIds,omitempty"`

	// The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null
	// values) and $orderby.
	ApproximateLastSignInDateTime nullable.Type[string] `json:"approximateLastSignInDateTime,omitempty"`

	// Set of commands sent to this device.
	Commands *[]Command `json:"commands,omitempty"`

	// The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	// Read-only.
	ComplianceExpirationDateTime nullable.Type[string] `json:"complianceExpirationDateTime,omitempty"`

	// User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
	DeviceCategory nullable.Type[string] `json:"deviceCategory,omitempty"`

	// Unique identifier set by Azure Device Registration Service at the time of registration. This ID is an alternate key
	// that can be used to reference the device object. Also supports $filter (eq, ne, not, startsWith).
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// For internal use only. Set to null.
	DeviceMetadata nullable.Type[string] `json:"deviceMetadata,omitempty"`

	// Ownership of the device. Intune sets this property. Possible values are: unknown, company, personal.
	DeviceOwnership nullable.Type[string] `json:"deviceOwnership,omitempty"`

	// Device template used to instantiate this device. Nullable. Read-only.
	DeviceTemplate *[]DeviceTemplate `json:"deviceTemplate,omitempty"`

	// For internal use only.
	DeviceVersion nullable.Type[int64] `json:"deviceVersion,omitempty"`

	// The display name for the device. Maximum length is 256 characters. Required. Supports $filter (eq, ne, not, ge, le,
	// in, startsWith, and eq on null values), $search, and $orderby.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The on-premises domain name of Microsoft Entra hybrid joined devices. Intune sets this property.
	DomainName nullable.Type[string] `json:"domainName,omitempty"`

	// Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate
	// device identifiers, or Windows Autopilot profile name. Intune sets this property.
	EnrollmentProfileName nullable.Type[string] `json:"enrollmentProfileName,omitempty"`

	// Enrollment type of the device. Intune sets this property. Possible values are: unknown, userEnrollment,
	// deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless,
	// windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement,
	// windowsAzureADJoinUsingDeviceAuth,appleUserEnrollment, appleUserEnrollmentWithServiceAccount. NOTE: This property
	// might return other values apart from those listed.
	EnrollmentType nullable.Type[string] `json:"enrollmentType,omitempty"`

	// Contains extension attributes 1-15 for the device. The individual extension attributes aren't selectable. These
	// properties are mastered in cloud and can be set during creation or update of a device object in Microsoft Entra ID.
	// Supports $filter (eq, not, startsWith, and eq on null values).
	ExtensionAttributes *OnPremisesExtensionAttributes `json:"extensionAttributes,omitempty"`

	// The collection of open extensions defined for the device. Read-only. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`

	// List of host names for the device.
	Hostnames *[]string `json:"hostnames,omitempty"`

	// true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only
	// be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq,
	// ne, not).
	IsCompliant nullable.Type[bool] `json:"isCompliant,omitempty"`

	// true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by
	// Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
	IsManaged nullable.Type[bool] `json:"isManaged,omitempty"`

	// Indicates whether the device is a member of a restricted management administrative unit. The default value is false.
	// Read-only. To manage a device that's a member of a restricted management administrative unit, the administrator or
	// calling app must be assigned a Microsoft Entra role at the scope of the restricted management administrative unit.
	IsManagementRestricted nullable.Type[bool] `json:"isManagementRestricted,omitempty"`

	// true if the device is rooted or jail-broken. This property can only be updated by Intune.
	IsRooted nullable.Type[bool] `json:"isRooted,omitempty"`

	// Form factor of the device. Only returned if the user signs in with a Microsoft account as part of Project Rome.
	Kind nullable.Type[string] `json:"kind,omitempty"`

	// Management channel of the device. Intune sets this property. Possible values are: eas, mdm, easMdm, intuneClient,
	// easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas,
	// unknown, jamf, googleCloudDevicePolicyController.
	ManagementType nullable.Type[string] `json:"managementType,omitempty"`

	// Manufacturer of the device. Read-only.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
	MdmAppId nullable.Type[string] `json:"mdmAppId,omitempty"`

	// Groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
	MemberOf *[]DirectoryObject `json:"memberOf,omitempty"`

	// List of OData IDs for `MemberOf` to bind to this entity
	MemberOf_ODataBind *[]string `json:"memberOf@odata.bind,omitempty"`

	// Model of the device. Read-only.
	Model nullable.Type[string] `json:"model,omitempty"`

	// Friendly name of the device. Only returned if user signs in with a Microsoft account as part of Project Rome.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and
	// time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
	OnPremisesLastSyncDateTime nullable.Type[string] `json:"onPremisesLastSyncDateTime,omitempty"`

	// The on-premises security identifier (SID) for the user who was synchronized from on-premises to the cloud. Read-only.
	// Returned only on $select. Supports $filter (eq).
	OnPremisesSecurityIdentifier nullable.Type[string] `json:"onPremisesSecurityIdentifier,omitempty"`

	// true if this object is synced from an on-premises directory; false if this object was originally synced from an
	// on-premises directory but is no longer synced; null if this object has never been synced from an on-premises
	// directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
	OnPremisesSyncEnabled nullable.Type[bool] `json:"onPremisesSyncEnabled,omitempty"`

	// The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on
	// null values).
	OperatingSystem nullable.Type[string] `json:"operatingSystem,omitempty"`

	// Operating system version of the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null
	// values).
	OperatingSystemVersion nullable.Type[string] `json:"operatingSystemVersion,omitempty"`

	// For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith, /$count eq 0, /$count ne 0.
	PhysicalIds *[]string `json:"physicalIds,omitempty"`

	// Platform of device. Only returned if the user signs in with a Microsoft account as part of Project Rome.
	Platform nullable.Type[string] `json:"platform,omitempty"`

	// The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
	ProfileType nullable.Type[string] `json:"profileType,omitempty"`

	// The user that cloud joined the device or registered their personal device. The registered owner is set at the time of
	// registration. Read-only. Nullable. Supports $expand.
	RegisteredOwners *[]DirectoryObject `json:"registeredOwners,omitempty"`

	// List of OData IDs for `RegisteredOwners` to bind to this entity
	RegisteredOwners_ODataBind *[]string `json:"registeredOwners@odata.bind,omitempty"`

	// Collection of registered users of the device. For cloud joined devices and registered personal devices, registered
	// users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports
	// $expand.
	RegisteredUsers *[]DirectoryObject `json:"registeredUsers,omitempty"`

	// List of OData IDs for `RegisteredUsers` to bind to this entity
	RegisteredUsers_ODataBind *[]string `json:"registeredUsers@odata.bind,omitempty"`

	// Date and time of when the device was registered. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	RegistrationDateTime nullable.Type[string] `json:"registrationDateTime,omitempty"`

	// Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
	Status nullable.Type[string] `json:"status,omitempty"`

	// List of labels applied to the device by the system. Supports $filter (/$count eq 0, /$count ne 0).
	SystemLabels *[]string `json:"systemLabels,omitempty"`

	// Groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
	TransitiveMemberOf *[]DirectoryObject `json:"transitiveMemberOf,omitempty"`

	// List of OData IDs for `TransitiveMemberOf` to bind to this entity
	TransitiveMemberOf_ODataBind *[]string `json:"transitiveMemberOf@odata.bind,omitempty"`

	// Type of trust for the joined device. Read-only. Possible values: Workplace (indicates bring your own personal
	// devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Microsoft Entra
	// ID). For more information, see Introduction to device management in Microsoft Entra ID. Supports $filter (eq, ne,
	// not, in).
	TrustType nullable.Type[string] `json:"trustType,omitempty"`

	// Represents the usage rights a device has been granted.
	UsageRights *[]UsageRight `json:"usageRights,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s Device) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s Device) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Device{}

func (s Device) MarshalJSON() ([]byte, error) {
	type wrapper Device
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Device: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Device: %+v", err)
	}

	delete(decoded, "approximateLastSignInDateTime")
	delete(decoded, "complianceExpirationDateTime")
	delete(decoded, "deviceTemplate")
	delete(decoded, "extensions")
	delete(decoded, "isCompliant")
	delete(decoded, "isManagementRestricted")
	delete(decoded, "manufacturer")
	delete(decoded, "mdmAppId")
	delete(decoded, "memberOf")
	delete(decoded, "model")
	delete(decoded, "onPremisesLastSyncDateTime")
	delete(decoded, "onPremisesSecurityIdentifier")
	delete(decoded, "onPremisesSyncEnabled")
	delete(decoded, "registeredOwners")
	delete(decoded, "registeredUsers")
	delete(decoded, "registrationDateTime")
	delete(decoded, "trustType")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.device"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Device: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Device{}

func (s *Device) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccountEnabled                nullable.Type[bool]            `json:"accountEnabled,omitempty"`
		AlternativeNames              *[]string                      `json:"alternativeNames,omitempty"`
		AlternativeSecurityIds        *[]AlternativeSecurityId       `json:"alternativeSecurityIds,omitempty"`
		ApproximateLastSignInDateTime nullable.Type[string]          `json:"approximateLastSignInDateTime,omitempty"`
		Commands                      *[]Command                     `json:"commands,omitempty"`
		ComplianceExpirationDateTime  nullable.Type[string]          `json:"complianceExpirationDateTime,omitempty"`
		DeviceCategory                nullable.Type[string]          `json:"deviceCategory,omitempty"`
		DeviceId                      nullable.Type[string]          `json:"deviceId,omitempty"`
		DeviceMetadata                nullable.Type[string]          `json:"deviceMetadata,omitempty"`
		DeviceOwnership               nullable.Type[string]          `json:"deviceOwnership,omitempty"`
		DeviceTemplate                *[]DeviceTemplate              `json:"deviceTemplate,omitempty"`
		DeviceVersion                 nullable.Type[int64]           `json:"deviceVersion,omitempty"`
		DisplayName                   nullable.Type[string]          `json:"displayName,omitempty"`
		DomainName                    nullable.Type[string]          `json:"domainName,omitempty"`
		EnrollmentProfileName         nullable.Type[string]          `json:"enrollmentProfileName,omitempty"`
		EnrollmentType                nullable.Type[string]          `json:"enrollmentType,omitempty"`
		ExtensionAttributes           *OnPremisesExtensionAttributes `json:"extensionAttributes,omitempty"`
		Hostnames                     *[]string                      `json:"hostnames,omitempty"`
		IsCompliant                   nullable.Type[bool]            `json:"isCompliant,omitempty"`
		IsManaged                     nullable.Type[bool]            `json:"isManaged,omitempty"`
		IsManagementRestricted        nullable.Type[bool]            `json:"isManagementRestricted,omitempty"`
		IsRooted                      nullable.Type[bool]            `json:"isRooted,omitempty"`
		Kind                          nullable.Type[string]          `json:"kind,omitempty"`
		ManagementType                nullable.Type[string]          `json:"managementType,omitempty"`
		Manufacturer                  nullable.Type[string]          `json:"manufacturer,omitempty"`
		MdmAppId                      nullable.Type[string]          `json:"mdmAppId,omitempty"`
		MemberOf_ODataBind            *[]string                      `json:"memberOf@odata.bind,omitempty"`
		Model                         nullable.Type[string]          `json:"model,omitempty"`
		Name                          nullable.Type[string]          `json:"name,omitempty"`
		OnPremisesLastSyncDateTime    nullable.Type[string]          `json:"onPremisesLastSyncDateTime,omitempty"`
		OnPremisesSecurityIdentifier  nullable.Type[string]          `json:"onPremisesSecurityIdentifier,omitempty"`
		OnPremisesSyncEnabled         nullable.Type[bool]            `json:"onPremisesSyncEnabled,omitempty"`
		OperatingSystem               nullable.Type[string]          `json:"operatingSystem,omitempty"`
		OperatingSystemVersion        nullable.Type[string]          `json:"operatingSystemVersion,omitempty"`
		PhysicalIds                   *[]string                      `json:"physicalIds,omitempty"`
		Platform                      nullable.Type[string]          `json:"platform,omitempty"`
		ProfileType                   nullable.Type[string]          `json:"profileType,omitempty"`
		RegisteredOwners_ODataBind    *[]string                      `json:"registeredOwners@odata.bind,omitempty"`
		RegisteredUsers_ODataBind     *[]string                      `json:"registeredUsers@odata.bind,omitempty"`
		RegistrationDateTime          nullable.Type[string]          `json:"registrationDateTime,omitempty"`
		Status                        nullable.Type[string]          `json:"status,omitempty"`
		SystemLabels                  *[]string                      `json:"systemLabels,omitempty"`
		TransitiveMemberOf_ODataBind  *[]string                      `json:"transitiveMemberOf@odata.bind,omitempty"`
		TrustType                     nullable.Type[string]          `json:"trustType,omitempty"`
		UsageRights                   *[]UsageRight                  `json:"usageRights,omitempty"`
		DeletedDateTime               nullable.Type[string]          `json:"deletedDateTime,omitempty"`
		Id                            *string                        `json:"id,omitempty"`
		ODataId                       *string                        `json:"@odata.id,omitempty"`
		ODataType                     *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccountEnabled = decoded.AccountEnabled
	s.AlternativeNames = decoded.AlternativeNames
	s.AlternativeSecurityIds = decoded.AlternativeSecurityIds
	s.ApproximateLastSignInDateTime = decoded.ApproximateLastSignInDateTime
	s.Commands = decoded.Commands
	s.ComplianceExpirationDateTime = decoded.ComplianceExpirationDateTime
	s.DeviceCategory = decoded.DeviceCategory
	s.DeviceId = decoded.DeviceId
	s.DeviceMetadata = decoded.DeviceMetadata
	s.DeviceOwnership = decoded.DeviceOwnership
	s.DeviceTemplate = decoded.DeviceTemplate
	s.DeviceVersion = decoded.DeviceVersion
	s.DisplayName = decoded.DisplayName
	s.DomainName = decoded.DomainName
	s.EnrollmentProfileName = decoded.EnrollmentProfileName
	s.EnrollmentType = decoded.EnrollmentType
	s.ExtensionAttributes = decoded.ExtensionAttributes
	s.Hostnames = decoded.Hostnames
	s.IsCompliant = decoded.IsCompliant
	s.IsManaged = decoded.IsManaged
	s.IsManagementRestricted = decoded.IsManagementRestricted
	s.IsRooted = decoded.IsRooted
	s.Kind = decoded.Kind
	s.ManagementType = decoded.ManagementType
	s.Manufacturer = decoded.Manufacturer
	s.MdmAppId = decoded.MdmAppId
	s.MemberOf_ODataBind = decoded.MemberOf_ODataBind
	s.Model = decoded.Model
	s.Name = decoded.Name
	s.OnPremisesLastSyncDateTime = decoded.OnPremisesLastSyncDateTime
	s.OnPremisesSecurityIdentifier = decoded.OnPremisesSecurityIdentifier
	s.OnPremisesSyncEnabled = decoded.OnPremisesSyncEnabled
	s.OperatingSystem = decoded.OperatingSystem
	s.OperatingSystemVersion = decoded.OperatingSystemVersion
	s.PhysicalIds = decoded.PhysicalIds
	s.Platform = decoded.Platform
	s.ProfileType = decoded.ProfileType
	s.RegisteredOwners_ODataBind = decoded.RegisteredOwners_ODataBind
	s.RegisteredUsers_ODataBind = decoded.RegisteredUsers_ODataBind
	s.RegistrationDateTime = decoded.RegistrationDateTime
	s.Status = decoded.Status
	s.SystemLabels = decoded.SystemLabels
	s.TransitiveMemberOf_ODataBind = decoded.TransitiveMemberOf_ODataBind
	s.TrustType = decoded.TrustType
	s.UsageRights = decoded.UsageRights
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Device into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["extensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Extensions into list []json.RawMessage: %+v", err)
		}

		output := make([]Extension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Extensions' for 'Device': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Extensions = &output
	}

	if v, ok := temp["memberOf"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MemberOf into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MemberOf' for 'Device': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MemberOf = &output
	}

	if v, ok := temp["registeredOwners"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RegisteredOwners into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RegisteredOwners' for 'Device': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RegisteredOwners = &output
	}

	if v, ok := temp["registeredUsers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RegisteredUsers into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RegisteredUsers' for 'Device': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RegisteredUsers = &output
	}

	if v, ok := temp["transitiveMemberOf"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling TransitiveMemberOf into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'TransitiveMemberOf' for 'Device': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TransitiveMemberOf = &output
	}

	return nil
}
