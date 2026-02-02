package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamworkDeviceConfiguration{}

type TeamworkDeviceConfiguration struct {
	// The camera configuration. Applicable only for Microsoft Teams Rooms-enabled devices.
	CameraConfiguration *TeamworkCameraConfiguration `json:"cameraConfiguration,omitempty"`

	// Identity of the user who created the device configuration document.
	CreatedBy IdentitySet `json:"createdBy"`

	// The UTC date and time when the device configuration document was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The display configuration.
	DisplayConfiguration *TeamworkDisplayConfiguration `json:"displayConfiguration,omitempty"`

	// The hardware configuration. Applicable only for Teams Rooms-enabled devices.
	HardwareConfiguration *TeamworkHardwareConfiguration `json:"hardwareConfiguration,omitempty"`

	// Identity of the user who last modified the device configuration.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The UTC date and time when the device configuration was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The microphone configuration. Applicable only for Teams Rooms-enabled devices.
	MicrophoneConfiguration *TeamworkMicrophoneConfiguration `json:"microphoneConfiguration,omitempty"`

	// Information related to software versions for the device, such as firmware, operating system, Teams client, and admin
	// agent.
	SoftwareVersions *TeamworkDeviceSoftwareVersions `json:"softwareVersions,omitempty"`

	// The speaker configuration. Applicable only for Teams Rooms-enabled devices.
	SpeakerConfiguration *TeamworkSpeakerConfiguration `json:"speakerConfiguration,omitempty"`

	// The system configuration. Not applicable for Teams Rooms-enabled devices.
	SystemConfiguration *TeamworkSystemConfiguration `json:"systemConfiguration,omitempty"`

	// The Teams client configuration. Applicable only for Teams Rooms-enabled devices.
	TeamsClientConfiguration *TeamworkTeamsClientConfiguration `json:"teamsClientConfiguration,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TeamworkDeviceConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamworkDeviceConfiguration{}

func (s TeamworkDeviceConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper TeamworkDeviceConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamworkDeviceConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamworkDeviceConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamworkDeviceConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamworkDeviceConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamworkDeviceConfiguration{}

func (s *TeamworkDeviceConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CameraConfiguration      *TeamworkCameraConfiguration      `json:"cameraConfiguration,omitempty"`
		CreatedDateTime          nullable.Type[string]             `json:"createdDateTime,omitempty"`
		DisplayConfiguration     *TeamworkDisplayConfiguration     `json:"displayConfiguration,omitempty"`
		HardwareConfiguration    *TeamworkHardwareConfiguration    `json:"hardwareConfiguration,omitempty"`
		LastModifiedDateTime     nullable.Type[string]             `json:"lastModifiedDateTime,omitempty"`
		MicrophoneConfiguration  *TeamworkMicrophoneConfiguration  `json:"microphoneConfiguration,omitempty"`
		SoftwareVersions         *TeamworkDeviceSoftwareVersions   `json:"softwareVersions,omitempty"`
		SpeakerConfiguration     *TeamworkSpeakerConfiguration     `json:"speakerConfiguration,omitempty"`
		SystemConfiguration      *TeamworkSystemConfiguration      `json:"systemConfiguration,omitempty"`
		TeamsClientConfiguration *TeamworkTeamsClientConfiguration `json:"teamsClientConfiguration,omitempty"`
		Id                       *string                           `json:"id,omitempty"`
		ODataId                  *string                           `json:"@odata.id,omitempty"`
		ODataType                *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CameraConfiguration = decoded.CameraConfiguration
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayConfiguration = decoded.DisplayConfiguration
	s.HardwareConfiguration = decoded.HardwareConfiguration
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.MicrophoneConfiguration = decoded.MicrophoneConfiguration
	s.SoftwareVersions = decoded.SoftwareVersions
	s.SpeakerConfiguration = decoded.SpeakerConfiguration
	s.SystemConfiguration = decoded.SystemConfiguration
	s.TeamsClientConfiguration = decoded.TeamsClientConfiguration
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamworkDeviceConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'TeamworkDeviceConfiguration': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'TeamworkDeviceConfiguration': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
