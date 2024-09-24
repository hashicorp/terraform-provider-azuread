package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleSuccessStatistics struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PermanentFail    nullable.Type[int64]  `json:"permanentFail,omitempty"`
	PermanentSuccess nullable.Type[int64]  `json:"permanentSuccess,omitempty"`
	RemoveFail       nullable.Type[int64]  `json:"removeFail,omitempty"`
	RemoveSuccess    nullable.Type[int64]  `json:"removeSuccess,omitempty"`
	RoleId           nullable.Type[string] `json:"roleId,omitempty"`
	RoleName         nullable.Type[string] `json:"roleName,omitempty"`
	TemporaryFail    nullable.Type[int64]  `json:"temporaryFail,omitempty"`
	TemporarySuccess nullable.Type[int64]  `json:"temporarySuccess,omitempty"`
	UnknownFail      nullable.Type[int64]  `json:"unknownFail,omitempty"`
}
