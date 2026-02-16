package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MailboxRestoreArtifact = GranularMailboxRestoreArtifact{}

type GranularMailboxRestoreArtifact struct {
	// .
	ArtifactCount nullable.Type[int64] `json:"artifactCount,omitempty"`

	// .
	SearchResponseId *string `json:"searchResponseId,omitempty"`

	// Fields inherited from MailboxRestoreArtifact

	// The newly restored folder identifier for the user.
	RestoredFolderId nullable.Type[string] `json:"restoredFolderId,omitempty"`

	// The new restored folder name.
	RestoredFolderName nullable.Type[string] `json:"restoredFolderName,omitempty"`

	// The number of items that are being restored in the folder.
	RestoredItemCount nullable.Type[int64] `json:"restoredItemCount,omitempty"`

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

func (s GranularMailboxRestoreArtifact) MailboxRestoreArtifact() BaseMailboxRestoreArtifactImpl {
	return BaseMailboxRestoreArtifactImpl{
		RestoredFolderId:   s.RestoredFolderId,
		RestoredFolderName: s.RestoredFolderName,
		RestoredItemCount:  s.RestoredItemCount,
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

func (s GranularMailboxRestoreArtifact) RestoreArtifactBase() BaseRestoreArtifactBaseImpl {
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

func (s GranularMailboxRestoreArtifact) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GranularMailboxRestoreArtifact{}

func (s GranularMailboxRestoreArtifact) MarshalJSON() ([]byte, error) {
	type wrapper GranularMailboxRestoreArtifact
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GranularMailboxRestoreArtifact: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GranularMailboxRestoreArtifact: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.granularMailboxRestoreArtifact"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GranularMailboxRestoreArtifact: %+v", err)
	}

	return encoded, nil
}
