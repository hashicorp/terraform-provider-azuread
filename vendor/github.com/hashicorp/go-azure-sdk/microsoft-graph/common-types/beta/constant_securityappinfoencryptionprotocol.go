package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAppInfoEncryptionProtocol string

const (
	SecurityAppInfoEncryptionProtocol_NotApplicable SecurityAppInfoEncryptionProtocol = "notApplicable"
	SecurityAppInfoEncryptionProtocol_NotSupported  SecurityAppInfoEncryptionProtocol = "notSupported"
	SecurityAppInfoEncryptionProtocol_Ssl3          SecurityAppInfoEncryptionProtocol = "ssl3"
	SecurityAppInfoEncryptionProtocol_Tls10         SecurityAppInfoEncryptionProtocol = "tls1_0"
	SecurityAppInfoEncryptionProtocol_Tls11         SecurityAppInfoEncryptionProtocol = "tls1_1"
	SecurityAppInfoEncryptionProtocol_Tls12         SecurityAppInfoEncryptionProtocol = "tls1_2"
	SecurityAppInfoEncryptionProtocol_Tls13         SecurityAppInfoEncryptionProtocol = "tls1_3"
	SecurityAppInfoEncryptionProtocol_Unknown       SecurityAppInfoEncryptionProtocol = "unknown"
)

func PossibleValuesForSecurityAppInfoEncryptionProtocol() []string {
	return []string{
		string(SecurityAppInfoEncryptionProtocol_NotApplicable),
		string(SecurityAppInfoEncryptionProtocol_NotSupported),
		string(SecurityAppInfoEncryptionProtocol_Ssl3),
		string(SecurityAppInfoEncryptionProtocol_Tls10),
		string(SecurityAppInfoEncryptionProtocol_Tls11),
		string(SecurityAppInfoEncryptionProtocol_Tls12),
		string(SecurityAppInfoEncryptionProtocol_Tls13),
		string(SecurityAppInfoEncryptionProtocol_Unknown),
	}
}

func (s *SecurityAppInfoEncryptionProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAppInfoEncryptionProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAppInfoEncryptionProtocol(input string) (*SecurityAppInfoEncryptionProtocol, error) {
	vals := map[string]SecurityAppInfoEncryptionProtocol{
		"notapplicable": SecurityAppInfoEncryptionProtocol_NotApplicable,
		"notsupported":  SecurityAppInfoEncryptionProtocol_NotSupported,
		"ssl3":          SecurityAppInfoEncryptionProtocol_Ssl3,
		"tls1_0":        SecurityAppInfoEncryptionProtocol_Tls10,
		"tls1_1":        SecurityAppInfoEncryptionProtocol_Tls11,
		"tls1_2":        SecurityAppInfoEncryptionProtocol_Tls12,
		"tls1_3":        SecurityAppInfoEncryptionProtocol_Tls13,
		"unknown":       SecurityAppInfoEncryptionProtocol_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAppInfoEncryptionProtocol(input)
	return &out, nil
}
