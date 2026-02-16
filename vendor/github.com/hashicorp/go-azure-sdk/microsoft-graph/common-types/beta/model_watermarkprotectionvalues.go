package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WatermarkProtectionValues struct {
	// Indicates whether to apply a watermark to any shared content.
	IsEnabledForContentSharing nullable.Type[bool] `json:"isEnabledForContentSharing,omitempty"`

	// Indicates whether to apply a watermark to everyone's video feed.
	IsEnabledForVideo nullable.Type[bool] `json:"isEnabledForVideo,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
