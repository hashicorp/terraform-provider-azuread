package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCOnPremisesConnection{}

type CloudPCOnPremisesConnection struct {
	// The fully qualified domain name (FQDN) of the Active Directory domain you want to join. Maximum length is 255.
	// Optional.
	AdDomainName nullable.Type[string] `json:"adDomainName,omitempty"`

	// The password associated with the username of an Active Directory account (adDomainUsername).
	AdDomainPassword nullable.Type[string] `json:"adDomainPassword,omitempty"`

	// The username of an Active Directory account (user or service account) that has permission to create computer objects
	// in Active Directory. Required format: admin@contoso.com. Optional.
	AdDomainUsername nullable.Type[string] `json:"adDomainUsername,omitempty"`

	// The interface URL of the partner service's resource that links to this Azure network connection. Returned only on
	// $select.
	AlternateResourceUrl nullable.Type[string] `json:"alternateResourceUrl,omitempty"`

	// Specifies how the provisioned Cloud PC joins to Microsoft Entra. It includes different types, one is Microsoft Entra
	// ID join, which means there's no on-premises Active Directory (AD) in the current tenant, and the Cloud PC device is
	// joined by Microsoft Entra. Another one is hybridAzureADJoin, which means there's also an on-premises Active Directory
	// (AD) in the current tenant and the Cloud PC device joins to on-premises Active Directory (AD) and Microsoft Entra.
	// The type also determines which types of users can be assigned and can sign into a Cloud PC. The azureADJoin type
	// indicates that cloud-only and hybrid users can be assigned and signed into the Cloud PC. hybridAzureADJoin indicates
	// only hybrid users can be assigned and signed into the Cloud PC. The default value is hybridAzureADJoin.
	ConnectionType *CloudPCOnPremisesConnectionType `json:"connectionType,omitempty"`

	// The display name for the Azure network connection.
	DisplayName *string `json:"displayName,omitempty"`

	HealthCheckStatus *CloudPCOnPremisesConnectionStatus `json:"healthCheckStatus,omitempty"`

	// Indicates the results of health checks performed on the on-premises connection. Read-only. Returned only on $select.
	// For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure
	// network connection, including healthCheckStatusDetail. Read-only.
	HealthCheckStatusDetail *CloudPCOnPremisesConnectionStatusDetail `json:"healthCheckStatusDetail,omitempty"`

	// When true, the Azure network connection is in use. When false, the connection isn't in use. You can't delete a
	// connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see
	// Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetail. Read-only.
	InUse nullable.Type[bool] `json:"inUse,omitempty"`

	// The organizational unit (OU) in which the computer account is created. If left null, the OU configured as the default
	// (a well-known computer object container) in the tenant's Active Directory domain (OU) is used. Optional.
	OrganizationalUnit nullable.Type[string] `json:"organizationalUnit,omitempty"`

	// The unique identifier of the target resource group used associated with the on-premises network connectivity for
	// Cloud PCs. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}'
	ResourceGroupId *string `json:"resourceGroupId,omitempty"`

	// The unique identifier of the target subnet used associated with the on-premises network connectivity for Cloud PCs.
	// Required format:
	// '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}'
	SubnetId *string `json:"subnetId,omitempty"`

	// The unique identifier of the Azure subscription associated with the tenant.
	SubscriptionId *string `json:"subscriptionId,omitempty"`

	// The name of the Azure subscription is used to create an Azure network connection. Read-only.
	SubscriptionName nullable.Type[string] `json:"subscriptionName,omitempty"`

	// The unique identifier of the target virtual network used associated with the on-premises network connectivity for
	// Cloud PCs. Required format:
	// '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}'
	VirtualNetworkId *string `json:"virtualNetworkId,omitempty"`

	// Indicates the resource location of the target virtual network. For example, the location can be eastus2, westeurope,
	// etc. Read-only (computed value).
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

	delete(decoded, "healthCheckStatusDetail")
	delete(decoded, "inUse")
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
