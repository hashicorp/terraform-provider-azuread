package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnLocalIdentifier string

const (
	VpnLocalIdentifier_ClientCertificateSubjectName VpnLocalIdentifier = "clientCertificateSubjectName"
	VpnLocalIdentifier_DeviceFQDN                   VpnLocalIdentifier = "deviceFQDN"
	VpnLocalIdentifier_Empty                        VpnLocalIdentifier = "empty"
)

func PossibleValuesForVpnLocalIdentifier() []string {
	return []string{
		string(VpnLocalIdentifier_ClientCertificateSubjectName),
		string(VpnLocalIdentifier_DeviceFQDN),
		string(VpnLocalIdentifier_Empty),
	}
}

func (s *VpnLocalIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnLocalIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnLocalIdentifier(input string) (*VpnLocalIdentifier, error) {
	vals := map[string]VpnLocalIdentifier{
		"clientcertificatesubjectname": VpnLocalIdentifier_ClientCertificateSubjectName,
		"devicefqdn":                   VpnLocalIdentifier_DeviceFQDN,
		"empty":                        VpnLocalIdentifier_Empty,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnLocalIdentifier(input)
	return &out, nil
}
