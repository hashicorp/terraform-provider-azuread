package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSourceDeviceImage struct {
	// The display name for the source image. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The fully qualified unique identifier (ID) of the source image resource in Azure. The ID format is:
	// '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}'.
	// Read-only.
	ResourceId *string `json:"resourceId,omitempty"`

	// The display name of the subscription that hosts the source image. Read-only.
	SubscriptionDisplayName nullable.Type[string] `json:"subscriptionDisplayName,omitempty"`

	// The unique identifier (ID) of the subscription that hosts the source image. Read-only.
	SubscriptionId nullable.Type[string] `json:"subscriptionId,omitempty"`
}

var _ json.Marshaler = CloudPCSourceDeviceImage{}

func (s CloudPCSourceDeviceImage) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCSourceDeviceImage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCSourceDeviceImage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCSourceDeviceImage: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "resourceId")
	delete(decoded, "subscriptionDisplayName")
	delete(decoded, "subscriptionId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCSourceDeviceImage: %+v", err)
	}

	return encoded, nil
}
