package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSpeakerCoachSettings struct {
	// The audience engagement related feedback types that students should receive from the Speaker Coach.
	AudienceEngagementSettings *EducationSpeakerCoachAudienceEngagementSettings `json:"audienceEngagementSettings,omitempty"`

	// The content related feedback types that students should receive from the Speaker Coach.
	ContentSettings *EducationSpeakerCoachContentSettings `json:"contentSettings,omitempty"`

	// The delivery related feedback types that students should receive from the Speaker Coach.
	DeliverySettings *EducationSpeakerCoachDeliverySettings `json:"deliverySettings,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
