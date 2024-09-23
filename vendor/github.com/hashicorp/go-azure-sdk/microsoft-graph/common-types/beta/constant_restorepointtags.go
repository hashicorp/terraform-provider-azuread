package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointTags string

const (
	RestorePointTags_FastRestore RestorePointTags = "fastRestore"
	RestorePointTags_None        RestorePointTags = "none"
)

func PossibleValuesForRestorePointTags() []string {
	return []string{
		string(RestorePointTags_FastRestore),
		string(RestorePointTags_None),
	}
}

func (s *RestorePointTags) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestorePointTags(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestorePointTags(input string) (*RestorePointTags, error) {
	vals := map[string]RestorePointTags{
		"fastrestore": RestorePointTags_FastRestore,
		"none":        RestorePointTags_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestorePointTags(input)
	return &out, nil
}
