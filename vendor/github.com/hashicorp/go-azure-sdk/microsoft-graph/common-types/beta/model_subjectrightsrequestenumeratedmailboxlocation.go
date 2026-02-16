package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SubjectRightsRequestMailboxLocation = SubjectRightsRequestEnumeratedMailboxLocation{}

type SubjectRightsRequestEnumeratedMailboxLocation struct {
	// Collection of mailboxes that should be included in the search. Includes the UPN of each mailbox, for example,
	// Monica.Thompson@contoso.com. Going forward, use the userPrincipalNames property. If you specify either upns or
	// userPrincipalNames, the same values are populated automatically to the other property.
	Upns *[]string `json:"upns,omitempty"`

	// Collection of mailboxes that should be included in the search. Includes the user principal name (UPN) of each
	// mailbox, for example, Monica.Thompson@contoso.com.
	UserPrincipalNames *[]string `json:"userPrincipalNames,omitempty"`

	// Fields inherited from SubjectRightsRequestMailboxLocation

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SubjectRightsRequestEnumeratedMailboxLocation) SubjectRightsRequestMailboxLocation() BaseSubjectRightsRequestMailboxLocationImpl {
	return BaseSubjectRightsRequestMailboxLocationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SubjectRightsRequestEnumeratedMailboxLocation{}

func (s SubjectRightsRequestEnumeratedMailboxLocation) MarshalJSON() ([]byte, error) {
	type wrapper SubjectRightsRequestEnumeratedMailboxLocation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SubjectRightsRequestEnumeratedMailboxLocation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SubjectRightsRequestEnumeratedMailboxLocation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.subjectRightsRequestEnumeratedMailboxLocation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SubjectRightsRequestEnumeratedMailboxLocation: %+v", err)
	}

	return encoded, nil
}
