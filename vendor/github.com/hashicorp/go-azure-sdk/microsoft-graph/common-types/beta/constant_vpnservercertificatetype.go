package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnServerCertificateType string

const (
	VpnServerCertificateType_Ecdsa256 VpnServerCertificateType = "ecdsa256"
	VpnServerCertificateType_Ecdsa384 VpnServerCertificateType = "ecdsa384"
	VpnServerCertificateType_Ecdsa521 VpnServerCertificateType = "ecdsa521"
	VpnServerCertificateType_Rsa      VpnServerCertificateType = "rsa"
)

func PossibleValuesForVpnServerCertificateType() []string {
	return []string{
		string(VpnServerCertificateType_Ecdsa256),
		string(VpnServerCertificateType_Ecdsa384),
		string(VpnServerCertificateType_Ecdsa521),
		string(VpnServerCertificateType_Rsa),
	}
}

func (s *VpnServerCertificateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnServerCertificateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnServerCertificateType(input string) (*VpnServerCertificateType, error) {
	vals := map[string]VpnServerCertificateType{
		"ecdsa256": VpnServerCertificateType_Ecdsa256,
		"ecdsa384": VpnServerCertificateType_Ecdsa384,
		"ecdsa521": VpnServerCertificateType_Ecdsa521,
		"rsa":      VpnServerCertificateType_Rsa,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnServerCertificateType(input)
	return &out, nil
}
