package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EndpointPrivilegeManagementProvisioningStatus{}

type EndpointPrivilegeManagementProvisioningStatus struct {
	// Indicates whether tenant has a valid Intune Endpoint Privilege Management license. Possible value are : 0 - notPaid,
	// 1 - paid, 2 - trial. See LicenseType enum for more details. Default notPaid .
	LicenseType *LicenseType `json:"licenseType,omitempty"`

	// Indicates whether tenant is onboarded to Microsoft Managed Platform - Cloud (MMPC). When set to true, implies tenant
	// is onboarded and when set to false, implies tenant is not onboarded. Default set to false.
	OnboardedToMicrosoftManagedPlatform *bool `json:"onboardedToMicrosoftManagedPlatform,omitempty"`

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

func (s EndpointPrivilegeManagementProvisioningStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EndpointPrivilegeManagementProvisioningStatus{}

func (s EndpointPrivilegeManagementProvisioningStatus) MarshalJSON() ([]byte, error) {
	type wrapper EndpointPrivilegeManagementProvisioningStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EndpointPrivilegeManagementProvisioningStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EndpointPrivilegeManagementProvisioningStatus: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.endpointPrivilegeManagementProvisioningStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EndpointPrivilegeManagementProvisioningStatus: %+v", err)
	}

	return encoded, nil
}
