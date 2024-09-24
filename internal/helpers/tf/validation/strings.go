// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validation

import (
	"fmt"
	"regexp"
	"strings"
)

// StringIsEmailAddress validates that the given string is a valid email address (foo@bar.com)
func StringIsEmailAddress(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		return nil, []error{fmt.Errorf("expected a string value for %q", k)}
	}

	if strings.TrimSpace(v) == "" {
		return nil, []error{fmt.Errorf("value must not be empty for %q", k)}
	}

	regExIsEmailAddress := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !regExIsEmailAddress.MatchString(v) {
		return nil, []error{fmt.Errorf("value must be a valid email address for %q", k)}
	}

	return
}
