package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalItemContentType string

const (
	ExternalConnectorsExternalItemContentType_Html ExternalConnectorsExternalItemContentType = "html"
	ExternalConnectorsExternalItemContentType_Text ExternalConnectorsExternalItemContentType = "text"
)

func PossibleValuesForExternalConnectorsExternalItemContentType() []string {
	return []string{
		string(ExternalConnectorsExternalItemContentType_Html),
		string(ExternalConnectorsExternalItemContentType_Text),
	}
}

func (s *ExternalConnectorsExternalItemContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsExternalItemContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsExternalItemContentType(input string) (*ExternalConnectorsExternalItemContentType, error) {
	vals := map[string]ExternalConnectorsExternalItemContentType{
		"html": ExternalConnectorsExternalItemContentType_Html,
		"text": ExternalConnectorsExternalItemContentType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsExternalItemContentType(input)
	return &out, nil
}
