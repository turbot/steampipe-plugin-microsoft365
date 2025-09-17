package microsoft365

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

func userColumns() []*plugin.Column {
	return commonColumns([]*plugin.Column{
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier of the user.", Transform: transform.FromMethod("GetId")},
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name for the user.", Transform: transform.FromMethod("GetDisplayName")},
		{Name: "user_principal_name", Type: proto.ColumnType_STRING, Description: "The user principal name (UPN) of the user.", Transform: transform.FromMethod("GetUserPrincipalName")},
		{Name: "mail", Type: proto.ColumnType_STRING, Description: "The SMTP address for the user.", Transform: transform.FromMethod("GetMail")},
		{Name: "mail_nickname", Type: proto.ColumnType_STRING, Description: "The mail alias for the user.", Transform: transform.FromMethod("GetMailNickname")},
		{Name: "given_name", Type: proto.ColumnType_STRING, Description: "The given name (first name) of the user.", Transform: transform.FromMethod("GetGivenName")},
		{Name: "surname", Type: proto.ColumnType_STRING, Description: "The surname (last name) of the user.", Transform: transform.FromMethod("GetSurname")},
		{Name: "job_title", Type: proto.ColumnType_STRING, Description: "The user's job title.", Transform: transform.FromMethod("GetJobTitle")},
		{Name: "department", Type: proto.ColumnType_STRING, Description: "The name for the department in which the user works.", Transform: transform.FromMethod("GetDepartment")},
		{Name: "company_name", Type: proto.ColumnType_STRING, Description: "The name of the company in which the user works.", Transform: transform.FromMethod("GetCompanyName")},
		{Name: "office_location", Type: proto.ColumnType_STRING, Description: "The office location in the user's place of business.", Transform: transform.FromMethod("GetOfficeLocation")},
		{Name: "business_phones", Type: proto.ColumnType_JSON, Description: "The telephone numbers for the user.", Transform: transform.FromMethod("GetBusinessPhones")},
		{Name: "mobile_phone", Type: proto.ColumnType_STRING, Description: "The primary cellular telephone number for the user.", Transform: transform.FromMethod("GetMobilePhone")},
		{Name: "fax_number", Type: proto.ColumnType_STRING, Description: "The fax number of the user.", Transform: transform.FromMethod("GetFaxNumber")},
		{Name: "account_enabled", Type: proto.ColumnType_BOOL, Description: "True if the account is enabled; otherwise, false.", Transform: transform.FromMethod("GetAccountEnabled")},
		{Name: "user_type", Type: proto.ColumnType_STRING, Description: "A string value that can be used to classify user types in your directory.", Transform: transform.FromMethod("GetUserType")},
		{Name: "creation_type", Type: proto.ColumnType_STRING, Description: "Indicates whether the user account was created through one of the following methods.", Transform: transform.FromMethod("GetCreationType")},
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the user was created.", Transform: transform.FromMethod("GetCreatedDateTime")},
		{Name: "last_password_change_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when this Azure AD user last changed their password.", Transform: transform.FromMethod("GetLastPasswordChangeDateTime")},
		{Name: "sign_in_sessions_valid_from_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid.", Transform: transform.FromMethod("GetSignInSessionsValidFromDateTime")},
		{Name: "preferred_language", Type: proto.ColumnType_STRING, Description: "The preferred language for the user.", Transform: transform.FromMethod("GetPreferredLanguage")},
		{Name: "preferred_data_location", Type: proto.ColumnType_STRING, Description: "The preferred data location for the user.", Transform: transform.FromMethod("GetPreferredDataLocation")},
		{Name: "usage_location", Type: proto.ColumnType_STRING, Description: "A two letter country code (ISO standard 3166).", Transform: transform.FromMethod("GetUsageLocation")},
		{Name: "age_group", Type: proto.ColumnType_STRING, Description: "Sets the age group of the user.", Transform: transform.FromMethod("GetAgeGroup")},
		{Name: "legal_age_group_classification", Type: proto.ColumnType_STRING, Description: "Used by enterprise applications to determine the legal age group of the user.", Transform: transform.FromMethod("GetLegalAgeGroupClassification")},
		{Name: "consent_provided_for_minor", Type: proto.ColumnType_STRING, Description: "Sets whether consent has been obtained for minors.", Transform: transform.FromMethod("GetConsentProvidedForMinor")},
		{Name: "external_user_state", Type: proto.ColumnType_STRING, Description: "For an external user invited to the tenant using the invitation API.", Transform: transform.FromMethod("GetExternalUserState")},
		{Name: "external_user_state_change_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Shows the timestamp for the latest change to the externalUserState property.", Transform: transform.FromMethod("GetExternalUserStateChangeDateTime")},
		{Name: "is_management_restricted", Type: proto.ColumnType_BOOL, Description: "Do not use – reserved for future use.", Transform: transform.FromMethod("GetIsManagementRestricted")},
		{Name: "is_resource_account", Type: proto.ColumnType_BOOL, Description: "Do not use – reserved for future use.", Transform: transform.FromMethod("GetIsResourceAccount")},
		{Name: "show_in_address_list", Type: proto.ColumnType_BOOL, Description: "True if the Outlook global address list should contain this user.", Transform: transform.FromMethod("GetShowInAddressList")},
		{Name: "on_premises_sync_enabled", Type: proto.ColumnType_BOOL, Description: "true if this object is synced from an on-premises directory.", Transform: transform.FromMethod("GetOnPremisesSyncEnabled")},
		{Name: "on_premises_last_sync_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Indicates the last time at which the object was synced with the on-premises directory.", Transform: transform.FromMethod("GetOnPremisesLastSyncDateTime")},
		{Name: "on_premises_distinguished_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises Active Directory distinguished name or DN.", Transform: transform.FromMethod("GetOnPremisesDistinguishedName")},
		{Name: "on_premises_domain_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory.", Transform: transform.FromMethod("GetOnPremisesDomainName")},
		{Name: "on_premises_immutable_id", Type: proto.ColumnType_STRING, Description: "This property is used to associate an on-premises Active Directory user account to their Azure AD user object.", Transform: transform.FromMethod("GetOnPremisesImmutableId")},
		{Name: "on_premises_sam_account_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises sAMAccountName synchronized from the on-premises directory.", Transform: transform.FromMethod("GetOnPremisesSamAccountName")},
		{Name: "on_premises_security_identifier", Type: proto.ColumnType_STRING, Description: "Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud.", Transform: transform.FromMethod("GetOnPremisesSecurityIdentifier")},
		{Name: "on_premises_user_principal_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises userPrincipalName synchronized from the on-premises directory.", Transform: transform.FromMethod("GetOnPremisesUserPrincipalName")},
		{Name: "proxy_addresses", Type: proto.ColumnType_JSON, Description: "For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com'].", Transform: transform.FromMethod("GetProxyAddresses")},
		{Name: "other_mails", Type: proto.ColumnType_JSON, Description: "A list of additional email addresses for the user.", Transform: transform.FromMethod("GetOtherMails")},
		{Name: "im_addresses", Type: proto.ColumnType_JSON, Description: "The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.", Transform: transform.FromMethod("GetImAddresses")},
		{Name: "about_me", Type: proto.ColumnType_STRING, Description: "A freeform text entry field for the user to describe themselves.", Hydrate: getMicrosoft365User, Transform: transform.FromMethod("GetAboutMe")},
		{Name: "interests", Type: proto.ColumnType_JSON, Description: "A list for the user to describe their interests.", Hydrate: getMicrosoft365User, Transform: transform.FromMethod("GetInterests")},
		{Name: "past_projects", Type: proto.ColumnType_JSON, Description: "A list for the user to enumerate their past projects.", Hydrate: getMicrosoft365User, Transform: transform.FromMethod("GetPastProjects")},
		{Name: "responsibilities", Type: proto.ColumnType_JSON, Description: "A list for the user to enumerate their responsibilities.", Hydrate: getMicrosoft365User, Transform: transform.FromMethod("GetResponsibilities")},
		{Name: "schools", Type: proto.ColumnType_JSON, Description: "A list for the user to enumerate the schools they have attended.", Hydrate: getMicrosoft365User, Transform: transform.FromMethod("GetSchools")},
		{Name: "skills", Type: proto.ColumnType_JSON, Description: "A list for the user to enumerate their skills.", Hydrate: getMicrosoft365User, Transform: transform.FromMethod("GetSkills")},
		{Name: "hire_date", Type: proto.ColumnType_TIMESTAMP, Description: "The hire date of the user.", Hydrate: getMicrosoft365User, Transform: transform.FromMethod("GetHireDate")},
		{Name: "employee_hire_date", Type: proto.ColumnType_TIMESTAMP, Description: "The hire date of the user.", Transform: transform.FromMethod("GetEmployeeHireDate")},
		{Name: "employee_leave_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time when the user left or will leave the organization.", Transform: transform.FromMethod("GetEmployeeLeaveDateTime")},
		{Name: "employee_id", Type: proto.ColumnType_STRING, Description: "The employee identifier assigned to the user by the organization.", Transform: transform.FromMethod("GetEmployeeId")},
		{Name: "employee_type", Type: proto.ColumnType_STRING, Description: "Captures enterprise worker type.", Transform: transform.FromMethod("GetEmployeeType")},
		{Name: "birthday", Type: proto.ColumnType_TIMESTAMP, Description: "The birthday of the user.", Hydrate: getMicrosoft365User, Transform: transform.FromMethod("GetBirthday")},
		{Name: "city", Type: proto.ColumnType_STRING, Description: "The city in which the user is located.", Transform: transform.FromMethod("GetCity")},
		{Name: "country", Type: proto.ColumnType_STRING, Description: "The country/region in which the user is located.", Transform: transform.FromMethod("GetCountry")},
		{Name: "state", Type: proto.ColumnType_STRING, Description: "The state or province in the user's address.", Transform: transform.FromMethod("GetState")},
		{Name: "street_address", Type: proto.ColumnType_STRING, Description: "The street address of the user's place of business.", Transform: transform.FromMethod("GetStreetAddress")},
		{Name: "postal_code", Type: proto.ColumnType_STRING, Description: "The postal code for the user's postal address.", Transform: transform.FromMethod("GetPostalCode")},
		{Name: "my_site", Type: proto.ColumnType_STRING, Description: "The URL for the user's personal site.", Hydrate: getMicrosoft365User, Transform: transform.FromMethod("GetMySite")},
		{Name: "preferred_name", Type: proto.ColumnType_STRING, Description: "The preferred name for the user.", Hydrate: getMicrosoft365User, Transform: transform.FromMethod("GetPreferredName")},
		{Name: "security_identifier", Type: proto.ColumnType_STRING, Description: "Security identifier (SID) of the user.", Transform: transform.FromMethod("GetSecurityIdentifier")},
		{Name: "password_policies", Type: proto.ColumnType_STRING, Description: "Specifies password policies for the user.", Transform: transform.FromMethod("GetPasswordPolicies")},
		// While querying the device_enrollment_limit we will encounter the Forbidden error if the tenant is not have the "Microsoft Intune" license.
		// 		Intune is bundled with higher-tier SKUs like:
		// 		Microsoft 365 Business Premium
		// 		Microsoft 365 E3 / E5
		// 		Enterprise Mobility + Security (EMS) E3 / E5
		// 		Microsoft Intune (standalone)
		//      Enterprise Mobility + Security (EMS) E3 / E5
		//		Microsoft Intune (standalone)
		// {Name: "device_enrollment_limit", Type: proto.ColumnType_INT, Description: "The limit on the maximum number of devices that the user is permitted to enroll.", Transform: transform.FromMethod("GetDeviceEnrollmentLimit")},

		// JSON Fields for complex objects
		{Name: "assigned_licenses", Type: proto.ColumnType_JSON, Description: "The licenses that are assigned to the user.", Transform: transform.FromMethod("UserAssignedLicenses")},
		{Name: "assigned_plans", Type: proto.ColumnType_JSON, Description: "The plans that are assigned to the user.", Transform: transform.FromMethod("UserAssignedPlans")},
		{Name: "provisioned_plans", Type: proto.ColumnType_JSON, Description: "The plans that are provisioned for the user.", Transform: transform.FromMethod("UserProvisionedPlans")},
		{Name: "password_profile", Type: proto.ColumnType_JSON, Description: "Specifies the password profile for the user.", Transform: transform.FromMethod("UserPasswordProfile")},
		{Name: "on_premises_extension_attributes", Type: proto.ColumnType_JSON, Description: "Contains extensionAttributes 1-15 for the user.", Transform: transform.FromMethod("UserOnPremisesExtensionAttributes")},
		{Name: "on_premises_provisioning_errors", Type: proto.ColumnType_JSON, Description: "Any errors synchronizing the user account to the on-premises directory.", Transform: transform.FromMethod("UserOnPremisesProvisioningErrors")},
		{Name: "service_provisioning_errors", Type: proto.ColumnType_JSON, Description: "Errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from a user object.", Transform: transform.FromMethod("UserServiceProvisioningErrors")},
		{Name: "identities", Type: proto.ColumnType_JSON, Description: "Represents the identities that can be used to sign in to this user account.", Transform: transform.FromMethod("UserIdentities")},
		{Name: "license_assignment_states", Type: proto.ColumnType_JSON, Description: "State of license assignments for this user.", Transform: transform.FromMethod("UserLicenseAssignmentStates")},
		{Name: "license_details", Type: proto.ColumnType_JSON, Description: "A collection of this user's license details.", Transform: transform.FromValue(), Hydrate: getUserLicenseDetails},

		// Mailbox Settings columns
		{Name: "user_purpose", Type: proto.ColumnType_STRING, Description: "The purpose of the mailbox.", Transform: transform.FromMethod("GetUserPurpose"), Hydrate: getUserMailboxSettings},
		{Name: "archive_folder", Type: proto.ColumnType_STRING, Description: "Folder ID of an archive folder for the user.", Transform: transform.FromMethod("GetArchiveFolder"), Hydrate: getUserMailboxSettings},
		{Name: "automatic_replies_setting", Type: proto.ColumnType_JSON, Description: "Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user.", Transform: transform.FromMethod("GetAutomaticRepliesSetting"), Hydrate: getUserMailboxSettings},
		{Name: "date_format", Type: proto.ColumnType_STRING, Description: "The date format for the user's mailbox.", Transform: transform.FromMethod("GetDateFormat"), Hydrate: getUserMailboxSettings},
		{Name: "time_format", Type: proto.ColumnType_STRING, Description: "The time format for the user's mailbox.", Transform: transform.FromMethod("GetTimeFormat"), Hydrate: getUserMailboxSettings},
		{Name: "time_zone", Type: proto.ColumnType_STRING, Description: "The default time zone for the user's mailbox.", Transform: transform.FromMethod("GetTimeZone"), Hydrate: getUserMailboxSettings},
		{Name: "language", Type: proto.ColumnType_JSON, Description: "The locale information for the user, including the preferred language and country/region.", Transform: transform.FromMethod("GetLanguage"), Hydrate: getUserMailboxSettings},
		{Name: "working_hours", Type: proto.ColumnType_JSON, Description: "The days of the week and hours in a specific time zone that the user works.", Transform: transform.FromMethod("GetWorkingHours"), Hydrate: getUserMailboxSettings},
		{Name: "delegate_meeting_message_delivery_options", Type: proto.ColumnType_STRING, Description: "If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive copies of meeting-related messages that are addressed to the user.", Transform: transform.FromMethod("GetDelegateMeetingMessageDeliveryOptions"), Hydrate: getUserMailboxSettings},

		// Authentication Registration Details columns
		{Name: "is_mfa_capable", Type: proto.ColumnType_BOOL, Description: "Whether the user is capable of using multi-factor authentication.", Transform: transform.FromMethod("GetIsMfaCapable"), Hydrate: getUserRegistrationDetails},
		{Name: "is_mfa_registered", Type: proto.ColumnType_BOOL, Description: "Whether the user is registered for multi-factor authentication.", Transform: transform.FromMethod("GetIsMfaRegistered"), Hydrate: getUserRegistrationDetails},
		{Name: "is_passwordless_capable", Type: proto.ColumnType_BOOL, Description: "Whether the user is capable of using passwordless authentication.", Transform: transform.FromMethod("GetIsPasswordlessCapable"), Hydrate: getUserRegistrationDetails},
		{Name: "is_sspr_capable", Type: proto.ColumnType_BOOL, Description: "Whether the user is capable of using self-service password reset.", Transform: transform.FromMethod("GetIsSsprCapable"), Hydrate: getUserRegistrationDetails},
		{Name: "is_sspr_enabled", Type: proto.ColumnType_BOOL, Description: "Whether self-service password reset is enabled for the user.", Transform: transform.FromMethod("GetIsSsprEnabled"), Hydrate: getUserRegistrationDetails},
		{Name: "is_sspr_registered", Type: proto.ColumnType_BOOL, Description: "Whether the user is registered for self-service password reset.", Transform: transform.FromMethod("GetIsSsprRegistered"), Hydrate: getUserRegistrationDetails},
		{Name: "methods_registered", Type: proto.ColumnType_JSON, Description: "List of authentication methods registered by the user.", Transform: transform.FromMethod("GetMethodsRegistered"), Hydrate: getUserRegistrationDetails},
		{Name: "system_preferred_authentication_methods", Type: proto.ColumnType_JSON, Description: "List of system-preferred authentication methods.", Transform: transform.FromMethod("GetSystemPreferredAuthenticationMethods"), Hydrate: getUserRegistrationDetails},
		{Name: "user_preferred_method_for_secondary_authentication", Type: proto.ColumnType_STRING, Description: "The user's preferred method for secondary authentication.", Transform: transform.FromMethod("GetUserPreferredMethodForSecondaryAuthentication"), Hydrate: getUserRegistrationDetails},
		{Name: "is_system_preferred_authentication_method_enabled", Type: proto.ColumnType_BOOL, Description: "Whether system-preferred authentication method is enabled.", Transform: transform.FromMethod("GetIsSystemPreferredAuthenticationMethodEnabled"), Hydrate: getUserRegistrationDetails},
		{Name: "registration_last_updated_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time when the user registration details were last updated.", Transform: transform.FromMethod("GetLastUpdatedDateTime"), Hydrate: getUserRegistrationDetails},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetDisplayName")},
		{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Odata query to search for resources."},
	})
}

//// TABLE DEFINITION

func tableMicrosoft365User(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_user",
		Description: "Users in Microsoft 365.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365Users,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "filter", Require: plugin.Optional},
				{Name: "id", Require: plugin.Optional},
				{Name: "display_name", Require: plugin.Optional},
				{Name: "user_principal_name", Require: plugin.Optional},
				{Name: "mail", Require: plugin.Optional},
				{Name: "account_enabled", Require: plugin.Optional},
				{Name: "user_type", Require: plugin.Optional},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365User,
			KeyColumns: plugin.AllColumns([]string{"id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Columns: userColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365Users(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_user.listMicrosoft365Users", "connection_error", err)
		return nil, err
	}

	input := &users.UsersRequestBuilderGetQueryParameters{}

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is 999 (API limit)
	pageSize := int64(999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}
	input.Top = Int32(int32(pageSize))

	var queryFilter string
	equalQuals := d.EqualsQuals

	// Select specific properties including license-related fields that are not returned by default
	// According to Microsoft Graph documentation, licenseAssignmentStates and licenseDetails need to be explicitly requested
	selectFields := buildUserSelectFields(ctx, d)
	input.Select = selectFields

	filter := buildUserQueryFilter(equalQuals)

	if equalQuals["filter"] != nil {
		queryFilter = equalQuals["filter"].GetStringValue()
	}

	if queryFilter != "" {
		filter = append(filter, queryFilter)
	}

	if len(filter) > 0 {
		joinStr := strings.Join(filter, " and ")
		input.Filter = &joinStr
	}

	options := &users.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.Users().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.Userable](result, adapter, models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("microsoft365_user.listMicrosoft365Users", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem models.Userable) bool {
		user := pageItem

		d.StreamListItem(ctx, &Microsoft365UserInfo{Userable: user, MailboxSettings: nil})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("microsoft365_user.listMicrosoft365Users", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365User(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	userID := d.EqualsQualString("id")
	if h.Item != nil {
		user := h.Item.(*Microsoft365UserInfo)
	
		userID = *user.GetId()
	}
	
	if userID == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_user.getMicrosoft365User", "connection_error", err)
		return nil, err
	}

	// Select specific properties including license-related fields that are not returned by default
	selectFields := buildUserSelectFields(ctx, d)

	// Appending the selected fields that are only be retrieved via GET API.
	options := &users.UserItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.UserItemRequestBuilderGetQueryParameters{
			Select: append(selectFields, []string{"hireDate","birthday","schools","responsibilities","pastProjects","preferredName","interests","aboutMe","mySite","skills"}...),
		},
	}

	result, err := client.Users().ByUserId(userID).Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	// Get mailbox settings for this user
	mailboxSettings, err := client.Users().ByUserId(userID).MailboxSettings().Get(ctx, nil)
	if err != nil {
		// If we can't get mailbox settings, return user without them
		return &Microsoft365UserInfo{Userable: result, MailboxSettings: nil}, nil
	}

	return &Microsoft365UserInfo{Userable: result, MailboxSettings: mailboxSettings}, nil
}

func buildUserQueryFilter(equalQuals plugin.KeyColumnEqualsQualMap) []string {
	filters := []string{}

	filterQuals := map[string]string{
		"display_name":        "string",
		"id":                  "string",
		"user_principal_name": "string",
		"mail":                "string",
		"account_enabled":     "bool",
		"user_type":           "string",
	}

	for qual, qualType := range filterQuals {
		switch qualType {
		case "string":
			if equalQuals[qual] != nil {
				filters = append(filters, fmt.Sprintf("%s eq '%s'", strcase.ToCamel(qual), equalQuals[qual].GetStringValue()))
			}
		case "bool":
			if equalQuals[qual] != nil {
				filters = append(filters, fmt.Sprintf("%s eq %t", strcase.ToCamel(qual), equalQuals[qual].GetBoolValue()))
			}
		}
	}
	return filters
}

//// HYDRATE FUNCTIONS

func getUserMailboxSettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	user := h.Item.(*Microsoft365UserInfo)
	userID := *user.GetId()

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_user.getUserMailboxSettings", "connection_error", err)
		return nil, err
	}

	// Get mailbox settings for this user
	mailboxSettings, err := client.Users().ByUserId(userID).MailboxSettings().Get(ctx, nil)
	if err != nil {
		// Return nil if we can't get mailbox settings (user might not have a mailbox)
		return nil, nil
	}

	return mailboxSettings, nil
}

func getUserRegistrationDetails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	user := h.Item.(*Microsoft365UserInfo)

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_user.getUserRegistrationDetails", "connection_error", err)
		return nil, err
	}

	// Get user registration details from Reports API
	// Note: This uses the Reports API endpoint for user registration details
	userRegistrationDetails, err := client.Reports().AuthenticationMethods().UserRegistrationDetails().Get(ctx, nil)
	if err != nil {
		logger.Error("microsoft365_user.getUserRegistrationDetails", "api_error", err)
		return nil, nil
	}

	// Find the registration details for this specific user
	if userRegistrationDetails.GetValue() != nil {
		for _, detail := range userRegistrationDetails.GetValue() {
			if detail.GetUserPrincipalName() != nil && *detail.GetUserPrincipalName() == *user.GetUserPrincipalName() {
				return detail, nil
			}
		}
	}

	// Return nil if no registration details found for this user
	return nil, nil
}

func getUserLicenseDetails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Get the user from the hydrate data
	userInfo := h.Item.(*Microsoft365UserInfo)
	user := userInfo.Userable

	if user.GetId() == nil {
		return nil, nil
	}

	userID := *user.GetId()

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_user.getUserLicenseDetails", "connection_error", err)
		return nil, err
	}

	// Get license details using the beta endpoint approach
	// This calls /users/{id}/licenseDetails which is available in beta
	licenseDetailsResult, err := client.Users().ByUserId(userID).LicenseDetails().Get(ctx, nil)
	if err != nil {
		logger.Error("microsoft365_user.getUserLicenseDetails", "api_error", err)
		return nil, nil // Return nil instead of error to avoid breaking the query
	}

	// Convert license details to JSON-serializable format
	if licenseDetailsResult.GetValue() != nil {
		details := []map[string]interface{}{}
		for _, detail := range licenseDetailsResult.GetValue() {
			detailData := map[string]interface{}{}

			if detail.GetId() != nil {
				detailData["id"] = *detail.GetId()
			}
			if detail.GetSkuId() != nil {
				detailData["sku_id"] = *detail.GetSkuId()
			}
			if detail.GetSkuPartNumber() != nil {
				detailData["sku_part_number"] = *detail.GetSkuPartNumber()
			}

			// Add service plans
			if detail.GetServicePlans() != nil {
				servicePlans := []map[string]interface{}{}
				for _, plan := range detail.GetServicePlans() {
					planData := map[string]interface{}{}
					if plan.GetServicePlanId() != nil {
						planData["service_plan_id"] = *plan.GetServicePlanId()
					}
					if plan.GetServicePlanName() != nil {
						planData["service_plan_name"] = *plan.GetServicePlanName()
					}
					if plan.GetProvisioningStatus() != nil {
						planData["provisioning_status"] = *plan.GetProvisioningStatus()
					}
					if plan.GetAppliesTo() != nil {
						planData["applies_to"] = *plan.GetAppliesTo()
					}
					servicePlans = append(servicePlans, planData)
				}
				detailData["service_plans"] = servicePlans
			}

			details = append(details, detailData)
		}
		return details, nil
	}

	return nil, nil
}

//// BUILD SELECT Input parameter

func buildUserSelectFields(ctx context.Context, d *plugin.QueryData) []string {
	// Always include basic fields that are typically needed
	fieldSelected := []string{
		"id",
		"displayName",
		"userPrincipalName",
		"mail",
	}

	// Get the queried columns from the query context
	givenColumns := d.QueryContext.Columns

	// Process each queried column
	for _, columnName := range givenColumns {
		// Skip special columns that are not Graph API fields
		if columnName == "title" || columnName == "filter" || columnName == "_ctx" || columnName == "tenant_id" || columnName == "sp_connection_name" || columnName == "sp_ctx" {
			continue
		}

		// Check if this column has a mapping in our userSelectFields map
		if selectValue, ok := userSelectFields[columnName]; ok {
			// Add the mapped Graph API field if not already included
			if !slices.Contains(fieldSelected, selectValue) {
				fieldSelected = append(fieldSelected, selectValue)
			}
		}
		// Note: We don't fall back to camelCase conversion anymore to avoid unsupported fields
	}

	return fieldSelected
}

// Common fields for list and get API call
var userSelectFields = map[string]string{
	// Basic user information
	"mail_nickname":                         "mailNickname",
	"given_name":                            "givenName",
	"surname":                               "surname",
	"job_title":                             "jobTitle",
	"department":                            "department",
	"company_name":                          "companyName",
	"office_location":                       "officeLocation",
	"business_phones":                       "businessPhones",
	"mobile_phone":                          "mobilePhone",
	"fax_number":                            "faxNumber",
	"account_enabled":                       "accountEnabled",
	"user_type":                             "userType",
	"creation_type":                         "creationType",
	"created_date_time":                     "createdDateTime",
	"last_password_change_date_time":        "lastPasswordChangeDateTime",
	"sign_in_sessions_valid_from_date_time": "signInSessionsValidFromDateTime",
	"preferred_language":                    "preferredLanguage",
	"preferred_data_location":               "preferredDataLocation",
	"usage_location":                        "usageLocation",
	"age_group":                             "ageGroup",
	"legal_age_group_classification":        "legalAgeGroupClassification",
	"consent_provided_for_minor":            "consentProvidedForMinor",
	"external_user_state":                   "externalUserState",
	"external_user_state_change_date_time":  "externalUserStateChangeDateTime",
	"is_management_restricted":              "isManagementRestricted",
	"is_resource_account":                   "isResourceAccount",
	"show_in_address_list":                  "showInAddressList",

	// On-premises information
	"on_premises_sync_enabled":        "onPremisesSyncEnabled",
	"on_premises_last_sync_date_time": "onPremisesLastSyncDateTime",
	"on_premises_distinguished_name":  "onPremisesDistinguishedName",
	"on_premises_domain_name":         "onPremisesDomainName",
	"on_premises_immutable_id":        "onPremisesImmutableId",
	"on_premises_sam_account_name":    "onPremisesSamAccountName",
	"on_premises_security_identifier": "onPremisesSecurityIdentifier",
	"on_premises_user_principal_name": "onPremisesUserPrincipalName",

	// Contact information
	"proxy_addresses": "proxyAddresses",
	"other_mails":     "otherMails",
	"im_addresses":    "imAddresses",

	// Personal information (available in v1.0 when using $select)
	"employee_hire_date":       "employeeHireDate",
	"employee_leave_date_time": "employeeLeaveDateTime",
	"employee_id":              "employeeId",
	"employee_type":            "employeeType",

	// Address information
	"city":           "city",
	"country":        "country",
	"state":          "state",
	"street_address": "streetAddress",
	"postal_code":    "postalCode",

	// Security and system information
	"security_identifier":     "securityIdentifier",
	"password_policies":       "passwordPolicies",

	// License and plan information (returned by list/get APIs)
	"assigned_licenses":                "assignedLicenses",
	"assigned_plans":                   "assignedPlans",
	"provisioned_plans":                "provisionedPlans",
	"password_profile":                 "passwordProfile",
	"on_premises_extension_attributes": "onPremisesExtensionAttributes",
	"on_premises_provisioning_errors":  "onPremisesProvisioningErrors",
	"service_provisioning_errors":      "serviceProvisioningErrors",
	"identities":                       "identities",
	"license_assignment_states":        "licenseAssignmentStates",

	// Note: license_details is excluded as it requires a separate API call via hydrate function
	// Note: mailbox settings columns are excluded as they require separate API calls via hydrate functions
}
