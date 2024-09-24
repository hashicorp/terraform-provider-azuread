package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteTargetEnvironment string

const (
	BrowserSiteTargetEnvironment_Configurable         BrowserSiteTargetEnvironment = "configurable"
	BrowserSiteTargetEnvironment_InternetExplorer11   BrowserSiteTargetEnvironment = "internetExplorer11"
	BrowserSiteTargetEnvironment_InternetExplorerMode BrowserSiteTargetEnvironment = "internetExplorerMode"
	BrowserSiteTargetEnvironment_MicrosoftEdge        BrowserSiteTargetEnvironment = "microsoftEdge"
	BrowserSiteTargetEnvironment_None                 BrowserSiteTargetEnvironment = "none"
)

func PossibleValuesForBrowserSiteTargetEnvironment() []string {
	return []string{
		string(BrowserSiteTargetEnvironment_Configurable),
		string(BrowserSiteTargetEnvironment_InternetExplorer11),
		string(BrowserSiteTargetEnvironment_InternetExplorerMode),
		string(BrowserSiteTargetEnvironment_MicrosoftEdge),
		string(BrowserSiteTargetEnvironment_None),
	}
}

func (s *BrowserSiteTargetEnvironment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSiteTargetEnvironment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSiteTargetEnvironment(input string) (*BrowserSiteTargetEnvironment, error) {
	vals := map[string]BrowserSiteTargetEnvironment{
		"configurable":         BrowserSiteTargetEnvironment_Configurable,
		"internetexplorer11":   BrowserSiteTargetEnvironment_InternetExplorer11,
		"internetexplorermode": BrowserSiteTargetEnvironment_InternetExplorerMode,
		"microsoftedge":        BrowserSiteTargetEnvironment_MicrosoftEdge,
		"none":                 BrowserSiteTargetEnvironment_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteTargetEnvironment(input)
	return &out, nil
}
