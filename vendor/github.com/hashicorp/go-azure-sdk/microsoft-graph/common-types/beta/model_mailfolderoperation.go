package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderOperation interface {
	Entity
	MailFolderOperation() BaseMailFolderOperationImpl
}

var _ MailFolderOperation = BaseMailFolderOperationImpl{}

type BaseMailFolderOperationImpl struct {
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

func (s BaseMailFolderOperationImpl) MailFolderOperation() BaseMailFolderOperationImpl {
	return s
}

func (s BaseMailFolderOperationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ MailFolderOperation = RawMailFolderOperationImpl{}

// RawMailFolderOperationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMailFolderOperationImpl struct {
	mailFolderOperation BaseMailFolderOperationImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawMailFolderOperationImpl) MailFolderOperation() BaseMailFolderOperationImpl {
	return s.mailFolderOperation
}

func (s RawMailFolderOperationImpl) Entity() BaseEntityImpl {
	return s.mailFolderOperation.Entity()
}

var _ json.Marshaler = BaseMailFolderOperationImpl{}

func (s BaseMailFolderOperationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseMailFolderOperationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseMailFolderOperationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseMailFolderOperationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mailFolderOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseMailFolderOperationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalMailFolderOperationImplementation(input []byte) (MailFolderOperation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MailFolderOperation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.updateAllMessagesReadStateOperation") {
		var out UpdateAllMessagesReadStateOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UpdateAllMessagesReadStateOperation: %+v", err)
		}
		return out, nil
	}

	var parent BaseMailFolderOperationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMailFolderOperationImpl: %+v", err)
	}

	return RawMailFolderOperationImpl{
		mailFolderOperation: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
