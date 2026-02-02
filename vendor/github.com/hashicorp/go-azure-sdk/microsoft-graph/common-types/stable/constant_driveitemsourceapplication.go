package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemSourceApplication string

const (
	DriveItemSourceApplication_Loki       DriveItemSourceApplication = "loki"
	DriveItemSourceApplication_Loop       DriveItemSourceApplication = "loop"
	DriveItemSourceApplication_Office     DriveItemSourceApplication = "office"
	DriveItemSourceApplication_OneDrive   DriveItemSourceApplication = "oneDrive"
	DriveItemSourceApplication_Other      DriveItemSourceApplication = "other"
	DriveItemSourceApplication_PowerPoint DriveItemSourceApplication = "powerPoint"
	DriveItemSourceApplication_SharePoint DriveItemSourceApplication = "sharePoint"
	DriveItemSourceApplication_Stream     DriveItemSourceApplication = "stream"
	DriveItemSourceApplication_Teams      DriveItemSourceApplication = "teams"
	DriveItemSourceApplication_Yammer     DriveItemSourceApplication = "yammer"
)

func PossibleValuesForDriveItemSourceApplication() []string {
	return []string{
		string(DriveItemSourceApplication_Loki),
		string(DriveItemSourceApplication_Loop),
		string(DriveItemSourceApplication_Office),
		string(DriveItemSourceApplication_OneDrive),
		string(DriveItemSourceApplication_Other),
		string(DriveItemSourceApplication_PowerPoint),
		string(DriveItemSourceApplication_SharePoint),
		string(DriveItemSourceApplication_Stream),
		string(DriveItemSourceApplication_Teams),
		string(DriveItemSourceApplication_Yammer),
	}
}

func (s *DriveItemSourceApplication) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDriveItemSourceApplication(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDriveItemSourceApplication(input string) (*DriveItemSourceApplication, error) {
	vals := map[string]DriveItemSourceApplication{
		"loki":       DriveItemSourceApplication_Loki,
		"loop":       DriveItemSourceApplication_Loop,
		"office":     DriveItemSourceApplication_Office,
		"onedrive":   DriveItemSourceApplication_OneDrive,
		"other":      DriveItemSourceApplication_Other,
		"powerpoint": DriveItemSourceApplication_PowerPoint,
		"sharepoint": DriveItemSourceApplication_SharePoint,
		"stream":     DriveItemSourceApplication_Stream,
		"teams":      DriveItemSourceApplication_Teams,
		"yammer":     DriveItemSourceApplication_Yammer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DriveItemSourceApplication(input)
	return &out, nil
}
