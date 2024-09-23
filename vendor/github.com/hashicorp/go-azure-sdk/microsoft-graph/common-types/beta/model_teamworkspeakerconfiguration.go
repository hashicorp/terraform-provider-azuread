package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSpeakerConfiguration struct {
	DefaultCommunicationSpeaker *TeamworkPeripheral `json:"defaultCommunicationSpeaker,omitempty"`
	DefaultSpeaker              *TeamworkPeripheral `json:"defaultSpeaker,omitempty"`

	// True if the communication speaker is optional. Used to compute the health state if the communication speaker is not
	// optional.
	IsCommunicationSpeakerOptional nullable.Type[bool] `json:"isCommunicationSpeakerOptional,omitempty"`

	// True if the configured speaker is optional. Used to compute the health state if the speaker is not optional.
	IsSpeakerOptional nullable.Type[bool] `json:"isSpeakerOptional,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Speakers *[]TeamworkPeripheral `json:"speakers,omitempty"`
}
