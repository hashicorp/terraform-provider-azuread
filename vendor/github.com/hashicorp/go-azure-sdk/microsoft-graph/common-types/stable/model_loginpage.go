package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = LoginPage{}

type LoginPage struct {
	// The HTML content of the login page.
	Content nullable.Type[string] `json:"content,omitempty"`

	// Identity of the user who created the login page.
	CreatedBy *EmailIdentity `json:"createdBy,omitempty"`

	// Date and time when the login page was created. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description about the login page.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name of the login page.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The content language of the login page.
	Language nullable.Type[string] `json:"language,omitempty"`

	// Identity of the user who last modified the login page.
	LastModifiedBy *EmailIdentity `json:"lastModifiedBy,omitempty"`

	// Date and time when the login page was last modified. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The source of the content. Possible values are: unknown, global, tenant, unknownFutureValue.
	Source *SimulationContentSource `json:"source,omitempty"`

	// The login page status. Possible values are: unknown, draft, ready, archive, delete, unknownFutureValue.
	Status *SimulationContentStatus `json:"status,omitempty"`

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

func (s LoginPage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = LoginPage{}

func (s LoginPage) MarshalJSON() ([]byte, error) {
	type wrapper LoginPage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LoginPage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LoginPage: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.loginPage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LoginPage: %+v", err)
	}

	return encoded, nil
}
