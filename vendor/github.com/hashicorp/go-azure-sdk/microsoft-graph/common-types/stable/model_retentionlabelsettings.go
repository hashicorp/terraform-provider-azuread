package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetentionLabelSettings struct {
	// Describes the item behavior during retention period. Possible values are: doNotRetain, retain, retainAsRecord,
	// retainAsRegulatoryRecord, unknownFutureValue. Read-only.
	BehaviorDuringRetentionPeriod *SecurityBehaviorDuringRetentionPeriod `json:"behaviorDuringRetentionPeriod,omitempty"`

	// Specifies whether updates to document content are allowed. Read-only.
	IsContentUpdateAllowed nullable.Type[bool] `json:"isContentUpdateAllowed,omitempty"`

	// Specifies whether the document deletion is allowed. Read-only.
	IsDeleteAllowed nullable.Type[bool] `json:"isDeleteAllowed,omitempty"`

	// Specifies whether you're allowed to change the retention label on the document. Read-only.
	IsLabelUpdateAllowed nullable.Type[bool] `json:"isLabelUpdateAllowed,omitempty"`

	// Specifies whether updates to the item metadata (for example, the Title field) are blocked. Read-only.
	IsMetadataUpdateAllowed nullable.Type[bool] `json:"isMetadataUpdateAllowed,omitempty"`

	// Specifies whether the item is locked. Read-write.
	IsRecordLocked nullable.Type[bool] `json:"isRecordLocked,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = RetentionLabelSettings{}

func (s RetentionLabelSettings) MarshalJSON() ([]byte, error) {
	type wrapper RetentionLabelSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RetentionLabelSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RetentionLabelSettings: %+v", err)
	}

	delete(decoded, "behaviorDuringRetentionPeriod")
	delete(decoded, "isContentUpdateAllowed")
	delete(decoded, "isDeleteAllowed")
	delete(decoded, "isLabelUpdateAllowed")
	delete(decoded, "isMetadataUpdateAllowed")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RetentionLabelSettings: %+v", err)
	}

	return encoded, nil
}
