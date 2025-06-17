package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCOnPremisesConnection{}

type CloudPCOnPremisesConnection struct {
	// The fully qualified domain name (FQDN) of the Active Directory domain you want to join. Optional.
	AdDomainName nullable.Type[string] `json:"adDomainName,omitempty"`

	// The password associated with adDomainUsername.
	AdDomainPassword nullable.Type[string] `json:"adDomainPassword,omitempty"`

	// The username of an Active Directory account (user or service account) that has permissions to create computer objects
	// in Active Directory. Required format: admin@contoso.com. Optional.
	AdDomainUsername nullable.Type[string] `json:"adDomainUsername,omitempty"`

	// The interface URL of the partner service's resource that links to this Azure network connection. Returned only on
	// $select.
	AlternateResourceUrl nullable.Type[string] `json:"alternateResourceUrl,omitempty"`

	// Specifies the method by which a provisioned Cloud PC is joined to Microsoft Entra. The azureADJoin option indicates
	// the absence of an on-premises Active Directory (AD) in the current tenant that results in the Cloud PC device only
	// joining to Microsoft Entra. The hybridAzureADJoin option indicates the presence of an on-premises AD in the current
	// tenant and that the Cloud PC joins both the on-premises AD and Microsoft Entra. The selected option also determines
	// the types of users who can be assigned and can sign into a Cloud PC. The azureADJoin option allows both cloud-only
	// and hybrid users to be assigned and sign in, whereas hybridAzureADJoin is restricted to hybrid users only. The
	// default value is hybridAzureADJoin. The possible values are: hybridAzureADJoin, azureADJoin, unknownFutureValue.
	ConnectionType *CloudPCOnPremisesConnectionType `json:"connectionType,omitempty"`

	// The display name for the Azure network connection.
	DisplayName *string `json:"displayName,omitempty"`

	// false if the regular health checks on the network/domain configuration are currently active. true if the checks are
	// paused. If you perform a create or update operation on a onPremisesNetworkConnection resource, this value is set to
	// false for 4 weeks. If you retry a health check on network/domain configuration, this value is set to false for two
	// weeks. If the onPremisesNetworkConnection resource is attached in a provisioningPolicy or used by a Cloud PC in the
	// past 4 weeks, healthCheckPaused is set to false. Read-only. Default is false.
	HealthCheckPaused nullable.Type[bool] `json:"healthCheckPaused,omitempty"`

	HealthCheckStatus *CloudPCOnPremisesConnectionStatus `json:"healthCheckStatus,omitempty"`

	// Indicates the results of health checks performed on the on-premises connection. Returned only on $select. For an
	// example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network
	// connection, including healthCheckStatusDetails. Read-only.
	HealthCheckStatusDetail *CloudPCOnPremisesConnectionStatusDetail `json:"healthCheckStatusDetail,omitempty"`

	// The details of the connection's health checks and the corresponding results. Returned only on $select. For an example
	// that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection,
	// including healthCheckStatusDetails. Read-only.
	HealthCheckStatusDetails *CloudPCOnPremisesConnectionStatusDetails `json:"healthCheckStatusDetails,omitempty"`

	// When true, the Azure network connection is in use. When false, the connection isn't in use. You can't delete a
	// connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see
	// Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetails. Read-only.
	InUse nullable.Type[bool] `json:"inUse,omitempty"`

	// Indicates whether a Cloud PC is using this on-premises network connection. true if at least one Cloud PC is using it.
	// Otherwise, false. Read-only. Default is false.
	InUseByCloudPC nullable.Type[bool] `json:"inUseByCloudPc,omitempty"`

	ManagedBy *CloudPCManagementService `json:"managedBy,omitempty"`

	// The organizational unit (OU) in which the computer account is created. If left null, the OU configured as the default
	// (a well-known computer object container) in your Active Directory domain (OU) is used. Optional.
	OrganizationalUnit nullable.Type[string] `json:"organizationalUnit,omitempty"`

	// The ID of the target resource group. Required format:
	// /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}.
	ResourceGroupId *string `json:"resourceGroupId,omitempty"`

	ScopeIds *[]string `json:"scopeIds,omitempty"`

	// The ID of the target subnet. Required format:
	// /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}.
	SubnetId *string `json:"subnetId,omitempty"`

	// The ID of the target Azure subscription associated with your tenant.
	SubscriptionId *string `json:"subscriptionId,omitempty"`

	// The name of the target Azure subscription. Read-only.
	SubscriptionName nullable.Type[string] `json:"subscriptionName,omitempty"`

	// Specifies the method by which a provisioned Cloud PC is joined to Microsoft Entra. The azureADJoin option indicates
	// the absence of an on-premises Active Directory (AD) in the current tenant that results in the Cloud PC device only
	// joining to Microsoft Entra. The hybridAzureADJoin option indicates the presence of an on-premises AD in the current
	// tenant and that the Cloud PC joins both the on-premises AD and Microsoft Entra. The selected option also determines
	// the types of users who can be assigned and can sign into a Cloud PC. The azureADJoin option allows both cloud-only
	// and hybrid users to be assigned and sign in, whereas hybridAzureADJoin is restricted to hybrid users only. The
	// default value is hybridAzureADJoin. The possible values are: hybridAzureADJoin, azureADJoin, unknownFutureValue. The
	// type property is deprecated and stopped returning data on January 31, 2024. Going forward, use the connectionType
	// property.
	Type *CloudPCOnPremisesConnectionType `json:"type,omitempty"`

	// The ID of the target virtual network. Required format:
	// /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}.
	VirtualNetworkId *string `json:"virtualNetworkId,omitempty"`

	// Indicates the resource location of the virtual target network. Read-only, computed value.
	VirtualNetworkLocation nullable.Type[string] `json:"virtualNetworkLocation,omitempty"`

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

func (s CloudPCOnPremisesConnection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCOnPremisesConnection{}

func (s CloudPCOnPremisesConnection) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCOnPremisesConnection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCOnPremisesConnection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCOnPremisesConnection: %+v", err)
	}

	delete(decoded, "healthCheckPaused")
	delete(decoded, "healthCheckStatusDetail")
	delete(decoded, "healthCheckStatusDetails")
	delete(decoded, "inUse")
	delete(decoded, "inUseByCloudPc")
	delete(decoded, "subscriptionName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcOnPremisesConnection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCOnPremisesConnection: %+v", err)
	}

	return encoded, nil
}
