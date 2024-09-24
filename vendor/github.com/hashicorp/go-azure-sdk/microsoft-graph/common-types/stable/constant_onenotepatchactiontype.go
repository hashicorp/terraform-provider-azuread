package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePatchActionType string

const (
	OnenotePatchActionType_Append  OnenotePatchActionType = "Append"
	OnenotePatchActionType_Delete  OnenotePatchActionType = "Delete"
	OnenotePatchActionType_Insert  OnenotePatchActionType = "Insert"
	OnenotePatchActionType_Prepend OnenotePatchActionType = "Prepend"
	OnenotePatchActionType_Replace OnenotePatchActionType = "Replace"
)

func PossibleValuesForOnenotePatchActionType() []string {
	return []string{
		string(OnenotePatchActionType_Append),
		string(OnenotePatchActionType_Delete),
		string(OnenotePatchActionType_Insert),
		string(OnenotePatchActionType_Prepend),
		string(OnenotePatchActionType_Replace),
	}
}

func (s *OnenotePatchActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnenotePatchActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnenotePatchActionType(input string) (*OnenotePatchActionType, error) {
	vals := map[string]OnenotePatchActionType{
		"append":  OnenotePatchActionType_Append,
		"delete":  OnenotePatchActionType_Delete,
		"insert":  OnenotePatchActionType_Insert,
		"prepend": OnenotePatchActionType_Prepend,
		"replace": OnenotePatchActionType_Replace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnenotePatchActionType(input)
	return &out, nil
}
