package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsSecretInformationWebServices string

const (
	AwsSecretInformationWebServices_CertificateAuthority AwsSecretInformationWebServices = "certificateAuthority"
	AwsSecretInformationWebServices_CertificateManager   AwsSecretInformationWebServices = "certificateManager"
	AwsSecretInformationWebServices_CloudHsm             AwsSecretInformationWebServices = "cloudHsm"
	AwsSecretInformationWebServices_SecretsManager       AwsSecretInformationWebServices = "secretsManager"
)

func PossibleValuesForAwsSecretInformationWebServices() []string {
	return []string{
		string(AwsSecretInformationWebServices_CertificateAuthority),
		string(AwsSecretInformationWebServices_CertificateManager),
		string(AwsSecretInformationWebServices_CloudHsm),
		string(AwsSecretInformationWebServices_SecretsManager),
	}
}

func (s *AwsSecretInformationWebServices) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsSecretInformationWebServices(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsSecretInformationWebServices(input string) (*AwsSecretInformationWebServices, error) {
	vals := map[string]AwsSecretInformationWebServices{
		"certificateauthority": AwsSecretInformationWebServices_CertificateAuthority,
		"certificatemanager":   AwsSecretInformationWebServices_CertificateManager,
		"cloudhsm":             AwsSecretInformationWebServices_CloudHsm,
		"secretsmanager":       AwsSecretInformationWebServices_SecretsManager,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsSecretInformationWebServices(input)
	return &out, nil
}
