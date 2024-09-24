package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebBrowserCookieSettings string

const (
	WebBrowserCookieSettings_AllowAlways              WebBrowserCookieSettings = "allowAlways"
	WebBrowserCookieSettings_AllowCurrentWebSite      WebBrowserCookieSettings = "allowCurrentWebSite"
	WebBrowserCookieSettings_AllowFromWebsitesVisited WebBrowserCookieSettings = "allowFromWebsitesVisited"
	WebBrowserCookieSettings_BlockAlways              WebBrowserCookieSettings = "blockAlways"
	WebBrowserCookieSettings_BrowserDefault           WebBrowserCookieSettings = "browserDefault"
)

func PossibleValuesForWebBrowserCookieSettings() []string {
	return []string{
		string(WebBrowserCookieSettings_AllowAlways),
		string(WebBrowserCookieSettings_AllowCurrentWebSite),
		string(WebBrowserCookieSettings_AllowFromWebsitesVisited),
		string(WebBrowserCookieSettings_BlockAlways),
		string(WebBrowserCookieSettings_BrowserDefault),
	}
}

func (s *WebBrowserCookieSettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebBrowserCookieSettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebBrowserCookieSettings(input string) (*WebBrowserCookieSettings, error) {
	vals := map[string]WebBrowserCookieSettings{
		"allowalways":              WebBrowserCookieSettings_AllowAlways,
		"allowcurrentwebsite":      WebBrowserCookieSettings_AllowCurrentWebSite,
		"allowfromwebsitesvisited": WebBrowserCookieSettings_AllowFromWebsitesVisited,
		"blockalways":              WebBrowserCookieSettings_BlockAlways,
		"browserdefault":           WebBrowserCookieSettings_BrowserDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebBrowserCookieSettings(input)
	return &out, nil
}
