package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessClientApp string

const (
	ConditionalAccessClientApp_All                         ConditionalAccessClientApp = "all"
	ConditionalAccessClientApp_Browser                     ConditionalAccessClientApp = "browser"
	ConditionalAccessClientApp_EasSupported                ConditionalAccessClientApp = "easSupported"
	ConditionalAccessClientApp_ExchangeActiveSync          ConditionalAccessClientApp = "exchangeActiveSync"
	ConditionalAccessClientApp_MobileAppsAndDesktopClients ConditionalAccessClientApp = "mobileAppsAndDesktopClients"
	ConditionalAccessClientApp_Other                       ConditionalAccessClientApp = "other"
)

func PossibleValuesForConditionalAccessClientApp() []string {
	return []string{
		string(ConditionalAccessClientApp_All),
		string(ConditionalAccessClientApp_Browser),
		string(ConditionalAccessClientApp_EasSupported),
		string(ConditionalAccessClientApp_ExchangeActiveSync),
		string(ConditionalAccessClientApp_MobileAppsAndDesktopClients),
		string(ConditionalAccessClientApp_Other),
	}
}

func (s *ConditionalAccessClientApp) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessClientApp(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessClientApp(input string) (*ConditionalAccessClientApp, error) {
	vals := map[string]ConditionalAccessClientApp{
		"all":                         ConditionalAccessClientApp_All,
		"browser":                     ConditionalAccessClientApp_Browser,
		"eassupported":                ConditionalAccessClientApp_EasSupported,
		"exchangeactivesync":          ConditionalAccessClientApp_ExchangeActiveSync,
		"mobileappsanddesktopclients": ConditionalAccessClientApp_MobileAppsAndDesktopClients,
		"other":                       ConditionalAccessClientApp_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessClientApp(input)
	return &out, nil
}
