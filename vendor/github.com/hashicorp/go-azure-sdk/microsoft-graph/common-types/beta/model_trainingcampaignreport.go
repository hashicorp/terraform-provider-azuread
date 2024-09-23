package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingCampaignReport struct {
	// The overview of the attack simulation and training campaign.
	CampaignUsers *[]UserSimulationDetails `json:"campaignUsers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Overview *TrainingCampaignReportOverview `json:"overview,omitempty"`
}
