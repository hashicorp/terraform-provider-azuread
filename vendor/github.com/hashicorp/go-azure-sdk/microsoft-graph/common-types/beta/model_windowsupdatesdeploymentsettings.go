package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentSettings struct {
	// Settings for governing whether content is applicable to a device.
	ContentApplicability *WindowsUpdatesContentApplicabilitySettings `json:"contentApplicability,omitempty"`

	// Settings for governing whether updates should be expedited.
	Expedite *WindowsUpdatesExpediteSettings `json:"expedite,omitempty"`

	// Settings for governing conditions to monitor and automated actions to take.
	Monitoring *WindowsUpdatesMonitoringSettings `json:"monitoring,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Settings for governing how and when the content is rolled out.
	Schedule *WindowsUpdatesScheduleSettings `json:"schedule,omitempty"`

	// Settings for governing end user update experience.
	UserExperience *WindowsUpdatesUserExperienceSettings `json:"userExperience,omitempty"`
}
