package directoryobject

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DirectoryObjectOperationPredicate struct {
}

func (p DirectoryObjectOperationPredicate) Matches(input stable.DirectoryObject) bool {

	return true
}

type ExtensionPropertyOperationPredicate struct {
}

func (p ExtensionPropertyOperationPredicate) Matches(input stable.ExtensionProperty) bool {

	return true
}
