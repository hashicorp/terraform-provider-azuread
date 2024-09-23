package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationUser{}

type EducationUser struct {
	// True if the account is enabled; otherwise, false. This property is required when a user is created. Supports
	// /$filter.
	AccountEnabled nullable.Type[bool] `json:"accountEnabled,omitempty"`

	// The licenses that are assigned to the user. Not nullable.
	AssignedLicenses *[]AssignedLicense `json:"assignedLicenses,omitempty"`

	// The plans that are assigned to the user. Read-only. Not nullable.
	AssignedPlans *[]AssignedPlan `json:"assignedPlans,omitempty"`

	// List of assignments for the user. Nullable.
	Assignments *[]EducationAssignment `json:"assignments,omitempty"`

	// The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this
	// property.
	BusinessPhones *[]string `json:"businessPhones,omitempty"`

	// Classes to which the user belongs. Nullable.
	Classes *[]EducationClass `json:"classes,omitempty"`

	// Entity who created the user.
	CreatedBy IdentitySet `json:"createdBy"`

	// The name for the department in which the user works. Supports /$filter.
	Department nullable.Type[string] `json:"department,omitempty"`

	// The name displayed in the address book for the user. Supports $filter and $orderby.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The type of external source this resource was generated from (automatically determined from externalSourceDetail).
	// Possible values are: sis, lms, or manual.
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`

	// The name of the external source this resource was generated from.
	ExternalSourceDetail nullable.Type[string] `json:"externalSourceDetail,omitempty"`

	// The given name (first name) of the user. Supports /$filter.
	GivenName nullable.Type[string] `json:"givenName,omitempty"`

	// The SMTP address for the user; for example, 'jeff@contoso.com'. Read-Only. Supports /$filter.
	Mail nullable.Type[string] `json:"mail,omitempty"`

	// The mail alias for the user. This property must be specified when a user is created. Supports /$filter.
	MailNickname nullable.Type[string] `json:"mailNickname,omitempty"`

	// Mail address of user. Note: type and postOfficeBox aren't supported for educationUser resources.
	MailingAddress *PhysicalAddress `json:"mailingAddress,omitempty"`

	// The middle name of user.
	MiddleName nullable.Type[string] `json:"middleName,omitempty"`

	// The primary cellular telephone number for the user.
	MobilePhone nullable.Type[string] `json:"mobilePhone,omitempty"`

	// The office location for the user.
	OfficeLocation nullable.Type[string] `json:"officeLocation,omitempty"`

	// Additional information used to associate the Microsoft Entra user with its Active Directory counterpart.
	OnPremisesInfo *EducationOnPremisesInfo `json:"onPremisesInfo,omitempty"`

	// Specifies password policies for the user. For more details, see the standard [user] resource.
	PasswordPolicies nullable.Type[string] `json:"passwordPolicies,omitempty"`

	// Specifies the password profile for the user. The profile contains the user's password. This property is required when
	// a user is created. For more details, see the standard [user] resource.
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`

	// The preferred language for the user. Should follow ISO 639-1 Code; for example, 'en-US'.
	PreferredLanguage nullable.Type[string] `json:"preferredLanguage,omitempty"`

	PrimaryRole *EducationUserRole `json:"primaryRole,omitempty"`

	// The plans that are provisioned for the user. Read-only. Not nullable.
	ProvisionedPlans *[]ProvisionedPlan `json:"provisionedPlans,omitempty"`

	RefreshTokensValidFromDateTime nullable.Type[string] `json:"refreshTokensValidFromDateTime,omitempty"`

	// Related records related to the user. Possible relationships are parent, relative, aide, doctor, guardian, child,
	// other, unknownFutureValue
	RelatedContacts *[]RelatedContact `json:"relatedContacts,omitempty"`

	// Address where user lives. Note: type and postOfficeBox aren't supported for educationUser resources.
	ResidenceAddress *PhysicalAddress `json:"residenceAddress,omitempty"`

	// When set, the grading rubric attached to the assignment.
	Rubrics *[]EducationRubric `json:"rubrics,omitempty"`

	// Schools to which the user belongs. Nullable.
	Schools *[]EducationSchool `json:"schools,omitempty"`

	ShowInAddressList nullable.Type[bool] `json:"showInAddressList,omitempty"`

	// If the primary role is student, this block contains student specific data.
	Student *EducationStudent `json:"student,omitempty"`

	// The user's surname (family name or last name). Supports /$filter.
	Surname nullable.Type[string] `json:"surname,omitempty"`

	// Classes for which the user is a teacher.
	TaughtClasses *[]EducationClass `json:"taughtClasses,omitempty"`

	// If the primary role is teacher, this block contains teacher specific data.
	Teacher *EducationTeacher `json:"teacher,omitempty"`

	// A two-letter country code ([ISO 3166 Alpha-2]). Required for users who are assigned licenses. Not nullable. Supports
	// /$filter.
	UsageLocation nullable.Type[string] `json:"usageLocation,omitempty"`

	User *User `json:"user,omitempty"`

	// The user principal name (UPN) for the user. Supports $filter and $orderby. For more details, see the standard [user]
	// resource.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest'. Supports
	// /$filter.
	UserType nullable.Type[string] `json:"userType,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationUser) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationUser{}

func (s EducationUser) MarshalJSON() ([]byte, error) {
	type wrapper EducationUser
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationUser: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationUser: %+v", err)
	}

	delete(decoded, "assignedPlans")
	delete(decoded, "provisionedPlans")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationUser"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationUser: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationUser{}

func (s *EducationUser) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccountEnabled                 nullable.Type[bool]      `json:"accountEnabled,omitempty"`
		AssignedLicenses               *[]AssignedLicense       `json:"assignedLicenses,omitempty"`
		AssignedPlans                  *[]AssignedPlan          `json:"assignedPlans,omitempty"`
		Assignments                    *[]EducationAssignment   `json:"assignments,omitempty"`
		BusinessPhones                 *[]string                `json:"businessPhones,omitempty"`
		Classes                        *[]EducationClass        `json:"classes,omitempty"`
		Department                     nullable.Type[string]    `json:"department,omitempty"`
		DisplayName                    nullable.Type[string]    `json:"displayName,omitempty"`
		ExternalSource                 *EducationExternalSource `json:"externalSource,omitempty"`
		ExternalSourceDetail           nullable.Type[string]    `json:"externalSourceDetail,omitempty"`
		GivenName                      nullable.Type[string]    `json:"givenName,omitempty"`
		Mail                           nullable.Type[string]    `json:"mail,omitempty"`
		MailNickname                   nullable.Type[string]    `json:"mailNickname,omitempty"`
		MailingAddress                 *PhysicalAddress         `json:"mailingAddress,omitempty"`
		MiddleName                     nullable.Type[string]    `json:"middleName,omitempty"`
		MobilePhone                    nullable.Type[string]    `json:"mobilePhone,omitempty"`
		OfficeLocation                 nullable.Type[string]    `json:"officeLocation,omitempty"`
		OnPremisesInfo                 *EducationOnPremisesInfo `json:"onPremisesInfo,omitempty"`
		PasswordPolicies               nullable.Type[string]    `json:"passwordPolicies,omitempty"`
		PasswordProfile                *PasswordProfile         `json:"passwordProfile,omitempty"`
		PreferredLanguage              nullable.Type[string]    `json:"preferredLanguage,omitempty"`
		PrimaryRole                    *EducationUserRole       `json:"primaryRole,omitempty"`
		ProvisionedPlans               *[]ProvisionedPlan       `json:"provisionedPlans,omitempty"`
		RefreshTokensValidFromDateTime nullable.Type[string]    `json:"refreshTokensValidFromDateTime,omitempty"`
		RelatedContacts                *[]RelatedContact        `json:"relatedContacts,omitempty"`
		ResidenceAddress               *PhysicalAddress         `json:"residenceAddress,omitempty"`
		Rubrics                        *[]EducationRubric       `json:"rubrics,omitempty"`
		Schools                        *[]EducationSchool       `json:"schools,omitempty"`
		ShowInAddressList              nullable.Type[bool]      `json:"showInAddressList,omitempty"`
		Student                        *EducationStudent        `json:"student,omitempty"`
		Surname                        nullable.Type[string]    `json:"surname,omitempty"`
		TaughtClasses                  *[]EducationClass        `json:"taughtClasses,omitempty"`
		Teacher                        *EducationTeacher        `json:"teacher,omitempty"`
		UsageLocation                  nullable.Type[string]    `json:"usageLocation,omitempty"`
		User                           *User                    `json:"user,omitempty"`
		UserPrincipalName              nullable.Type[string]    `json:"userPrincipalName,omitempty"`
		UserType                       nullable.Type[string]    `json:"userType,omitempty"`
		Id                             *string                  `json:"id,omitempty"`
		ODataId                        *string                  `json:"@odata.id,omitempty"`
		ODataType                      *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccountEnabled = decoded.AccountEnabled
	s.AssignedLicenses = decoded.AssignedLicenses
	s.AssignedPlans = decoded.AssignedPlans
	s.Assignments = decoded.Assignments
	s.BusinessPhones = decoded.BusinessPhones
	s.Classes = decoded.Classes
	s.Department = decoded.Department
	s.DisplayName = decoded.DisplayName
	s.ExternalSource = decoded.ExternalSource
	s.ExternalSourceDetail = decoded.ExternalSourceDetail
	s.GivenName = decoded.GivenName
	s.Mail = decoded.Mail
	s.MailNickname = decoded.MailNickname
	s.MailingAddress = decoded.MailingAddress
	s.MiddleName = decoded.MiddleName
	s.MobilePhone = decoded.MobilePhone
	s.OfficeLocation = decoded.OfficeLocation
	s.OnPremisesInfo = decoded.OnPremisesInfo
	s.PasswordPolicies = decoded.PasswordPolicies
	s.PasswordProfile = decoded.PasswordProfile
	s.PreferredLanguage = decoded.PreferredLanguage
	s.PrimaryRole = decoded.PrimaryRole
	s.ProvisionedPlans = decoded.ProvisionedPlans
	s.RefreshTokensValidFromDateTime = decoded.RefreshTokensValidFromDateTime
	s.RelatedContacts = decoded.RelatedContacts
	s.ResidenceAddress = decoded.ResidenceAddress
	s.Rubrics = decoded.Rubrics
	s.Schools = decoded.Schools
	s.ShowInAddressList = decoded.ShowInAddressList
	s.Student = decoded.Student
	s.Surname = decoded.Surname
	s.TaughtClasses = decoded.TaughtClasses
	s.Teacher = decoded.Teacher
	s.UsageLocation = decoded.UsageLocation
	s.User = decoded.User
	s.UserPrincipalName = decoded.UserPrincipalName
	s.UserType = decoded.UserType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationUser into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EducationUser': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
