package application

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGetsByIdsRequest struct {
	Ids   *[]string `json:"ids,omitempty"`
	Types *[]string `json:"types,omitempty"`
}
