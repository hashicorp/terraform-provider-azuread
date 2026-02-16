package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationAttackTechnique string

const (
	SimulationAttackTechnique_AttachmentMalware    SimulationAttackTechnique = "attachmentMalware"
	SimulationAttackTechnique_CredentialHarvesting SimulationAttackTechnique = "credentialHarvesting"
	SimulationAttackTechnique_DriveByUrl           SimulationAttackTechnique = "driveByUrl"
	SimulationAttackTechnique_LinkInAttachment     SimulationAttackTechnique = "linkInAttachment"
	SimulationAttackTechnique_LinkToMalwareFile    SimulationAttackTechnique = "linkToMalwareFile"
	SimulationAttackTechnique_Unknown              SimulationAttackTechnique = "unknown"
)

func PossibleValuesForSimulationAttackTechnique() []string {
	return []string{
		string(SimulationAttackTechnique_AttachmentMalware),
		string(SimulationAttackTechnique_CredentialHarvesting),
		string(SimulationAttackTechnique_DriveByUrl),
		string(SimulationAttackTechnique_LinkInAttachment),
		string(SimulationAttackTechnique_LinkToMalwareFile),
		string(SimulationAttackTechnique_Unknown),
	}
}

func (s *SimulationAttackTechnique) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSimulationAttackTechnique(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSimulationAttackTechnique(input string) (*SimulationAttackTechnique, error) {
	vals := map[string]SimulationAttackTechnique{
		"attachmentmalware":    SimulationAttackTechnique_AttachmentMalware,
		"credentialharvesting": SimulationAttackTechnique_CredentialHarvesting,
		"drivebyurl":           SimulationAttackTechnique_DriveByUrl,
		"linkinattachment":     SimulationAttackTechnique_LinkInAttachment,
		"linktomalwarefile":    SimulationAttackTechnique_LinkToMalwareFile,
		"unknown":              SimulationAttackTechnique_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationAttackTechnique(input)
	return &out, nil
}
