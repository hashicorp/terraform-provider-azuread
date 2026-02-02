package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RemoteDesktopSecurityConfiguration{}

type RemoteDesktopSecurityConfiguration struct {
	ApprovedClientApps *[]ApprovedClientApp `json:"approvedClientApps,omitempty"`

	// Determines if Microsoft Entra ID RDS authentication protocol for RDP is enabled.
	IsRemoteDesktopProtocolEnabled *bool `json:"isRemoteDesktopProtocolEnabled,omitempty"`

	// The collection of target device groups that are associated with the RDS security configuration that will be enabled
	// for SSO when a client connects to the target device over RDP using the new Microsoft Entra ID RDS authentication
	// protocol.
	TargetDeviceGroups *[]TargetDeviceGroup `json:"targetDeviceGroups,omitempty"`

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

func (s RemoteDesktopSecurityConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RemoteDesktopSecurityConfiguration{}

func (s RemoteDesktopSecurityConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper RemoteDesktopSecurityConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RemoteDesktopSecurityConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RemoteDesktopSecurityConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.remoteDesktopSecurityConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RemoteDesktopSecurityConfiguration: %+v", err)
	}

	return encoded, nil
}
