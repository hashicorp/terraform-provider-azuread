package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSharedCookieSourceEnvironment string

const (
	BrowserSharedCookieSourceEnvironment_Both               BrowserSharedCookieSourceEnvironment = "both"
	BrowserSharedCookieSourceEnvironment_InternetExplorer11 BrowserSharedCookieSourceEnvironment = "internetExplorer11"
	BrowserSharedCookieSourceEnvironment_MicrosoftEdge      BrowserSharedCookieSourceEnvironment = "microsoftEdge"
)

func PossibleValuesForBrowserSharedCookieSourceEnvironment() []string {
	return []string{
		string(BrowserSharedCookieSourceEnvironment_Both),
		string(BrowserSharedCookieSourceEnvironment_InternetExplorer11),
		string(BrowserSharedCookieSourceEnvironment_MicrosoftEdge),
	}
}

func (s *BrowserSharedCookieSourceEnvironment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSharedCookieSourceEnvironment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSharedCookieSourceEnvironment(input string) (*BrowserSharedCookieSourceEnvironment, error) {
	vals := map[string]BrowserSharedCookieSourceEnvironment{
		"both":               BrowserSharedCookieSourceEnvironment_Both,
		"internetexplorer11": BrowserSharedCookieSourceEnvironment_InternetExplorer11,
		"microsoftedge":      BrowserSharedCookieSourceEnvironment_MicrosoftEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSharedCookieSourceEnvironment(input)
	return &out, nil
}
