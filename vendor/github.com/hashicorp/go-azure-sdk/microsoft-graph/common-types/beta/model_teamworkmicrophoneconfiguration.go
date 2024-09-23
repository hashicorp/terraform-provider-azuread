package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkMicrophoneConfiguration struct {
	DefaultMicrophone *TeamworkPeripheral `json:"defaultMicrophone,omitempty"`

	// True if the configured microphone is optional. False if the microphone is not optional and the health state of the
	// device should be computed.
	IsMicrophoneOptional nullable.Type[bool] `json:"isMicrophoneOptional,omitempty"`

	Microphones *[]TeamworkPeripheral `json:"microphones,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
