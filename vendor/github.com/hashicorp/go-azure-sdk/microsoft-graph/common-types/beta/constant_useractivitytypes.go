package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserActivityTypes string

const (
	UserActivityTypes_DownloadFile UserActivityTypes = "downloadFile"
	UserActivityTypes_DownloadText UserActivityTypes = "downloadText"
	UserActivityTypes_None         UserActivityTypes = "none"
	UserActivityTypes_UploadFile   UserActivityTypes = "uploadFile"
	UserActivityTypes_UploadText   UserActivityTypes = "uploadText"
)

func PossibleValuesForUserActivityTypes() []string {
	return []string{
		string(UserActivityTypes_DownloadFile),
		string(UserActivityTypes_DownloadText),
		string(UserActivityTypes_None),
		string(UserActivityTypes_UploadFile),
		string(UserActivityTypes_UploadText),
	}
}

func (s *UserActivityTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserActivityTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserActivityTypes(input string) (*UserActivityTypes, error) {
	vals := map[string]UserActivityTypes{
		"downloadfile": UserActivityTypes_DownloadFile,
		"downloadtext": UserActivityTypes_DownloadText,
		"none":         UserActivityTypes_None,
		"uploadfile":   UserActivityTypes_UploadFile,
		"uploadtext":   UserActivityTypes_UploadText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserActivityTypes(input)
	return &out, nil
}
