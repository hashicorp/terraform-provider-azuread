package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MailFolderOperation = UpdateAllMessagesReadStateOperation{}

type UpdateAllMessagesReadStateOperation struct {

	// Fields inherited from MailFolderOperation

	// The location of the long-running operation.
	ResourceLocation nullable.Type[string] `json:"resourceLocation,omitempty"`

	// The status of the long-running operation. The possible values are: notStarted, running, succeeded, failed,
	// unknownFutureValue.
	Status *MailFolderOperationStatus `json:"status,omitempty"`

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

func (s UpdateAllMessagesReadStateOperation) MailFolderOperation() BaseMailFolderOperationImpl {
	return BaseMailFolderOperationImpl{
		ResourceLocation: s.ResourceLocation,
		Status:           s.Status,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s UpdateAllMessagesReadStateOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UpdateAllMessagesReadStateOperation{}

func (s UpdateAllMessagesReadStateOperation) MarshalJSON() ([]byte, error) {
	type wrapper UpdateAllMessagesReadStateOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UpdateAllMessagesReadStateOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UpdateAllMessagesReadStateOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.updateAllMessagesReadStateOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UpdateAllMessagesReadStateOperation: %+v", err)
	}

	return encoded, nil
}
