package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArtifactQuery struct {
	// The type of artifact to search. The possible values are: message, unknownFutureValue.
	ArtifactType *RestorableArtifact `json:"artifactType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies criteria to retrieve artifacts.
	QueryExpression nullable.Type[string] `json:"queryExpression,omitempty"`
}
