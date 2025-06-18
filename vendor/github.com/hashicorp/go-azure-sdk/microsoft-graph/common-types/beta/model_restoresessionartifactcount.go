package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreSessionArtifactCount struct {
	// The number of artifacts whose restoration completed.
	Completed nullable.Type[int64] `json:"completed,omitempty"`

	// The number of artifacts whose restoration failed.
	Failed nullable.Type[int64] `json:"failed,omitempty"`

	// The number of artifacts whose restoration is in progress.
	InProgress nullable.Type[int64] `json:"inProgress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of artifacts present in the restore session.
	Total nullable.Type[int64] `json:"total,omitempty"`
}
