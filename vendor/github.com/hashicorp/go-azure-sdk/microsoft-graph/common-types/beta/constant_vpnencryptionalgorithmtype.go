package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnEncryptionAlgorithmType string

const (
	VpnEncryptionAlgorithmType_Aes128           VpnEncryptionAlgorithmType = "aes128"
	VpnEncryptionAlgorithmType_Aes128Gcm        VpnEncryptionAlgorithmType = "aes128Gcm"
	VpnEncryptionAlgorithmType_Aes192           VpnEncryptionAlgorithmType = "aes192"
	VpnEncryptionAlgorithmType_Aes192Gcm        VpnEncryptionAlgorithmType = "aes192Gcm"
	VpnEncryptionAlgorithmType_Aes256           VpnEncryptionAlgorithmType = "aes256"
	VpnEncryptionAlgorithmType_Aes256Gcm        VpnEncryptionAlgorithmType = "aes256Gcm"
	VpnEncryptionAlgorithmType_ChaCha20Poly1305 VpnEncryptionAlgorithmType = "chaCha20Poly1305"
	VpnEncryptionAlgorithmType_Des              VpnEncryptionAlgorithmType = "des"
	VpnEncryptionAlgorithmType_TripleDes        VpnEncryptionAlgorithmType = "tripleDes"
)

func PossibleValuesForVpnEncryptionAlgorithmType() []string {
	return []string{
		string(VpnEncryptionAlgorithmType_Aes128),
		string(VpnEncryptionAlgorithmType_Aes128Gcm),
		string(VpnEncryptionAlgorithmType_Aes192),
		string(VpnEncryptionAlgorithmType_Aes192Gcm),
		string(VpnEncryptionAlgorithmType_Aes256),
		string(VpnEncryptionAlgorithmType_Aes256Gcm),
		string(VpnEncryptionAlgorithmType_ChaCha20Poly1305),
		string(VpnEncryptionAlgorithmType_Des),
		string(VpnEncryptionAlgorithmType_TripleDes),
	}
}

func (s *VpnEncryptionAlgorithmType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnEncryptionAlgorithmType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnEncryptionAlgorithmType(input string) (*VpnEncryptionAlgorithmType, error) {
	vals := map[string]VpnEncryptionAlgorithmType{
		"aes128":           VpnEncryptionAlgorithmType_Aes128,
		"aes128gcm":        VpnEncryptionAlgorithmType_Aes128Gcm,
		"aes192":           VpnEncryptionAlgorithmType_Aes192,
		"aes192gcm":        VpnEncryptionAlgorithmType_Aes192Gcm,
		"aes256":           VpnEncryptionAlgorithmType_Aes256,
		"aes256gcm":        VpnEncryptionAlgorithmType_Aes256Gcm,
		"chacha20poly1305": VpnEncryptionAlgorithmType_ChaCha20Poly1305,
		"des":              VpnEncryptionAlgorithmType_Des,
		"tripledes":        VpnEncryptionAlgorithmType_TripleDes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnEncryptionAlgorithmType(input)
	return &out, nil
}
