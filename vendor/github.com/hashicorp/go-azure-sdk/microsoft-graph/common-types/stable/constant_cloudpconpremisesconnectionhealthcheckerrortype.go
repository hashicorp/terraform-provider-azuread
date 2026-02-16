package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionHealthCheckErrorType string

const (
	CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckAccessDenied                                           CloudPCOnPremisesConnectionHealthCheckErrorType = "adJoinCheckAccessDenied"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckAccountLockedOrDisabled                                CloudPCOnPremisesConnectionHealthCheckErrorType = "adJoinCheckAccountLockedOrDisabled"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckAccountQuotaExceeded                                   CloudPCOnPremisesConnectionHealthCheckErrorType = "adJoinCheckAccountQuotaExceeded"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckComputerObjectAlreadyExists                            CloudPCOnPremisesConnectionHealthCheckErrorType = "adJoinCheckComputerObjectAlreadyExists"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckCredentialsExpired                                     CloudPCOnPremisesConnectionHealthCheckErrorType = "adJoinCheckCredentialsExpired"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckFqdnNotFound                                           CloudPCOnPremisesConnectionHealthCheckErrorType = "adJoinCheckFqdnNotFound"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckIncorrectCredentials                                   CloudPCOnPremisesConnectionHealthCheckErrorType = "adJoinCheckIncorrectCredentials"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckOrganizationalUnitIncorrectFormat                      CloudPCOnPremisesConnectionHealthCheckErrorType = "adJoinCheckOrganizationalUnitIncorrectFormat"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckOrganizationalUnitNotFound                             CloudPCOnPremisesConnectionHealthCheckErrorType = "adJoinCheckOrganizationalUnitNotFound"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckServerNotOperational                                   CloudPCOnPremisesConnectionHealthCheckErrorType = "adJoinCheckServerNotOperational"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckUnknownError                                           CloudPCOnPremisesConnectionHealthCheckErrorType = "adJoinCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckConnectDisabled                             CloudPCOnPremisesConnectionHealthCheckErrorType = "azureAdDeviceSyncCheckConnectDisabled"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckDeviceNotFound                              CloudPCOnPremisesConnectionHealthCheckErrorType = "azureAdDeviceSyncCheckDeviceNotFound"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckDurationExceeded                            CloudPCOnPremisesConnectionHealthCheckErrorType = "azureAdDeviceSyncCheckDurationExceeded"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckLongSyncCircle                              CloudPCOnPremisesConnectionHealthCheckErrorType = "azureAdDeviceSyncCheckLongSyncCircle"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckScpNotConfigured                            CloudPCOnPremisesConnectionHealthCheckErrorType = "azureAdDeviceSyncCheckScpNotConfigured"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckTransientServiceError                       CloudPCOnPremisesConnectionHealthCheckErrorType = "azureAdDeviceSyncCheckTransientServiceError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckUnknownError                                CloudPCOnPremisesConnectionHealthCheckErrorType = "azureAdDeviceSyncCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_DnsCheckFqdnNotFound                                              CloudPCOnPremisesConnectionHealthCheckErrorType = "dnsCheckFqdnNotFound"
	CloudPCOnPremisesConnectionHealthCheckErrorType_DnsCheckNameWithInvalidCharacter                                  CloudPCOnPremisesConnectionHealthCheckErrorType = "dnsCheckNameWithInvalidCharacter"
	CloudPCOnPremisesConnectionHealthCheckErrorType_DnsCheckUnknownError                                              CloudPCOnPremisesConnectionHealthCheckErrorType = "dnsCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckAzureADUrlNotAllowListed                 CloudPCOnPremisesConnectionHealthCheckErrorType = "endpointConnectivityCheckAzureADUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckCloudPCUrlNotAllowListed                 CloudPCOnPremisesConnectionHealthCheckErrorType = "endpointConnectivityCheckCloudPcUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckIntuneUrlNotAllowListed                  CloudPCOnPremisesConnectionHealthCheckErrorType = "endpointConnectivityCheckIntuneUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckLocaleUrlNotAllowListed                  CloudPCOnPremisesConnectionHealthCheckErrorType = "endpointConnectivityCheckLocaleUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckUnknownError                             CloudPCOnPremisesConnectionHealthCheckErrorType = "endpointConnectivityCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckWVDUrlNotAllowListed                     CloudPCOnPremisesConnectionHealthCheckErrorType = "endpointConnectivityCheckWVDUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorAllocateResourceFailed                         CloudPCOnPremisesConnectionHealthCheckErrorType = "internalServerErrorAllocateResourceFailed"
	CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorDeploymentCanceled                             CloudPCOnPremisesConnectionHealthCheckErrorType = "internalServerErrorDeploymentCanceled"
	CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorUnableToRunDscScript                           CloudPCOnPremisesConnectionHealthCheckErrorType = "internalServerErrorUnableToRunDscScript"
	CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorVMDeploymentTimeout                            CloudPCOnPremisesConnectionHealthCheckErrorType = "internalServerErrorVMDeploymentTimeout"
	CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerUnknownError                                        CloudPCOnPremisesConnectionHealthCheckErrorType = "internalServerUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoResourceGroupNetworkContributorRole              CloudPCOnPremisesConnectionHealthCheckErrorType = "permissionCheckNoResourceGroupNetworkContributorRole"
	CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoResourceGroupOwnerRole                           CloudPCOnPremisesConnectionHealthCheckErrorType = "permissionCheckNoResourceGroupOwnerRole"
	CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoSubscriptionReaderRole                           CloudPCOnPremisesConnectionHealthCheckErrorType = "permissionCheckNoSubscriptionReaderRole"
	CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoVNetContributorRole                              CloudPCOnPremisesConnectionHealthCheckErrorType = "permissionCheckNoVNetContributorRole"
	CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoWindows365NetworkInterfaceContributorRole        CloudPCOnPremisesConnectionHealthCheckErrorType = "permissionCheckNoWindows365NetworkInterfaceContributorRole"
	CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoWindows365NetworkUserRole                        CloudPCOnPremisesConnectionHealthCheckErrorType = "permissionCheckNoWindows365NetworkUserRole"
	CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckTransientServiceError                              CloudPCOnPremisesConnectionHealthCheckErrorType = "permissionCheckTransientServiceError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckUnknownError                                       CloudPCOnPremisesConnectionHealthCheckErrorType = "permissionCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckAzurePolicyViolation                     CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckAzurePolicyViolation"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckDeploymentQuotaLimitReached              CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckDeploymentQuotaLimitReached"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckGeneralSubscriptionError                 CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckGeneralSubscriptionError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckIntuneCustomWindowsRestrictionViolation  CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckIntuneCustomWindowsRestrictionViolation"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckIntuneDefaultWindowsRestrictionViolation CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckIntuneDefaultWindowsRestrictionViolation"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckNoIntuneReaderRoleError                  CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckNoIntuneReaderRoleError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckNoSubnetIP                               CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckNoSubnetIP"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupBeingDeleted                CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckResourceGroupBeingDeleted"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupInvalid                     CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckResourceGroupInvalid"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupLockedForDelete             CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckResourceGroupLockedForDelete"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupLockedForReadonly           CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckResourceGroupLockedForReadonly"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubnetDelegationFailed                   CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckSubnetDelegationFailed"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubnetInvalid                            CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckSubnetInvalid"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubnetWithExternalResources              CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckSubnetWithExternalResources"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubscriptionDisabled                     CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckSubscriptionDisabled"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubscriptionNotFound                     CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckSubscriptionNotFound"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubscriptionTransferred                  CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckSubscriptionTransferred"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckTransientServiceError                    CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckTransientServiceError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckUnknownError                             CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckUnsupportedVNetRegion                    CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckUnsupportedVNetRegion"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckVNetBeingMoved                           CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckVNetBeingMoved"
	CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckVNetInvalid                              CloudPCOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckVNetInvalid"
	CloudPCOnPremisesConnectionHealthCheckErrorType_SsoCheckKerberosConfigurationError                                CloudPCOnPremisesConnectionHealthCheckErrorType = "ssoCheckKerberosConfigurationError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckStunUrlNotAllowListed                         CloudPCOnPremisesConnectionHealthCheckErrorType = "udpConnectivityCheckStunUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckTurnUrlNotAllowListed                         CloudPCOnPremisesConnectionHealthCheckErrorType = "udpConnectivityCheckTurnUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckUnknownError                                  CloudPCOnPremisesConnectionHealthCheckErrorType = "udpConnectivityCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckUrlsNotAllowListed                            CloudPCOnPremisesConnectionHealthCheckErrorType = "udpConnectivityCheckUrlsNotAllowListed"
)

func PossibleValuesForCloudPCOnPremisesConnectionHealthCheckErrorType() []string {
	return []string{
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckAccessDenied),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckAccountLockedOrDisabled),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckAccountQuotaExceeded),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckComputerObjectAlreadyExists),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckCredentialsExpired),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckFqdnNotFound),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckIncorrectCredentials),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckOrganizationalUnitIncorrectFormat),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckOrganizationalUnitNotFound),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckServerNotOperational),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckConnectDisabled),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckDeviceNotFound),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckDurationExceeded),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckLongSyncCircle),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckScpNotConfigured),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckTransientServiceError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_DnsCheckFqdnNotFound),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_DnsCheckNameWithInvalidCharacter),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_DnsCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckAzureADUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckCloudPCUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckIntuneUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckLocaleUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckWVDUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorAllocateResourceFailed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorDeploymentCanceled),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorUnableToRunDscScript),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorVMDeploymentTimeout),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoResourceGroupNetworkContributorRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoResourceGroupOwnerRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoSubscriptionReaderRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoVNetContributorRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoWindows365NetworkInterfaceContributorRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoWindows365NetworkUserRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckTransientServiceError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckAzurePolicyViolation),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckDeploymentQuotaLimitReached),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckGeneralSubscriptionError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckIntuneCustomWindowsRestrictionViolation),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckIntuneDefaultWindowsRestrictionViolation),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckNoIntuneReaderRoleError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckNoSubnetIP),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupBeingDeleted),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupInvalid),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupLockedForDelete),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupLockedForReadonly),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubnetDelegationFailed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubnetInvalid),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubnetWithExternalResources),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubscriptionDisabled),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubscriptionNotFound),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubscriptionTransferred),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckTransientServiceError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckUnsupportedVNetRegion),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckVNetBeingMoved),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckVNetInvalid),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_SsoCheckKerberosConfigurationError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckStunUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckTurnUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckUrlsNotAllowListed),
	}
}

func (s *CloudPCOnPremisesConnectionHealthCheckErrorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCOnPremisesConnectionHealthCheckErrorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCOnPremisesConnectionHealthCheckErrorType(input string) (*CloudPCOnPremisesConnectionHealthCheckErrorType, error) {
	vals := map[string]CloudPCOnPremisesConnectionHealthCheckErrorType{
		"adjoincheckaccessdenied":                                           CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckAccessDenied,
		"adjoincheckaccountlockedordisabled":                                CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckAccountLockedOrDisabled,
		"adjoincheckaccountquotaexceeded":                                   CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckAccountQuotaExceeded,
		"adjoincheckcomputerobjectalreadyexists":                            CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckComputerObjectAlreadyExists,
		"adjoincheckcredentialsexpired":                                     CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckCredentialsExpired,
		"adjoincheckfqdnnotfound":                                           CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckFqdnNotFound,
		"adjoincheckincorrectcredentials":                                   CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckIncorrectCredentials,
		"adjoincheckorganizationalunitincorrectformat":                      CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckOrganizationalUnitIncorrectFormat,
		"adjoincheckorganizationalunitnotfound":                             CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckOrganizationalUnitNotFound,
		"adjoincheckservernotoperational":                                   CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckServerNotOperational,
		"adjoincheckunknownerror":                                           CloudPCOnPremisesConnectionHealthCheckErrorType_AdJoinCheckUnknownError,
		"azureaddevicesynccheckconnectdisabled":                             CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckConnectDisabled,
		"azureaddevicesynccheckdevicenotfound":                              CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckDeviceNotFound,
		"azureaddevicesynccheckdurationexceeded":                            CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckDurationExceeded,
		"azureaddevicesyncchecklongsynccircle":                              CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckLongSyncCircle,
		"azureaddevicesynccheckscpnotconfigured":                            CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckScpNotConfigured,
		"azureaddevicesyncchecktransientserviceerror":                       CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckTransientServiceError,
		"azureaddevicesynccheckunknownerror":                                CloudPCOnPremisesConnectionHealthCheckErrorType_AzureAdDeviceSyncCheckUnknownError,
		"dnscheckfqdnnotfound":                                              CloudPCOnPremisesConnectionHealthCheckErrorType_DnsCheckFqdnNotFound,
		"dnschecknamewithinvalidcharacter":                                  CloudPCOnPremisesConnectionHealthCheckErrorType_DnsCheckNameWithInvalidCharacter,
		"dnscheckunknownerror":                                              CloudPCOnPremisesConnectionHealthCheckErrorType_DnsCheckUnknownError,
		"endpointconnectivitycheckazureadurlnotallowlisted":                 CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckAzureADUrlNotAllowListed,
		"endpointconnectivitycheckcloudpcurlnotallowlisted":                 CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckCloudPCUrlNotAllowListed,
		"endpointconnectivitycheckintuneurlnotallowlisted":                  CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckIntuneUrlNotAllowListed,
		"endpointconnectivitychecklocaleurlnotallowlisted":                  CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckLocaleUrlNotAllowListed,
		"endpointconnectivitycheckunknownerror":                             CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckUnknownError,
		"endpointconnectivitycheckwvdurlnotallowlisted":                     CloudPCOnPremisesConnectionHealthCheckErrorType_EndpointConnectivityCheckWVDUrlNotAllowListed,
		"internalservererrorallocateresourcefailed":                         CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorAllocateResourceFailed,
		"internalservererrordeploymentcanceled":                             CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorDeploymentCanceled,
		"internalservererrorunabletorundscscript":                           CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorUnableToRunDscScript,
		"internalservererrorvmdeploymenttimeout":                            CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerErrorVMDeploymentTimeout,
		"internalserverunknownerror":                                        CloudPCOnPremisesConnectionHealthCheckErrorType_InternalServerUnknownError,
		"permissionchecknoresourcegroupnetworkcontributorrole":              CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoResourceGroupNetworkContributorRole,
		"permissionchecknoresourcegroupownerrole":                           CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoResourceGroupOwnerRole,
		"permissionchecknosubscriptionreaderrole":                           CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoSubscriptionReaderRole,
		"permissionchecknovnetcontributorrole":                              CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoVNetContributorRole,
		"permissionchecknowindows365networkinterfacecontributorrole":        CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoWindows365NetworkInterfaceContributorRole,
		"permissionchecknowindows365networkuserrole":                        CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckNoWindows365NetworkUserRole,
		"permissionchecktransientserviceerror":                              CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckTransientServiceError,
		"permissioncheckunknownerror":                                       CloudPCOnPremisesConnectionHealthCheckErrorType_PermissionCheckUnknownError,
		"resourceavailabilitycheckazurepolicyviolation":                     CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckAzurePolicyViolation,
		"resourceavailabilitycheckdeploymentquotalimitreached":              CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckDeploymentQuotaLimitReached,
		"resourceavailabilitycheckgeneralsubscriptionerror":                 CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckGeneralSubscriptionError,
		"resourceavailabilitycheckintunecustomwindowsrestrictionviolation":  CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckIntuneCustomWindowsRestrictionViolation,
		"resourceavailabilitycheckintunedefaultwindowsrestrictionviolation": CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckIntuneDefaultWindowsRestrictionViolation,
		"resourceavailabilitychecknointunereaderroleerror":                  CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckNoIntuneReaderRoleError,
		"resourceavailabilitychecknosubnetip":                               CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckNoSubnetIP,
		"resourceavailabilitycheckresourcegroupbeingdeleted":                CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupBeingDeleted,
		"resourceavailabilitycheckresourcegroupinvalid":                     CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupInvalid,
		"resourceavailabilitycheckresourcegrouplockedfordelete":             CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupLockedForDelete,
		"resourceavailabilitycheckresourcegrouplockedforreadonly":           CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckResourceGroupLockedForReadonly,
		"resourceavailabilitychecksubnetdelegationfailed":                   CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubnetDelegationFailed,
		"resourceavailabilitychecksubnetinvalid":                            CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubnetInvalid,
		"resourceavailabilitychecksubnetwithexternalresources":              CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubnetWithExternalResources,
		"resourceavailabilitychecksubscriptiondisabled":                     CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubscriptionDisabled,
		"resourceavailabilitychecksubscriptionnotfound":                     CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubscriptionNotFound,
		"resourceavailabilitychecksubscriptiontransferred":                  CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckSubscriptionTransferred,
		"resourceavailabilitychecktransientserviceerror":                    CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckTransientServiceError,
		"resourceavailabilitycheckunknownerror":                             CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckUnknownError,
		"resourceavailabilitycheckunsupportedvnetregion":                    CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckUnsupportedVNetRegion,
		"resourceavailabilitycheckvnetbeingmoved":                           CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckVNetBeingMoved,
		"resourceavailabilitycheckvnetinvalid":                              CloudPCOnPremisesConnectionHealthCheckErrorType_ResourceAvailabilityCheckVNetInvalid,
		"ssocheckkerberosconfigurationerror":                                CloudPCOnPremisesConnectionHealthCheckErrorType_SsoCheckKerberosConfigurationError,
		"udpconnectivitycheckstunurlnotallowlisted":                         CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckStunUrlNotAllowListed,
		"udpconnectivitycheckturnurlnotallowlisted":                         CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckTurnUrlNotAllowListed,
		"udpconnectivitycheckunknownerror":                                  CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckUnknownError,
		"udpconnectivitycheckurlsnotallowlisted":                            CloudPCOnPremisesConnectionHealthCheckErrorType_UdpConnectivityCheckUrlsNotAllowListed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOnPremisesConnectionHealthCheckErrorType(input)
	return &out, nil
}
