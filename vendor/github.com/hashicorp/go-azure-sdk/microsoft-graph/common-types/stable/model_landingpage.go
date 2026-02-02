package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = LandingPage{}

type LandingPage struct {
	// Identity of the user who created the landing page.
	CreatedBy *EmailIdentity `json:"createdBy,omitempty"`

	// Date and time when the landing page was created. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description of the landing page as defined by the user.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The detail information for a landing page associated with a simulation during its creation.
	Details *[]LandingPageDetail `json:"details,omitempty"`

	// The display name of the landing page.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Email identity of the user who last modified the landing page.
	LastModifiedBy *EmailIdentity `json:"lastModifiedBy,omitempty"`

	// Date and time when the landing page was last modified. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Content locale.
	Locale nullable.Type[string] `json:"locale,omitempty"`

	// The source of the content. Possible values are: unknown, global, tenant, unknownFutureValue.
	Source *SimulationContentSource `json:"source,omitempty"`

	// The status of the simulation. Possible values are: unknown, draft, ready, archive, delete, unknownFutureValue.
	Status *SimulationContentStatus `json:"status,omitempty"`

	// Supported locales.
	SupportedLocales *[]string `json:"supportedLocales,omitempty"`

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

func (s LandingPage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = LandingPage{}

func (s LandingPage) MarshalJSON() ([]byte, error) {
	type wrapper LandingPage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LandingPage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LandingPage: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.landingPage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LandingPage: %+v", err)
	}

	return encoded, nil
}
