package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MicrosoftTunnelSite{}

type MicrosoftTunnelSite struct {
	// The site's description (optional)
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the site. This property is required when a site is created.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The site's Internal Network Access Probe URL
	InternalNetworkProbeUrl nullable.Type[string] `json:"internalNetworkProbeUrl,omitempty"`

	// The MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
	MicrosoftTunnelConfiguration *MicrosoftTunnelConfiguration `json:"microsoftTunnelConfiguration,omitempty"`

	// A list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
	MicrosoftTunnelServers *[]MicrosoftTunnelServer `json:"microsoftTunnelServers,omitempty"`

	// The site's public domain name or IP address
	PublicAddress nullable.Type[string] `json:"publicAddress,omitempty"`

	// List of Scope Tags for this Entity instance
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The site's automatic upgrade setting. True for automatic upgrades, false for manual control
	UpgradeAutomatically *bool `json:"upgradeAutomatically,omitempty"`

	// The site provides the state of when an upgrade is available
	UpgradeAvailable *bool `json:"upgradeAvailable,omitempty"`

	// The site's upgrade window end time of day
	UpgradeWindowEndTime nullable.Type[string] `json:"upgradeWindowEndTime,omitempty"`

	// The site's upgrade window start time of day
	UpgradeWindowStartTime nullable.Type[string] `json:"upgradeWindowStartTime,omitempty"`

	// The site's timezone represented as a minute offset from UTC
	UpgradeWindowUtcOffsetInMinutes nullable.Type[int64] `json:"upgradeWindowUtcOffsetInMinutes,omitempty"`

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

func (s MicrosoftTunnelSite) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MicrosoftTunnelSite{}

func (s MicrosoftTunnelSite) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftTunnelSite
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftTunnelSite: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftTunnelSite: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftTunnelSite"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftTunnelSite: %+v", err)
	}

	return encoded, nil
}
