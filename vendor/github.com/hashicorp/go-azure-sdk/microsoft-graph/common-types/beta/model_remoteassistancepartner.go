package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RemoteAssistancePartner{}

type RemoteAssistancePartner struct {
	// Display name of the partner.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Timestamp of the last request sent to Intune by the TEM partner.
	LastConnectionDateTime *string `json:"lastConnectionDateTime,omitempty"`

	// When the OnboardingStatus is Onboarding, This is the date time when the onboarding request expires.
	OnboardingRequestExpiryDateTime *string `json:"onboardingRequestExpiryDateTime,omitempty"`

	// The current TeamViewer connector status
	OnboardingStatus *RemoteAssistanceOnboardingStatus `json:"onboardingStatus,omitempty"`

	// URL of the partner's onboarding portal, where an administrator can configure their Remote Assistance service.
	OnboardingUrl nullable.Type[string] `json:"onboardingUrl,omitempty"`

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

func (s RemoteAssistancePartner) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RemoteAssistancePartner{}

func (s RemoteAssistancePartner) MarshalJSON() ([]byte, error) {
	type wrapper RemoteAssistancePartner
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RemoteAssistancePartner: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RemoteAssistancePartner: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.remoteAssistancePartner"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RemoteAssistancePartner: %+v", err)
	}

	return encoded, nil
}
