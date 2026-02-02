package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = VirtualEndpoint{}

type VirtualEndpoint struct {
	// A collection of Cloud PC audit events.
	AuditEvents *[]CloudPCAuditEvent `json:"auditEvents,omitempty"`

	// A collection of cloud-managed virtual desktops.
	CloudPCs *[]CloudPC `json:"cloudPCs,omitempty"`

	// A collection of device image resources on Cloud PC.
	DeviceImages *[]CloudPCDeviceImage `json:"deviceImages,omitempty"`

	// A collection of gallery image resources on Cloud PC.
	GalleryImages *[]CloudPCGalleryImage `json:"galleryImages,omitempty"`

	// A defined collection of Azure resource information that can be used to establish Azure network connections for Cloud
	// PCs.
	OnPremisesConnections *[]CloudPCOnPremisesConnection `json:"onPremisesConnections,omitempty"`

	// A collection of Cloud PC provisioning policies.
	ProvisioningPolicies *[]CloudPCProvisioningPolicy `json:"provisioningPolicies,omitempty"`

	// A collection of Cloud PC user settings.
	UserSettings *[]CloudPCUserSetting `json:"userSettings,omitempty"`

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

func (s VirtualEndpoint) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VirtualEndpoint{}

func (s VirtualEndpoint) MarshalJSON() ([]byte, error) {
	type wrapper VirtualEndpoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualEndpoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEndpoint: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualEndpoint"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualEndpoint: %+v", err)
	}

	return encoded, nil
}
