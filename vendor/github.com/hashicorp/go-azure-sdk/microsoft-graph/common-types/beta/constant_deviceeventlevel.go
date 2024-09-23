package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEventLevel string

const (
	DeviceEventLevel_Critical    DeviceEventLevel = "critical"
	DeviceEventLevel_Error       DeviceEventLevel = "error"
	DeviceEventLevel_Information DeviceEventLevel = "information"
	DeviceEventLevel_None        DeviceEventLevel = "none"
	DeviceEventLevel_Verbose     DeviceEventLevel = "verbose"
	DeviceEventLevel_Warning     DeviceEventLevel = "warning"
)

func PossibleValuesForDeviceEventLevel() []string {
	return []string{
		string(DeviceEventLevel_Critical),
		string(DeviceEventLevel_Error),
		string(DeviceEventLevel_Information),
		string(DeviceEventLevel_None),
		string(DeviceEventLevel_Verbose),
		string(DeviceEventLevel_Warning),
	}
}

func (s *DeviceEventLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEventLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEventLevel(input string) (*DeviceEventLevel, error) {
	vals := map[string]DeviceEventLevel{
		"critical":    DeviceEventLevel_Critical,
		"error":       DeviceEventLevel_Error,
		"information": DeviceEventLevel_Information,
		"none":        DeviceEventLevel_None,
		"verbose":     DeviceEventLevel_Verbose,
		"warning":     DeviceEventLevel_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEventLevel(input)
	return &out, nil
}
