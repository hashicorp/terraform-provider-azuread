package application

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveKeyRequest struct {
	KeyId *string `json:"keyId,omitempty"`
	Proof *string `json:"proof,omitempty"`
}
