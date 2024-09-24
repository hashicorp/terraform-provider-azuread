package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FirewallCertificateRevocationListCheckMethodType string

const (
	FirewallCertificateRevocationListCheckMethodType_Attempt       FirewallCertificateRevocationListCheckMethodType = "attempt"
	FirewallCertificateRevocationListCheckMethodType_DeviceDefault FirewallCertificateRevocationListCheckMethodType = "deviceDefault"
	FirewallCertificateRevocationListCheckMethodType_None          FirewallCertificateRevocationListCheckMethodType = "none"
	FirewallCertificateRevocationListCheckMethodType_Require       FirewallCertificateRevocationListCheckMethodType = "require"
)

func PossibleValuesForFirewallCertificateRevocationListCheckMethodType() []string {
	return []string{
		string(FirewallCertificateRevocationListCheckMethodType_Attempt),
		string(FirewallCertificateRevocationListCheckMethodType_DeviceDefault),
		string(FirewallCertificateRevocationListCheckMethodType_None),
		string(FirewallCertificateRevocationListCheckMethodType_Require),
	}
}

func (s *FirewallCertificateRevocationListCheckMethodType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFirewallCertificateRevocationListCheckMethodType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFirewallCertificateRevocationListCheckMethodType(input string) (*FirewallCertificateRevocationListCheckMethodType, error) {
	vals := map[string]FirewallCertificateRevocationListCheckMethodType{
		"attempt":       FirewallCertificateRevocationListCheckMethodType_Attempt,
		"devicedefault": FirewallCertificateRevocationListCheckMethodType_DeviceDefault,
		"none":          FirewallCertificateRevocationListCheckMethodType_None,
		"require":       FirewallCertificateRevocationListCheckMethodType_Require,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FirewallCertificateRevocationListCheckMethodType(input)
	return &out, nil
}
