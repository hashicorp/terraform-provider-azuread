package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomerVoiceSettings struct {
	// Controls whether phishing protection is run on forms created by users, blocking the creation of forms if common
	// phishing questions are detected.
	IsInOrgFormsPhishingScanEnabled nullable.Type[bool] `json:"isInOrgFormsPhishingScanEnabled,omitempty"`

	// Controls whether the names of users who fill out forms are recorded.
	IsRecordIdentityByDefaultEnabled nullable.Type[bool] `json:"isRecordIdentityByDefaultEnabled,omitempty"`

	// Controls whether only users inside your organization can submit a response.
	IsRestrictedSurveyAccessEnabled nullable.Type[bool] `json:"isRestrictedSurveyAccessEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
