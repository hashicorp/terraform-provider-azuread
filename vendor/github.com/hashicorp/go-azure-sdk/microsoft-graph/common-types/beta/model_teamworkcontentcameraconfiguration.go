package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkContentCameraConfiguration struct {
	// True if the content camera is inverted.
	IsContentCameraInverted nullable.Type[bool] `json:"isContentCameraInverted,omitempty"`

	// True if the content camera is optional.
	IsContentCameraOptional nullable.Type[bool] `json:"isContentCameraOptional,omitempty"`

	// True if the content enhancement is enabled.
	IsContentEnhancementEnabled nullable.Type[bool] `json:"isContentEnhancementEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
