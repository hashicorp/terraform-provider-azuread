package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TrainingLanguageDetail{}

type TrainingLanguageDetail struct {
	// Language specific content for the training.
	Content *string `json:"content,omitempty"`

	// Identity of the user who created the language details.
	CreatedBy *EmailIdentity `json:"createdBy,omitempty"`

	// Date and time when the language details were created. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description as defined by the user.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name as defined by the user.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether the training has a default language.
	IsDefaultLangauge nullable.Type[bool] `json:"isDefaultLangauge,omitempty"`

	// Identity of the user who last modified the details.
	LastModifiedBy *EmailIdentity `json:"lastModifiedBy,omitempty"`

	// Date and time when the trainingLanguageDetail was last modified. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Content locale for the training detail.
	Locale *string `json:"locale,omitempty"`

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

func (s TrainingLanguageDetail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TrainingLanguageDetail{}

func (s TrainingLanguageDetail) MarshalJSON() ([]byte, error) {
	type wrapper TrainingLanguageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TrainingLanguageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TrainingLanguageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.trainingLanguageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TrainingLanguageDetail: %+v", err)
	}

	return encoded, nil
}
