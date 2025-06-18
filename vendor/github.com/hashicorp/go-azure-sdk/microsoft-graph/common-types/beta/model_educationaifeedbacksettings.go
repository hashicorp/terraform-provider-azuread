package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAiFeedbackSettings struct {
	// The audience engagement related feedback types that students should receive from the AI feedback.
	AudienceEngagementSettings *EducationAiFeedbackAudienceEngagementSettings `json:"audienceEngagementSettings,omitempty"`

	// The content related feedback types that students should receive from the AI feedback.
	ContentSettings *EducationAiFeedbackContentSettings `json:"contentSettings,omitempty"`

	// The delivery related feedback types that students should receive from the AI feedback.
	DeliverySettings *EducationAiFeedbackDeliverySettings `json:"deliverySettings,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
