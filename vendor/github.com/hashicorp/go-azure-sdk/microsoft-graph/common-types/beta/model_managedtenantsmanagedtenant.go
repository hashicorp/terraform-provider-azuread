package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagedTenant{}

type ManagedTenantsManagedTenant struct {
	// Aggregate view of device compliance policies across managed tenants.
	AggregatedPolicyCompliances *[]ManagedTenantsAggregatedPolicyCompliance `json:"aggregatedPolicyCompliances,omitempty"`

	AppPerformances *[]ManagedTenantsAppPerformance `json:"appPerformances,omitempty"`

	// The collection of audit events across managed tenants.
	AuditEvents *[]ManagedTenantsAuditEvent `json:"auditEvents,omitempty"`

	// The collection of cloud PC connections across managed tenants.
	CloudPCConnections *[]ManagedTenantsCloudPCConnection `json:"cloudPcConnections,omitempty"`

	// The collection of cloud PC devices across managed tenants.
	CloudPCDevices *[]ManagedTenantsCloudPCDevice `json:"cloudPcDevices,omitempty"`

	// Overview of cloud PC information across managed tenants.
	CloudPCsOverview *[]ManagedTenantsCloudPCOverview `json:"cloudPcsOverview,omitempty"`

	// Aggregate view of conditional access policy coverage across managed tenants.
	ConditionalAccessPolicyCoverages *[]ManagedTenantsConditionalAccessPolicyCoverage `json:"conditionalAccessPolicyCoverages,omitempty"`

	// Summary information for user registration for multi-factor authentication and self service password reset across
	// managed tenants.
	CredentialUserRegistrationsSummaries *[]ManagedTenantsCredentialUserRegistrationsSummary `json:"credentialUserRegistrationsSummaries,omitempty"`

	DeviceAppPerformances *[]ManagedTenantsDeviceAppPerformance `json:"deviceAppPerformances,omitempty"`

	// Summary information for device compliance policy setting states across managed tenants.
	DeviceCompliancePolicySettingStateSummaries *[]ManagedTenantsDeviceCompliancePolicySettingStateSummary `json:"deviceCompliancePolicySettingStateSummaries,omitempty"`

	DeviceHealthStatuses *[]ManagedTenantsDeviceHealthStatus `json:"deviceHealthStatuses,omitempty"`

	// Trend insights for device compliance across managed tenants.
	ManagedDeviceComplianceTrends *[]ManagedTenantsManagedDeviceComplianceTrend `json:"managedDeviceComplianceTrends,omitempty"`

	// The collection of compliance for managed devices across managed tenants.
	ManagedDeviceCompliances *[]ManagedTenantsManagedDeviceCompliance `json:"managedDeviceCompliances,omitempty"`

	ManagedTenantAlertLogs            *[]ManagedTenantsManagedTenantAlertLog            `json:"managedTenantAlertLogs,omitempty"`
	ManagedTenantAlertRuleDefinitions *[]ManagedTenantsManagedTenantAlertRuleDefinition `json:"managedTenantAlertRuleDefinitions,omitempty"`
	ManagedTenantAlertRules           *[]ManagedTenantsManagedTenantAlertRule           `json:"managedTenantAlertRules,omitempty"`
	ManagedTenantAlerts               *[]ManagedTenantsManagedTenantAlert               `json:"managedTenantAlerts,omitempty"`
	ManagedTenantApiNotifications     *[]ManagedTenantsManagedTenantApiNotification     `json:"managedTenantApiNotifications,omitempty"`
	ManagedTenantEmailNotifications   *[]ManagedTenantsManagedTenantEmailNotification   `json:"managedTenantEmailNotifications,omitempty"`
	ManagedTenantTicketingEndpoints   *[]ManagedTenantsManagedTenantTicketingEndpoint   `json:"managedTenantTicketingEndpoints,omitempty"`

	// The tenant level status of management actions across managed tenants.
	ManagementActionTenantDeploymentStatuses *[]ManagedTenantsManagementActionTenantDeploymentStatus `json:"managementActionTenantDeploymentStatuses,omitempty"`

	// The collection of baseline management actions across managed tenants.
	ManagementActions *[]ManagedTenantsManagementAction `json:"managementActions,omitempty"`

	// The collection of baseline management intents across managed tenants.
	ManagementIntents *[]ManagedTenantsManagementIntent `json:"managementIntents,omitempty"`

	ManagementTemplateCollectionTenantSummaries *[]ManagedTenantsManagementTemplateCollectionTenantSummary `json:"managementTemplateCollectionTenantSummaries,omitempty"`
	ManagementTemplateCollections               *[]ManagedTenantsManagementTemplateCollection              `json:"managementTemplateCollections,omitempty"`
	ManagementTemplateStepTenantSummaries       *[]ManagedTenantsManagementTemplateStepTenantSummary       `json:"managementTemplateStepTenantSummaries,omitempty"`
	ManagementTemplateStepVersions              *[]ManagedTenantsManagementTemplateStepVersion             `json:"managementTemplateStepVersions,omitempty"`
	ManagementTemplateSteps                     *[]ManagedTenantsManagementTemplateStep                    `json:"managementTemplateSteps,omitempty"`

	// The collection of baseline management templates across managed tenants.
	ManagementTemplates *[]ManagedTenantsManagementTemplate `json:"managementTemplates,omitempty"`

	// The collection of role assignments to a signed-in user for a managed tenant.
	MyRoles *[]ManagedTenantsMyRole `json:"myRoles,omitempty"`

	// The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
	TenantGroups *[]ManagedTenantsTenantGroup `json:"tenantGroups,omitempty"`

	// The collection of tenant tags across managed tenants.
	TenantTags *[]ManagedTenantsTenantTag `json:"tenantTags,omitempty"`

	// The collection of tenants associated with the managing entity.
	Tenants *[]ManagedTenantsTenant `json:"tenants,omitempty"`

	// The collection of tenant level customized information across managed tenants.
	TenantsCustomizedInformation *[]ManagedTenantsTenantCustomizedInformation `json:"tenantsCustomizedInformation,omitempty"`

	// The collection tenant level detailed information across managed tenants.
	TenantsDetailedInformation *[]ManagedTenantsTenantDetailedInformation `json:"tenantsDetailedInformation,omitempty"`

	// The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
	WindowsDeviceMalwareStates *[]ManagedTenantsWindowsDeviceMalwareState `json:"windowsDeviceMalwareStates,omitempty"`

	// The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
	WindowsProtectionStates *[]ManagedTenantsWindowsProtectionState `json:"windowsProtectionStates,omitempty"`

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

func (s ManagedTenantsManagedTenant) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagedTenant{}

func (s ManagedTenantsManagedTenant) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagedTenant
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagedTenant: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedTenant: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managedTenant"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagedTenant: %+v", err)
	}

	return encoded, nil
}
