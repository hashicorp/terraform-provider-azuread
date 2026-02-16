package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentitySet = CommunicationsIdentitySet{}

type CommunicationsIdentitySet struct {
	// The application instance associated with this action.
	ApplicationInstance Identity `json:"applicationInstance"`

	// An identity the participant would like to present itself as to the other participants in the call.
	AssertedIdentity Identity `json:"assertedIdentity"`

	// The Azure Communication Services user associated with this action.
	AzureCommunicationServicesUser Identity `json:"azureCommunicationServicesUser"`

	// The encrypted user associated with this action.
	Encrypted Identity `json:"encrypted"`

	// Type of endpoint that the participant uses. Possible values are: default, voicemail, skypeForBusiness,
	// skypeForBusinessVoipPhone, unknownFutureValue.
	EndpointType *EndpointType `json:"endpointType,omitempty"`

	// The guest user associated with this action.
	Guest Identity `json:"guest"`

	// The Skype for Business on-premises user associated with this action.
	OnPremises Identity `json:"onPremises"`

	// The phone user associated with this action.
	Phone Identity `json:"phone"`

	// Fields inherited from IdentitySet

	// Optional. The application associated with this action.
	Application Identity `json:"application"`

	// Optional. The device associated with this action.
	Device Identity `json:"device"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Optional. The user associated with this action.
	User Identity `json:"user"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CommunicationsIdentitySet) IdentitySet() BaseIdentitySetImpl {
	return BaseIdentitySetImpl{
		Application: s.Application,
		Device:      s.Device,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		User:        s.User,
	}
}

var _ json.Marshaler = CommunicationsIdentitySet{}

func (s CommunicationsIdentitySet) MarshalJSON() ([]byte, error) {
	type wrapper CommunicationsIdentitySet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CommunicationsIdentitySet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CommunicationsIdentitySet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.communicationsIdentitySet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CommunicationsIdentitySet: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CommunicationsIdentitySet{}

func (s *CommunicationsIdentitySet) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EndpointType *EndpointType `json:"endpointType,omitempty"`
		ODataId      *string       `json:"@odata.id,omitempty"`
		ODataType    *string       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EndpointType = decoded.EndpointType
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CommunicationsIdentitySet into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["application"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Application' for 'CommunicationsIdentitySet': %+v", err)
		}
		s.Application = impl
	}

	if v, ok := temp["applicationInstance"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ApplicationInstance' for 'CommunicationsIdentitySet': %+v", err)
		}
		s.ApplicationInstance = impl
	}

	if v, ok := temp["assertedIdentity"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AssertedIdentity' for 'CommunicationsIdentitySet': %+v", err)
		}
		s.AssertedIdentity = impl
	}

	if v, ok := temp["azureCommunicationServicesUser"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AzureCommunicationServicesUser' for 'CommunicationsIdentitySet': %+v", err)
		}
		s.AzureCommunicationServicesUser = impl
	}

	if v, ok := temp["device"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Device' for 'CommunicationsIdentitySet': %+v", err)
		}
		s.Device = impl
	}

	if v, ok := temp["encrypted"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Encrypted' for 'CommunicationsIdentitySet': %+v", err)
		}
		s.Encrypted = impl
	}

	if v, ok := temp["guest"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Guest' for 'CommunicationsIdentitySet': %+v", err)
		}
		s.Guest = impl
	}

	if v, ok := temp["onPremises"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'OnPremises' for 'CommunicationsIdentitySet': %+v", err)
		}
		s.OnPremises = impl
	}

	if v, ok := temp["phone"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Phone' for 'CommunicationsIdentitySet': %+v", err)
		}
		s.Phone = impl
	}

	if v, ok := temp["user"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'User' for 'CommunicationsIdentitySet': %+v", err)
		}
		s.User = impl
	}

	return nil
}
