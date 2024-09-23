package group

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEvaluatesDynamicMembershipRequest struct {
	MemberId       nullable.Type[string] `json:"memberId,omitempty"`
	MembershipRule nullable.Type[string] `json:"membershipRule,omitempty"`
}
