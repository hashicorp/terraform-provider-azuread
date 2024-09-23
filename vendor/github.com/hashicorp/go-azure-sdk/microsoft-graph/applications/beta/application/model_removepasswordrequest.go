package application

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemovePasswordRequest struct {
	KeyId *string `json:"keyId,omitempty"`
}
