package msgraph

import (
	"encoding/json"
	goerrors "errors"

	"github.com/manicminer/hamilton/odata"
)

// StringNullWhenEmpty is a string type that marshals its JSON representation as null when set to its zero value.
// Can be used with a pointer reference with the `omitempty` tag to omit a field when the pointer is nil, but send a
// JSON null value when the string is empty.
type StringNullWhenEmpty string

func (s StringNullWhenEmpty) MarshalJSON() ([]byte, error) {
	if s == "" {
		return []byte("null"), nil
	}
	return json.Marshal(string(s))
}

type AccessPackageCatalogStatus = string

const (
	AccessPackageCatalogStatusPublished AccessPackageCatalogStatus = "Published"
)

type AccessPackageCatalogType = string

const (
	AccessPackageCatalogTypeServiceDefault AccessPackageCatalogType = "ServiceDefault"
	AccessPackageCatalogTypeUserManaged    AccessPackageCatalogType = "UserManaged"
)

type AccessPackageResourceOriginSystem = string

const (
	AccessPackageResourceOriginSystemAadApplication   AccessPackageResourceOriginSystem = "AadApplication"
	AccessPackageResourceOriginSystemAadGroup         AccessPackageResourceOriginSystem = "AadGroup"
	AccessPackageResourceOriginSystemSharePointOnline AccessPackageResourceOriginSystem = "SharePointOnline"
)

type AccessPackageResourceRequestState = string

const (
	AccessPackageResourceRequestStateDelivered AccessPackageResourceRequestState = "Delivered"
)

type AccessPackageResourceRequestType = string

const (
	AccessPackageResourceRequestTypeAdminAdd     AccessPackageResourceRequestType = "AdminAdd"
	AccessPackageResourceRequestTypeAdmminRemove AccessPackageResourceRequestType = "AdminRemove"
)

type AccessPackageResourceType = string

const (
	AccessPackageResourceTypeApplication          AccessPackageResourceType = "Application"
	AccessPackageResourceTypeSharePointOnlineSite AccessPackageResourceType = "SharePoint Online Site"
)

type AgeGroup = StringNullWhenEmpty

const (
	AgeGroupNone     AgeGroup = ""
	AgeGroupAdult    AgeGroup = "Adult"
	AgeGroupMinor    AgeGroup = "Minor"
	AgeGroupNotAdult AgeGroup = "NotAdult"
)

type ApplicationExtensionDataType = string

const (
	ApplicationExtensionDataTypeBinary       ApplicationExtensionDataType = "Binary"
	ApplicationExtensionDataTypeBoolean      ApplicationExtensionDataType = "Boolean"
	ApplicationExtensionDataTypeDateTime     ApplicationExtensionDataType = "DateTime"
	ApplicationExtensionDataTypeInteger      ApplicationExtensionDataType = "Integer"
	ApplicationExtensionDataTypeLargeInteger ApplicationExtensionDataType = "LargeInteger"
	ApplicationExtensionDataTypeString       ApplicationExtensionDataType = "String"
)

type ApplicationExtensionTargetObject = string

const (
	ApplicationExtensionTargetObjectApplication  ApplicationExtensionTargetObject = "Application"
	ApplicationExtensionTargetObjectDevice       ApplicationExtensionTargetObject = "Device"
	ApplicationExtensionTargetObjectGroup        ApplicationExtensionTargetObject = "Group"
	ApplicationExtensionTargetObjectOrganization ApplicationExtensionTargetObject = "Organization"
	ApplicationExtensionTargetObjectUser         ApplicationExtensionTargetObject = "User"
)

type ApplicationTemplateCategory = string

const (
	ApplicationTemplateCategoryCollaboration      ApplicationTemplateCategory = "Collaboration"
	ApplicationTemplateCategoryBusinessManagement ApplicationTemplateCategory = "Business Management"
	ApplicationTemplateCategoryConsumer           ApplicationTemplateCategory = "Consumer"
	ApplicationTemplateCategoryContentManagement  ApplicationTemplateCategory = "Content management"
	ApplicationTemplateCategoryCRM                ApplicationTemplateCategory = "CRM"
	ApplicationTemplateCategoryDataServices       ApplicationTemplateCategory = "Data services"
	ApplicationTemplateCategoryDeveloperServices  ApplicationTemplateCategory = "Developer services"
	ApplicationTemplateCategoryECommerce          ApplicationTemplateCategory = "E-commerce"
	ApplicationTemplateCategoryEducation          ApplicationTemplateCategory = "Education"
	ApplicationTemplateCategoryERP                ApplicationTemplateCategory = "ERP"
	ApplicationTemplateCategoryFinance            ApplicationTemplateCategory = "Finance"
	ApplicationTemplateCategoryHealth             ApplicationTemplateCategory = "Health"
	ApplicationTemplateCategoryHumanResources     ApplicationTemplateCategory = "Human resources"
	ApplicationTemplateCategoryITInfrastructure   ApplicationTemplateCategory = "IT infrastructure"
	ApplicationTemplateCategoryMail               ApplicationTemplateCategory = "Mail"
	ApplicationTemplateCategoryManagement         ApplicationTemplateCategory = "Management"
	ApplicationTemplateCategoryMarketing          ApplicationTemplateCategory = "Marketing"
	ApplicationTemplateCategoryMedia              ApplicationTemplateCategory = "Media"
	ApplicationTemplateCategoryProductivity       ApplicationTemplateCategory = "Productivity"
	ApplicationTemplateCategoryProjectManagement  ApplicationTemplateCategory = "Project management"
	ApplicationTemplateCategoryTelecommunications ApplicationTemplateCategory = "Telecommunications"
	ApplicationTemplateCategoryTools              ApplicationTemplateCategory = "Tools"
	ApplicationTemplateCategoryTravel             ApplicationTemplateCategory = "Travel"
	ApplicationTemplateCategoryWebDesignHosting   ApplicationTemplateCategory = "Web design & hosting"
)

type AppRoleAllowedMemberType = string

const (
	AppRoleAllowedMemberTypeApplication AppRoleAllowedMemberType = "Application"
	AppRoleAllowedMemberTypeUser        AppRoleAllowedMemberType = "User"
)

type ApprovalMode = string

const (
	ApprovalModeNoApproval  ApprovalMode = "NoApproval"
	ApprovalModeSerial      ApprovalMode = "Serial"
	ApprovalModeSingleStage ApprovalMode = "SingleStage"
)

type AttestationLevel = string

const (
	AttestationLevelAttested    AttestationLevel = "attested"
	AttestationLevelNotAttested AttestationLevel = "notAttested"
)

type AuthenticationMethodFeature = string

const (
	AuthenticationMethodFeatureSsprRegistered      AuthenticationMethodFeature = "ssprRegistered"
	AuthenticationMethodFeatureSsprEnabled         AuthenticationMethodFeature = "ssprEnabled"
	AuthenticationMethodFeatureSsprCapable         AuthenticationMethodFeature = "ssprCapable"
	AuthenticationMethodFeaturePasswordlessCapable AuthenticationMethodFeature = "passwordlessCapable"
	AuthenticationMethodFeatureMfaCapable          AuthenticationMethodFeature = "mfaCapable"
)

type AuthenticationMethodKeyStrength = string

const (
	AuthenticationMethodKeyStrengthNormal  AuthenticationMethodKeyStrength = "normal"
	AuthenticationMethodKeyStrengthWeak    AuthenticationMethodKeyStrength = "weak"
	AuthenticationMethodKeyStrengthUnknown AuthenticationMethodKeyStrength = "unknown"
)

type AuthenticationPhoneType = string

const (
	AuthenticationPhoneTypeMobile          AuthenticationPhoneType = "mobile"
	AuthenticationPhoneTypeAlternateMobile AuthenticationPhoneType = "alternateMobile"
	AuthenticationPhoneTypeOffice          AuthenticationPhoneType = "office"
)

type BodyType = string

const (
	BodyTypeText BodyType = "text"
	BodyTypeHtml BodyType = "html"
)

type ConsentProvidedForMinor = StringNullWhenEmpty

const (
	ConsentProvidedForMinorNone        ConsentProvidedForMinor = ""
	ConsentProvidedForMinorDenied      ConsentProvidedForMinor = "Denied"
	ConsentProvidedForMinorGranted     ConsentProvidedForMinor = "Granted"
	ConsentProvidedForMinorNotRequired ConsentProvidedForMinor = "NotRequired"
)

type CredentialUsageSummaryPeriod = string

const (
	CredentialUsageSummaryPeriod30 CredentialUsageSummaryPeriod = "D30"
	CredentialUsageSummaryPeriod7  CredentialUsageSummaryPeriod = "D7"
	CredentialUsageSummaryPeriod1  CredentialUsageSummaryPeriod = "D1"
)

type ConditionalAccessPolicyState = string

const (
	ConditionalAccessPolicyStateEnabled                           ConditionalAccessPolicyState = "enabled"
	ConditionalAccessPolicyStateDisabled                          ConditionalAccessPolicyState = "disabled"
	ConditionalAccessPolicyStateEnabledForReportingButNotEnforced ConditionalAccessPolicyState = "enabledForReportingButNotEnforced"
)

type ExtensionSchemaTargetType = string

const (
	ExtensionSchemaTargetTypeAdministrativeUnit ExtensionSchemaTargetType = "AdministrativeUnit"
	ExtensionSchemaTargetTypeContact            ExtensionSchemaTargetType = "Contact"
	ExtensionSchemaTargetTypeDevice             ExtensionSchemaTargetType = "Device"
	ExtensionSchemaTargetTypeEvent              ExtensionSchemaTargetType = "Event"
	ExtensionSchemaTargetTypePost               ExtensionSchemaTargetType = "Post"
	ExtensionSchemaTargetTypeGroup              ExtensionSchemaTargetType = "Group"
	ExtensionSchemaTargetTypeMessage            ExtensionSchemaTargetType = "Message"
	ExtensionSchemaTargetTypeOrganization       ExtensionSchemaTargetType = "Organization"
	ExtensionSchemaTargetTypeUser               ExtensionSchemaTargetType = "User"
)

type ExtensionSchemaPropertyDataType = string

const (
	ExtensionSchemaPropertyDataBinary   ExtensionSchemaPropertyDataType = "Binary"
	ExtensionSchemaPropertyDataBoolean  ExtensionSchemaPropertyDataType = "Boolean"
	ExtensionSchemaPropertyDataDateTime ExtensionSchemaPropertyDataType = "DateTime"
	ExtensionSchemaPropertyDataInteger  ExtensionSchemaPropertyDataType = "Integer"
	ExtensionSchemaPropertyDataString   ExtensionSchemaPropertyDataType = "String"
)

type FeatureType = string

const (
	FeatureTypeRegistration       FeatureType = "registration"
	FeatureTypeReset              FeatureType = "reset"
	FeatureTypeUnknownFutureValue FeatureType = "unknownFutureValue"
)

type GroupType = string

const (
	GroupTypeUnified GroupType = "Unified"
)

type GroupMembershipClaim = string

const (
	GroupMembershipClaimAll              GroupMembershipClaim = "All"
	GroupMembershipClaimNone             GroupMembershipClaim = "None"
	GroupMembershipClaimApplicationGroup GroupMembershipClaim = "ApplicationGroup"
	GroupMembershipClaimDirectoryRole    GroupMembershipClaim = "DirectoryRole"
	GroupMembershipClaimSecurityGroup    GroupMembershipClaim = "SecurityGroup"
)

type GroupResourceBehaviorOption = string

const (
	GroupResourceBehaviorOptionAllowOnlyMembersToPost   GroupResourceBehaviorOption = "AllowOnlyMembersToPost"
	GroupResourceBehaviorOptionHideGroupInOutlook       GroupResourceBehaviorOption = "HideGroupInOutlook"
	GroupResourceBehaviorOptionSubscribeNewGroupMembers GroupResourceBehaviorOption = "SubscribeNewGroupMembers"
	GroupResourceBehaviorOptionWelcomeEmailDisabled     GroupResourceBehaviorOption = "WelcomeEmailDisabled"
)

type GroupResourceProvisioningOption = string

const (
	GroupResourceProvisioningOptionTeam GroupResourceProvisioningOption = "Team"
)

type GroupTheme = StringNullWhenEmpty

const (
	GroupThemeNone   GroupTheme = ""
	GroupThemeBlue   GroupTheme = "Blue"
	GroupThemeGreen  GroupTheme = "Green"
	GroupThemeOrange GroupTheme = "Orange"
	GroupThemePink   GroupTheme = "Pink"
	GroupThemePurple GroupTheme = "Purple"
	GroupThemeRed    GroupTheme = "Red"
	GroupThemeTeal   GroupTheme = "Teal"
)

type GroupVisibility = string

const (
	GroupVisibilityHiddenMembership GroupVisibility = "Hiddenmembership"
	GroupVisibilityPrivate          GroupVisibility = "Private"
	GroupVisibilityPublic           GroupVisibility = "Public"
)

type InvitedUserType = string

const (
	InvitedUserTypeGuest  InvitedUserType = "Guest"
	InvitedUserTypeMember InvitedUserType = "Member"
)

type KeyCredentialType = string

const (
	KeyCredentialTypeAsymmetricX509Cert  KeyCredentialType = "AsymmetricX509Cert"
	KeyCredentialTypeX509CertAndPassword KeyCredentialType = "X509CertAndPassword"
)

type KeyCredentialUsage = string

const (
	KeyCredentialUsageSign   KeyCredentialUsage = "Sign"
	KeyCredentialUsageVerify KeyCredentialUsage = "Verify"
)

type Members []DirectoryObject

func (o Members) MarshalJSON() ([]byte, error) {
	members := make([]odata.Id, len(o))
	for i, v := range o {
		if v.ODataId == nil {
			return nil, goerrors.New("marshaling Members: encountered DirectoryObject with nil ODataId")
		}
		members[i] = *v.ODataId
	}
	return json.Marshal(members)
}

func (o *Members) UnmarshalJSON(data []byte) error {
	var members []odata.Id
	if err := json.Unmarshal(data, &members); err != nil {
		return err
	}
	for _, v := range members {
		*o = append(*o, DirectoryObject{ODataId: &v})
	}
	return nil
}

type MethodUsabilityReason string

const (
	MethodUsabilityReasonEnabledByPolicy  MethodUsabilityReason = "enabledByPolicy"
	MethodUsabilityReasonDisabledByPolicy MethodUsabilityReason = "disabledByPolicy"
	MethodUsabilityReasonExpired          MethodUsabilityReason = "expired"
	MethodUsabilityReasonNotYetValid      MethodUsabilityReason = "notYetValid"
	MethodUsabilityReasonOneTimeUsed      MethodUsabilityReason = "oneTimeUsed"
)

type Owners []DirectoryObject

func (o Owners) MarshalJSON() ([]byte, error) {
	owners := make([]odata.Id, len(o))
	for i, v := range o {
		if v.ODataId == nil {
			return nil, goerrors.New("marshaling Owners: encountered DirectoryObject with nil ODataId")
		}
		owners[i] = *v.ODataId
	}
	return json.Marshal(owners)
}

func (o *Owners) UnmarshalJSON(data []byte) error {
	var owners []odata.Id
	if err := json.Unmarshal(data, &owners); err != nil {
		return err
	}
	for _, v := range owners {
		*o = append(*o, DirectoryObject{ODataId: &v})
	}
	return nil
}

type PermissionScopeType = string

const (
	PermissionScopeTypeAdmin PermissionScopeType = "Admin"
	PermissionScopeTypeUser  PermissionScopeType = "User"
)

type PreferredSingleSignOnMode = StringNullWhenEmpty

const (
	PreferredSingleSignOnModeNone         PreferredSingleSignOnMode = ""
	PreferredSingleSignOnModeNotSupported PreferredSingleSignOnMode = "notSupported"
	PreferredSingleSignOnModeOidc         PreferredSingleSignOnMode = "oidc"
	PreferredSingleSignOnModePassword     PreferredSingleSignOnMode = "password"
	PreferredSingleSignOnModeSaml         PreferredSingleSignOnMode = "saml"
)

type RegistrationAuthMethod = string

const (
	RegistrationAuthMethodEmail                RegistrationAuthMethod = "email"
	RegistrationAuthMethodMobilePhone          RegistrationAuthMethod = "mobilePhone"
	RegistrationAuthMethodOfficePhone          RegistrationAuthMethod = "officePhone"
	RegistrationAuthMethodSecurityQuestion     RegistrationAuthMethod = "securityQuestion"
	RegistrationAuthMethodAppNotification      RegistrationAuthMethod = "appNotification"
	RegistrationAuthMethodAppCode              RegistrationAuthMethod = "appCode"
	RegistrationAuthMethodAlternateMobilePhone RegistrationAuthMethod = "alternateMobilePhone"
	RegistrationAuthMethodFido                 RegistrationAuthMethod = "fido"
	RegistrationAuthMethodAppPassword          RegistrationAuthMethod = "appPassword"
	RegistrationAuthMethodUnknownFutureValue   RegistrationAuthMethod = "unknownFutureValue"
)

type RegistrationStatus = string

const (
	RegistrationStatusRegistered    RegistrationStatus = "registered"
	RegistrationStatusEnabled       RegistrationStatus = "enabled"
	RegistrationStatusCapable       RegistrationStatus = "capable"
	RegistrationStatusMfaRegistered RegistrationStatus = "mfaRegistered"
)

type RequestorSettingsScopeType = string

const (
	RequestorSettingsScopeTypeAllConfiguredConnectedOrganizationSubjects RequestorSettingsScopeType = "AllConfiguredConnectedOrganizationSubjects"
	RequestorSettingsScopeTypeAllExistingConnectedOrganizationSubjects   RequestorSettingsScopeType = "AllExistingConnectedOrganizationSubjects"
	RequestorSettingsScopeTypeAllExistingDirectoryMemberUsers            RequestorSettingsScopeType = "AllExistingDirectoryMemberUsers"
	RequestorSettingsScopeTypeAllExistingDirectorySubjects               RequestorSettingsScopeType = "AllExistingDirectorySubjects"
	RequestorSettingsScopeTypeAllExternalSubjects                        RequestorSettingsScopeType = "AllExternalSubjects"
	RequestorSettingsScopeTypeNoSubjects                                 RequestorSettingsScopeType = "NoSubjects"
	RequestorSettingsScopeTypeSpecificConnectedOrganizationSubjects      RequestorSettingsScopeType = "SpecificConnectedOrganizationSubjects"
	RequestorSettingsScopeTypeSpecificDirectorySubjects                  RequestorSettingsScopeType = "SpecificDirectorySubjects"
)

type ResourceAccessType = string

const (
	ResourceAccessTypeRole  ResourceAccessType = "Role"
	ResourceAccessTypeScope ResourceAccessType = "Scope"
)

type SchemaExtensionStatus = string

const (
	SchemaExtensionStatusInDevelopment SchemaExtensionStatus = "InDevelopment"
	SchemaExtensionStatusAvailable     SchemaExtensionStatus = "Available"
	SchemaExtensionStatusDeprecated    SchemaExtensionStatus = "Deprecated"
)

type SchemaExtensionProperties interface {
	UnmarshalJSON([]byte) error
}

type SchemaExtensionMap map[string]interface{}

func (m *SchemaExtensionMap) UnmarshalJSON(data []byte) error {
	type sem SchemaExtensionMap
	m2 := (*sem)(m)
	return json.Unmarshal(data, m2)
}

type SignInAudience = string

const (
	SignInAudienceAzureADMyOrg                       SignInAudience = "AzureADMyOrg"
	SignInAudienceAzureADMultipleOrgs                SignInAudience = "AzureADMultipleOrgs"
	SignInAudienceAzureADandPersonalMicrosoftAccount SignInAudience = "AzureADandPersonalMicrosoftAccount"
	SignInAudiencePersonalMicrosoftAccount           SignInAudience = "PersonalMicrosoftAccount"
)

type UsageAuthMethod = string

const (
	UsageAuthMethodEmail                 UsageAuthMethod = "email"
	UsageAuthMethodMobileSMS             UsageAuthMethod = "mobileSMS"
	UsageAuthMethodMobileCall            UsageAuthMethod = "mobileCall"
	UsageAuthMethodOfficePhone           UsageAuthMethod = "officePhone"
	UsageAuthMethodSecurityQuestion      UsageAuthMethod = "securityQuestion"
	UsageAuthMethodAppNotification       UsageAuthMethod = "appNotification"
	UsageAuthMethodAppCode               UsageAuthMethod = "appCode"
	UsageAuthMethodAlternativeMobileCall UsageAuthMethod = "alternateMobileCall"
	UsageAuthMethodFido                  UsageAuthMethod = "fido"
	UsageAuthMethodAppPassword           UsageAuthMethod = "appPassword"
	UsageAuthMethodUnknownFutureValue    UsageAuthMethod = "unknownFutureValue"
)

type IncludedUserRoles = string

const (
	IncludedUserRolesAll             IncludedUserRoles = "all"
	IncludedUserRolesPrivilegedAdmin IncludedUserRoles = "privilegedAdmin"
	IncludedUserRolesAdmin           IncludedUserRoles = "admin"
	IncludedUserRolesUser            IncludedUserRoles = "user"
)

type IncludedUserTypes = string

const (
	IncludedUserTypesAll    IncludedUserTypes = "all"
	IncludedUserTypesMember IncludedUserTypes = "member"
	IncludedUserTypesGuest  IncludedUserTypes = "guest"
)
