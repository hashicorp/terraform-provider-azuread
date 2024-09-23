package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessNetworkAccessRoot{}

type NetworkaccessNetworkAccessRoot struct {
	Alerts *[]NetworkaccessAlert `json:"alerts,omitempty"`

	// Connectivity represents all the connectivity components in Global Secure Access.
	Connectivity *NetworkaccessConnectivity `json:"connectivity,omitempty"`

	// A filtering policy defines the specific traffic that is allowed or blocked through the Global Secure Access services
	// for a filtering profile.
	FilteringPolicies *[]NetworkaccessFilteringPolicy `json:"filteringPolicies,omitempty"`

	// A filtering profile associates network access policies with Microsoft Entra ID Conditional Access policies, so that
	// access policies can be applied to users and groups.
	FilteringProfiles *[]NetworkaccessFilteringProfile `json:"filteringProfiles,omitempty"`

	// A forwarding policy defines the specific traffic that is routed through the Global Secure Access Service. It's then
	// added to a forwarding profile.
	ForwardingPolicies *[]NetworkaccessForwardingPolicy `json:"forwardingPolicies,omitempty"`

	// A forwarding profile determines which types of traffic are routed through the Global Secure Access services and which
	// ones are skipped. The handling of specific traffic is determined by the forwarding policies that are added to the
	// forwarding profile.
	ForwardingProfiles *[]NetworkaccessForwardingProfile `json:"forwardingProfiles,omitempty"`

	// Represents network connections that are routed through Global Secure Access.
	Logs *NetworkaccessLogs `json:"logs,omitempty"`

	// Represents the status of the Global Secure Access services for the tenant.
	Reports *NetworkaccessReports `json:"reports,omitempty"`

	// Global Secure Access settings.
	Settings *NetworkaccessSettings `json:"settings,omitempty"`

	// Represents the status of the Global Secure Access services for the tenant.
	TenantStatus *NetworkaccessTenantStatus `json:"tenantStatus,omitempty"`

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

func (s NetworkaccessNetworkAccessRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessNetworkAccessRoot{}

func (s NetworkaccessNetworkAccessRoot) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessNetworkAccessRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessNetworkAccessRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessNetworkAccessRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.networkAccessRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessNetworkAccessRoot: %+v", err)
	}

	return encoded, nil
}
