package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningAction string

const (
	ProvisioningAction_Create       ProvisioningAction = "create"
	ProvisioningAction_Delete       ProvisioningAction = "delete"
	ProvisioningAction_Disable      ProvisioningAction = "disable"
	ProvisioningAction_Other        ProvisioningAction = "other"
	ProvisioningAction_StagedDelete ProvisioningAction = "stagedDelete"
	ProvisioningAction_Update       ProvisioningAction = "update"
)

func PossibleValuesForProvisioningAction() []string {
	return []string{
		string(ProvisioningAction_Create),
		string(ProvisioningAction_Delete),
		string(ProvisioningAction_Disable),
		string(ProvisioningAction_Other),
		string(ProvisioningAction_StagedDelete),
		string(ProvisioningAction_Update),
	}
}

func (s *ProvisioningAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningAction(input string) (*ProvisioningAction, error) {
	vals := map[string]ProvisioningAction{
		"create":       ProvisioningAction_Create,
		"delete":       ProvisioningAction_Delete,
		"disable":      ProvisioningAction_Disable,
		"other":        ProvisioningAction_Other,
		"stageddelete": ProvisioningAction_StagedDelete,
		"update":       ProvisioningAction_Update,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningAction(input)
	return &out, nil
}
