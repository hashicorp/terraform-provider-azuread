package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalItemContentType string

const (
	ExternalItemContentType_Html ExternalItemContentType = "html"
	ExternalItemContentType_Text ExternalItemContentType = "text"
)

func PossibleValuesForExternalItemContentType() []string {
	return []string{
		string(ExternalItemContentType_Html),
		string(ExternalItemContentType_Text),
	}
}

func (s *ExternalItemContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalItemContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalItemContentType(input string) (*ExternalItemContentType, error) {
	vals := map[string]ExternalItemContentType{
		"html": ExternalItemContentType_Html,
		"text": ExternalItemContentType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalItemContentType(input)
	return &out, nil
}
