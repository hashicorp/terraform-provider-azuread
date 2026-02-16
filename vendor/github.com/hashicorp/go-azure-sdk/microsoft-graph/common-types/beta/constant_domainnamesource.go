package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainNameSource string

const (
	DomainNameSource_FullDomainName    DomainNameSource = "fullDomainName"
	DomainNameSource_NetBiosDomainName DomainNameSource = "netBiosDomainName"
)

func PossibleValuesForDomainNameSource() []string {
	return []string{
		string(DomainNameSource_FullDomainName),
		string(DomainNameSource_NetBiosDomainName),
	}
}

func (s *DomainNameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDomainNameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDomainNameSource(input string) (*DomainNameSource, error) {
	vals := map[string]DomainNameSource{
		"fulldomainname":    DomainNameSource_FullDomainName,
		"netbiosdomainname": DomainNameSource_NetBiosDomainName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DomainNameSource(input)
	return &out, nil
}
