package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VppTokenActionResult = VppTokenRevokeLicensesActionResult{}

type VppTokenRevokeLicensesActionResult struct {
	// Possible types of reasons for an Apple Volume Purchase Program token action failure.
	ActionFailureReason *VppTokenActionFailureReason `json:"actionFailureReason,omitempty"`

	// A count of the number of licenses that failed to revoke.
	FailedLicensesCount *int64 `json:"failedLicensesCount,omitempty"`

	// A count of the number of licenses that were attempted to revoke.
	TotalLicensesCount *int64 `json:"totalLicensesCount,omitempty"`

	// Fields inherited from VppTokenActionResult

	// Action name
	ActionName nullable.Type[string] `json:"actionName,omitempty"`

	ActionState *ActionState `json:"actionState,omitempty"`

	// Time the action state was last updated
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Time the action was initiated
	StartDateTime *string `json:"startDateTime,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s VppTokenRevokeLicensesActionResult) VppTokenActionResult() BaseVppTokenActionResultImpl {
	return BaseVppTokenActionResultImpl{
		ActionName:          s.ActionName,
		ActionState:         s.ActionState,
		LastUpdatedDateTime: s.LastUpdatedDateTime,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
		StartDateTime:       s.StartDateTime,
	}
}

var _ json.Marshaler = VppTokenRevokeLicensesActionResult{}

func (s VppTokenRevokeLicensesActionResult) MarshalJSON() ([]byte, error) {
	type wrapper VppTokenRevokeLicensesActionResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VppTokenRevokeLicensesActionResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VppTokenRevokeLicensesActionResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.vppTokenRevokeLicensesActionResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VppTokenRevokeLicensesActionResult: %+v", err)
	}

	return encoded, nil
}
