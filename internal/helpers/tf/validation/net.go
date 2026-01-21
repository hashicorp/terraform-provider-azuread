// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package validation

import (
	"fmt"
	"net/netip"
)

func StringIsIpPrefix(i interface{}, k string) (warnings []string, errors []error) {
	if warnings, errors = StringIsNotEmpty(i, k); len(errors) > 0 {
		return warnings, errors
	}

	if _, err := netip.ParsePrefix(i.(string)); err != nil {
		return nil, []error{fmt.Errorf("expected %q to be a valid IPv4 or IPv6 prefix", k)}
	}

	return
}

func PrefixLengthAtLeast(minLength int) func(interface{}, string) ([]string, []error) {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		if warnings, errors = StringIsNotEmpty(i, k); len(errors) > 0 {
			return warnings, errors
		}

		prefix, err := netip.ParsePrefix(i.(string))
		if err != nil {
			return nil, []error{fmt.Errorf("expected %q to be a valid IPv4 or IPv6 prefix", k)}
		}

		if prefixLength := prefix.Bits(); prefixLength < minLength {
			return nil, []error{fmt.Errorf("expected %q to have a prefix length at least %d, got %d", k, minLength, prefixLength)}
		}

		return
	}
}

func PrefixLengthAtMost(maxLength int) func(interface{}, string) ([]string, []error) {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		if warnings, errors = StringIsNotEmpty(i, k); len(errors) > 0 {
			return warnings, errors
		}

		prefix, err := netip.ParsePrefix(i.(string))
		if err != nil {
			return nil, []error{fmt.Errorf("expected %q to be a valid IPv4 or IPv6 prefix", k)}
		}

		if prefixLength := prefix.Bits(); prefixLength > maxLength {
			return nil, []error{fmt.Errorf("expected %q to have a prefix length at most %d, got %d", k, maxLength, prefixLength)}
		}

		return
	}
}

func PrefixLengthBetween(minLength, maxLength int) func(interface{}, string) ([]string, []error) {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		if warnings, errors = StringIsNotEmpty(i, k); len(errors) > 0 {
			return warnings, errors
		}

		prefix, err := netip.ParsePrefix(i.(string))
		if err != nil {
			return nil, []error{fmt.Errorf("expected %q to be a valid IPv4 or IPv6 prefix", k)}
		}

		if prefixLength := prefix.Bits(); prefixLength < minLength {
			return nil, []error{fmt.Errorf("expected %q to have a prefix length at least %d, got %d", k, minLength, prefixLength)}
		}

		if prefixLength := prefix.Bits(); prefixLength > maxLength {
			return nil, []error{fmt.Errorf("expected %q to have a prefix length at most %d, got %d", k, maxLength, prefixLength)}
		}

		return
	}
}
