package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPC{}

type CloudPC struct {
	// The Microsoft Entra device ID of the Cloud PC.
	AadDeviceId nullable.Type[string] `json:"aadDeviceId,omitempty"`

	// The allotment name divides tenant licenses into smaller batches or groups that help restrict the number of licenses
	// available for use in a specific assignment. When the provisioningType is dedicated, the allotment name is null.
	// Read-only.
	AllotmentDisplayName nullable.Type[string] `json:"allotmentDisplayName,omitempty"`

	// The connection setting of the Cloud PC. Possible values: enableSingleSignOn. Read Only.
	ConnectionSetting *CloudPCConnectionSetting `json:"connectionSetting,omitempty"`

	ConnectionSettings *CloudPCConnectionSettings `json:"connectionSettings,omitempty"`

	// The connectivity health check result of a Cloud PC, including the updated timestamp and whether the Cloud PC can be
	// connected.
	ConnectivityResult *CloudPCConnectivityResult `json:"connectivityResult,omitempty"`

	// The name of the geographical region where the Cloud PC is currently provisioned. For example, westus3, eastus2, and
	// southeastasia. Read-only.
	DeviceRegionName nullable.Type[string] `json:"deviceRegionName,omitempty"`

	// The disaster recovery status of the Cloud PC, including the primary region, secondary region, and capability type.
	// The default value is null that indicates that the disaster recovery setting is disabled. To receive a response with
	// the disasterRecoveryCapability property, $select and $filter it by disasterRecoveryCapability/{subProperty} in the
	// request URL. For more details, see Example 4: List Cloud PCs filtered by disaster recovery capability type.
	// Read-only.
	DisasterRecoveryCapability *CloudPCDisasterRecoveryCapability `json:"disasterRecoveryCapability,omitempty"`

	// The disk encryption applied to the Cloud PC. Possible values: notAvailable, notEncrypted,
	// encryptedUsingPlatformManagedKey, encryptedUsingCustomerManagedKey, and unknownFutureValue.
	DiskEncryptionState *CloudPCDiskEncryptionState `json:"diskEncryptionState,omitempty"`

	// The display name of the Cloud PC.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time when the grace period ends and reprovisioning or deprovisioning happens. Required only if the
	// status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	GracePeriodEndDateTime nullable.Type[string] `json:"gracePeriodEndDateTime,omitempty"`

	// Name of the OS image that's on the Cloud PC.
	ImageDisplayName nullable.Type[string] `json:"imageDisplayName,omitempty"`

	// The last login result of the Cloud PC. For example, { 'time': '2014-01-01T00:00:00Z'}.
	LastLoginResult *CloudPCLoginResult `json:"lastLoginResult,omitempty"`

	// The last modified date and time of the Cloud PC. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The last remote action result of the enterprise Cloud PCs. The supported remote actions are: Reboot, Rename,
	// Reprovision, Restore, Troubleshoot.
	LastRemoteActionResult *CloudPCRemoteActionResult `json:"lastRemoteActionResult,omitempty"`

	// The Intune device ID of the Cloud PC.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// The Intune device name of the Cloud PC.
	ManagedDeviceName nullable.Type[string] `json:"managedDeviceName,omitempty"`

	// The Azure network connection that is applied during the provisioning of Cloud PCs.
	OnPremisesConnectionName nullable.Type[string] `json:"onPremisesConnectionName,omitempty"`

	// The version of the operating system (OS) to provision on Cloud PCs. Possible values are: windows10, windows11,
	// unknownFutureValue.
	OsVersion *CloudPCOperatingSystem `json:"osVersion,omitempty"`

	// The results of every partner agent's installation status on Cloud PC.
	PartnerAgentInstallResults *[]CloudPCPartnerAgentInstallResult `json:"partnerAgentInstallResults,omitempty"`

	// The power state of a Cloud PC. The possible values are: running, poweredOff, unknown. This property only supports
	// shift work Cloud PCs.
	PowerState *CloudPCPowerState `json:"powerState,omitempty"`

	// The provisioning policy ID of the Cloud PC.
	ProvisioningPolicyId nullable.Type[string] `json:"provisioningPolicyId,omitempty"`

	// The provisioning policy that is applied during the provisioning of Cloud PCs.
	ProvisioningPolicyName nullable.Type[string] `json:"provisioningPolicyName,omitempty"`

	// The type of licenses to be used when provisioning Cloud PCs using this policy. Possible values are: dedicated,
	// shared, unknownFutureValue,sharedByUser, sharedByEntraGroup. You must use the Prefer: include-unknown-enum-members
	// request header to get the following values from this evolvable enum: sharedByUser, sharedByEntraGroup. The default
	// value is dedicated. CAUTION: The shared member is deprecated and will stop returning on April 30, 2027； in the
	// future, use the sharedByUser member.
	ProvisioningType *CloudPCProvisioningType `json:"provisioningType,omitempty"`

	ScopeIds *[]string `json:"scopeIds,omitempty"`

	// The service plan ID of the Cloud PC.
	ServicePlanId nullable.Type[string] `json:"servicePlanId,omitempty"`

	// The service plan name of the Cloud PC.
	ServicePlanName nullable.Type[string] `json:"servicePlanName,omitempty"`

	// The service plan type of the Cloud PC.
	ServicePlanType *CloudPCServicePlanType `json:"servicePlanType,omitempty"`

	Status *CloudPCStatus `json:"status,omitempty"`

	// Indicates the detailed status associated with Cloud PC, including error/warning code, error/warning message,
	// additionalInformation. For example, { 'code': 'internalServerError', 'message': 'There was an error during the Cloud
	// PC upgrade. Please contact support.', 'additionalInformation': null }.
	StatusDetail *CloudPCStatusDetail `json:"statusDetail,omitempty"`

	// The details of the Cloud PC status. For example, { 'code': 'internalServerError', 'message': 'There was an error
	// during the Cloud PC upgrade. Please contact support.', 'additionalInformation': null }. This property is deprecated
	// and will no longer be supported effective August 31, 2024. Use statusDetail instead.
	StatusDetails *CloudPCStatusDetails `json:"statusDetails,omitempty"`

	// The account type of the user on provisioned Cloud PCs. Possible values are: standardUser, administrator,
	// unknownFutureValue.
	UserAccountType *CloudPCUserAccountType `json:"userAccountType,omitempty"`

	// The user principal name (UPN) of the user assigned to the Cloud PC.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s CloudPC) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPC{}

func (s CloudPC) MarshalJSON() ([]byte, error) {
	type wrapper CloudPC
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPC: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPC: %+v", err)
	}

	delete(decoded, "allotmentDisplayName")
	delete(decoded, "deviceRegionName")
	delete(decoded, "disasterRecoveryCapability")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPC"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPC: %+v", err)
	}

	return encoded, nil
}
