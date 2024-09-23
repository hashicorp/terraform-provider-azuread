package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PagePromotionType string

const (
	PagePromotionType_MicrosoftReserved PagePromotionType = "microsoftReserved"
	PagePromotionType_NewsPost          PagePromotionType = "newsPost"
	PagePromotionType_Page              PagePromotionType = "page"
)

func PossibleValuesForPagePromotionType() []string {
	return []string{
		string(PagePromotionType_MicrosoftReserved),
		string(PagePromotionType_NewsPost),
		string(PagePromotionType_Page),
	}
}

func (s *PagePromotionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePagePromotionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePagePromotionType(input string) (*PagePromotionType, error) {
	vals := map[string]PagePromotionType{
		"microsoftreserved": PagePromotionType_MicrosoftReserved,
		"newspost":          PagePromotionType_NewsPost,
		"page":              PagePromotionType_Page,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PagePromotionType(input)
	return &out, nil
}
