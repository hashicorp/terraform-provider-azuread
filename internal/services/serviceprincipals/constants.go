// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

const (
	DelegatedPermissionGrantConsentTypeAllPrincipals = "AllPrincipals"
	DelegatedPermissionGrantConsentTypePrincipal     = "Principal"
)

const (
	KeyCredentialTypeAsymmetricX509Cert  = "AsymmetricX509Cert"
	KeyCredentialTypeX509CertAndPassword = "X509CertAndPassword"
)

var possibleValuesForKeyCredentialType = []string{KeyCredentialTypeAsymmetricX509Cert, KeyCredentialTypeX509CertAndPassword}

const (
	PreferredSingleSignOnModeNone         = ""
	PreferredSingleSignOnModeNotSupported = "notSupported"
	PreferredSingleSignOnModeOidc         = "oidc"
	PreferredSingleSignOnModePassword     = "password"
	PreferredSingleSignOnModeSaml         = "saml"
)

var possibleValuesForPreferredSingleSignOnMode = []string{
	PreferredSingleSignOnModeNone,
	PreferredSingleSignOnModeNotSupported,
	PreferredSingleSignOnModeOidc,
	PreferredSingleSignOnModePassword,
	PreferredSingleSignOnModeSaml,
}
