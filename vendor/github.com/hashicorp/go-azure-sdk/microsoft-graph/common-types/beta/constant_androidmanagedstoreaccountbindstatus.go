package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountBindStatus string

const (
	AndroidManagedStoreAccountBindStatus_Bound             AndroidManagedStoreAccountBindStatus = "bound"
	AndroidManagedStoreAccountBindStatus_BoundAndValidated AndroidManagedStoreAccountBindStatus = "boundAndValidated"
	AndroidManagedStoreAccountBindStatus_NotBound          AndroidManagedStoreAccountBindStatus = "notBound"
	AndroidManagedStoreAccountBindStatus_Unbinding         AndroidManagedStoreAccountBindStatus = "unbinding"
)

func PossibleValuesForAndroidManagedStoreAccountBindStatus() []string {
	return []string{
		string(AndroidManagedStoreAccountBindStatus_Bound),
		string(AndroidManagedStoreAccountBindStatus_BoundAndValidated),
		string(AndroidManagedStoreAccountBindStatus_NotBound),
		string(AndroidManagedStoreAccountBindStatus_Unbinding),
	}
}

func (s *AndroidManagedStoreAccountBindStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreAccountBindStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreAccountBindStatus(input string) (*AndroidManagedStoreAccountBindStatus, error) {
	vals := map[string]AndroidManagedStoreAccountBindStatus{
		"bound":             AndroidManagedStoreAccountBindStatus_Bound,
		"boundandvalidated": AndroidManagedStoreAccountBindStatus_BoundAndValidated,
		"notbound":          AndroidManagedStoreAccountBindStatus_NotBound,
		"unbinding":         AndroidManagedStoreAccountBindStatus_Unbinding,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAccountBindStatus(input)
	return &out, nil
}
