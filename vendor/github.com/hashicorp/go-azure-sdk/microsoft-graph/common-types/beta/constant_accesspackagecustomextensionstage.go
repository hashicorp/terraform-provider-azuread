package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageCustomExtensionStage string

const (
	AccessPackageCustomExtensionStage_AssignmentFourteenDaysBeforeExpiration AccessPackageCustomExtensionStage = "assignmentFourteenDaysBeforeExpiration"
	AccessPackageCustomExtensionStage_AssignmentOneDayBeforeExpiration       AccessPackageCustomExtensionStage = "assignmentOneDayBeforeExpiration"
	AccessPackageCustomExtensionStage_AssignmentRequestApproved              AccessPackageCustomExtensionStage = "assignmentRequestApproved"
	AccessPackageCustomExtensionStage_AssignmentRequestCreated               AccessPackageCustomExtensionStage = "assignmentRequestCreated"
	AccessPackageCustomExtensionStage_AssignmentRequestGranted               AccessPackageCustomExtensionStage = "assignmentRequestGranted"
	AccessPackageCustomExtensionStage_AssignmentRequestRemoved               AccessPackageCustomExtensionStage = "assignmentRequestRemoved"
)

func PossibleValuesForAccessPackageCustomExtensionStage() []string {
	return []string{
		string(AccessPackageCustomExtensionStage_AssignmentFourteenDaysBeforeExpiration),
		string(AccessPackageCustomExtensionStage_AssignmentOneDayBeforeExpiration),
		string(AccessPackageCustomExtensionStage_AssignmentRequestApproved),
		string(AccessPackageCustomExtensionStage_AssignmentRequestCreated),
		string(AccessPackageCustomExtensionStage_AssignmentRequestGranted),
		string(AccessPackageCustomExtensionStage_AssignmentRequestRemoved),
	}
}

func (s *AccessPackageCustomExtensionStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageCustomExtensionStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageCustomExtensionStage(input string) (*AccessPackageCustomExtensionStage, error) {
	vals := map[string]AccessPackageCustomExtensionStage{
		"assignmentfourteendaysbeforeexpiration": AccessPackageCustomExtensionStage_AssignmentFourteenDaysBeforeExpiration,
		"assignmentonedaybeforeexpiration":       AccessPackageCustomExtensionStage_AssignmentOneDayBeforeExpiration,
		"assignmentrequestapproved":              AccessPackageCustomExtensionStage_AssignmentRequestApproved,
		"assignmentrequestcreated":               AccessPackageCustomExtensionStage_AssignmentRequestCreated,
		"assignmentrequestgranted":               AccessPackageCustomExtensionStage_AssignmentRequestGranted,
		"assignmentrequestremoved":               AccessPackageCustomExtensionStage_AssignmentRequestRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageCustomExtensionStage(input)
	return &out, nil
}
