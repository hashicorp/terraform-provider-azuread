package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ErrorCode string

const (
	ErrorCode_Deleted      ErrorCode = "deleted"
	ErrorCode_NoError      ErrorCode = "noError"
	ErrorCode_NotFound     ErrorCode = "notFound"
	ErrorCode_Unauthorized ErrorCode = "unauthorized"
)

func PossibleValuesForErrorCode() []string {
	return []string{
		string(ErrorCode_Deleted),
		string(ErrorCode_NoError),
		string(ErrorCode_NotFound),
		string(ErrorCode_Unauthorized),
	}
}

func (s *ErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseErrorCode(input string) (*ErrorCode, error) {
	vals := map[string]ErrorCode{
		"deleted":      ErrorCode_Deleted,
		"noerror":      ErrorCode_NoError,
		"notfound":     ErrorCode_NotFound,
		"unauthorized": ErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ErrorCode(input)
	return &out, nil
}
