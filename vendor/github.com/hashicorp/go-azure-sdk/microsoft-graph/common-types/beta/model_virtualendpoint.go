package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = VirtualEndpoint{}

type VirtualEndpoint struct {
	// Cloud PC audit event.
	AuditEvents *[]CloudPCAuditEvent `json:"auditEvents,omitempty"`

	// Bulk actions applied to a Cloud PC.
	BulkActions *[]CloudPCBulkAction `json:"bulkActions,omitempty"`

	// Cloud managed virtual desktops.
	CloudPCs *[]CloudPC `json:"cloudPCs,omitempty"`

	// Cloud PC organization mapping between public and US Government Community Cloud (GCC) organizations.
	CrossCloudGovernmentOrganizationMapping *CloudPCCrossCloudGovernmentOrganizationMapping `json:"crossCloudGovernmentOrganizationMapping,omitempty"`

	// The image resource on Cloud PC.
	DeviceImages *[]CloudPCDeviceImage `json:"deviceImages,omitempty"`

	// The external partner settings on a Cloud PC.
	ExternalPartnerSettings *[]CloudPCExternalPartnerSetting `json:"externalPartnerSettings,omitempty"`

	// Front-line service plans for a Cloud PC.
	FrontLineServicePlans *[]CloudPCFrontLineServicePlan `json:"frontLineServicePlans,omitempty"`

	// The gallery image resource on Cloud PC.
	GalleryImages *[]CloudPCGalleryImage `json:"galleryImages,omitempty"`

	// A defined collection of Azure resource information that can be used to establish on-premises network connectivity for
	// Cloud PCs.
	OnPremisesConnections *[]CloudPCOnPremisesConnection `json:"onPremisesConnections,omitempty"`

	// The Cloud PC organization settings for a tenant.
	OrganizationSettings *CloudPCOrganizationSettings `json:"organizationSettings,omitempty"`

	// Cloud PC provisioning policy.
	ProvisioningPolicies *[]CloudPCProvisioningPolicy `json:"provisioningPolicies,omitempty"`

	// Cloud PC related reports.
	Reports *CloudPCReports `json:"reports,omitempty"`

	// Cloud PC service plans.
	ServicePlans *[]CloudPCServicePlan `json:"servicePlans,omitempty"`

	// Cloud PC snapshots.
	Snapshots *[]CloudPCSnapshot `json:"snapshots,omitempty"`

	// Cloud PC supported regions.
	SupportedRegions *[]CloudPCSupportedRegion `json:"supportedRegions,omitempty"`

	// Cloud PC user settings.
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

var _ json.Unmarshaler = &VirtualEndpoint{}

func (s *VirtualEndpoint) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AuditEvents                             *[]CloudPCAuditEvent                            `json:"auditEvents,omitempty"`
		CloudPCs                                *[]CloudPC                                      `json:"cloudPCs,omitempty"`
		CrossCloudGovernmentOrganizationMapping *CloudPCCrossCloudGovernmentOrganizationMapping `json:"crossCloudGovernmentOrganizationMapping,omitempty"`
		DeviceImages                            *[]CloudPCDeviceImage                           `json:"deviceImages,omitempty"`
		ExternalPartnerSettings                 *[]CloudPCExternalPartnerSetting                `json:"externalPartnerSettings,omitempty"`
		FrontLineServicePlans                   *[]CloudPCFrontLineServicePlan                  `json:"frontLineServicePlans,omitempty"`
		GalleryImages                           *[]CloudPCGalleryImage                          `json:"galleryImages,omitempty"`
		OnPremisesConnections                   *[]CloudPCOnPremisesConnection                  `json:"onPremisesConnections,omitempty"`
		OrganizationSettings                    *CloudPCOrganizationSettings                    `json:"organizationSettings,omitempty"`
		ProvisioningPolicies                    *[]CloudPCProvisioningPolicy                    `json:"provisioningPolicies,omitempty"`
		Reports                                 *CloudPCReports                                 `json:"reports,omitempty"`
		ServicePlans                            *[]CloudPCServicePlan                           `json:"servicePlans,omitempty"`
		Snapshots                               *[]CloudPCSnapshot                              `json:"snapshots,omitempty"`
		SupportedRegions                        *[]CloudPCSupportedRegion                       `json:"supportedRegions,omitempty"`
		UserSettings                            *[]CloudPCUserSetting                           `json:"userSettings,omitempty"`
		Id                                      *string                                         `json:"id,omitempty"`
		ODataId                                 *string                                         `json:"@odata.id,omitempty"`
		ODataType                               *string                                         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AuditEvents = decoded.AuditEvents
	s.CloudPCs = decoded.CloudPCs
	s.CrossCloudGovernmentOrganizationMapping = decoded.CrossCloudGovernmentOrganizationMapping
	s.DeviceImages = decoded.DeviceImages
	s.ExternalPartnerSettings = decoded.ExternalPartnerSettings
	s.FrontLineServicePlans = decoded.FrontLineServicePlans
	s.GalleryImages = decoded.GalleryImages
	s.OnPremisesConnections = decoded.OnPremisesConnections
	s.OrganizationSettings = decoded.OrganizationSettings
	s.ProvisioningPolicies = decoded.ProvisioningPolicies
	s.Reports = decoded.Reports
	s.ServicePlans = decoded.ServicePlans
	s.Snapshots = decoded.Snapshots
	s.SupportedRegions = decoded.SupportedRegions
	s.UserSettings = decoded.UserSettings
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling VirtualEndpoint into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["bulkActions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling BulkActions into list []json.RawMessage: %+v", err)
		}

		output := make([]CloudPCBulkAction, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalCloudPCBulkActionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'BulkActions' for 'VirtualEndpoint': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.BulkActions = &output
	}

	return nil
}
