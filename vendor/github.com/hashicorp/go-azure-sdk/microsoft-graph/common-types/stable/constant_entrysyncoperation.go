package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntrySyncOperation string

const (
	EntrySyncOperation_Add    EntrySyncOperation = "Add"
	EntrySyncOperation_Delete EntrySyncOperation = "Delete"
	EntrySyncOperation_None   EntrySyncOperation = "None"
	EntrySyncOperation_Update EntrySyncOperation = "Update"
)

func PossibleValuesForEntrySyncOperation() []string {
	return []string{
		string(EntrySyncOperation_Add),
		string(EntrySyncOperation_Delete),
		string(EntrySyncOperation_None),
		string(EntrySyncOperation_Update),
	}
}

func (s *EntrySyncOperation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntrySyncOperation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntrySyncOperation(input string) (*EntrySyncOperation, error) {
	vals := map[string]EntrySyncOperation{
		"add":    EntrySyncOperation_Add,
		"delete": EntrySyncOperation_Delete,
		"none":   EntrySyncOperation_None,
		"update": EntrySyncOperation_Update,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntrySyncOperation(input)
	return &out, nil
}
