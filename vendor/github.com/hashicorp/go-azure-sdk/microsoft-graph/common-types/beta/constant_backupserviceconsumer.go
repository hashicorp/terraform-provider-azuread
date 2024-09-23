package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupServiceConsumer string

const (
	BackupServiceConsumer_Firstparty BackupServiceConsumer = "firstparty"
	BackupServiceConsumer_Thirdparty BackupServiceConsumer = "thirdparty"
	BackupServiceConsumer_Unknown    BackupServiceConsumer = "unknown"
)

func PossibleValuesForBackupServiceConsumer() []string {
	return []string{
		string(BackupServiceConsumer_Firstparty),
		string(BackupServiceConsumer_Thirdparty),
		string(BackupServiceConsumer_Unknown),
	}
}

func (s *BackupServiceConsumer) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBackupServiceConsumer(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBackupServiceConsumer(input string) (*BackupServiceConsumer, error) {
	vals := map[string]BackupServiceConsumer{
		"firstparty": BackupServiceConsumer_Firstparty,
		"thirdparty": BackupServiceConsumer_Thirdparty,
		"unknown":    BackupServiceConsumer_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupServiceConsumer(input)
	return &out, nil
}
