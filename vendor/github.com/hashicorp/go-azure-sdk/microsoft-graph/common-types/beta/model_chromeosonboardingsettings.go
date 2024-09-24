package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ChromeOSOnboardingSettings{}

type ChromeOSOnboardingSettings struct {
	// The ChromebookTenant's LastDirectorySyncDateTime
	LastDirectorySyncDateTime nullable.Type[string] `json:"lastDirectorySyncDateTime,omitempty"`

	// The ChromebookTenant's LastModifiedDateTime
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The onboarding status of the tenant.
	OnboardingStatus *OnboardingStatus `json:"onboardingStatus,omitempty"`

	// The ChromebookTenant's OwnerUserPrincipalName
	OwnerUserPrincipalName nullable.Type[string] `json:"ownerUserPrincipalName,omitempty"`

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

func (s ChromeOSOnboardingSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChromeOSOnboardingSettings{}

func (s ChromeOSOnboardingSettings) MarshalJSON() ([]byte, error) {
	type wrapper ChromeOSOnboardingSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChromeOSOnboardingSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChromeOSOnboardingSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.chromeOSOnboardingSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChromeOSOnboardingSettings: %+v", err)
	}

	return encoded, nil
}
