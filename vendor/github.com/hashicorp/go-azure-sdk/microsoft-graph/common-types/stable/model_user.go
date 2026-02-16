package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = User{}

type User struct {
	// A freeform text entry field for the user to describe themselves. Returned only on $select.
	AboutMe nullable.Type[string] `json:"aboutMe,omitempty"`

	// true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on
	// $select. Supports $filter (eq, ne, not, and in).
	AccountEnabled nullable.Type[bool] `json:"accountEnabled,omitempty"`

	// The user's activities across devices. Read-only. Nullable.
	Activities *[]UserActivity `json:"activities,omitempty"`

	// Sets the age group of the user. Allowed values: null, Minor, NotAdult, and Adult. For more information, see legal age
	// group property definitions. Returned only on $select. Supports $filter (eq, ne, not, and in).
	AgeGroup nullable.Type[string] `json:"ageGroup,omitempty"`

	// The user's terms of use acceptance statuses. Read-only. Nullable.
	AgreementAcceptances *[]AgreementAcceptance `json:"agreementAcceptances,omitempty"`

	// Represents the app roles a user is granted for an application. Supports $expand.
	AppRoleAssignments *[]AppRoleAssignment `json:"appRoleAssignments,omitempty"`

	// The licenses that are assigned to the user, including inherited (group-based) licenses. This property doesn't
	// differentiate between directly assigned and inherited licenses. Use the licenseAssignmentStates property to identify
	// the directly assigned and inherited licenses. Not nullable. Returned only on $select. Supports $filter (eq, not,
	// /$count eq 0, /$count ne 0).
	AssignedLicenses *[]AssignedLicense `json:"assignedLicenses,omitempty"`

	// The plans that are assigned to the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq and
	// not).
	AssignedPlans *[]AssignedPlan `json:"assignedPlans,omitempty"`

	// The authentication methods that are supported for the user.
	Authentication *Authentication `json:"authentication,omitempty"`

	AuthorizationInfo *AuthorizationInfo `json:"authorizationInfo,omitempty"`

	// The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always
	// in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z. Returned only on $select.
	Birthday *string `json:"birthday,omitempty"`

	// The telephone numbers for the user. NOTE: Although it's a string collection, only one number can be set for this
	// property. Read-only for users synced from the on-premises directory. Returned by default. Supports $filter (eq, not,
	// ge, le, startsWith).
	BusinessPhones *[]string `json:"businessPhones,omitempty"`

	// The user's primary calendar. Read-only.
	Calendar *Calendar `json:"calendar,omitempty"`

	// The user's calendar groups. Read-only. Nullable.
	CalendarGroups *[]CalendarGroup `json:"calendarGroups,omitempty"`

	// The calendar view for the calendar. Read-only. Nullable.
	CalendarView *[]Event `json:"calendarView,omitempty"`

	// The user's calendars. Read-only. Nullable.
	Calendars *[]Calendar `json:"calendars,omitempty"`

	Chats *[]Chat `json:"chats,omitempty"`

	// The city where the user is located. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq,
	// ne, not, ge, le, in, startsWith, and eq on null values).
	City nullable.Type[string] `json:"city,omitempty"`

	CloudClipboard *CloudClipboardRoot `json:"cloudClipboard,omitempty"`

	// The name of the company that the user is associated with. This property can be useful for describing the company that
	// a guest comes from. The maximum length is 64 characters.Returned only on $select. Supports $filter (eq, ne, not, ge,
	// le, in, startsWith, and eq on null values).
	CompanyName nullable.Type[string] `json:"companyName,omitempty"`

	// Sets whether consent was obtained for minors. Allowed values: null, Granted, Denied, and NotRequired. For more
	// information, see legal age group property definitions. Returned only on $select. Supports $filter (eq, ne, not, and
	// in).
	ConsentProvidedForMinor nullable.Type[string] `json:"consentProvidedForMinor,omitempty"`

	// The user's contacts folders. Read-only. Nullable.
	ContactFolders *[]ContactFolder `json:"contactFolders,omitempty"`

	// The user's contacts. Read-only. Nullable.
	Contacts *[]Contact `json:"contacts,omitempty"`

	// The country or region where the user is located; for example, US or UK. Maximum length is 128 characters. Returned
	// only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	Country nullable.Type[string] `json:"country,omitempty"`

	// The date and time the user was created, in ISO 8601 format and UTC. The value can't be modified and is automatically
	// populated when the entity is created. Nullable. For on-premises users, the value represents when they were first
	// created in Microsoft Entra ID. Property is null for some users created before June 2018 and on-premises users that
	// were synced to Microsoft Entra ID before June 2018. Read-only. Returned only on $select. Supports $filter (eq, ne,
	// not , ge, le, in).
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Directory objects that the user created. Read-only. Nullable.
	CreatedObjects *[]DirectoryObject `json:"createdObjects,omitempty"`

	// List of OData IDs for `CreatedObjects` to bind to this entity
	CreatedObjects_ODataBind *[]string `json:"createdObjects@odata.bind,omitempty"`

	// Indicates whether the user account was created through one of the following methods: As a regular school or work
	// account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant
	// (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through
	// self-service sign-up by a guest signing up through a link that is part of a user flow (SelfServiceSignUp).
	// Read-only.Returned only on $select. Supports $filter (eq, ne, not, in).
	CreationType nullable.Type[string] `json:"creationType,omitempty"`

	// An open complex type that holds the value of a custom security attribute that is assigned to a directory object.
	// Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith). The filter value is case-sensitive.
	// To read this property, the calling app must be assigned the CustomSecAttributeAssignment.Read.All permission. To
	// write this property, the calling app must be assigned the CustomSecAttributeAssignment.ReadWrite.All permissions. To
	// read or write this property in delegated scenarios, the admin must be assigned the Attribute Assignment Administrator
	// role.
	CustomSecurityAttributes *CustomSecurityAttributeValue `json:"customSecurityAttributes,omitempty"`

	// The name of the department in which the user works. Maximum length is 64 characters. Returned only on $select.
	// Supports $filter (eq, ne, not , ge, le, in, and eq on null values).
	Department nullable.Type[string] `json:"department,omitempty"`

	// The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
	DeviceEnrollmentLimit *int64 `json:"deviceEnrollmentLimit,omitempty"`

	// The list of troubleshooting events for this user.
	DeviceManagementTroubleshootingEvents *[]DeviceManagementTroubleshootingEvent `json:"deviceManagementTroubleshootingEvents,omitempty"`

	// The users and contacts that report to the user. (The users and contacts that have their manager property set to this
	// user.) Read-only. Nullable. Supports $expand.
	DirectReports *[]DirectoryObject `json:"directReports,omitempty"`

	// List of OData IDs for `DirectReports` to bind to this entity
	DirectReports_ODataBind *[]string `json:"directReports@odata.bind,omitempty"`

	// The name displayed in the address book for the user. This value is usually the combination of the user's first name,
	// middle initial, and family name. This property is required when a user is created and it can't be cleared during
	// updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in,
	// startsWith, and eq on null values), $orderby, and $search.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The user's OneDrive. Read-only.
	Drive *Drive `json:"drive,omitempty"`

	// A collection of drives available for this user. Read-only.
	Drives *[]Drive `json:"drives,omitempty"`

	EmployeeExperience *EmployeeExperienceUser `json:"employeeExperience,omitempty"`

	// The date and time when the user was hired or will start work in a future hire. Returned only on $select. Supports
	// $filter (eq, ne, not , ge, le, in).
	EmployeeHireDate nullable.Type[string] `json:"employeeHireDate,omitempty"`

	// The employee identifier assigned to the user by the organization. The maximum length is 16 characters. Returned only
	// on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
	EmployeeId nullable.Type[string] `json:"employeeId,omitempty"`

	// The date and time when the user left or will leave the organization. To read this property, the calling app must be
	// assigned the User-LifeCycleInfo.Read.All permission. To write this property, the calling app must be assigned the
	// User.Read.All and User-LifeCycleInfo.ReadWrite.All permissions. To read this property in delegated scenarios, the
	// admin needs at least one of the following Microsoft Entra roles: Lifecycle Workflows Administrator (least privilege),
	// Global Reader. To write this property in delegated scenarios, the admin needs the Global Administrator role. Supports
	// $filter (eq, ne, not , ge, le, in). For more information, see Configure the employeeLeaveDateTime property for a
	// user.
	EmployeeLeaveDateTime nullable.Type[string] `json:"employeeLeaveDateTime,omitempty"`

	// Represents organization data (for example, division and costCenter) associated with a user. Returned only on $select.
	// Supports $filter (eq, ne, not , ge, le, in).
	EmployeeOrgData *EmployeeOrgData `json:"employeeOrgData,omitempty"`

	// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Returned only on $select.
	// Supports $filter (eq, ne, not , ge, le, in, startsWith).
	EmployeeType nullable.Type[string] `json:"employeeType,omitempty"`

	// The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
	Events *[]Event `json:"events,omitempty"`

	// The collection of open extensions defined for the user. Read-only. Supports $expand. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`

	// For a guest invited to the tenant using the invitation API, this property represents the invited user's invitation
	// status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only
	// on $select. Supports $filter (eq, ne, not , in).
	ExternalUserState nullable.Type[string] `json:"externalUserState,omitempty"`

	// Shows the timestamp for the latest change to the externalUserState property. Returned only on $select. Supports
	// $filter (eq, ne, not , in).
	ExternalUserStateChangeDateTime nullable.Type[string] `json:"externalUserStateChangeDateTime,omitempty"`

	// The fax number of the user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq
	// on null values).
	FaxNumber nullable.Type[string] `json:"faxNumber,omitempty"`

	FollowedSites *[]Site `json:"followedSites,omitempty"`

	// The given name (first name) of the user. Maximum length is 64 characters. Returned by default. Supports $filter (eq,
	// ne, not , ge, le, in, startsWith, and eq on null values).
	GivenName nullable.Type[string] `json:"givenName,omitempty"`

	// The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is
	// always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z. Returned only on $select. Note:
	// This property is specific to SharePoint in Microsoft 365. We recommend using the native employeeHireDate property to
	// set and update hire date values using Microsoft Graph APIs.
	HireDate *string `json:"hireDate,omitempty"`

	// Represents the identities that can be used to sign in to this user account. Microsoft (also known as a local
	// account), organizations, or social identity providers such as Facebook, Google, and Microsoft can provide identity
	// and tie it to a user account. It might contain multiple items with the same signInType value. Returned only on
	// $select. Supports $filter (eq) with limitations.
	Identities *[]ObjectIdentity `json:"identities,omitempty"`

	// The instant message voice-over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only.
	// Returned only on $select. Supports $filter (eq, not, ge, le, startsWith).
	ImAddresses *[]string `json:"imAddresses,omitempty"`

	// Relevance classification of the user's messages based on explicit designations that override inferred relevance or
	// importance.
	InferenceClassification *InferenceClassification `json:"inferenceClassification,omitempty"`

	// Represents relationships between a user and items such as OneDrive for work or school documents, calculated using
	// advanced analytics and machine learning techniques. Read-only. Nullable.
	Insights *ItemInsights `json:"insights,omitempty"`

	// A list for the user to describe their interests. Returned only on $select.
	Interests *[]string `json:"interests,omitempty"`

	IsManagementRestricted nullable.Type[bool] `json:"isManagementRestricted,omitempty"`

	// Don't use – reserved for future use.
	IsResourceAccount nullable.Type[bool] `json:"isResourceAccount,omitempty"`

	// The user's job title. Maximum length is 128 characters. Returned by default. Supports $filter (eq, ne, not , ge, le,
	// in, startsWith, and eq on null values).
	JobTitle nullable.Type[string] `json:"jobTitle,omitempty"`

	JoinedTeams *[]Team `json:"joinedTeams,omitempty"`

	// The time when this Microsoft Entra user last changed their password or when their password was created, whichever
	// date the latest action was performed. The date and time information uses ISO 8601 format and is always in UTC. For
	// example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
	LastPasswordChangeDateTime nullable.Type[string] `json:"lastPasswordChangeDateTime,omitempty"`

	// Used by enterprise applications to determine the legal age group of the user. This property is read-only and
	// calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, Undefined,
	// MinorWithOutParentalConsent, MinorWithParentalConsent, MinorNoParentalConsentRequired, NotAdult, and Adult. For more
	// information, see legal age group property definitions. Returned only on $select.
	LegalAgeGroupClassification nullable.Type[string] `json:"legalAgeGroupClassification,omitempty"`

	// State of license assignments for this user. Also indicates licenses that are directly assigned or the user inherited
	// through group memberships. Read-only. Returned only on $select.
	LicenseAssignmentStates *[]LicenseAssignmentState `json:"licenseAssignmentStates,omitempty"`

	// A collection of this user's license details. Read-only.
	LicenseDetails *[]LicenseDetails `json:"licenseDetails,omitempty"`

	// The SMTP address for the user, for example, jeff@contoso.com. Changes to this property update the user's
	// proxyAddresses collection to include the value as an SMTP address. This property can't contain accent characters.
	// NOTE: We don't recommend updating this property for Azure AD B2C user profiles. Use the otherMails property instead.
	// Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith, and eq on null values).
	Mail nullable.Type[string] `json:"mail,omitempty"`

	// The user's mail folders. Read-only. Nullable.
	MailFolders *[]MailFolder `json:"mailFolders,omitempty"`

	// The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters.
	// Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	MailNickname nullable.Type[string] `json:"mailNickname,omitempty"`

	// Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies
	// to incoming messages, locale, and time zone. Returned only on $select.
	MailboxSettings *MailboxSettings `json:"mailboxSettings,omitempty"`

	// Zero or more managed app registrations that belong to the user.
	ManagedAppRegistrations *[]ManagedAppRegistration `json:"managedAppRegistrations,omitempty"`

	// The managed devices associated with the user.
	ManagedDevices *[]ManagedDevice `json:"managedDevices,omitempty"`

	// The user or contact that is this user's manager. Read-only. Supports $expand.
	Manager *DirectoryObject `json:"manager,omitempty"`

	// OData ID for `Manager` to bind to this entity
	Manager_ODataBind *string `json:"manager@odata.bind,omitempty"`

	// The groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
	MemberOf *[]DirectoryObject `json:"memberOf,omitempty"`

	// List of OData IDs for `MemberOf` to bind to this entity
	MemberOf_ODataBind *[]string `json:"memberOf@odata.bind,omitempty"`

	// The messages in a mailbox or folder. Read-only. Nullable.
	Messages *[]Message `json:"messages,omitempty"`

	// The primary cellular telephone number for the user. Read-only for users synced from the on-premises directory.
	// Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq
	// on null values) and $search.
	MobilePhone nullable.Type[string] `json:"mobilePhone,omitempty"`

	// The URL for the user's site. Returned only on $select.
	MySite nullable.Type[string] `json:"mySite,omitempty"`

	OAuth2PermissionGrants *[]OAuth2PermissionGrant `json:"oauth2PermissionGrants,omitempty"`

	// The office location in the user's place of business. Returned by default. Supports $filter (eq, ne, not, ge, le, in,
	// startsWith, and eq on null values).
	OfficeLocation nullable.Type[string] `json:"officeLocation,omitempty"`

	// Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who
	// are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned
	// only on $select.
	OnPremisesDistinguishedName nullable.Type[string] `json:"onPremisesDistinguishedName,omitempty"`

	// Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The
	// property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via
	// Microsoft Entra Connect. Read-only. Returned only on $select.
	OnPremisesDomainName nullable.Type[string] `json:"onPremisesDomainName,omitempty"`

	// Contains extensionAttributes1-15 for the user. These extension attributes are also known as Exchange custom
	// attributes 1-15. Each attribute can store up to 1024 characters. For an onPremisesSyncEnabled user, the source of
	// authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where
	// onPremisesSyncEnabled is false), these properties can be set during the creation or update of a user object. For a
	// cloud-only user previously synced from on-premises Active Directory, these properties are read-only in Microsoft
	// Graph but can be fully managed through the Exchange Admin Center or the Exchange Online V2 module in PowerShell.
	// Returned only on $select. Supports $filter (eq, ne, not, in).
	OnPremisesExtensionAttributes *OnPremisesExtensionAttributes `json:"onPremisesExtensionAttributes,omitempty"`

	// This property is used to associate an on-premises Active Directory user account to their Microsoft Entra user object.
	// This property must be specified when creating a new user account in the Graph if you're using a federated domain for
	// the user's userPrincipalName (UPN) property. NOTE: The $ and _ characters can't be used when specifying this
	// property. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).
	OnPremisesImmutableId nullable.Type[string] `json:"onPremisesImmutableId,omitempty"`

	// Indicates the last time at which the object was synced with the on-premises directory; for example:
	// 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in
	// UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports
	// $filter (eq, ne, not, ge, le, in).
	OnPremisesLastSyncDateTime nullable.Type[string] `json:"onPremisesLastSyncDateTime,omitempty"`

	// Errors when using Microsoft synchronization product during provisioning. Returned only on $select. Supports $filter
	// (eq, not, ge, le).
	OnPremisesProvisioningErrors *[]OnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`

	// Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated
	// for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect.
	// Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
	OnPremisesSamAccountName nullable.Type[string] `json:"onPremisesSamAccountName,omitempty"`

	// Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud.
	// Read-only. Returned only on $select. Supports $filter (eq including on null values).
	OnPremisesSecurityIdentifier nullable.Type[string] `json:"onPremisesSecurityIdentifier,omitempty"`

	// true if this user object is currently being synced from an on-premises Active Directory (AD); otherwise the user
	// isn't being synced and can be managed in Microsoft Entra ID. Read-only. Returned only on $select. Supports $filter
	// (eq, ne, not, in, and eq on null values).
	OnPremisesSyncEnabled nullable.Type[bool] `json:"onPremisesSyncEnabled,omitempty"`

	// Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only
	// populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra
	// Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
	OnPremisesUserPrincipalName nullable.Type[string] `json:"onPremisesUserPrincipalName,omitempty"`

	Onenote *Onenote `json:"onenote,omitempty"`

	// Information about a meeting, including the URL used to join a meeting, the attendees list, and the description.
	OnlineMeetings *[]OnlineMeeting `json:"onlineMeetings,omitempty"`

	// A list of other email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com']. Can store up
	// to 250 values, each with a limit of 250 characters. NOTE: This property can't contain accent characters. Returned
	// only on $select. Supports $filter (eq, not, ge, le, in, startsWith, endsWith, /$count eq 0, /$count ne 0).
	OtherMails *[]string `json:"otherMails,omitempty"`

	Outlook *OutlookUser `json:"outlook,omitempty"`

	// Devices the user owns. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0, /$count eq 1,
	// /$count ne 1).
	OwnedDevices *[]DirectoryObject `json:"ownedDevices,omitempty"`

	// List of OData IDs for `OwnedDevices` to bind to this entity
	OwnedDevices_ODataBind *[]string `json:"ownedDevices@odata.bind,omitempty"`

	// Directory objects the user owns. Read-only. Nullable. Supports $expand, $select nested in $expand, and $filter
	// (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
	OwnedObjects *[]DirectoryObject `json:"ownedObjects,omitempty"`

	// List of OData IDs for `OwnedObjects` to bind to this entity
	OwnedObjects_ODataBind *[]string `json:"ownedObjects@odata.bind,omitempty"`

	// Specifies password policies for the user. This value is an enumeration with one possible value being
	// DisableStrongPassword, which allows weaker passwords than the default policy to be specified.
	// DisablePasswordExpiration can also be specified. The two might be specified together; for example:
	// DisablePasswordExpiration, DisableStrongPassword. Returned only on $select. For more information on the default
	// password policies, see Microsoft Entra password policies. Supports $filter (ne, not, and eq on null values).
	PasswordPolicies nullable.Type[string] `json:"passwordPolicies,omitempty"`

	// Specifies the password profile for the user. The profile contains the user's password. This property is required when
	// a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies
	// property. By default, a strong password is required. Returned only on $select. Supports $filter (eq, ne, not, in, and
	// eq on null values). To update this property: User-PasswordProfile.ReadWrite.All is the least privileged permission to
	// update this property. In delegated scenarios, the User Administrator Microsoft Entra role is the least privileged
	// admin role supported to update this property for nonadmin users. Privileged Authentication Administrator is the least
	// privileged role that's allowed to update this property for all administrators in the tenant. In general, the
	// signed-in user must have a higher privileged administrator role as indicated in Who can reset passwords. In app-only
	// scenarios, the calling app must be assigned a supported permission and at least the User Administrator Microsoft
	// Entra role.
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`

	// A list for the user to enumerate their past projects. Returned only on $select.
	PastProjects *[]string `json:"pastProjects,omitempty"`

	// People that are relevant to the user. Read-only. Nullable.
	People *[]Person `json:"people,omitempty"`

	// List all resource-specific permission grants of a user.
	PermissionGrants *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`

	// The user's profile photo. Read-only.
	Photo *ProfilePhoto `json:"photo,omitempty"`

	// The collection of the user's profile photos in different sizes. Read-only.
	Photos *[]ProfilePhoto `json:"photos,omitempty"`

	// Entry-point to the Planner resource that might exist for a user. Read-only.
	Planner *PlannerUser `json:"planner,omitempty"`

	// The postal code for the user's postal address. The postal code is specific to the user's country or region. In the
	// United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Returned only on
	// $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	PostalCode nullable.Type[string] `json:"postalCode,omitempty"`

	// The preferred data location for the user. For more information, see OneDrive Online Multi-Geo.
	PreferredDataLocation nullable.Type[string] `json:"preferredDataLocation,omitempty"`

	// The preferred language for the user. The preferred language format is based on RFC 4646. The name is a combination of
	// an ISO 639 two-letter lowercase culture code associated with the language, and an ISO 3166 two-letter uppercase
	// subculture code associated with the country or region. Example: 'en-US', or 'es-ES'. Returned by default. Supports
	// $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values)
	PreferredLanguage nullable.Type[string] `json:"preferredLanguage,omitempty"`

	// The preferred name for the user. Not Supported. This attribute returns an empty string.Returned only on $select.
	PreferredName nullable.Type[string] `json:"preferredName,omitempty"`

	Presence *Presence  `json:"presence,omitempty"`
	Print    *UserPrint `json:"print,omitempty"`

	// The plans that are provisioned for the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq,
	// not, ge, le).
	ProvisionedPlans *[]ProvisionedPlan `json:"provisionedPlans,omitempty"`

	// For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. Changes to the mail property update this
	// collection to include the value as an SMTP address. For more information, see mail and proxyAddresses properties. The
	// proxy address prefixed with SMTP (capitalized) is the primary proxy address, while those addresses prefixed with smtp
	// are the secondary proxy addresses. For Azure AD B2C accounts, this property has a limit of 10 unique addresses.
	// Read-only in Microsoft Graph; you can update this property only through the Microsoft 365 admin center. Not nullable.
	// Returned only on $select. Supports $filter (eq, not, ge, le, startsWith, endsWith, /$count eq 0, /$count ne 0).
	ProxyAddresses *[]string `json:"proxyAddresses,omitempty"`

	// Devices that are registered for the user. Read-only. Nullable. Supports $expand and returns up to 100 objects.
	RegisteredDevices *[]DirectoryObject `json:"registeredDevices,omitempty"`

	// List of OData IDs for `RegisteredDevices` to bind to this entity
	RegisteredDevices_ODataBind *[]string `json:"registeredDevices@odata.bind,omitempty"`

	// A list for the user to enumerate their responsibilities. Returned only on $select.
	Responsibilities *[]string `json:"responsibilities,omitempty"`

	// A list for the user to enumerate the schools they attended. Returned only on $select.
	Schools *[]string `json:"schools,omitempty"`

	ScopedRoleMemberOf *[]ScopedRoleMembership `json:"scopedRoleMemberOf,omitempty"`

	// Security identifier (SID) of the user, used in Windows scenarios. Read-only. Returned by default. Supports $select
	// and $filter (eq, not, ge, le, startsWith).
	SecurityIdentifier nullable.Type[string] `json:"securityIdentifier,omitempty"`

	// Errors published by a federated service describing a nontransient, service-specific error regarding the properties or
	// link from a user object. Supports $filter (eq, not, for isResolved and serviceInstance).
	ServiceProvisioningErrors *[]ServiceProvisioningError `json:"serviceProvisioningErrors,omitempty"`

	Settings *UserSettings `json:"settings,omitempty"`

	// Do not use in Microsoft Graph. Manage this property through the Microsoft 365 admin center instead. Represents
	// whether the user should be included in the Outlook global address list. See Known issue.
	ShowInAddressList nullable.Type[bool] `json:"showInAddressList,omitempty"`

	// Get the last signed-in date and request ID of the sign-in for a given user. Read-only.Returned only on $select.
	// Supports $filter (eq, ne, not, ge, le) but not with any other filterable properties. Note: Details for this property
	// require a Microsoft Entra ID P1 or P2 license and the AuditLog.Read.All permission.This property isn't returned for a
	// user who never signed in or last signed in before April 2020.
	SignInActivity *SignInActivity `json:"signInActivity,omitempty"`

	// Any refresh tokens or session tokens (session cookies) issued before this time are invalid. Applications get an error
	// when using an invalid refresh or session token to acquire a delegated access token (to access APIs such as Microsoft
	// Graph). If this happens, the application needs to acquire a new refresh token by requesting the authorized endpoint.
	// Read-only. Use revokeSignInSessions to reset. Returned only on $select.
	SignInSessionsValidFromDateTime nullable.Type[string] `json:"signInSessionsValidFromDateTime,omitempty"`

	// A list for the user to enumerate their skills. Returned only on $select.
	Skills *[]string `json:"skills,omitempty"`

	// The identifier that relates the user to the working time schedule triggers. Read-Only. Nullable
	Solutions *UserSolutionRoot `json:"solutions,omitempty"`

	// The users and groups responsible for this guest's privileges in the tenant and keeping the guest's information and
	// access updated. (HTTP Methods: GET, POST, DELETE.). Supports $expand.
	Sponsors *[]DirectoryObject `json:"sponsors,omitempty"`

	// List of OData IDs for `Sponsors` to bind to this entity
	Sponsors_ODataBind *[]string `json:"sponsors@odata.bind,omitempty"`

	// The state or province in the user's address. Maximum length is 128 characters. Returned only on $select. Supports
	// $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	State nullable.Type[string] `json:"state,omitempty"`

	// The street address of the user's place of business. Maximum length is 1,024 characters. Returned only on $select.
	// Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	StreetAddress nullable.Type[string] `json:"streetAddress,omitempty"`

	// The user's surname (family name or last name). Maximum length is 64 characters. Returned by default. Supports $filter
	// (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	Surname nullable.Type[string] `json:"surname,omitempty"`

	// A container for Microsoft Teams features available for the user. Read-only. Nullable.
	Teamwork *UserTeamwork `json:"teamwork,omitempty"`

	// Represents the To Do services available to a user.
	Todo *Todo `json:"todo,omitempty"`

	// The groups, including nested groups, and directory roles that a user is a member of. Nullable.
	TransitiveMemberOf *[]DirectoryObject `json:"transitiveMemberOf,omitempty"`

	// List of OData IDs for `TransitiveMemberOf` to bind to this entity
	TransitiveMemberOf_ODataBind *[]string `json:"transitiveMemberOf@odata.bind,omitempty"`

	// A two-letter country code (ISO standard 3166). Required for users that are assigned licenses due to legal
	// requirements to check for availability of services in countries/regions. Examples include: US, JP, and GB. Not
	// nullable. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	UsageLocation nullable.Type[string] `json:"usageLocation,omitempty"`

	// The user principal name (UPN) of the user. The UPN is an Internet-style sign-in name for the user based on the
	// Internet standard RFC 822. By convention, this value should map to the user's email name. The general format is
	// alias@domain, where the domain must be present in the tenant's collection of verified domains. This property is
	// required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains
	// property of organization.NOTE: This property can't contain accent characters. Only the following characters are
	// allowed A - Z, a - z, 0 - 9, ' . - _ ! # ^ ~. For the complete list of allowed characters, see username policies.
	// Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith) and $orderby.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// A string value that can be used to classify user types in your directory. The possible values are Member and Guest.
	// Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). NOTE: For more information about
	// the permissions for members and guests, see What are the default user permissions in Microsoft Entra ID?
	UserType nullable.Type[string] `json:"userType,omitempty"`

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

func (s User) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s User) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = User{}

func (s User) MarshalJSON() ([]byte, error) {
	type wrapper User
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling User: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling User: %+v", err)
	}

	delete(decoded, "activities")
	delete(decoded, "agreementAcceptances")
	delete(decoded, "assignedPlans")
	delete(decoded, "calendar")
	delete(decoded, "calendarGroups")
	delete(decoded, "calendarView")
	delete(decoded, "calendars")
	delete(decoded, "contactFolders")
	delete(decoded, "contacts")
	delete(decoded, "createdDateTime")
	delete(decoded, "createdObjects")
	delete(decoded, "creationType")
	delete(decoded, "directReports")
	delete(decoded, "drive")
	delete(decoded, "drives")
	delete(decoded, "events")
	delete(decoded, "extensions")
	delete(decoded, "imAddresses")
	delete(decoded, "insights")
	delete(decoded, "licenseAssignmentStates")
	delete(decoded, "licenseDetails")
	delete(decoded, "mailFolders")
	delete(decoded, "manager")
	delete(decoded, "memberOf")
	delete(decoded, "messages")
	delete(decoded, "onPremisesDistinguishedName")
	delete(decoded, "onPremisesDomainName")
	delete(decoded, "onPremisesLastSyncDateTime")
	delete(decoded, "onPremisesSamAccountName")
	delete(decoded, "onPremisesSecurityIdentifier")
	delete(decoded, "onPremisesSyncEnabled")
	delete(decoded, "onPremisesUserPrincipalName")
	delete(decoded, "ownedDevices")
	delete(decoded, "ownedObjects")
	delete(decoded, "people")
	delete(decoded, "photo")
	delete(decoded, "photos")
	delete(decoded, "planner")
	delete(decoded, "provisionedPlans")
	delete(decoded, "registeredDevices")
	delete(decoded, "securityIdentifier")
	delete(decoded, "signInActivity")
	delete(decoded, "signInSessionsValidFromDateTime")
	delete(decoded, "teamwork")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.user"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling User: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &User{}

func (s *User) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AboutMe                         nullable.Type[string]              `json:"aboutMe,omitempty"`
		AccountEnabled                  nullable.Type[bool]                `json:"accountEnabled,omitempty"`
		Activities                      *[]UserActivity                    `json:"activities,omitempty"`
		AgeGroup                        nullable.Type[string]              `json:"ageGroup,omitempty"`
		AgreementAcceptances            *[]AgreementAcceptance             `json:"agreementAcceptances,omitempty"`
		AppRoleAssignments              *[]AppRoleAssignment               `json:"appRoleAssignments,omitempty"`
		AssignedLicenses                *[]AssignedLicense                 `json:"assignedLicenses,omitempty"`
		AssignedPlans                   *[]AssignedPlan                    `json:"assignedPlans,omitempty"`
		Authentication                  *Authentication                    `json:"authentication,omitempty"`
		AuthorizationInfo               *AuthorizationInfo                 `json:"authorizationInfo,omitempty"`
		Birthday                        *string                            `json:"birthday,omitempty"`
		BusinessPhones                  *[]string                          `json:"businessPhones,omitempty"`
		Calendar                        *Calendar                          `json:"calendar,omitempty"`
		CalendarGroups                  *[]CalendarGroup                   `json:"calendarGroups,omitempty"`
		CalendarView                    *[]Event                           `json:"calendarView,omitempty"`
		Calendars                       *[]Calendar                        `json:"calendars,omitempty"`
		Chats                           *[]Chat                            `json:"chats,omitempty"`
		City                            nullable.Type[string]              `json:"city,omitempty"`
		CloudClipboard                  *CloudClipboardRoot                `json:"cloudClipboard,omitempty"`
		CompanyName                     nullable.Type[string]              `json:"companyName,omitempty"`
		ConsentProvidedForMinor         nullable.Type[string]              `json:"consentProvidedForMinor,omitempty"`
		ContactFolders                  *[]ContactFolder                   `json:"contactFolders,omitempty"`
		Contacts                        *[]Contact                         `json:"contacts,omitempty"`
		Country                         nullable.Type[string]              `json:"country,omitempty"`
		CreatedDateTime                 nullable.Type[string]              `json:"createdDateTime,omitempty"`
		CreatedObjects_ODataBind        *[]string                          `json:"createdObjects@odata.bind,omitempty"`
		CreationType                    nullable.Type[string]              `json:"creationType,omitempty"`
		CustomSecurityAttributes        *CustomSecurityAttributeValue      `json:"customSecurityAttributes,omitempty"`
		Department                      nullable.Type[string]              `json:"department,omitempty"`
		DeviceEnrollmentLimit           *int64                             `json:"deviceEnrollmentLimit,omitempty"`
		DirectReports_ODataBind         *[]string                          `json:"directReports@odata.bind,omitempty"`
		DisplayName                     nullable.Type[string]              `json:"displayName,omitempty"`
		Drive                           *Drive                             `json:"drive,omitempty"`
		Drives                          *[]Drive                           `json:"drives,omitempty"`
		EmployeeExperience              *EmployeeExperienceUser            `json:"employeeExperience,omitempty"`
		EmployeeHireDate                nullable.Type[string]              `json:"employeeHireDate,omitempty"`
		EmployeeId                      nullable.Type[string]              `json:"employeeId,omitempty"`
		EmployeeLeaveDateTime           nullable.Type[string]              `json:"employeeLeaveDateTime,omitempty"`
		EmployeeOrgData                 *EmployeeOrgData                   `json:"employeeOrgData,omitempty"`
		EmployeeType                    nullable.Type[string]              `json:"employeeType,omitempty"`
		Events                          *[]Event                           `json:"events,omitempty"`
		ExternalUserState               nullable.Type[string]              `json:"externalUserState,omitempty"`
		ExternalUserStateChangeDateTime nullable.Type[string]              `json:"externalUserStateChangeDateTime,omitempty"`
		FaxNumber                       nullable.Type[string]              `json:"faxNumber,omitempty"`
		FollowedSites                   *[]Site                            `json:"followedSites,omitempty"`
		GivenName                       nullable.Type[string]              `json:"givenName,omitempty"`
		HireDate                        *string                            `json:"hireDate,omitempty"`
		Identities                      *[]ObjectIdentity                  `json:"identities,omitempty"`
		ImAddresses                     *[]string                          `json:"imAddresses,omitempty"`
		InferenceClassification         *InferenceClassification           `json:"inferenceClassification,omitempty"`
		Insights                        *ItemInsights                      `json:"insights,omitempty"`
		Interests                       *[]string                          `json:"interests,omitempty"`
		IsManagementRestricted          nullable.Type[bool]                `json:"isManagementRestricted,omitempty"`
		IsResourceAccount               nullable.Type[bool]                `json:"isResourceAccount,omitempty"`
		JobTitle                        nullable.Type[string]              `json:"jobTitle,omitempty"`
		JoinedTeams                     *[]Team                            `json:"joinedTeams,omitempty"`
		LastPasswordChangeDateTime      nullable.Type[string]              `json:"lastPasswordChangeDateTime,omitempty"`
		LegalAgeGroupClassification     nullable.Type[string]              `json:"legalAgeGroupClassification,omitempty"`
		LicenseAssignmentStates         *[]LicenseAssignmentState          `json:"licenseAssignmentStates,omitempty"`
		LicenseDetails                  *[]LicenseDetails                  `json:"licenseDetails,omitempty"`
		Mail                            nullable.Type[string]              `json:"mail,omitempty"`
		MailNickname                    nullable.Type[string]              `json:"mailNickname,omitempty"`
		MailboxSettings                 *MailboxSettings                   `json:"mailboxSettings,omitempty"`
		ManagedDevices                  *[]ManagedDevice                   `json:"managedDevices,omitempty"`
		Manager_ODataBind               *string                            `json:"manager@odata.bind,omitempty"`
		MemberOf_ODataBind              *[]string                          `json:"memberOf@odata.bind,omitempty"`
		MobilePhone                     nullable.Type[string]              `json:"mobilePhone,omitempty"`
		MySite                          nullable.Type[string]              `json:"mySite,omitempty"`
		OAuth2PermissionGrants          *[]OAuth2PermissionGrant           `json:"oauth2PermissionGrants,omitempty"`
		OfficeLocation                  nullable.Type[string]              `json:"officeLocation,omitempty"`
		OnPremisesDistinguishedName     nullable.Type[string]              `json:"onPremisesDistinguishedName,omitempty"`
		OnPremisesDomainName            nullable.Type[string]              `json:"onPremisesDomainName,omitempty"`
		OnPremisesExtensionAttributes   *OnPremisesExtensionAttributes     `json:"onPremisesExtensionAttributes,omitempty"`
		OnPremisesImmutableId           nullable.Type[string]              `json:"onPremisesImmutableId,omitempty"`
		OnPremisesLastSyncDateTime      nullable.Type[string]              `json:"onPremisesLastSyncDateTime,omitempty"`
		OnPremisesProvisioningErrors    *[]OnPremisesProvisioningError     `json:"onPremisesProvisioningErrors,omitempty"`
		OnPremisesSamAccountName        nullable.Type[string]              `json:"onPremisesSamAccountName,omitempty"`
		OnPremisesSecurityIdentifier    nullable.Type[string]              `json:"onPremisesSecurityIdentifier,omitempty"`
		OnPremisesSyncEnabled           nullable.Type[bool]                `json:"onPremisesSyncEnabled,omitempty"`
		OnPremisesUserPrincipalName     nullable.Type[string]              `json:"onPremisesUserPrincipalName,omitempty"`
		Onenote                         *Onenote                           `json:"onenote,omitempty"`
		OnlineMeetings                  *[]OnlineMeeting                   `json:"onlineMeetings,omitempty"`
		OtherMails                      *[]string                          `json:"otherMails,omitempty"`
		Outlook                         *OutlookUser                       `json:"outlook,omitempty"`
		OwnedDevices_ODataBind          *[]string                          `json:"ownedDevices@odata.bind,omitempty"`
		OwnedObjects_ODataBind          *[]string                          `json:"ownedObjects@odata.bind,omitempty"`
		PasswordPolicies                nullable.Type[string]              `json:"passwordPolicies,omitempty"`
		PasswordProfile                 *PasswordProfile                   `json:"passwordProfile,omitempty"`
		PastProjects                    *[]string                          `json:"pastProjects,omitempty"`
		People                          *[]Person                          `json:"people,omitempty"`
		PermissionGrants                *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`
		Photo                           *ProfilePhoto                      `json:"photo,omitempty"`
		Photos                          *[]ProfilePhoto                    `json:"photos,omitempty"`
		Planner                         *PlannerUser                       `json:"planner,omitempty"`
		PostalCode                      nullable.Type[string]              `json:"postalCode,omitempty"`
		PreferredDataLocation           nullable.Type[string]              `json:"preferredDataLocation,omitempty"`
		PreferredLanguage               nullable.Type[string]              `json:"preferredLanguage,omitempty"`
		PreferredName                   nullable.Type[string]              `json:"preferredName,omitempty"`
		Presence                        *Presence                          `json:"presence,omitempty"`
		Print                           *UserPrint                         `json:"print,omitempty"`
		ProvisionedPlans                *[]ProvisionedPlan                 `json:"provisionedPlans,omitempty"`
		ProxyAddresses                  *[]string                          `json:"proxyAddresses,omitempty"`
		RegisteredDevices_ODataBind     *[]string                          `json:"registeredDevices@odata.bind,omitempty"`
		Responsibilities                *[]string                          `json:"responsibilities,omitempty"`
		Schools                         *[]string                          `json:"schools,omitempty"`
		ScopedRoleMemberOf              *[]ScopedRoleMembership            `json:"scopedRoleMemberOf,omitempty"`
		SecurityIdentifier              nullable.Type[string]              `json:"securityIdentifier,omitempty"`
		Settings                        *UserSettings                      `json:"settings,omitempty"`
		ShowInAddressList               nullable.Type[bool]                `json:"showInAddressList,omitempty"`
		SignInActivity                  *SignInActivity                    `json:"signInActivity,omitempty"`
		SignInSessionsValidFromDateTime nullable.Type[string]              `json:"signInSessionsValidFromDateTime,omitempty"`
		Skills                          *[]string                          `json:"skills,omitempty"`
		Solutions                       *UserSolutionRoot                  `json:"solutions,omitempty"`
		Sponsors_ODataBind              *[]string                          `json:"sponsors@odata.bind,omitempty"`
		State                           nullable.Type[string]              `json:"state,omitempty"`
		StreetAddress                   nullable.Type[string]              `json:"streetAddress,omitempty"`
		Surname                         nullable.Type[string]              `json:"surname,omitempty"`
		Teamwork                        *UserTeamwork                      `json:"teamwork,omitempty"`
		Todo                            *Todo                              `json:"todo,omitempty"`
		TransitiveMemberOf_ODataBind    *[]string                          `json:"transitiveMemberOf@odata.bind,omitempty"`
		UsageLocation                   nullable.Type[string]              `json:"usageLocation,omitempty"`
		UserPrincipalName               nullable.Type[string]              `json:"userPrincipalName,omitempty"`
		UserType                        nullable.Type[string]              `json:"userType,omitempty"`
		DeletedDateTime                 nullable.Type[string]              `json:"deletedDateTime,omitempty"`
		Id                              *string                            `json:"id,omitempty"`
		ODataId                         *string                            `json:"@odata.id,omitempty"`
		ODataType                       *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AboutMe = decoded.AboutMe
	s.AccountEnabled = decoded.AccountEnabled
	s.Activities = decoded.Activities
	s.AgeGroup = decoded.AgeGroup
	s.AgreementAcceptances = decoded.AgreementAcceptances
	s.AppRoleAssignments = decoded.AppRoleAssignments
	s.AssignedLicenses = decoded.AssignedLicenses
	s.AssignedPlans = decoded.AssignedPlans
	s.Authentication = decoded.Authentication
	s.AuthorizationInfo = decoded.AuthorizationInfo
	s.Birthday = decoded.Birthday
	s.BusinessPhones = decoded.BusinessPhones
	s.Calendar = decoded.Calendar
	s.CalendarGroups = decoded.CalendarGroups
	s.CalendarView = decoded.CalendarView
	s.Calendars = decoded.Calendars
	s.Chats = decoded.Chats
	s.City = decoded.City
	s.CloudClipboard = decoded.CloudClipboard
	s.CompanyName = decoded.CompanyName
	s.ConsentProvidedForMinor = decoded.ConsentProvidedForMinor
	s.ContactFolders = decoded.ContactFolders
	s.Contacts = decoded.Contacts
	s.Country = decoded.Country
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CreatedObjects_ODataBind = decoded.CreatedObjects_ODataBind
	s.CreationType = decoded.CreationType
	s.CustomSecurityAttributes = decoded.CustomSecurityAttributes
	s.Department = decoded.Department
	s.DeviceEnrollmentLimit = decoded.DeviceEnrollmentLimit
	s.DirectReports_ODataBind = decoded.DirectReports_ODataBind
	s.DisplayName = decoded.DisplayName
	s.Drive = decoded.Drive
	s.Drives = decoded.Drives
	s.EmployeeExperience = decoded.EmployeeExperience
	s.EmployeeHireDate = decoded.EmployeeHireDate
	s.EmployeeId = decoded.EmployeeId
	s.EmployeeLeaveDateTime = decoded.EmployeeLeaveDateTime
	s.EmployeeOrgData = decoded.EmployeeOrgData
	s.EmployeeType = decoded.EmployeeType
	s.Events = decoded.Events
	s.ExternalUserState = decoded.ExternalUserState
	s.ExternalUserStateChangeDateTime = decoded.ExternalUserStateChangeDateTime
	s.FaxNumber = decoded.FaxNumber
	s.FollowedSites = decoded.FollowedSites
	s.GivenName = decoded.GivenName
	s.HireDate = decoded.HireDate
	s.Identities = decoded.Identities
	s.ImAddresses = decoded.ImAddresses
	s.InferenceClassification = decoded.InferenceClassification
	s.Insights = decoded.Insights
	s.Interests = decoded.Interests
	s.IsManagementRestricted = decoded.IsManagementRestricted
	s.IsResourceAccount = decoded.IsResourceAccount
	s.JobTitle = decoded.JobTitle
	s.JoinedTeams = decoded.JoinedTeams
	s.LastPasswordChangeDateTime = decoded.LastPasswordChangeDateTime
	s.LegalAgeGroupClassification = decoded.LegalAgeGroupClassification
	s.LicenseAssignmentStates = decoded.LicenseAssignmentStates
	s.LicenseDetails = decoded.LicenseDetails
	s.Mail = decoded.Mail
	s.MailNickname = decoded.MailNickname
	s.MailboxSettings = decoded.MailboxSettings
	s.ManagedDevices = decoded.ManagedDevices
	s.Manager_ODataBind = decoded.Manager_ODataBind
	s.MemberOf_ODataBind = decoded.MemberOf_ODataBind
	s.MobilePhone = decoded.MobilePhone
	s.MySite = decoded.MySite
	s.OAuth2PermissionGrants = decoded.OAuth2PermissionGrants
	s.OfficeLocation = decoded.OfficeLocation
	s.OnPremisesDistinguishedName = decoded.OnPremisesDistinguishedName
	s.OnPremisesDomainName = decoded.OnPremisesDomainName
	s.OnPremisesExtensionAttributes = decoded.OnPremisesExtensionAttributes
	s.OnPremisesImmutableId = decoded.OnPremisesImmutableId
	s.OnPremisesLastSyncDateTime = decoded.OnPremisesLastSyncDateTime
	s.OnPremisesProvisioningErrors = decoded.OnPremisesProvisioningErrors
	s.OnPremisesSamAccountName = decoded.OnPremisesSamAccountName
	s.OnPremisesSecurityIdentifier = decoded.OnPremisesSecurityIdentifier
	s.OnPremisesSyncEnabled = decoded.OnPremisesSyncEnabled
	s.OnPremisesUserPrincipalName = decoded.OnPremisesUserPrincipalName
	s.Onenote = decoded.Onenote
	s.OnlineMeetings = decoded.OnlineMeetings
	s.OtherMails = decoded.OtherMails
	s.Outlook = decoded.Outlook
	s.OwnedDevices_ODataBind = decoded.OwnedDevices_ODataBind
	s.OwnedObjects_ODataBind = decoded.OwnedObjects_ODataBind
	s.PasswordPolicies = decoded.PasswordPolicies
	s.PasswordProfile = decoded.PasswordProfile
	s.PastProjects = decoded.PastProjects
	s.People = decoded.People
	s.PermissionGrants = decoded.PermissionGrants
	s.Photo = decoded.Photo
	s.Photos = decoded.Photos
	s.Planner = decoded.Planner
	s.PostalCode = decoded.PostalCode
	s.PreferredDataLocation = decoded.PreferredDataLocation
	s.PreferredLanguage = decoded.PreferredLanguage
	s.PreferredName = decoded.PreferredName
	s.Presence = decoded.Presence
	s.Print = decoded.Print
	s.ProvisionedPlans = decoded.ProvisionedPlans
	s.ProxyAddresses = decoded.ProxyAddresses
	s.RegisteredDevices_ODataBind = decoded.RegisteredDevices_ODataBind
	s.Responsibilities = decoded.Responsibilities
	s.Schools = decoded.Schools
	s.ScopedRoleMemberOf = decoded.ScopedRoleMemberOf
	s.SecurityIdentifier = decoded.SecurityIdentifier
	s.Settings = decoded.Settings
	s.ShowInAddressList = decoded.ShowInAddressList
	s.SignInActivity = decoded.SignInActivity
	s.SignInSessionsValidFromDateTime = decoded.SignInSessionsValidFromDateTime
	s.Skills = decoded.Skills
	s.Solutions = decoded.Solutions
	s.Sponsors_ODataBind = decoded.Sponsors_ODataBind
	s.State = decoded.State
	s.StreetAddress = decoded.StreetAddress
	s.Surname = decoded.Surname
	s.Teamwork = decoded.Teamwork
	s.Todo = decoded.Todo
	s.TransitiveMemberOf_ODataBind = decoded.TransitiveMemberOf_ODataBind
	s.UsageLocation = decoded.UsageLocation
	s.UserPrincipalName = decoded.UserPrincipalName
	s.UserType = decoded.UserType
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling User into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdObjects"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CreatedObjects into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CreatedObjects' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CreatedObjects = &output
	}

	if v, ok := temp["deviceManagementTroubleshootingEvents"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DeviceManagementTroubleshootingEvents into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementTroubleshootingEvent, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementTroubleshootingEventImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DeviceManagementTroubleshootingEvents' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DeviceManagementTroubleshootingEvents = &output
	}

	if v, ok := temp["directReports"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DirectReports into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DirectReports' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DirectReports = &output
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
				return fmt.Errorf("unmarshaling index %d field 'Extensions' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Extensions = &output
	}

	if v, ok := temp["mailFolders"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MailFolders into list []json.RawMessage: %+v", err)
		}

		output := make([]MailFolder, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMailFolderImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MailFolders' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MailFolders = &output
	}

	if v, ok := temp["managedAppRegistrations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ManagedAppRegistrations into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedAppRegistration, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedAppRegistrationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ManagedAppRegistrations' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ManagedAppRegistrations = &output
	}

	if v, ok := temp["manager"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Manager' for 'User': %+v", err)
		}
		s.Manager = &impl
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
				return fmt.Errorf("unmarshaling index %d field 'MemberOf' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MemberOf = &output
	}

	if v, ok := temp["messages"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Messages into list []json.RawMessage: %+v", err)
		}

		output := make([]Message, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMessageImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Messages' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Messages = &output
	}

	if v, ok := temp["ownedDevices"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling OwnedDevices into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'OwnedDevices' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.OwnedDevices = &output
	}

	if v, ok := temp["ownedObjects"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling OwnedObjects into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'OwnedObjects' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.OwnedObjects = &output
	}

	if v, ok := temp["registeredDevices"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RegisteredDevices into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RegisteredDevices' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RegisteredDevices = &output
	}

	if v, ok := temp["serviceProvisioningErrors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ServiceProvisioningErrors into list []json.RawMessage: %+v", err)
		}

		output := make([]ServiceProvisioningError, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalServiceProvisioningErrorImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ServiceProvisioningErrors' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ServiceProvisioningErrors = &output
	}

	if v, ok := temp["sponsors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Sponsors into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Sponsors' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Sponsors = &output
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
				return fmt.Errorf("unmarshaling index %d field 'TransitiveMemberOf' for 'User': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TransitiveMemberOf = &output
	}

	return nil
}
