package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomExtensionData = AccessPackageAssignmentRequestCallbackData{}

type AccessPackageAssignmentRequestCallbackData struct {
	// Details for the callback.
	CustomExtensionStageInstanceDetail nullable.Type[string] `json:"customExtensionStageInstanceDetail,omitempty"`

	// Unique identifier of the callout to the custom extension.
	CustomExtensionStageInstanceId nullable.Type[string] `json:"customExtensionStageInstanceId,omitempty"`

	// Indicates the stage at which the custom callout extension is executed. The possible values are:
	// assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved,
	// assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
	Stage *AccessPackageCustomExtensionStage `json:"stage,omitempty"`

	// Allow the extension to be able to deny or cancel the request submitted by the requestor. The supported values are
	// Denied and Canceled. This property can only be set for an assignmentRequestCreated stage.
	State nullable.Type[string] `json:"state,omitempty"`

	// Fields inherited from CustomExtensionData

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessPackageAssignmentRequestCallbackData) CustomExtensionData() BaseCustomExtensionDataImpl {
	return BaseCustomExtensionDataImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageAssignmentRequestCallbackData{}

func (s AccessPackageAssignmentRequestCallbackData) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageAssignmentRequestCallbackData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageAssignmentRequestCallbackData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageAssignmentRequestCallbackData: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageAssignmentRequestCallbackData"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageAssignmentRequestCallbackData: %+v", err)
	}

	return encoded, nil
}
