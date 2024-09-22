// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validation

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
)

func ISO639Language(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		return nil, []error{fmt.Errorf("expected a string value for %q", k)}
	}

	if _, err := language.Parse(v); err != nil && strings.Contains(err.Error(), "not well-formed") {
		return nil, []error{fmt.Errorf("value is not a well-formed language for %q", k)}
	}

	return
}
