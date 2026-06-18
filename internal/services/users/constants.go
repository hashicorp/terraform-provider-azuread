// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package users

// userResourceName is the lock key used to serialize mutations against a single user object. It is
// intentionally scoped to the parent user (not to any individual child resource) so that every
// resource which mutates the same user - such as multiple azuread_user_license resources - serializes
// against the same key. See tf.LockByName.
const userResourceName = "azuread_user"

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
