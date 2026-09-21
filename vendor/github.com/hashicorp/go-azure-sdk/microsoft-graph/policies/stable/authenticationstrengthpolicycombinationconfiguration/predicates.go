package authenticationstrengthpolicycombinationconfiguration

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AuthenticationCombinationConfigurationOperationPredicate struct {
}

func (p AuthenticationCombinationConfigurationOperationPredicate) Matches(input stable.AuthenticationCombinationConfiguration) bool {

	return true
}
