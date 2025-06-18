package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TlsClientRegistrationMetadata string

const (
	TlsClientRegistrationMetadata_Tlsclientauthsandns    TlsClientRegistrationMetadata = "tls_client_auth_san_dns"
	TlsClientRegistrationMetadata_Tlsclientauthsanemail  TlsClientRegistrationMetadata = "tls_client_auth_san_email"
	TlsClientRegistrationMetadata_Tlsclientauthsanip     TlsClientRegistrationMetadata = "tls_client_auth_san_ip"
	TlsClientRegistrationMetadata_Tlsclientauthsanuri    TlsClientRegistrationMetadata = "tls_client_auth_san_uri"
	TlsClientRegistrationMetadata_Tlsclientauthsubjectdn TlsClientRegistrationMetadata = "tls_client_auth_subject_dn"
)

func PossibleValuesForTlsClientRegistrationMetadata() []string {
	return []string{
		string(TlsClientRegistrationMetadata_Tlsclientauthsandns),
		string(TlsClientRegistrationMetadata_Tlsclientauthsanemail),
		string(TlsClientRegistrationMetadata_Tlsclientauthsanip),
		string(TlsClientRegistrationMetadata_Tlsclientauthsanuri),
		string(TlsClientRegistrationMetadata_Tlsclientauthsubjectdn),
	}
}

func (s *TlsClientRegistrationMetadata) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTlsClientRegistrationMetadata(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTlsClientRegistrationMetadata(input string) (*TlsClientRegistrationMetadata, error) {
	vals := map[string]TlsClientRegistrationMetadata{
		"tls_client_auth_san_dns":    TlsClientRegistrationMetadata_Tlsclientauthsandns,
		"tls_client_auth_san_email":  TlsClientRegistrationMetadata_Tlsclientauthsanemail,
		"tls_client_auth_san_ip":     TlsClientRegistrationMetadata_Tlsclientauthsanip,
		"tls_client_auth_san_uri":    TlsClientRegistrationMetadata_Tlsclientauthsanuri,
		"tls_client_auth_subject_dn": TlsClientRegistrationMetadata_Tlsclientauthsubjectdn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TlsClientRegistrationMetadata(input)
	return &out, nil
}
