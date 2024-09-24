package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePatchInsertPosition string

const (
	OnenotePatchInsertPosition_After  OnenotePatchInsertPosition = "After"
	OnenotePatchInsertPosition_Before OnenotePatchInsertPosition = "Before"
)

func PossibleValuesForOnenotePatchInsertPosition() []string {
	return []string{
		string(OnenotePatchInsertPosition_After),
		string(OnenotePatchInsertPosition_Before),
	}
}

func (s *OnenotePatchInsertPosition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnenotePatchInsertPosition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnenotePatchInsertPosition(input string) (*OnenotePatchInsertPosition, error) {
	vals := map[string]OnenotePatchInsertPosition{
		"after":  OnenotePatchInsertPosition_After,
		"before": OnenotePatchInsertPosition_Before,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnenotePatchInsertPosition(input)
	return &out, nil
}
