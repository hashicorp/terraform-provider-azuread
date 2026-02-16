package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnAttributeCollectionHandler = OnAttributeCollectionExternalUsersSelfServiceSignUp{}

type OnAttributeCollectionExternalUsersSelfServiceSignUp struct {
	// Required. The configuration for how attributes are displayed in the sign-up experience defined by a user flow, like
	// the externalUsersSelfServiceSignupEventsFlow, specifically on the attribute collection page.
	AttributeCollectionPage AuthenticationAttributeCollectionPage `json:"attributeCollectionPage"`

	Attributes *[]IdentityUserFlowAttribute `json:"attributes,omitempty"`

	// Fields inherited from OnAttributeCollectionHandler

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OnAttributeCollectionExternalUsersSelfServiceSignUp) OnAttributeCollectionHandler() BaseOnAttributeCollectionHandlerImpl {
	return BaseOnAttributeCollectionHandlerImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnAttributeCollectionExternalUsersSelfServiceSignUp{}

func (s OnAttributeCollectionExternalUsersSelfServiceSignUp) MarshalJSON() ([]byte, error) {
	type wrapper OnAttributeCollectionExternalUsersSelfServiceSignUp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnAttributeCollectionExternalUsersSelfServiceSignUp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnAttributeCollectionExternalUsersSelfServiceSignUp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onAttributeCollectionExternalUsersSelfServiceSignUp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnAttributeCollectionExternalUsersSelfServiceSignUp: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OnAttributeCollectionExternalUsersSelfServiceSignUp{}

func (s *OnAttributeCollectionExternalUsersSelfServiceSignUp) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AttributeCollectionPage AuthenticationAttributeCollectionPage `json:"attributeCollectionPage"`
		ODataId                 *string                               `json:"@odata.id,omitempty"`
		ODataType               *string                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AttributeCollectionPage = decoded.AttributeCollectionPage
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OnAttributeCollectionExternalUsersSelfServiceSignUp into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attributes"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Attributes into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentityUserFlowAttribute, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentityUserFlowAttributeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Attributes' for 'OnAttributeCollectionExternalUsersSelfServiceSignUp': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Attributes = &output
	}

	return nil
}
