package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectFlowTypes string

const (
	ObjectFlowTypes_Add    ObjectFlowTypes = "Add"
	ObjectFlowTypes_Delete ObjectFlowTypes = "Delete"
	ObjectFlowTypes_None   ObjectFlowTypes = "None"
	ObjectFlowTypes_Update ObjectFlowTypes = "Update"
)

func PossibleValuesForObjectFlowTypes() []string {
	return []string{
		string(ObjectFlowTypes_Add),
		string(ObjectFlowTypes_Delete),
		string(ObjectFlowTypes_None),
		string(ObjectFlowTypes_Update),
	}
}

func (s *ObjectFlowTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObjectFlowTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObjectFlowTypes(input string) (*ObjectFlowTypes, error) {
	vals := map[string]ObjectFlowTypes{
		"add":    ObjectFlowTypes_Add,
		"delete": ObjectFlowTypes_Delete,
		"none":   ObjectFlowTypes_None,
		"update": ObjectFlowTypes_Update,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectFlowTypes(input)
	return &out, nil
}
