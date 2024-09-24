package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FormsSettings struct {
	// Controls whether users can add images from Bing to forms.
	IsBingImageSearchEnabled nullable.Type[bool] `json:"isBingImageSearchEnabled,omitempty"`

	// Controls whether users can send a link to a form to an external user.
	IsExternalSendFormEnabled nullable.Type[bool] `json:"isExternalSendFormEnabled,omitempty"`

	// Controls whether users can collaborate on a form layout and structure with an external user.
	IsExternalShareCollaborationEnabled nullable.Type[bool] `json:"isExternalShareCollaborationEnabled,omitempty"`

	// Controls whether users can share form results with external users.
	IsExternalShareResultEnabled nullable.Type[bool] `json:"isExternalShareResultEnabled,omitempty"`

	// Controls whether users can share form templates with external users.
	IsExternalShareTemplateEnabled nullable.Type[bool] `json:"isExternalShareTemplateEnabled,omitempty"`

	// Controls whether phishing protection is run on forms created by users, blocking the creation of forms if common
	// phishing questions are detected.
	IsInOrgFormsPhishingScanEnabled nullable.Type[bool] `json:"isInOrgFormsPhishingScanEnabled,omitempty"`

	// Controls whether the names of users who fill out forms are recorded.
	IsRecordIdentityByDefaultEnabled nullable.Type[bool] `json:"isRecordIdentityByDefaultEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
