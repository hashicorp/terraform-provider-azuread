package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PageLayoutType string

const (
	PageLayoutType_Article           PageLayoutType = "article"
	PageLayoutType_Home              PageLayoutType = "home"
	PageLayoutType_MicrosoftReserved PageLayoutType = "microsoftReserved"
	PageLayoutType_NewsLink          PageLayoutType = "newsLink"
	PageLayoutType_VideoNewsLink     PageLayoutType = "videoNewsLink"
)

func PossibleValuesForPageLayoutType() []string {
	return []string{
		string(PageLayoutType_Article),
		string(PageLayoutType_Home),
		string(PageLayoutType_MicrosoftReserved),
		string(PageLayoutType_NewsLink),
		string(PageLayoutType_VideoNewsLink),
	}
}

func (s *PageLayoutType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePageLayoutType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePageLayoutType(input string) (*PageLayoutType, error) {
	vals := map[string]PageLayoutType{
		"article":           PageLayoutType_Article,
		"home":              PageLayoutType_Home,
		"microsoftreserved": PageLayoutType_MicrosoftReserved,
		"newslink":          PageLayoutType_NewsLink,
		"videonewslink":     PageLayoutType_VideoNewsLink,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PageLayoutType(input)
	return &out, nil
}
