package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCOrganizationSettings{}

type CloudPCOrganizationSettings struct {
	// Specifies whether new Cloud PCs will be automatically enrolled in Microsoft Endpoint Manager (MEM). The default value
	// is false.
	EnableMEMAutoEnroll nullable.Type[bool] `json:"enableMEMAutoEnroll,omitempty"`

	// True if the provisioned Cloud PC can be accessed by single sign-on. False indicates that the provisioned Cloud PC
	// doesn't support this feature. Default value is false. Windows 365 users can use single sign-on to authenticate to
	// Microsoft Entra ID with passwordless options (for example, FIDO keys) to access their Cloud PC. Optional.
	EnableSingleSignOn nullable.Type[bool] `json:"enableSingleSignOn,omitempty"`

	// The version of the operating system (OS) to provision on Cloud PCs. The possible values are: windows10, windows11,
	// unknownFutureValue.
	OsVersion *CloudPCOperatingSystem `json:"osVersion,omitempty"`

	// The account type of the user on provisioned Cloud PCs. The possible values are: standardUser, administrator,
	// unknownFutureValue.
	UserAccountType *CloudPCUserAccountType `json:"userAccountType,omitempty"`

	// Represents the Cloud PC organization settings for a tenant. A tenant has only one cloudPcOrganizationSettings object.
	// The default language value en-US.
	WindowsSettings *CloudPCWindowsSettings `json:"windowsSettings,omitempty"`

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

func (s CloudPCOrganizationSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCOrganizationSettings{}

func (s CloudPCOrganizationSettings) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCOrganizationSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCOrganizationSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCOrganizationSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcOrganizationSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCOrganizationSettings: %+v", err)
	}

	return encoded, nil
}
