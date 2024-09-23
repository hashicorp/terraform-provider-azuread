package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PartnersBillingManifest{}

type PartnersBillingManifest struct {
	// The total file count for this partner tenant ID.
	BlobCount *int64 `json:"blobCount,omitempty"`

	// A collection of blob objects that contain details of all the files for the partner tenant ID.
	Blobs *[]PartnersBillingBlob `json:"blobs,omitempty"`

	// The date and time when a manifest resource was created. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The billing data file format. The possible value is: compressedJSONLines. Each blob is a compressed file and data in
	// the file is in JSON lines format. Decompress the file to access the data.
	DataFormat *string `json:"dataFormat,omitempty"`

	// Version of data represented by the manifest. Any change in eTag indicates a new data version.
	ETag *string `json:"eTag,omitempty"`

	// Indicates the division of data. If a given partition has more than the supported number, the data is split into
	// multiple files, each file representing a specific partitionValue. By default, the data in the file is partitioned by
	// the number of line items.
	PartitionType *string `json:"partitionType,omitempty"`

	// The Microsoft Entra tenant ID of the partner.
	PartnerTenantId *string `json:"partnerTenantId,omitempty"`

	// The root directory that contains all the files.
	RootDirectory *string `json:"rootDirectory,omitempty"`

	// The SAS token for accessing the directory or an individual file in the directory.
	SasToken *string `json:"sasToken,omitempty"`

	// The version of the manifest schema.
	SchemaVersion *string `json:"schemaVersion,omitempty"`

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

func (s PartnersBillingManifest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PartnersBillingManifest{}

func (s PartnersBillingManifest) MarshalJSON() ([]byte, error) {
	type wrapper PartnersBillingManifest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PartnersBillingManifest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnersBillingManifest: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.partners.billing.manifest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PartnersBillingManifest: %+v", err)
	}

	return encoded, nil
}
