package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreArtifactBase interface {
	Entity
	RestoreArtifactBase() BaseRestoreArtifactBaseImpl
}

var _ RestoreArtifactBase = BaseRestoreArtifactBaseImpl{}

type BaseRestoreArtifactBaseImpl struct {
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

func (s BaseRestoreArtifactBaseImpl) RestoreArtifactBase() BaseRestoreArtifactBaseImpl {
	return s
}

func (s BaseRestoreArtifactBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ RestoreArtifactBase = RawRestoreArtifactBaseImpl{}

// RawRestoreArtifactBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRestoreArtifactBaseImpl struct {
	restoreArtifactBase BaseRestoreArtifactBaseImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawRestoreArtifactBaseImpl) RestoreArtifactBase() BaseRestoreArtifactBaseImpl {
	return s.restoreArtifactBase
}

func (s RawRestoreArtifactBaseImpl) Entity() BaseEntityImpl {
	return s.restoreArtifactBase.Entity()
}

var _ json.Marshaler = BaseRestoreArtifactBaseImpl{}

func (s BaseRestoreArtifactBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseRestoreArtifactBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseRestoreArtifactBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseRestoreArtifactBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.restoreArtifactBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseRestoreArtifactBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalRestoreArtifactBaseImplementation(input []byte) (RestoreArtifactBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RestoreArtifactBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.driveRestoreArtifact") {
		var out DriveRestoreArtifact
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveRestoreArtifact: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxRestoreArtifact") {
		var out MailboxRestoreArtifact
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxRestoreArtifact: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.siteRestoreArtifact") {
		var out SiteRestoreArtifact
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SiteRestoreArtifact: %+v", err)
		}
		return out, nil
	}

	var parent BaseRestoreArtifactBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRestoreArtifactBaseImpl: %+v", err)
	}

	return RawRestoreArtifactBaseImpl{
		restoreArtifactBase: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
