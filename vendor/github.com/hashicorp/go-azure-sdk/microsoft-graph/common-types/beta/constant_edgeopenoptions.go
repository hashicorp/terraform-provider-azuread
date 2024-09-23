package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeOpenOptions string

const (
	EdgeOpenOptions_NewTabPage    EdgeOpenOptions = "newTabPage"
	EdgeOpenOptions_NotConfigured EdgeOpenOptions = "notConfigured"
	EdgeOpenOptions_PreviousPages EdgeOpenOptions = "previousPages"
	EdgeOpenOptions_SpecificPages EdgeOpenOptions = "specificPages"
	EdgeOpenOptions_StartPage     EdgeOpenOptions = "startPage"
)

func PossibleValuesForEdgeOpenOptions() []string {
	return []string{
		string(EdgeOpenOptions_NewTabPage),
		string(EdgeOpenOptions_NotConfigured),
		string(EdgeOpenOptions_PreviousPages),
		string(EdgeOpenOptions_SpecificPages),
		string(EdgeOpenOptions_StartPage),
	}
}

func (s *EdgeOpenOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdgeOpenOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdgeOpenOptions(input string) (*EdgeOpenOptions, error) {
	vals := map[string]EdgeOpenOptions{
		"newtabpage":    EdgeOpenOptions_NewTabPage,
		"notconfigured": EdgeOpenOptions_NotConfigured,
		"previouspages": EdgeOpenOptions_PreviousPages,
		"specificpages": EdgeOpenOptions_SpecificPages,
		"startpage":     EdgeOpenOptions_StartPage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdgeOpenOptions(input)
	return &out, nil
}
