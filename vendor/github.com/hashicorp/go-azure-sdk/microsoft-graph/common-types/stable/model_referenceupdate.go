package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferenceUpdate struct {
	ODataId   *string               `json:"@odata.id,omitempty"`
	ODataType nullable.Type[string] `json:"@odata.type,omitempty"`
}
