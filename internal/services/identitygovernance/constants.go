package identitygovernance

const (
	AccessReviewRecurrenceTypeAnnual     = "annual"
	AccessReviewRecurrenceTypeHalfYearly = "halfyearly"
	AccessReviewRecurrenceTypeMonthly    = "monthly"
	AccessReviewRecurrenceTypeQuarterly  = "quarterly"
	AccessReviewRecurrenceTypeWeekly     = "weekly"
)

var possibleValuesForAccessReviewRecurrenceType = []string{
	AccessReviewRecurrenceTypeAnnual,
	AccessReviewRecurrenceTypeHalfYearly,
	AccessReviewRecurrenceTypeMonthly,
	AccessReviewRecurrenceTypeQuarterly,
	AccessReviewRecurrenceTypeWeekly,
}

const (
	AccessReviewReviewerTypeManager   = "Manager"
	AccessReviewReviewerTypeReviewers = "Reviewers"
	AccessReviewReviewerTypeSelf      = "Self"
)

var possibleValuesForAccessReviewReviewerType = []string{
	AccessReviewReviewerTypeManager,
	AccessReviewReviewerTypeReviewers,
	AccessReviewReviewerTypeSelf,
}

const (
	RequestorScopeTypeAllConfiguredConnectedOrganizationSubjects = "AllConfiguredConnectedOrganizationSubjects"
	RequestorScopeTypeAllExistingConnectedOrganizationSubjects   = "AllExistingConnectedOrganizationSubjects"
	RequestorScopeTypeAllExistingDirectoryMemberUsers            = "AllExistingDirectoryMemberUsers"
	RequestorScopeTypeAllExistingDirectorySubjects               = "AllExistingDirectorySubjects"
	RequestorScopeTypeAllExternalSubjects                        = "AllExternalSubjects"
	RequestorScopeTypeNoSubjects                                 = "NoSubjects"
	RequestorScopeTypeSpecificConnectedOrganizationSubjects      = "SpecificConnectedOrganizationSubjects"
	RequestorScopeTypeSpecificDirectorySubjects                  = "SpecificDirectorySubjects"
)

var possibleValuesForRequestorScopeType = []string{
	RequestorScopeTypeAllConfiguredConnectedOrganizationSubjects,
	RequestorScopeTypeAllExistingConnectedOrganizationSubjects,
	RequestorScopeTypeAllExistingDirectoryMemberUsers,
	RequestorScopeTypeAllExistingDirectorySubjects,
	RequestorScopeTypeAllExternalSubjects,
	RequestorScopeTypeNoSubjects,
	RequestorScopeTypeSpecificConnectedOrganizationSubjects,
	RequestorScopeTypeSpecificDirectorySubjects,
}

const (
	CatalogStatusPublished   = "published"
	CatalogStatusUnpublished = "unpublished"
)

const (
	PrivilegedAccessGroupScheduleRequestStatusCanceled                = "Canceled"
	PrivilegedAccessGroupScheduleRequestStatusDenied                  = "Denied"
	PrivilegedAccessGroupScheduleRequestStatusFailed                  = "Failed"
	PrivilegedAccessGroupScheduleRequestStatusGranted                 = "Granted"
	PrivilegedAccessGroupScheduleRequestStatusPendingAdminDecision    = "PendingAdminDecision"
	PrivilegedAccessGroupScheduleRequestStatusPendingApproval         = "PendingApproval"
	PrivilegedAccessGroupScheduleRequestStatusPendingProvisioning     = "PendingProvisioning"
	PrivilegedAccessGroupScheduleRequestStatusPendingScheduleCreation = "PendingScheduleCreation"
	PrivilegedAccessGroupScheduleRequestStatusProvisioned             = "Provisioned"
	PrivilegedAccessGroupScheduleRequestStatusRevoked                 = "Revoked"
	PrivilegedAccessGroupScheduleRequestStatusScheduleCreated         = "ScheduleCreated"
)
