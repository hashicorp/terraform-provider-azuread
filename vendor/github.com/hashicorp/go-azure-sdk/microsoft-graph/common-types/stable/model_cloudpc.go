package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPC{}

type CloudPC struct {
	// The Microsoft Entra device ID for the Cloud PC, also known as the Azure Active Directory (Azure AD) device ID, that
	// consists of 32 characters in a GUID format. Generated on a VM joined to Microsoft Entra ID. Read-only.
	AadDeviceId nullable.Type[string] `json:"aadDeviceId,omitempty"`

	// The display name for the Cloud PC. Maximum length is 64 characters. Read-only. You can use the cloudPC: rename API to
	// modify the Cloud PC name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time when the grace period ends and reprovisioning or deprovisioning happen. Required only if the status
	// is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	GracePeriodEndDateTime nullable.Type[string] `json:"gracePeriodEndDateTime,omitempty"`

	// The name of the operating system image used for the Cloud PC. Maximum length is 50 characters. Only letters (A-Z,
	// a-z), numbers (0-9), and special characters (-,,.) are allowed for this property. The property value can't begin or
	// end with an underscore. Read-only.
	ImageDisplayName nullable.Type[string] `json:"imageDisplayName,omitempty"`

	// The last modified date and time of the Cloud PC. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The Intune enrolled device ID for the Cloud PC that consists of 32 characters in a GUID format. The managedDeviceId
	// property of Windows 365 Business Cloud PCs is always null as Windows 365 Business Cloud PCs aren't Intune-enrolled
	// automatically by Windows 365. Read-only.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// The Intune enrolled device name for the Cloud PC. The managedDeviceName property of Windows 365 Business Cloud PCs is
	// always null as Windows 365 Business Cloud PCs aren't Intune-enrolled automatically by Windows 365. Read-only.
	ManagedDeviceName nullable.Type[string] `json:"managedDeviceName,omitempty"`

	// The on-premises connection that applied during the provisioning of Cloud PCs. Read-only.
	OnPremisesConnectionName nullable.Type[string] `json:"onPremisesConnectionName,omitempty"`

	// The provisioning policy ID for the Cloud PC that consists of 32 characters in a GUID format. A policy defines the
	// type of Cloud PC the user wants to create. Read-only.
	ProvisioningPolicyId nullable.Type[string] `json:"provisioningPolicyId,omitempty"`

	// The provisioning policy that applied during the provisioning of Cloud PCs. Maximum length is 120 characters.
	// Read-only.
	ProvisioningPolicyName nullable.Type[string] `json:"provisioningPolicyName,omitempty"`

	// The type of licenses to be used when provisioning Cloud PCs using this policy. Possible values are: dedicated,
	// shared, unknownFutureValue. The default value is dedicated.
	ProvisioningType *CloudPCProvisioningType `json:"provisioningType,omitempty"`

	// The service plan ID for the Cloud PC that consists of 32 characters in a GUID format. For more information about
	// service plans, see Product names and service plan identifiers for licensing. Read-only.
	ServicePlanId nullable.Type[string] `json:"servicePlanId,omitempty"`

	// The service plan name for the customer-facing Cloud PC entity. Read-only.
	ServicePlanName nullable.Type[string] `json:"servicePlanName,omitempty"`

	// The user principal name (UPN) of the user assigned to the Cloud PC. Maximum length is 113 characters. For more
	// information on username policies, see Password policies and account restrictions in Microsoft Entra ID. Read-only.
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

	delete(decoded, "aadDeviceId")
	delete(decoded, "displayName")
	delete(decoded, "imageDisplayName")
	delete(decoded, "managedDeviceId")
	delete(decoded, "managedDeviceName")
	delete(decoded, "onPremisesConnectionName")
	delete(decoded, "provisioningPolicyId")
	delete(decoded, "provisioningPolicyName")
	delete(decoded, "servicePlanId")
	delete(decoded, "servicePlanName")
	delete(decoded, "userPrincipalName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPC"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPC: %+v", err)
	}

	return encoded, nil
}
