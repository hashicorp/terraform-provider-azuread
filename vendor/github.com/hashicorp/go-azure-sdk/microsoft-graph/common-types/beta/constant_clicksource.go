package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClickSource string

const (
	ClickSource_PhishingUrl ClickSource = "phishingUrl"
	ClickSource_QrCode      ClickSource = "qrCode"
	ClickSource_Unknown     ClickSource = "unknown"
)

func PossibleValuesForClickSource() []string {
	return []string{
		string(ClickSource_PhishingUrl),
		string(ClickSource_QrCode),
		string(ClickSource_Unknown),
	}
}

func (s *ClickSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClickSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClickSource(input string) (*ClickSource, error) {
	vals := map[string]ClickSource{
		"phishingurl": ClickSource_PhishingUrl,
		"qrcode":      ClickSource_QrCode,
		"unknown":     ClickSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClickSource(input)
	return &out, nil
}
