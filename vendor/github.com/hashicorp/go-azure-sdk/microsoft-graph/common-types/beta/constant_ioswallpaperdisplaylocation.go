package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosWallpaperDisplayLocation string

const (
	IosWallpaperDisplayLocation_HomeScreen         IosWallpaperDisplayLocation = "homeScreen"
	IosWallpaperDisplayLocation_LockAndHomeScreens IosWallpaperDisplayLocation = "lockAndHomeScreens"
	IosWallpaperDisplayLocation_LockScreen         IosWallpaperDisplayLocation = "lockScreen"
	IosWallpaperDisplayLocation_NotConfigured      IosWallpaperDisplayLocation = "notConfigured"
)

func PossibleValuesForIosWallpaperDisplayLocation() []string {
	return []string{
		string(IosWallpaperDisplayLocation_HomeScreen),
		string(IosWallpaperDisplayLocation_LockAndHomeScreens),
		string(IosWallpaperDisplayLocation_LockScreen),
		string(IosWallpaperDisplayLocation_NotConfigured),
	}
}

func (s *IosWallpaperDisplayLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosWallpaperDisplayLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosWallpaperDisplayLocation(input string) (*IosWallpaperDisplayLocation, error) {
	vals := map[string]IosWallpaperDisplayLocation{
		"homescreen":         IosWallpaperDisplayLocation_HomeScreen,
		"lockandhomescreens": IosWallpaperDisplayLocation_LockAndHomeScreens,
		"lockscreen":         IosWallpaperDisplayLocation_LockScreen,
		"notconfigured":      IosWallpaperDisplayLocation_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosWallpaperDisplayLocation(input)
	return &out, nil
}
