package directoryrole

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGetsByIdsRequest struct {
	Ids   *[]string `json:"ids,omitempty"`
	Types *[]string `json:"types,omitempty"`
}
