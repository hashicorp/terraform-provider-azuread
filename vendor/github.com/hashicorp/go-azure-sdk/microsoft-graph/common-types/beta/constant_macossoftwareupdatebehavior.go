package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateBehavior string

const (
	MacOSSoftwareUpdateBehavior_Default       MacOSSoftwareUpdateBehavior = "default"
	MacOSSoftwareUpdateBehavior_DownloadOnly  MacOSSoftwareUpdateBehavior = "downloadOnly"
	MacOSSoftwareUpdateBehavior_InstallASAP   MacOSSoftwareUpdateBehavior = "installASAP"
	MacOSSoftwareUpdateBehavior_InstallLater  MacOSSoftwareUpdateBehavior = "installLater"
	MacOSSoftwareUpdateBehavior_NotConfigured MacOSSoftwareUpdateBehavior = "notConfigured"
	MacOSSoftwareUpdateBehavior_NotifyOnly    MacOSSoftwareUpdateBehavior = "notifyOnly"
)

func PossibleValuesForMacOSSoftwareUpdateBehavior() []string {
	return []string{
		string(MacOSSoftwareUpdateBehavior_Default),
		string(MacOSSoftwareUpdateBehavior_DownloadOnly),
		string(MacOSSoftwareUpdateBehavior_InstallASAP),
		string(MacOSSoftwareUpdateBehavior_InstallLater),
		string(MacOSSoftwareUpdateBehavior_NotConfigured),
		string(MacOSSoftwareUpdateBehavior_NotifyOnly),
	}
}

func (s *MacOSSoftwareUpdateBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateBehavior(input string) (*MacOSSoftwareUpdateBehavior, error) {
	vals := map[string]MacOSSoftwareUpdateBehavior{
		"default":       MacOSSoftwareUpdateBehavior_Default,
		"downloadonly":  MacOSSoftwareUpdateBehavior_DownloadOnly,
		"installasap":   MacOSSoftwareUpdateBehavior_InstallASAP,
		"installlater":  MacOSSoftwareUpdateBehavior_InstallLater,
		"notconfigured": MacOSSoftwareUpdateBehavior_NotConfigured,
		"notifyonly":    MacOSSoftwareUpdateBehavior_NotifyOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateBehavior(input)
	return &out, nil
}
