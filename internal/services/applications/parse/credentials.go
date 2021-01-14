package parse

import (
	"fmt"
	"strings"
)

type CredentialId struct {
	ObjectId string
	KeyType  string
	KeyId    string
}

func NewCredentialID(objectId, keyType, keyId string) CredentialId {
	return CredentialId{
		ObjectId: objectId,
		KeyType:  keyType,
		KeyId:    keyId,
	}
}

func (id CredentialId) String() string {
	return id.ObjectId + "/" + id.KeyType + "/" + id.KeyId
}

func CertificateID(idString string) (*CredentialId, error) {
	id, err := ObjectSubResourceID(idString, "certificate")
	if err != nil {
		return nil, fmt.Errorf("unable to parse Certificate ID: %v", err)
	}

	return &CredentialId{
		ObjectId: id.objectId,
		KeyType:  id.Type,
		KeyId:    id.subId,
	}, nil
}

func PasswordID(idString string) (*CredentialId, error) {
	id, err := ObjectSubResourceID(idString, "password")
	if err != nil {
		return nil, fmt.Errorf("unable to parse Password ID: %v", err)
	}

	return &CredentialId{
		ObjectId: id.objectId,
		KeyType:  id.Type,
		KeyId:    id.subId,
	}, nil
}

func OldPasswordID(id string) (*CredentialId, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Password ID expected to be in the format {objectId}/{keyId} - but got %q", id)
	}

	newId := parts[0] + "/password/" + parts[1]
	return PasswordID(newId)
}
