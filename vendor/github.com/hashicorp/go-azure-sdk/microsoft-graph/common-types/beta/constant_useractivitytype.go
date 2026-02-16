package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserActivityType string

const (
	UserActivityType_DownloadFile UserActivityType = "downloadFile"
	UserActivityType_DownloadText UserActivityType = "downloadText"
	UserActivityType_UploadFile   UserActivityType = "uploadFile"
	UserActivityType_UploadText   UserActivityType = "uploadText"
)

func PossibleValuesForUserActivityType() []string {
	return []string{
		string(UserActivityType_DownloadFile),
		string(UserActivityType_DownloadText),
		string(UserActivityType_UploadFile),
		string(UserActivityType_UploadText),
	}
}

func (s *UserActivityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserActivityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserActivityType(input string) (*UserActivityType, error) {
	vals := map[string]UserActivityType{
		"downloadfile": UserActivityType_DownloadFile,
		"downloadtext": UserActivityType_DownloadText,
		"uploadfile":   UserActivityType_UploadFile,
		"uploadtext":   UserActivityType_UploadText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserActivityType(input)
	return &out, nil
}
