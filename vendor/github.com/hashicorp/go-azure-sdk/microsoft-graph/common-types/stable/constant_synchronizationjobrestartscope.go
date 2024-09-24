package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobRestartScope string

const (
	SynchronizationJobRestartScope_ConnectorDataStore SynchronizationJobRestartScope = "ConnectorDataStore"
	SynchronizationJobRestartScope_Escrows            SynchronizationJobRestartScope = "Escrows"
	SynchronizationJobRestartScope_ForceDeletes       SynchronizationJobRestartScope = "ForceDeletes"
	SynchronizationJobRestartScope_Full               SynchronizationJobRestartScope = "Full"
	SynchronizationJobRestartScope_None               SynchronizationJobRestartScope = "None"
	SynchronizationJobRestartScope_QuarantineState    SynchronizationJobRestartScope = "QuarantineState"
	SynchronizationJobRestartScope_Watermark          SynchronizationJobRestartScope = "Watermark"
)

func PossibleValuesForSynchronizationJobRestartScope() []string {
	return []string{
		string(SynchronizationJobRestartScope_ConnectorDataStore),
		string(SynchronizationJobRestartScope_Escrows),
		string(SynchronizationJobRestartScope_ForceDeletes),
		string(SynchronizationJobRestartScope_Full),
		string(SynchronizationJobRestartScope_None),
		string(SynchronizationJobRestartScope_QuarantineState),
		string(SynchronizationJobRestartScope_Watermark),
	}
}

func (s *SynchronizationJobRestartScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationJobRestartScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationJobRestartScope(input string) (*SynchronizationJobRestartScope, error) {
	vals := map[string]SynchronizationJobRestartScope{
		"connectordatastore": SynchronizationJobRestartScope_ConnectorDataStore,
		"escrows":            SynchronizationJobRestartScope_Escrows,
		"forcedeletes":       SynchronizationJobRestartScope_ForceDeletes,
		"full":               SynchronizationJobRestartScope_Full,
		"none":               SynchronizationJobRestartScope_None,
		"quarantinestate":    SynchronizationJobRestartScope_QuarantineState,
		"watermark":          SynchronizationJobRestartScope_Watermark,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationJobRestartScope(input)
	return &out, nil
}
