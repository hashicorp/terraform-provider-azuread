package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateStore string

const (
	CertificateStore_Machine CertificateStore = "machine"
	CertificateStore_User    CertificateStore = "user"
)

func PossibleValuesForCertificateStore() []string {
	return []string{
		string(CertificateStore_Machine),
		string(CertificateStore_User),
	}
}

func (s *CertificateStore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateStore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateStore(input string) (*CertificateStore, error) {
	vals := map[string]CertificateStore{
		"machine": CertificateStore_Machine,
		"user":    CertificateStore_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateStore(input)
	return &out, nil
}
