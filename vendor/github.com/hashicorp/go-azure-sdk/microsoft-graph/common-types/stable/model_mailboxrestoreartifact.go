package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxRestoreArtifact interface {
	Entity
	RestoreArtifactBase
	MailboxRestoreArtifact() BaseMailboxRestoreArtifactImpl
}

var _ MailboxRestoreArtifact = BaseMailboxRestoreArtifactImpl{}

type BaseMailboxRestoreArtifactImpl struct {
	// The new restored folder identifier for the user.
	RestoredFolderId nullable.Type[string] `json:"restoredFolderId,omitempty"`

	// The new restored folder name.
	RestoredFolderName nullable.Type[string] `json:"restoredFolderName,omitempty"`

	// Fields inherited from RestoreArtifactBase

	// The time when restoration of restore artifact is completed.
	CompletionDateTime nullable.Type[string] `json:"completionDateTime,omitempty"`

	// Indicates the restoration destination. The possible values are: new, inPlace, unknownFutureValue.
	DestinationType *DestinationType `json:"destinationType,omitempty"`

	// Contains error details if the restore session fails or completes with an error.
	Error *PublicError `json:"error,omitempty"`

	// Represents the date and time when an artifact is protected by a protectionPolicy and can be restored.
	RestorePoint *RestorePoint `json:"restorePoint,omitempty"`

	// The time when restoration of restore artifact is started.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// The individual restoration status of the restore artifact. The possible values are: added, scheduling, scheduled,
	// inProgress, succeeded, failed, unknownFutureValue.
	Status *ArtifactRestoreStatus `json:"status,omitempty"`

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

func (s BaseMailboxRestoreArtifactImpl) MailboxRestoreArtifact() BaseMailboxRestoreArtifactImpl {
	return s
}

func (s BaseMailboxRestoreArtifactImpl) RestoreArtifactBase() BaseRestoreArtifactBaseImpl {
	return BaseRestoreArtifactBaseImpl{
		CompletionDateTime: s.CompletionDateTime,
		DestinationType:    s.DestinationType,
		Error:              s.Error,
		RestorePoint:       s.RestorePoint,
		StartDateTime:      s.StartDateTime,
		Status:             s.Status,
		Id:                 s.Id,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

func (s BaseMailboxRestoreArtifactImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ MailboxRestoreArtifact = RawMailboxRestoreArtifactImpl{}

// RawMailboxRestoreArtifactImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMailboxRestoreArtifactImpl struct {
	mailboxRestoreArtifact BaseMailboxRestoreArtifactImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawMailboxRestoreArtifactImpl) MailboxRestoreArtifact() BaseMailboxRestoreArtifactImpl {
	return s.mailboxRestoreArtifact
}

func (s RawMailboxRestoreArtifactImpl) RestoreArtifactBase() BaseRestoreArtifactBaseImpl {
	return s.mailboxRestoreArtifact.RestoreArtifactBase()
}

func (s RawMailboxRestoreArtifactImpl) Entity() BaseEntityImpl {
	return s.mailboxRestoreArtifact.Entity()
}

var _ json.Marshaler = BaseMailboxRestoreArtifactImpl{}

func (s BaseMailboxRestoreArtifactImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseMailboxRestoreArtifactImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseMailboxRestoreArtifactImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseMailboxRestoreArtifactImpl: %+v", err)
	}

	delete(decoded, "restoredFolderName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mailboxRestoreArtifact"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseMailboxRestoreArtifactImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalMailboxRestoreArtifactImplementation(input []byte) (MailboxRestoreArtifact, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MailboxRestoreArtifact into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.granularMailboxRestoreArtifact") {
		var out GranularMailboxRestoreArtifact
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GranularMailboxRestoreArtifact: %+v", err)
		}
		return out, nil
	}

	var parent BaseMailboxRestoreArtifactImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMailboxRestoreArtifactImpl: %+v", err)
	}

	return RawMailboxRestoreArtifactImpl{
		mailboxRestoreArtifact: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
