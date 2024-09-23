package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpEncryption string

const (
	GcpEncryption_Customer GcpEncryption = "customer"
	GcpEncryption_Google   GcpEncryption = "google"
)

func PossibleValuesForGcpEncryption() []string {
	return []string{
		string(GcpEncryption_Customer),
		string(GcpEncryption_Google),
	}
}

func (s *GcpEncryption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGcpEncryption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGcpEncryption(input string) (*GcpEncryption, error) {
	vals := map[string]GcpEncryption{
		"customer": GcpEncryption_Customer,
		"google":   GcpEncryption_Google,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GcpEncryption(input)
	return &out, nil
}
