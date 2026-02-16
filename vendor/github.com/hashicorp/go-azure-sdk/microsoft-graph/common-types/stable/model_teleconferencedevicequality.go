package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeleconferenceDeviceQuality struct {
	// A unique identifier for all the participant calls in a conference or a unique identifier for two participant calls in
	// P2P call. This needs to be copied over from Microsoft.Graph.Call.CallChainId.
	CallChainId *string `json:"callChainId,omitempty"`

	// A geo-region where the service is deployed, such as ProdNoam.
	CloudServiceDeploymentEnvironment nullable.Type[string] `json:"cloudServiceDeploymentEnvironment,omitempty"`

	// A unique deployment identifier assigned by Azure.
	CloudServiceDeploymentId nullable.Type[string] `json:"cloudServiceDeploymentId,omitempty"`

	// The Azure deployed cloud service instance name, such as FrontEndIN3.
	CloudServiceInstanceName nullable.Type[string] `json:"cloudServiceInstanceName,omitempty"`

	// The Azure deployed cloud service name, such as contoso.cloudapp.net.
	CloudServiceName nullable.Type[string] `json:"cloudServiceName,omitempty"`

	// Any additional description, such as VTC Bldg 30/21.
	DeviceDescription *string `json:"deviceDescription,omitempty"`

	// The user media agent name, such as Cisco SX80.
	DeviceName *string `json:"deviceName,omitempty"`

	// A unique identifier for a specific media leg of a participant in a conference. One participant can have multiple
	// media leg identifiers if retargeting happens. CVI partner assigns this value.
	MediaLegId *string `json:"mediaLegId,omitempty"`

	// The list of media qualities in a media session (call), such as audio quality, video quality, and/or screen sharing
	// quality.
	MediaQualityList *[]TeleconferenceDeviceMediaQuality `json:"mediaQualityList,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A unique identifier for a specific participant in a conference. The CVI partner needs to copy over
	// Call.MyParticipantId to this property.
	ParticipantId *string `json:"participantId,omitempty"`
}

var _ json.Unmarshaler = &TeleconferenceDeviceQuality{}

func (s *TeleconferenceDeviceQuality) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CallChainId                       *string               `json:"callChainId,omitempty"`
		CloudServiceDeploymentEnvironment nullable.Type[string] `json:"cloudServiceDeploymentEnvironment,omitempty"`
		CloudServiceDeploymentId          nullable.Type[string] `json:"cloudServiceDeploymentId,omitempty"`
		CloudServiceInstanceName          nullable.Type[string] `json:"cloudServiceInstanceName,omitempty"`
		CloudServiceName                  nullable.Type[string] `json:"cloudServiceName,omitempty"`
		DeviceDescription                 *string               `json:"deviceDescription,omitempty"`
		DeviceName                        *string               `json:"deviceName,omitempty"`
		MediaLegId                        *string               `json:"mediaLegId,omitempty"`
		ODataId                           *string               `json:"@odata.id,omitempty"`
		ODataType                         *string               `json:"@odata.type,omitempty"`
		ParticipantId                     *string               `json:"participantId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CallChainId = decoded.CallChainId
	s.CloudServiceDeploymentEnvironment = decoded.CloudServiceDeploymentEnvironment
	s.CloudServiceDeploymentId = decoded.CloudServiceDeploymentId
	s.CloudServiceInstanceName = decoded.CloudServiceInstanceName
	s.CloudServiceName = decoded.CloudServiceName
	s.DeviceDescription = decoded.DeviceDescription
	s.DeviceName = decoded.DeviceName
	s.MediaLegId = decoded.MediaLegId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ParticipantId = decoded.ParticipantId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeleconferenceDeviceQuality into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["mediaQualityList"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MediaQualityList into list []json.RawMessage: %+v", err)
		}

		output := make([]TeleconferenceDeviceMediaQuality, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalTeleconferenceDeviceMediaQualityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MediaQualityList' for 'TeleconferenceDeviceQuality': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MediaQualityList = &output
	}

	return nil
}
