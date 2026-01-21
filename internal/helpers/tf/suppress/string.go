// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package suppress

import (
	"strings"

	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

func CaseDifference(_, old, new string, _ *pluginsdk.ResourceData) bool {
	return strings.EqualFold(old, new)
}
