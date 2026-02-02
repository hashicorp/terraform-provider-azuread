package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateDestinationStore string

const (
	CertificateDestinationStore_ComputerCertStoreIntermediate CertificateDestinationStore = "computerCertStoreIntermediate"
	CertificateDestinationStore_ComputerCertStoreRoot         CertificateDestinationStore = "computerCertStoreRoot"
	CertificateDestinationStore_UserCertStoreIntermediate     CertificateDestinationStore = "userCertStoreIntermediate"
)

func PossibleValuesForCertificateDestinationStore() []string {
	return []string{
		string(CertificateDestinationStore_ComputerCertStoreIntermediate),
		string(CertificateDestinationStore_ComputerCertStoreRoot),
		string(CertificateDestinationStore_UserCertStoreIntermediate),
	}
}

func (s *CertificateDestinationStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateDestinationStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateDestinationStore(input string) (*CertificateDestinationStore, error) {
	vals := map[string]CertificateDestinationStore{
		"computercertstoreintermediate": CertificateDestinationStore_ComputerCertStoreIntermediate,
		"computercertstoreroot":         CertificateDestinationStore_ComputerCertStoreRoot,
		"usercertstoreintermediate":     CertificateDestinationStore_UserCertStoreIntermediate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateDestinationStore(input)
	return &out, nil
}
