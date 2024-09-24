package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsSettings struct {
	// When TRUE, indicates Tenant attach is configured properly and System Center Configuration Manager (SCCM) tenant
	// attached devices will show up in endpoint analytics reporting. When FALSE, indicates Tenant attach is not configured.
	// FALSE by default.
	ConfigurationManagerDataConnectorConfigured *bool `json:"configurationManagerDataConnectorConfigured,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
