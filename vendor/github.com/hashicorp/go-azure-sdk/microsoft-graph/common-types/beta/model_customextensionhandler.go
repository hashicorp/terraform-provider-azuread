package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CustomExtensionHandler{}

type CustomExtensionHandler struct {
	// Indicates which custom workflow extension is executed at this stage. Nullable. Supports $expand.
	CustomExtension *CustomAccessPackageWorkflowExtension `json:"customExtension,omitempty"`

	// Indicates the stage of the access package assignment request workflow when the access package custom extension runs.
	// The possible values are: assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted,
	// assignmentRequestRemoved, assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration,
	// unknownFutureValue.
	Stage *AccessPackageCustomExtensionStage `json:"stage,omitempty"`

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

func (s CustomExtensionHandler) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CustomExtensionHandler{}

func (s CustomExtensionHandler) MarshalJSON() ([]byte, error) {
	type wrapper CustomExtensionHandler
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomExtensionHandler: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomExtensionHandler: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customExtensionHandler"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomExtensionHandler: %+v", err)
	}

	return encoded, nil
}
