package administrativeunit

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateValidatesPropertyRequest struct {
	DisplayName      nullable.Type[string] `json:"displayName,omitempty"`
	EntityType       nullable.Type[string] `json:"entityType,omitempty"`
	MailNickname     nullable.Type[string] `json:"mailNickname,omitempty"`
	OnBehalfOfUserId nullable.Type[string] `json:"onBehalfOfUserId,omitempty"`
}
