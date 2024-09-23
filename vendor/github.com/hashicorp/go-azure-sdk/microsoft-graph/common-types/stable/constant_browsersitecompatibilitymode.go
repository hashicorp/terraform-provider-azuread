package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteCompatibilityMode string

const (
	BrowserSiteCompatibilityMode_Default                     BrowserSiteCompatibilityMode = "default"
	BrowserSiteCompatibilityMode_InternetExplorer10          BrowserSiteCompatibilityMode = "internetExplorer10"
	BrowserSiteCompatibilityMode_InternetExplorer11          BrowserSiteCompatibilityMode = "internetExplorer11"
	BrowserSiteCompatibilityMode_InternetExplorer5           BrowserSiteCompatibilityMode = "internetExplorer5"
	BrowserSiteCompatibilityMode_InternetExplorer7           BrowserSiteCompatibilityMode = "internetExplorer7"
	BrowserSiteCompatibilityMode_InternetExplorer7Enterprise BrowserSiteCompatibilityMode = "internetExplorer7Enterprise"
	BrowserSiteCompatibilityMode_InternetExplorer8           BrowserSiteCompatibilityMode = "internetExplorer8"
	BrowserSiteCompatibilityMode_InternetExplorer8Enterprise BrowserSiteCompatibilityMode = "internetExplorer8Enterprise"
	BrowserSiteCompatibilityMode_InternetExplorer9           BrowserSiteCompatibilityMode = "internetExplorer9"
)

func PossibleValuesForBrowserSiteCompatibilityMode() []string {
	return []string{
		string(BrowserSiteCompatibilityMode_Default),
		string(BrowserSiteCompatibilityMode_InternetExplorer10),
		string(BrowserSiteCompatibilityMode_InternetExplorer11),
		string(BrowserSiteCompatibilityMode_InternetExplorer5),
		string(BrowserSiteCompatibilityMode_InternetExplorer7),
		string(BrowserSiteCompatibilityMode_InternetExplorer7Enterprise),
		string(BrowserSiteCompatibilityMode_InternetExplorer8),
		string(BrowserSiteCompatibilityMode_InternetExplorer8Enterprise),
		string(BrowserSiteCompatibilityMode_InternetExplorer9),
	}
}

func (s *BrowserSiteCompatibilityMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSiteCompatibilityMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSiteCompatibilityMode(input string) (*BrowserSiteCompatibilityMode, error) {
	vals := map[string]BrowserSiteCompatibilityMode{
		"default":                     BrowserSiteCompatibilityMode_Default,
		"internetexplorer10":          BrowserSiteCompatibilityMode_InternetExplorer10,
		"internetexplorer11":          BrowserSiteCompatibilityMode_InternetExplorer11,
		"internetexplorer5":           BrowserSiteCompatibilityMode_InternetExplorer5,
		"internetexplorer7":           BrowserSiteCompatibilityMode_InternetExplorer7,
		"internetexplorer7enterprise": BrowserSiteCompatibilityMode_InternetExplorer7Enterprise,
		"internetexplorer8":           BrowserSiteCompatibilityMode_InternetExplorer8,
		"internetexplorer8enterprise": BrowserSiteCompatibilityMode_InternetExplorer8Enterprise,
		"internetexplorer9":           BrowserSiteCompatibilityMode_InternetExplorer9,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteCompatibilityMode(input)
	return &out, nil
}
