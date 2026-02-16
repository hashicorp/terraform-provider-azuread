package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ZebraFotaArtifact{}

type ZebraFotaArtifact struct {
	// The version of the Board Support Package (BSP. E.g.: 01.18.02.00)
	BoardSupportPackageVersion *string `json:"boardSupportPackageVersion,omitempty"`

	// Artifact description. (e.g.: `LifeGuard Update 98 (released 24-September-2021)
	Description *string `json:"description,omitempty"`

	// Applicable device model (e.g.: TC8300)
	DeviceModel *string `json:"deviceModel,omitempty"`

	// Artifact OS version (e.g.: 8.1.0)
	OsVersion *string `json:"osVersion,omitempty"`

	// Artifact patch version (e.g.: U00)
	PatchVersion *string `json:"patchVersion,omitempty"`

	// Artifact release notes URL (e.g.: https://www.zebra.com/<filename.pdf>)
	ReleaseNotesUrl *string `json:"releaseNotesUrl,omitempty"`

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

func (s ZebraFotaArtifact) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ZebraFotaArtifact{}

func (s ZebraFotaArtifact) MarshalJSON() ([]byte, error) {
	type wrapper ZebraFotaArtifact
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ZebraFotaArtifact: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ZebraFotaArtifact: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.zebraFotaArtifact"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ZebraFotaArtifact: %+v", err)
	}

	return encoded, nil
}
