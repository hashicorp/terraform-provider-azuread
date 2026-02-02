package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateState string

const (
	MacOSSoftwareUpdateState_Available                   MacOSSoftwareUpdateState = "available"
	MacOSSoftwareUpdateState_CommandFailed               MacOSSoftwareUpdateState = "commandFailed"
	MacOSSoftwareUpdateState_DownloadFailed              MacOSSoftwareUpdateState = "downloadFailed"
	MacOSSoftwareUpdateState_DownloadInsufficientNetwork MacOSSoftwareUpdateState = "downloadInsufficientNetwork"
	MacOSSoftwareUpdateState_DownloadInsufficientPower   MacOSSoftwareUpdateState = "downloadInsufficientPower"
	MacOSSoftwareUpdateState_DownloadInsufficientSpace   MacOSSoftwareUpdateState = "downloadInsufficientSpace"
	MacOSSoftwareUpdateState_Downloaded                  MacOSSoftwareUpdateState = "downloaded"
	MacOSSoftwareUpdateState_Downloading                 MacOSSoftwareUpdateState = "downloading"
	MacOSSoftwareUpdateState_Idle                        MacOSSoftwareUpdateState = "idle"
	MacOSSoftwareUpdateState_InstallFailed               MacOSSoftwareUpdateState = "installFailed"
	MacOSSoftwareUpdateState_InstallInsufficientPower    MacOSSoftwareUpdateState = "installInsufficientPower"
	MacOSSoftwareUpdateState_InstallInsufficientSpace    MacOSSoftwareUpdateState = "installInsufficientSpace"
	MacOSSoftwareUpdateState_Installing                  MacOSSoftwareUpdateState = "installing"
	MacOSSoftwareUpdateState_Scheduled                   MacOSSoftwareUpdateState = "scheduled"
	MacOSSoftwareUpdateState_Success                     MacOSSoftwareUpdateState = "success"
)

func PossibleValuesForMacOSSoftwareUpdateState() []string {
	return []string{
		string(MacOSSoftwareUpdateState_Available),
		string(MacOSSoftwareUpdateState_CommandFailed),
		string(MacOSSoftwareUpdateState_DownloadFailed),
		string(MacOSSoftwareUpdateState_DownloadInsufficientNetwork),
		string(MacOSSoftwareUpdateState_DownloadInsufficientPower),
		string(MacOSSoftwareUpdateState_DownloadInsufficientSpace),
		string(MacOSSoftwareUpdateState_Downloaded),
		string(MacOSSoftwareUpdateState_Downloading),
		string(MacOSSoftwareUpdateState_Idle),
		string(MacOSSoftwareUpdateState_InstallFailed),
		string(MacOSSoftwareUpdateState_InstallInsufficientPower),
		string(MacOSSoftwareUpdateState_InstallInsufficientSpace),
		string(MacOSSoftwareUpdateState_Installing),
		string(MacOSSoftwareUpdateState_Scheduled),
		string(MacOSSoftwareUpdateState_Success),
	}
}

func (s *MacOSSoftwareUpdateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateState(input string) (*MacOSSoftwareUpdateState, error) {
	vals := map[string]MacOSSoftwareUpdateState{
		"available":                   MacOSSoftwareUpdateState_Available,
		"commandfailed":               MacOSSoftwareUpdateState_CommandFailed,
		"downloadfailed":              MacOSSoftwareUpdateState_DownloadFailed,
		"downloadinsufficientnetwork": MacOSSoftwareUpdateState_DownloadInsufficientNetwork,
		"downloadinsufficientpower":   MacOSSoftwareUpdateState_DownloadInsufficientPower,
		"downloadinsufficientspace":   MacOSSoftwareUpdateState_DownloadInsufficientSpace,
		"downloaded":                  MacOSSoftwareUpdateState_Downloaded,
		"downloading":                 MacOSSoftwareUpdateState_Downloading,
		"idle":                        MacOSSoftwareUpdateState_Idle,
		"installfailed":               MacOSSoftwareUpdateState_InstallFailed,
		"installinsufficientpower":    MacOSSoftwareUpdateState_InstallInsufficientPower,
		"installinsufficientspace":    MacOSSoftwareUpdateState_InstallInsufficientSpace,
		"installing":                  MacOSSoftwareUpdateState_Installing,
		"scheduled":                   MacOSSoftwareUpdateState_Scheduled,
		"success":                     MacOSSoftwareUpdateState_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateState(input)
	return &out, nil
}
