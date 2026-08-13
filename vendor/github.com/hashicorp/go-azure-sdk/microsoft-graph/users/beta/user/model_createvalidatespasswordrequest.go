package user

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateValidatesPasswordRequest struct {
	Password *string `json:"password,omitempty"`
}
