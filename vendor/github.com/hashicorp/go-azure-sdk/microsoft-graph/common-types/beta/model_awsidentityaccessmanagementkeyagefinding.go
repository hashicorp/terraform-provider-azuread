package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Finding = AwsIdentityAccessManagementKeyAgeFinding{}

type AwsIdentityAccessManagementKeyAgeFinding struct {
	AccessKey             *AwsAccessKey          `json:"accessKey,omitempty"`
	ActionSummary         *ActionSummary         `json:"actionSummary,omitempty"`
	AwsAccessKeyDetails   *AwsAccessKeyDetails   `json:"awsAccessKeyDetails,omitempty"`
	PermissionsCreepIndex *PermissionsCreepIndex `json:"permissionsCreepIndex,omitempty"`
	Status                *IamStatus             `json:"status,omitempty"`

	// Fields inherited from Finding

	// Defines when the finding was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

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

func (s AwsIdentityAccessManagementKeyAgeFinding) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s AwsIdentityAccessManagementKeyAgeFinding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AwsIdentityAccessManagementKeyAgeFinding{}

func (s AwsIdentityAccessManagementKeyAgeFinding) MarshalJSON() ([]byte, error) {
	type wrapper AwsIdentityAccessManagementKeyAgeFinding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsIdentityAccessManagementKeyAgeFinding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsIdentityAccessManagementKeyAgeFinding: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsIdentityAccessManagementKeyAgeFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsIdentityAccessManagementKeyAgeFinding: %+v", err)
	}

	return encoded, nil
}
