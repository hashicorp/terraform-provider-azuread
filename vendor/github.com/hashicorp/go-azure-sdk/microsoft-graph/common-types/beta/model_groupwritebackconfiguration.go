package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WritebackConfiguration = GroupWritebackConfiguration{}

type GroupWritebackConfiguration struct {
	// Indicates the target on-premises group type the cloud object is written back as. Nullable. The possible values are:
	// universalDistributionGroup, universalSecurityGroup, universalMailEnabledSecurityGroup.If the cloud group is a unified
	// (Microsoft 365) group, this property can be one of the following: universalDistributionGroup, universalSecurityGroup,
	// universalMailEnabledSecurityGroup. Microsoft Entra security groups can be written back as universalSecurityGroup. If
	// isEnabled or the NewUnifiedGroupWritebackDefault group setting is true but this property isn't explicitly configured:
	// Microsoft 365 groups are written back as universalDistributionGroup by defaultSecurity groups are written back as
	// universalSecurityGroup by default
	OnPremisesGroupType nullable.Type[string] `json:"onPremisesGroupType,omitempty"`

	// Fields inherited from WritebackConfiguration

	// Indicates whether writeback of cloud groups to on-premise Active Directory is enabled. Default value is true for
	// Microsoft 365 groups and false for security groups.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s GroupWritebackConfiguration) WritebackConfiguration() BaseWritebackConfigurationImpl {
	return BaseWritebackConfigurationImpl{
		IsEnabled: s.IsEnabled,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupWritebackConfiguration{}

func (s GroupWritebackConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper GroupWritebackConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupWritebackConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupWritebackConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupWritebackConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupWritebackConfiguration: %+v", err)
	}

	return encoded, nil
}
