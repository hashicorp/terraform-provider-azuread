package parse

import "fmt"

type AdministrativeUnitMemberId struct {
	ObjectSubResourceId
	AdministrativeUnitId string
	MemberId             string
}

func NewAdministrativeUnitMemberID(groupId, memberId string) AdministrativeUnitMemberId {
	return AdministrativeUnitMemberId{
		ObjectSubResourceId:  NewObjectSubResourceID(groupId, "member", memberId),
		AdministrativeUnitId: groupId,
		MemberId:             memberId,
	}
}

func AdministrativeUnitMemberID(idString string) (*AdministrativeUnitMemberId, error) {
	id, err := ObjectSubResourceID(idString, "member")
	if err != nil {
		return nil, fmt.Errorf("unable to parse Member ID: %v", err)
	}

	return &AdministrativeUnitMemberId{
		ObjectSubResourceId:  *id,
		AdministrativeUnitId: id.objectId,
		MemberId:             id.subId,
	}, nil
}
