package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreArtifactBase = DriveRestoreArtifact{}

type DriveRestoreArtifact struct {
	// The new site identifier if destinationType is new, and the input site ID if the destinationType is inPlace.
	RestoredSiteId nullable.Type[string] `json:"restoredSiteId,omitempty"`

	// The name of the restored site.
	RestoredSiteName nullable.Type[string] `json:"restoredSiteName,omitempty"`

	// The web URL of the restored site.
	RestoredSiteWebUrl nullable.Type[string] `json:"restoredSiteWebUrl,omitempty"`

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

func (s DriveRestoreArtifact) RestoreArtifactBase() BaseRestoreArtifactBaseImpl {
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

func (s DriveRestoreArtifact) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DriveRestoreArtifact{}

func (s DriveRestoreArtifact) MarshalJSON() ([]byte, error) {
	type wrapper DriveRestoreArtifact
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DriveRestoreArtifact: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DriveRestoreArtifact: %+v", err)
	}

	delete(decoded, "restoredSiteName")
	delete(decoded, "restoredSiteWebUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.driveRestoreArtifact"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DriveRestoreArtifact: %+v", err)
	}

	return encoded, nil
}
