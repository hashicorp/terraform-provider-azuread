// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package users

const (
	AgeGroupAdult    = "Adult"
	AgeGroupMinor    = "Minor"
	AgeGroupNotAdult = "NotAdult"
)

var possibleValuesForAgeGroup = []string{AgeGroupAdult, AgeGroupMinor, AgeGroupNotAdult}

const (
	ConsentProvidedForMinorDenied      = "Denied"
	ConsentProvidedForMinorGranted     = "Granted"
	ConsentProvidedForMinorNotRequired = "NotRequired"
)

var possibleValuesForConsentProvidedForMinor = []string{ConsentProvidedForMinorDenied, ConsentProvidedForMinorGranted, ConsentProvidedForMinorNotRequired}
