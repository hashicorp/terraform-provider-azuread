package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentProcessingErrorType string

const (
	ContentProcessingErrorType_Permanent ContentProcessingErrorType = "permanent"
	ContentProcessingErrorType_Transient ContentProcessingErrorType = "transient"
)

func PossibleValuesForContentProcessingErrorType() []string {
	return []string{
		string(ContentProcessingErrorType_Permanent),
		string(ContentProcessingErrorType_Transient),
	}
}

func (s *ContentProcessingErrorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContentProcessingErrorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContentProcessingErrorType(input string) (*ContentProcessingErrorType, error) {
	vals := map[string]ContentProcessingErrorType{
		"permanent": ContentProcessingErrorType_Permanent,
		"transient": ContentProcessingErrorType_Transient,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentProcessingErrorType(input)
	return &out, nil
}
