package application

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckMemberGroupsRequest struct {
	GroupIds *[]string `json:"groupIds,omitempty"`
}
