package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkPeripheralsHealth struct {
	// The health details about the communication speaker.
	CommunicationSpeakerHealth *TeamworkPeripheralHealth `json:"communicationSpeakerHealth,omitempty"`

	// The health details about the content camera.
	ContentCameraHealth *TeamworkPeripheralHealth `json:"contentCameraHealth,omitempty"`

	// The health details about displays.
	DisplayHealthCollection *[]TeamworkPeripheralHealth `json:"displayHealthCollection,omitempty"`

	// The health details about the microphone.
	MicrophoneHealth *TeamworkPeripheralHealth `json:"microphoneHealth,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The health details about the room camera.
	RoomCameraHealth *TeamworkPeripheralHealth `json:"roomCameraHealth,omitempty"`

	// The health details about the speaker.
	SpeakerHealth *TeamworkPeripheralHealth `json:"speakerHealth,omitempty"`
}
