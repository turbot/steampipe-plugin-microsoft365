package office365

import "github.com/microsoftgraph/msgraph-sdk-go/models"

type Office365CalendarEventInfo struct {
	models.Eventable
}

func (event *Office365CalendarEventInfo) EventAttendees() []map[string]interface{} {
	if event.GetAttendees() == nil {
		return nil
	}

	attendees := []map[string]interface{}{}
	for _, a := range event.GetAttendees() {
		attendeeInfo := map[string]interface{}{}
		if a.GetStatus() != nil {
			responseStatusInfo := map[string]interface{}{}
			if a.GetStatus().GetResponse() != nil {
				responseStatusInfo["response"] = a.GetStatus().GetResponse().String()
			}
			if a.GetStatus().GetTime() != nil {
				responseStatusInfo["time"] = *a.GetStatus().GetTime()
			}
			attendeeInfo["status"] = attendeeInfo
		}
		if a.GetProposedNewTime() != nil {
			data := map[string]interface{}{}
			if a.GetProposedNewTime().GetEnd() != nil {
				endTimeInfo := map[string]interface{}{}
				if event.GetEnd().GetDateTime() != nil {
					endTimeInfo["dateTime"] = *event.GetEnd().GetDateTime()
				}
				if event.GetEnd().GetDateTime() != nil {
					endTimeInfo["timeZone"] = *event.GetEnd().GetTimeZone()
				}
				data["end"] = endTimeInfo
			}
			if a.GetProposedNewTime().GetStart() != nil {
				startTimeInfo := map[string]interface{}{}
				if event.GetStart().GetDateTime() != nil {
					startTimeInfo["dateTime"] = *event.GetStart().GetDateTime()
				}
				if event.GetStart().GetDateTime() != nil {
					startTimeInfo["timeZone"] = *event.GetStart().GetTimeZone()
				}
				data["start"] = startTimeInfo
			}
			attendeeInfo["proposedNewTime"] = data
		}

		attendees = append(attendees, attendeeInfo)
	}
	return attendees
}

func (event *Office365CalendarEventInfo) EventBody() map[string]interface{} {
	if event.GetBody() == nil {
		return nil
	}

	bodyInfo := map[string]interface{}{}
	if event.GetBody().GetContent() != nil {
		bodyInfo["content"] = *event.GetBody().GetContent()
	}
	if event.GetBody().GetContentType() != nil {
		bodyInfo["contentType"] = *event.GetBody().GetContentType()
	}
	return bodyInfo
}

func (event *Office365CalendarEventInfo) EventEnd() map[string]interface{} {
	if event.GetEnd() == nil {
		return nil
	}

	endTimeInfo := map[string]interface{}{}
	if event.GetEnd().GetDateTime() != nil {
		endTimeInfo["dateTime"] = *event.GetEnd().GetDateTime()
	}
	if event.GetEnd().GetDateTime() != nil {
		endTimeInfo["timeZone"] = *event.GetEnd().GetTimeZone()
	}
	return endTimeInfo
}

func (event *Office365CalendarEventInfo) EventLocation() map[string]interface{} {
	if event.GetLocation() == nil {
		return nil
	}

	locationInfo := map[string]interface{}{}
	if event.GetLocation().GetDisplayName() != nil {
		locationInfo["displayName"] = *event.GetLocation().GetDisplayName()
	}
	if event.GetLocation().GetDisplayName() != nil {
		locationInfo["displayName"] = *event.GetLocation().GetDisplayName()
	}
	if event.GetLocation().GetLocationEmailAddress() != nil {
		locationInfo["locationEmailAddress"] = *event.GetLocation().GetLocationEmailAddress()
	}
	if event.GetLocation().GetLocationType() != nil {
		locationInfo["locationType"] = event.GetLocation().GetLocationType().String()
	}
	if event.GetLocation().GetLocationUri() != nil {
		locationInfo["locationUri"] = *event.GetLocation().GetLocationUri()
	}
	if event.GetLocation().GetType() != nil {
		locationInfo["type"] = *event.GetLocation().GetType()
	}
	if event.GetLocation().GetUniqueId() != nil {
		locationInfo["uniqueId"] = *event.GetLocation().GetUniqueId()
	}
	if event.GetLocation().GetUniqueIdType() != nil {
		locationInfo["uniqueIdType"] = event.GetLocation().GetUniqueIdType().String()
	}

	addressInfo := map[string]interface{}{}
	if event.GetLocation().GetAddress() != nil {
		if event.GetLocation().GetAddress().GetCity() != nil {
			addressInfo["city"] = *event.GetLocation().GetAddress().GetCity()
		}
		if event.GetLocation().GetAddress().GetCountryOrRegion() != nil {
			addressInfo["countryOrRegion"] = *event.GetLocation().GetAddress().GetCountryOrRegion()
		}
		if event.GetLocation().GetAddress().GetPostalCode() != nil {
			addressInfo["postalCode"] = *event.GetLocation().GetAddress().GetPostalCode()
		}
		if event.GetLocation().GetAddress().GetState() != nil {
			addressInfo["state"] = *event.GetLocation().GetAddress().GetState()
		}
		if event.GetLocation().GetAddress().GetStreet() != nil {
			addressInfo["street"] = *event.GetLocation().GetAddress().GetStreet()
		}
	}
	locationInfo["address"] = addressInfo

	coordinatesInfo := map[string]interface{}{}
	if event.GetLocation().GetCoordinates() != nil {
		if event.GetLocation().GetCoordinates().GetAccuracy() != nil {
			coordinatesInfo["accuracy"] = *event.GetLocation().GetCoordinates().GetAccuracy()
		}
		if event.GetLocation().GetCoordinates().GetAltitude() != nil {
			coordinatesInfo["altitude"] = *event.GetLocation().GetCoordinates().GetAltitude()
		}
		if event.GetLocation().GetCoordinates().GetAltitudeAccuracy() != nil {
			coordinatesInfo["altitudeAccuracy"] = *event.GetLocation().GetCoordinates().GetAltitudeAccuracy()
		}
		if event.GetLocation().GetCoordinates().GetLatitude() != nil {
			coordinatesInfo["latitude"] = *event.GetLocation().GetCoordinates().GetLatitude()
		}
		if event.GetLocation().GetCoordinates().GetLongitude() != nil {
			coordinatesInfo["longitude"] = *event.GetLocation().GetCoordinates().GetLongitude()
		}
	}
	locationInfo["coordinates"] = coordinatesInfo

	return locationInfo
}

func (event *Office365CalendarEventInfo) EventLocations() []map[string]interface{} {
	if event.GetLocations() == nil {
		return nil
	}

	locInfo := []map[string]interface{}{}
	for _, l := range event.GetLocations() {
		locationInfo := map[string]interface{}{}
		if l.GetDisplayName() != nil {
			locationInfo["displayName"] = *l.GetDisplayName()
		}
		if l.GetDisplayName() != nil {
			locationInfo["displayName"] = *l.GetDisplayName()
		}
		if l.GetLocationEmailAddress() != nil {
			locationInfo["locationEmailAddress"] = *l.GetLocationEmailAddress()
		}
		if l.GetLocationType() != nil {
			locationInfo["locationType"] = l.GetLocationType().String()
		}
		if l.GetLocationUri() != nil {
			locationInfo["locationUri"] = *l.GetLocationUri()
		}
		if l.GetType() != nil {
			locationInfo["type"] = *l.GetType()
		}
		if l.GetUniqueId() != nil {
			locationInfo["uniqueId"] = *l.GetUniqueId()
		}
		if l.GetUniqueIdType() != nil {
			locationInfo["uniqueIdType"] = l.GetUniqueIdType().String()
		}

		addressInfo := map[string]interface{}{}
		if l.GetAddress() != nil {
			if l.GetAddress().GetCity() != nil {
				addressInfo["city"] = *l.GetAddress().GetCity()
			}
			if l.GetAddress().GetCountryOrRegion() != nil {
				addressInfo["countryOrRegion"] = *l.GetAddress().GetCountryOrRegion()
			}
			if l.GetAddress().GetPostalCode() != nil {
				addressInfo["postalCode"] = *l.GetAddress().GetPostalCode()
			}
			if l.GetAddress().GetState() != nil {
				addressInfo["state"] = *l.GetAddress().GetState()
			}
			if l.GetAddress().GetStreet() != nil {
				addressInfo["street"] = *l.GetAddress().GetStreet()
			}
		}
		locationInfo["address"] = addressInfo

		coordinatesInfo := map[string]interface{}{}
		if l.GetCoordinates() != nil {
			if l.GetCoordinates().GetAccuracy() != nil {
				coordinatesInfo["accuracy"] = *l.GetCoordinates().GetAccuracy()
			}
			if l.GetCoordinates().GetAltitude() != nil {
				coordinatesInfo["altitude"] = *l.GetCoordinates().GetAltitude()
			}
			if l.GetCoordinates().GetAltitudeAccuracy() != nil {
				coordinatesInfo["altitudeAccuracy"] = *l.GetCoordinates().GetAltitudeAccuracy()
			}
			if l.GetCoordinates().GetLatitude() != nil {
				coordinatesInfo["latitude"] = *l.GetCoordinates().GetLatitude()
			}
			if l.GetCoordinates().GetLongitude() != nil {
				coordinatesInfo["longitude"] = *l.GetCoordinates().GetLongitude()
			}
		}
		locationInfo["coordinates"] = coordinatesInfo

		locInfo = append(locInfo, locationInfo)
	}
	return locInfo
}

func (event *Office365CalendarEventInfo) EventOnlineMeeting() map[string]interface{} {
	if event.GetOnlineMeeting() == nil {
		return nil
	}

	onlineMeetingInfo := map[string]interface{}{
		"tollFreeNumbers": event.GetOnlineMeeting().GetTollFreeNumbers(),
	}
	if event.GetOnlineMeeting().GetConferenceId() != nil {
		onlineMeetingInfo["conferenceId"] = *event.GetOnlineMeeting().GetConferenceId()
	}
	if event.GetOnlineMeeting().GetJoinUrl() != nil {
		onlineMeetingInfo["joinUrl"] = *event.GetOnlineMeeting().GetJoinUrl()
	}
	if event.GetOnlineMeeting().GetQuickDial() != nil {
		onlineMeetingInfo["quickDial"] = *event.GetOnlineMeeting().GetQuickDial()
	}
	if event.GetOnlineMeeting().GetTollNumber() != nil {
		onlineMeetingInfo["tollNumber"] = *event.GetOnlineMeeting().GetTollNumber()
	}
	if event.GetOnlineMeeting().GetPhones() != nil {
		phones := []map[string]interface{}{}
		for _, p := range event.GetOnlineMeeting().GetPhones() {
			data := map[string]interface{}{}
			if p.GetLanguage() != nil {
				data["language"] = *p.GetLanguage()
			}
			if p.GetNumber() != nil {
				data["number"] = *p.GetNumber()
			}
			if p.GetRegion() != nil {
				data["region"] = *p.GetRegion()
			}
			if p.GetType() != nil {
				data["type"] = p.GetType().String()
			}
			phones = append(phones, data)
		}
		onlineMeetingInfo["phones"] = phones
	}
	return onlineMeetingInfo
}

func (event *Office365CalendarEventInfo) EventOrganizer() map[string]interface{} {
	if event.GetOrganizer() == nil {
		return nil
	}

	organizerInfo := map[string]interface{}{}
	if event.GetOrganizer().GetType() != nil {
		organizerInfo["type"] = *event.GetOrganizer().GetType()
	}

	if event.GetOrganizer().GetEmailAddress() != nil {
		emailAddressInfo := map[string]interface{}{}
		if event.GetOrganizer().GetEmailAddress().GetAddress() != nil {
			emailAddressInfo["address"] = *event.GetOrganizer().GetEmailAddress().GetAddress()
		}
		if event.GetOrganizer().GetEmailAddress().GetName() != nil {
			emailAddressInfo["name"] = *event.GetOrganizer().GetEmailAddress().GetName()
		}
		organizerInfo["emailAddress"] = emailAddressInfo
	}
	return organizerInfo
}

func (event *Office365CalendarEventInfo) EventRecurrence() map[string]interface{} {
	if event.GetRecurrence() == nil {
		return nil
	}

	recurrenceInfo := map[string]interface{}{}

	if event.GetRecurrence().GetPattern() != nil {
		patternInfo := map[string]interface{}{
			"daysOfWeek": event.GetRecurrence().GetPattern().GetDaysOfWeek(),
		}
		if event.GetRecurrence().GetPattern().GetDayOfMonth() != nil {
			patternInfo["dayOfMonth"] = *event.GetRecurrence().GetPattern().GetDayOfMonth()
		}
		if event.GetRecurrence().GetPattern().GetFirstDayOfWeek() != nil {
			patternInfo["firstDayOfWeek"] = event.GetRecurrence().GetPattern().GetFirstDayOfWeek().String()
		}
		if event.GetRecurrence().GetPattern().GetIndex() != nil {
			patternInfo["index"] = event.GetRecurrence().GetPattern().GetIndex().String()
		}
		if event.GetRecurrence().GetPattern().GetInterval() != nil {
			patternInfo["interval"] = *event.GetRecurrence().GetPattern().GetInterval()
		}
		if event.GetRecurrence().GetPattern().GetMonth() != nil {
			patternInfo["month"] = *event.GetRecurrence().GetPattern().GetMonth()
		}
		if event.GetRecurrence().GetPattern().GetType() != nil {
			patternInfo["type"] = *event.GetRecurrence().GetPattern().GetType()
		}
		recurrenceInfo["pattern"] = patternInfo
	}
	if event.GetRecurrence().GetRange() != nil {
		rangeInfo := map[string]interface{}{
			"type": event.GetRecurrence().GetRange().GetType().String(),
		}
		if event.GetRecurrence().GetRange().GetEndDate() != nil {
			rangeInfo["endDate"] = *event.GetRecurrence().GetRange().GetEndDate()
		}
		if event.GetRecurrence().GetRange().GetNumberOfOccurrences() != nil {
			rangeInfo["numberOfOccurrences"] = *event.GetRecurrence().GetRange().GetNumberOfOccurrences()
		}
		if event.GetRecurrence().GetRange().GetRecurrenceTimeZone() != nil {
			rangeInfo["recurrenceTimeZone"] = event.GetRecurrence().GetRange().GetRecurrenceTimeZone()
		}
		if event.GetRecurrence().GetRange().GetStartDate() != nil {
			rangeInfo["startDate"] = *event.GetRecurrence().GetRange().GetStartDate()
		}
		recurrenceInfo["range"] = rangeInfo
	}

	return recurrenceInfo
}

func (event *Office365CalendarEventInfo) EventResponseStatus() map[string]interface{} {
	if event.GetResponseStatus() == nil {
		return nil
	}

	responseStatusInfo := map[string]interface{}{}
	if event.GetResponseStatus().GetResponse() != nil {
		responseStatusInfo["response"] = event.GetResponseStatus().GetResponse().String()
	}

	if event.GetResponseStatus().GetTime() != nil {
		responseStatusInfo["time"] = *event.GetResponseStatus().GetTime()
	}
	return responseStatusInfo
}

func (event *Office365CalendarEventInfo) EventStart() map[string]interface{} {
	if event.GetStart() == nil {
		return nil
	}

	startTimeInfo := map[string]interface{}{}
	if event.GetStart().GetDateTime() != nil {
		startTimeInfo["dateTime"] = *event.GetStart().GetDateTime()
	}
	if event.GetStart().GetDateTime() != nil {
		startTimeInfo["timeZone"] = *event.GetStart().GetTimeZone()
	}
	return startTimeInfo
}

// type ADConditionalAccessPolicyInfo struct {
// 	models.ConditionalAccessPolicyable
// }

// type ADGroupInfo struct {
// 	models.Groupable
// }

// type ADServicePrincipalInfo struct {
// 	models.ServicePrincipalable
// }

// type ADSignInReportInfo struct {
// 	models.SignInable
// }

// type ADUserInfo struct {
// 	models.Userable
// }

// func (application *ADApplicationInfo) ApplicationAPI() map[string]interface{} {
// 	if application.GetApi() == nil {
// 		return nil
// 	}

// 	apiData := map[string]interface{}{
// 		"knownClientApplications": application.GetApi().GetKnownClientApplications(),
// 	}

// 	if application.GetApi().GetAcceptMappedClaims() != nil {
// 		apiData["acceptMappedClaims"] = *application.GetApi().GetAcceptMappedClaims()
// 	}
// 	if application.GetApi().GetRequestedAccessTokenVersion() != nil {
// 		apiData["requestedAccessTokenVersion"] = *application.GetApi().GetRequestedAccessTokenVersion()
// 	}

// 	oauth2PermissionScopes := []map[string]interface{}{}
// 	for _, p := range application.GetApi().GetOauth2PermissionScopes() {
// 		data := map[string]interface{}{}
// 		if p.GetAdminConsentDescription() != nil {
// 			data["adminConsentDescription"] = *p.GetAdminConsentDescription()
// 		}
// 		if p.GetAdminConsentDisplayName() != nil {
// 			data["adminConsentDisplayName"] = *p.GetAdminConsentDisplayName()
// 		}
// 		if p.GetId() != nil {
// 			data["id"] = *p.GetId()
// 		}
// 		if p.GetIsEnabled() != nil {
// 			data["isEnabled"] = *p.GetIsEnabled()
// 		}
// 		if p.GetOrigin() != nil {
// 			data["origin"] = *p.GetOrigin()
// 		}
// 		if p.GetType() != nil {
// 			data["type"] = *p.GetType()
// 		}
// 		if p.GetUserConsentDescription() != nil {
// 			data["userConsentDescription"] = p.GetUserConsentDescription()
// 		}
// 		if p.GetUserConsentDisplayName() != nil {
// 			data["userConsentDisplayName"] = p.GetUserConsentDisplayName()
// 		}
// 		if p.GetValue() != nil {
// 			data["value"] = *p.GetValue()
// 		}
// 		oauth2PermissionScopes = append(oauth2PermissionScopes, data)
// 	}
// 	apiData["oauth2PermissionScopes"] = oauth2PermissionScopes

// 	preAuthorizedApplications := []map[string]interface{}{}
// 	for _, p := range application.GetApi().GetPreAuthorizedApplications() {
// 		data := map[string]interface{}{
// 			"delegatedPermissionIds": p.GetDelegatedPermissionIds(),
// 		}
// 		if p.GetAppId() != nil {
// 			data["appId"] = *p.GetAppId()
// 		}
// 		preAuthorizedApplications = append(preAuthorizedApplications, data)
// 	}
// 	apiData["preAuthorizedApplications"] = preAuthorizedApplications

// 	return apiData
// }

// func (application *ADApplicationInfo) ApplicationInfo() map[string]interface{} {
// 	if application.GetInfo() == nil {
// 		return nil
// 	}

// 	return map[string]interface{}{
// 		"logoUrl":             application.GetInfo().GetLogoUrl(),
// 		"marketingUrl":        application.GetInfo().GetMarketingUrl(),
// 		"privacyStatementUrl": application.GetInfo().GetPrivacyStatementUrl(),
// 		"supportUrl":          application.GetInfo().GetSupportUrl(),
// 		"termsOfServiceUrl":   application.GetInfo().GetTermsOfServiceUrl(),
// 	}
// }

// func (application *ADApplicationInfo) ApplicationKeyCredentials() []map[string]interface{} {
// 	if application.GetKeyCredentials() == nil {
// 		return nil
// 	}

// 	keyCredentials := []map[string]interface{}{}
// 	for _, p := range application.GetKeyCredentials() {
// 		keyCredentialData := map[string]interface{}{}
// 		if p.GetDisplayName() != nil {
// 			keyCredentialData["displayName"] = *p.GetDisplayName()
// 		}
// 		if p.GetEndDateTime() != nil {
// 			keyCredentialData["endDateTime"] = *p.GetEndDateTime()
// 		}
// 		if p.GetKeyId() != nil {
// 			keyCredentialData["keyId"] = *p.GetKeyId()
// 		}
// 		if p.GetStartDateTime() != nil {
// 			keyCredentialData["startDateTime"] = *p.GetStartDateTime()
// 		}
// 		if p.GetType() != nil {
// 			keyCredentialData["type"] = *p.GetType()
// 		}
// 		if p.GetUsage() != nil {
// 			keyCredentialData["usage"] = *p.GetUsage()
// 		}
// 		if p.GetCustomKeyIdentifier() != nil {
// 			keyCredentialData["customKeyIdentifier"] = p.GetCustomKeyIdentifier()
// 		}
// 		if p.GetKey() != nil {
// 			keyCredentialData["key"] = p.GetKey()
// 		}
// 		keyCredentials = append(keyCredentials, keyCredentialData)
// 	}

// 	return keyCredentials
// }

// func (application *ADApplicationInfo) ApplicationParentalControlSettings() map[string]interface{} {
// 	if application.GetParentalControlSettings() == nil {
// 		return nil
// 	}

// 	parentalControlSettingData := map[string]interface{}{
// 		"countriesBlockedForMinors": application.GetParentalControlSettings().GetCountriesBlockedForMinors(),
// 	}
// 	if application.GetParentalControlSettings().GetLegalAgeGroupRule() != nil {
// 		parentalControlSettingData["legalAgeGroupRule"] = *application.GetParentalControlSettings().GetLegalAgeGroupRule()
// 	}

// 	return parentalControlSettingData
// }

// func (application *ADApplicationInfo) ApplicationPasswordCredentials() []map[string]interface{} {
// 	if application.GetPasswordCredentials() == nil {
// 		return nil
// 	}

// 	passwordCredentials := []map[string]interface{}{}
// 	for _, p := range application.GetPasswordCredentials() {
// 		passwordCredentialData := map[string]interface{}{}
// 		if p.GetDisplayName() != nil {
// 			passwordCredentialData["displayName"] = *p.GetDisplayName()
// 		}
// 		if p.GetHint() != nil {
// 			passwordCredentialData["hint"] = *p.GetHint()
// 		}
// 		if p.GetSecretText() != nil {
// 			passwordCredentialData["secretText"] = *p.GetSecretText()
// 		}
// 		if p.GetKeyId() != nil {
// 			passwordCredentialData["keyId"] = *p.GetKeyId()
// 		}
// 		if p.GetEndDateTime() != nil {
// 			passwordCredentialData["endDateTime"] = *p.GetEndDateTime()
// 		}
// 		if p.GetStartDateTime() != nil {
// 			passwordCredentialData["startDateTime"] = *p.GetStartDateTime()
// 		}
// 		if p.GetCustomKeyIdentifier() != nil {
// 			passwordCredentialData["customKeyIdentifier"] = p.GetCustomKeyIdentifier()
// 		}
// 		passwordCredentials = append(passwordCredentials, passwordCredentialData)
// 	}

// 	return passwordCredentials
// }

// func (application *ADApplicationInfo) ApplicationSpa() map[string]interface{} {
// 	if application.GetSpa() == nil {
// 		return nil
// 	}

// 	return map[string]interface{}{
// 		"redirectUris": application.GetSpa().GetRedirectUris(),
// 	}
// }

// func (application *ADApplicationInfo) ApplicationWeb() map[string]interface{} {
// 	if application.GetWeb() == nil {
// 		return nil
// 	}

// 	webData := map[string]interface{}{}
// 	if application.GetWeb().GetHomePageUrl() != nil {
// 		webData["homePageUrl"] = *application.GetWeb().GetHomePageUrl()
// 	}
// 	if application.GetWeb().GetLogoutUrl() != nil {
// 		webData["logoutUrl"] = *application.GetWeb().GetLogoutUrl()
// 	}
// 	if application.GetWeb().GetRedirectUris() != nil {
// 		webData["redirectUris"] = application.GetWeb().GetRedirectUris()
// 	}
// 	if application.GetWeb().GetImplicitGrantSettings() != nil {
// 		implicitGrantSettingsData := map[string]*bool{}

// 		if application.GetWeb().GetImplicitGrantSettings().GetEnableAccessTokenIssuance() != nil {
// 			implicitGrantSettingsData["enableAccessTokenIssuance"] = application.GetWeb().GetImplicitGrantSettings().GetEnableAccessTokenIssuance()
// 		}
// 		if application.GetWeb().GetImplicitGrantSettings().GetEnableIdTokenIssuance() != nil {
// 			implicitGrantSettingsData["enableIdTokenIssuance"] = application.GetWeb().GetImplicitGrantSettings().GetEnableIdTokenIssuance()
// 		}
// 		webData["implicitGrantSettings"] = implicitGrantSettingsData
// 	}

// 	return webData
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicyConditionsApplications() map[string]interface{} {
// 	if conditionalAccessPolicy.GetConditions() == nil {
// 		return nil
// 	}

// 	if conditionalAccessPolicy.GetConditions().GetApplications() == nil {
// 		return nil
// 	}

// 	return map[string]interface{}{
// 		"excludeApplications":                         conditionalAccessPolicy.GetConditions().GetApplications().GetExcludeApplications(),
// 		"includeApplications":                         conditionalAccessPolicy.GetConditions().GetApplications().GetIncludeApplications(),
// 		"includeAuthenticationContextClassReferences": conditionalAccessPolicy.GetConditions().GetApplications().GetIncludeAuthenticationContextClassReferences(),
// 		"includeUserActions":                          conditionalAccessPolicy.GetConditions().GetApplications().GetIncludeUserActions(),
// 	}
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicyConditionsClientAppTypes() []string {
// 	if conditionalAccessPolicy.GetConditions() == nil {
// 		return nil
// 	}
// 	return conditionalAccessPolicy.GetConditions().GetClientAppTypes()
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicyConditionsLocations() map[string]interface{} {
// 	if conditionalAccessPolicy.GetConditions() == nil {
// 		return nil
// 	}

// 	if conditionalAccessPolicy.GetConditions().GetLocations() == nil {
// 		return nil
// 	}

// 	return map[string]interface{}{
// 		"excludeLocations": conditionalAccessPolicy.GetConditions().GetLocations().GetExcludeLocations(),
// 		"includeLocations": conditionalAccessPolicy.GetConditions().GetLocations().GetIncludeLocations(),
// 	}
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicyConditionsPlatforms() map[string]interface{} {
// 	if conditionalAccessPolicy.GetConditions() == nil {
// 		return nil
// 	}

// 	if conditionalAccessPolicy.GetConditions().GetPlatforms() == nil {
// 		return nil
// 	}

// 	return map[string]interface{}{
// 		"excludePlatforms": conditionalAccessPolicy.GetConditions().GetPlatforms().GetExcludePlatforms(),
// 		"includePlatforms": conditionalAccessPolicy.GetConditions().GetPlatforms().GetIncludePlatforms(),
// 	}
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicyConditionsSignInRiskLevels() []string {
// 	if conditionalAccessPolicy.GetConditions() == nil {
// 		return nil
// 	}
// 	return conditionalAccessPolicy.GetConditions().GetSignInRiskLevels()
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicyConditionsUsers() map[string]interface{} {
// 	if conditionalAccessPolicy.GetConditions() == nil {
// 		return nil
// 	}

// 	if conditionalAccessPolicy.GetConditions().GetUsers() == nil {
// 		return nil
// 	}

// 	return map[string]interface{}{
// 		"excludeGroups": conditionalAccessPolicy.GetConditions().GetUsers().GetExcludeGroups(),
// 		"excludeRoles":  conditionalAccessPolicy.GetConditions().GetUsers().GetExcludeRoles(),
// 		"excludeUsers":  conditionalAccessPolicy.GetConditions().GetUsers().GetExcludeUsers(),
// 		"includeGroups": conditionalAccessPolicy.GetConditions().GetUsers().GetIncludeGroups(),
// 		"includeRoles":  conditionalAccessPolicy.GetConditions().GetUsers().GetIncludeRoles(),
// 		"includeUsers":  conditionalAccessPolicy.GetConditions().GetUsers().GetIncludeUsers(),
// 	}
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicyConditionsUserRiskLevels() []string {
// 	if conditionalAccessPolicy.GetConditions() == nil {
// 		return nil
// 	}
// 	return conditionalAccessPolicy.GetConditions().GetUserRiskLevels()
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicyGrantControlsBuiltInControls() []string {
// 	if conditionalAccessPolicy.GetGrantControls() == nil {
// 		return nil
// 	}
// 	return conditionalAccessPolicy.GetGrantControls().GetBuiltInControls()
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicyGrantControlsCustomAuthenticationFactors() []string {
// 	if conditionalAccessPolicy.GetGrantControls() == nil {
// 		return nil
// 	}
// 	return conditionalAccessPolicy.GetGrantControls().GetCustomAuthenticationFactors()
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicyGrantControlsOperator() *string {
// 	if conditionalAccessPolicy.GetGrantControls() == nil {
// 		return nil
// 	}
// 	return conditionalAccessPolicy.GetGrantControls().GetOperator()
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicyGrantControlsTermsOfUse() []string {
// 	if conditionalAccessPolicy.GetGrantControls() == nil {
// 		return nil
// 	}
// 	return conditionalAccessPolicy.GetGrantControls().GetTermsOfUse()
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicySessionControlsApplicationEnforcedRestrictions() map[string]interface{} {
// 	if conditionalAccessPolicy.GetSessionControls() == nil {
// 		return nil
// 	}
// 	if conditionalAccessPolicy.GetSessionControls().GetApplicationEnforcedRestrictions() == nil {
// 		return nil
// 	}

// 	data := map[string]interface{}{}
// 	if conditionalAccessPolicy.GetSessionControls().GetApplicationEnforcedRestrictions().GetIsEnabled() != nil {
// 		data["isEnabled"] = conditionalAccessPolicy.GetSessionControls().GetApplicationEnforcedRestrictions().GetIsEnabled()
// 	}
// 	if conditionalAccessPolicy.GetSessionControls().GetApplicationEnforcedRestrictions().GetType() != nil {
// 		data["type"] = conditionalAccessPolicy.GetSessionControls().GetApplicationEnforcedRestrictions().GetType()
// 	}
// 	return data
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicySessionControlsCloudAppSecurity() map[string]interface{} {
// 	if conditionalAccessPolicy.GetSessionControls() == nil {
// 		return nil
// 	}
// 	if conditionalAccessPolicy.GetSessionControls().GetCloudAppSecurity() == nil {
// 		return nil
// 	}

// 	data := map[string]interface{}{}
// 	if conditionalAccessPolicy.GetSessionControls().GetCloudAppSecurity().GetIsEnabled() != nil {
// 		data["isEnabled"] = conditionalAccessPolicy.GetSessionControls().GetCloudAppSecurity().GetIsEnabled()
// 	}
// 	if conditionalAccessPolicy.GetSessionControls().GetCloudAppSecurity().GetCloudAppSecurityType() != nil {
// 		data["cloudAppSecurityType"] = conditionalAccessPolicy.GetSessionControls().GetCloudAppSecurity().GetCloudAppSecurityType()
// 	}
// 	return data
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicySessionControlsPersistentBrowser() map[string]interface{} {
// 	if conditionalAccessPolicy.GetSessionControls() == nil {
// 		return nil
// 	}
// 	if conditionalAccessPolicy.GetSessionControls().GetPersistentBrowser() == nil {
// 		return nil
// 	}

// 	data := map[string]interface{}{}
// 	if conditionalAccessPolicy.GetSessionControls().GetPersistentBrowser().GetIsEnabled() != nil {
// 		data["isEnabled"] = conditionalAccessPolicy.GetSessionControls().GetPersistentBrowser().GetIsEnabled()
// 	}
// 	if conditionalAccessPolicy.GetSessionControls().GetPersistentBrowser().GetMode() != nil {
// 		data["mode"] = conditionalAccessPolicy.GetSessionControls().GetPersistentBrowser().GetMode()
// 	}
// 	return data
// }

// func (conditionalAccessPolicy *ADConditionalAccessPolicyInfo) ConditionalAccessPolicySessionControlsSignInFrequency() map[string]interface{} {
// 	if conditionalAccessPolicy.GetSessionControls() == nil {
// 		return nil
// 	}
// 	if conditionalAccessPolicy.GetSessionControls().GetSignInFrequency() == nil {
// 		return nil
// 	}

// 	data := map[string]interface{}{}
// 	if conditionalAccessPolicy.GetSessionControls().GetSignInFrequency().GetIsEnabled() != nil {
// 		data["isEnabled"] = conditionalAccessPolicy.GetSessionControls().GetSignInFrequency().GetIsEnabled()
// 	}
// 	if conditionalAccessPolicy.GetSessionControls().GetSignInFrequency().GetValue() != nil {
// 		data["value"] = conditionalAccessPolicy.GetSessionControls().GetSignInFrequency().GetValue()
// 	}
// 	return data
// }

// func (group *ADGroupInfo) GroupAssignedLabels() []map[string]*string {
// 	if group.GetAssignedLabels() == nil {
// 		return nil
// 	}

// 	assignedLabels := []map[string]*string{}
// 	for _, i := range group.GetAssignedLabels() {
// 		label := map[string]*string{
// 			"labelId":     i.GetLabelId(),
// 			"displayName": i.GetDisplayName(),
// 		}
// 		assignedLabels = append(assignedLabels, label)
// 	}
// 	return assignedLabels
// }

// func (servicePrincipal *ADServicePrincipalInfo) ServicePrincipalAddIns() []map[string]interface{} {
// 	if servicePrincipal.GetAddIns() == nil {
// 		return nil
// 	}

// 	addIns := []map[string]interface{}{}
// 	for _, p := range servicePrincipal.GetAddIns() {
// 		addInData := map[string]interface{}{}
// 		if p.GetId() != nil {
// 			addInData["id"] = *p.GetId()
// 		}
// 		if p.GetType() != nil {
// 			addInData["type"] = *p.GetType()
// 		}

// 		addInProperties := []map[string]interface{}{}
// 		for _, k := range p.GetProperties() {
// 			addInPropertyData := map[string]interface{}{}
// 			if k.GetKey() != nil {
// 				addInPropertyData["key"] = *k.GetKey()
// 			}
// 			if k.GetValue() != nil {
// 				addInPropertyData["value"] = *k.GetValue()
// 			}
// 			addInProperties = append(addInProperties, addInPropertyData)
// 		}
// 		addInData["properties"] = addInProperties

// 		addIns = append(addIns, addInData)
// 	}
// 	return addIns
// }

// func (servicePrincipal *ADServicePrincipalInfo) ServicePrincipalAppRoles() []map[string]interface{} {
// 	if servicePrincipal.GetAppRoles() == nil {
// 		return nil
// 	}

// 	appRoles := []map[string]interface{}{}
// 	for _, p := range servicePrincipal.GetAppRoles() {
// 		appRoleData := map[string]interface{}{
// 			"allowedMemberTypes": p.GetAllowedMemberTypes(),
// 		}
// 		if p.GetDescription() != nil {
// 			appRoleData["description"] = *p.GetDescription()
// 		}
// 		if p.GetDisplayName() != nil {
// 			appRoleData["displayName"] = *p.GetDisplayName()
// 		}
// 		if p.GetId() != nil {
// 			appRoleData["id"] = *p.GetId()
// 		}
// 		if p.GetIsEnabled() != nil {
// 			appRoleData["isEnabled"] = *p.GetIsEnabled()
// 		}
// 		if p.GetOrigin() != nil {
// 			appRoleData["origin"] = *p.GetOrigin()
// 		}
// 		if p.GetValue() != nil {
// 			appRoleData["value"] = *p.GetValue()
// 		}
// 		appRoles = append(appRoles, appRoleData)
// 	}
// 	return appRoles
// }

// func (servicePrincipal *ADServicePrincipalInfo) ServicePrincipalInfo() map[string]interface{} {
// 	if servicePrincipal.GetInfo() == nil {
// 		return nil
// 	}

// 	return map[string]interface{}{
// 		"logoUrl":             servicePrincipal.GetInfo().GetLogoUrl(),
// 		"marketingUrl":        servicePrincipal.GetInfo().GetMarketingUrl(),
// 		"privacyStatementUrl": servicePrincipal.GetInfo().GetPrivacyStatementUrl(),
// 		"supportUrl":          servicePrincipal.GetInfo().GetSupportUrl(),
// 		"termsOfServiceUrl":   servicePrincipal.GetInfo().GetTermsOfServiceUrl(),
// 	}
// }

// func (servicePrincipal *ADServicePrincipalInfo) ServicePrincipalKeyCredentials() []map[string]interface{} {
// 	if servicePrincipal.GetKeyCredentials() == nil {
// 		return nil
// 	}

// 	keyCredentials := []map[string]interface{}{}
// 	for _, p := range servicePrincipal.GetKeyCredentials() {
// 		keyCredentialData := map[string]interface{}{}
// 		if p.GetDisplayName() != nil {
// 			keyCredentialData["displayName"] = *p.GetDisplayName()
// 		}
// 		if p.GetEndDateTime() != nil {
// 			keyCredentialData["endDateTime"] = *p.GetEndDateTime()
// 		}
// 		if p.GetKeyId() != nil {
// 			keyCredentialData["keyId"] = *p.GetKeyId()
// 		}
// 		if p.GetStartDateTime() != nil {
// 			keyCredentialData["startDateTime"] = *p.GetStartDateTime()
// 		}
// 		if p.GetType() != nil {
// 			keyCredentialData["type"] = *p.GetType()
// 		}
// 		if p.GetUsage() != nil {
// 			keyCredentialData["usage"] = *p.GetUsage()
// 		}
// 		if p.GetCustomKeyIdentifier() != nil {
// 			keyCredentialData["customKeyIdentifier"] = p.GetCustomKeyIdentifier()
// 		}
// 		if p.GetKey() != nil {
// 			keyCredentialData["key"] = p.GetKey()
// 		}
// 		keyCredentials = append(keyCredentials, keyCredentialData)
// 	}
// 	return keyCredentials
// }

// func (servicePrincipal *ADServicePrincipalInfo) ServicePrincipalOauth2PermissionScopes() []map[string]interface{} {
// 	if servicePrincipal.GetOauth2PermissionScopes() == nil {
// 		return nil
// 	}

// 	oauth2PermissionScopes := []map[string]interface{}{}
// 	for _, p := range servicePrincipal.GetOauth2PermissionScopes() {
// 		data := map[string]interface{}{}
// 		if p.GetAdminConsentDescription() != nil {
// 			data["adminConsentDescription"] = *p.GetAdminConsentDescription()
// 		}
// 		if p.GetAdminConsentDisplayName() != nil {
// 			data["adminConsentDisplayName"] = *p.GetAdminConsentDisplayName()
// 		}
// 		if p.GetId() != nil {
// 			data["id"] = *p.GetId()
// 		}
// 		if p.GetIsEnabled() != nil {
// 			data["isEnabled"] = *p.GetIsEnabled()
// 		}
// 		if p.GetType() != nil {
// 			data["type"] = *p.GetType()
// 		}
// 		if p.GetOrigin() != nil {
// 			data["origin"] = *p.GetOrigin()
// 		}
// 		if p.GetUserConsentDescription() != nil {
// 			data["userConsentDescription"] = p.GetUserConsentDescription()
// 		}
// 		if p.GetUserConsentDisplayName() != nil {
// 			data["userConsentDisplayName"] = p.GetUserConsentDisplayName()
// 		}
// 		if p.GetValue() != nil {
// 			data["value"] = p.GetValue()
// 		}
// 		oauth2PermissionScopes = append(oauth2PermissionScopes, data)
// 	}
// 	return oauth2PermissionScopes
// }

// func (servicePrincipal *ADServicePrincipalInfo) ServicePrincipalPasswordCredentials() []map[string]interface{} {
// 	if servicePrincipal.GetPasswordCredentials() == nil {
// 		return nil
// 	}

// 	passwordCredentials := []map[string]interface{}{}
// 	for _, p := range servicePrincipal.GetPasswordCredentials() {
// 		passwordCredentialData := map[string]interface{}{}
// 		if p.GetDisplayName() != nil {
// 			passwordCredentialData["displayName"] = *p.GetDisplayName()
// 		}
// 		if p.GetHint() != nil {
// 			passwordCredentialData["hint"] = *p.GetHint()
// 		}
// 		if p.GetSecretText() != nil {
// 			passwordCredentialData["secretText"] = *p.GetSecretText()
// 		}
// 		if p.GetKeyId() != nil {
// 			passwordCredentialData["keyId"] = *p.GetKeyId()
// 		}
// 		if p.GetEndDateTime() != nil {
// 			passwordCredentialData["endDateTime"] = *p.GetEndDateTime()
// 		}
// 		if p.GetStartDateTime() != nil {
// 			passwordCredentialData["startDateTime"] = *p.GetStartDateTime()
// 		}
// 		if p.GetCustomKeyIdentifier() != nil {
// 			passwordCredentialData["customKeyIdentifier"] = p.GetCustomKeyIdentifier()
// 		}
// 		passwordCredentials = append(passwordCredentials, passwordCredentialData)
// 	}
// 	return passwordCredentials
// }

// func (signIn *ADSignInReportInfo) SignInAppliedConditionalAccessPolicies() []map[string]interface{} {
// 	if signIn.GetAppliedConditionalAccessPolicies() == nil {
// 		return nil
// 	}

// 	policies := []map[string]interface{}{}
// 	for _, p := range signIn.GetAppliedConditionalAccessPolicies() {
// 		policyData := map[string]interface{}{
// 			"enforcedGrantControls":   p.GetEnforcedGrantControls(),
// 			"enforcedSessionControls": p.GetEnforcedSessionControls(),
// 		}
// 		if p.GetDisplayName() != nil {
// 			policyData["displayName"] = *p.GetDisplayName()
// 		}
// 		if p.GetId() != nil {
// 			policyData["id"] = *p.GetId()
// 		}
// 		if p.GetResult() != nil {
// 			policyData["result"] = p.GetResult()
// 		}
// 		policies = append(policies, policyData)
// 	}

// 	return policies
// }

// func (signIn *ADSignInReportInfo) SignInDeviceDetail() map[string]interface{} {
// 	if signIn.GetDeviceDetail() == nil {
// 		return nil
// 	}

// 	deviceDetailInfo := map[string]interface{}{}
// 	if signIn.GetDeviceDetail().GetBrowser() != nil {
// 		deviceDetailInfo["browser"] = *signIn.GetDeviceDetail().GetBrowser()
// 	}
// 	if signIn.GetDeviceDetail().GetDeviceId() != nil {
// 		deviceDetailInfo["deviceId"] = *signIn.GetDeviceDetail().GetDeviceId()
// 	}
// 	if signIn.GetDeviceDetail().GetDisplayName() != nil {
// 		deviceDetailInfo["displayName"] = *signIn.GetDeviceDetail().GetDisplayName()
// 	}
// 	if signIn.GetDeviceDetail().GetIsCompliant() != nil {
// 		deviceDetailInfo["isCompliant"] = *signIn.GetDeviceDetail().GetIsCompliant()
// 	}
// 	if signIn.GetDeviceDetail().GetIsManaged() != nil {
// 		deviceDetailInfo["isManaged"] = *signIn.GetDeviceDetail().GetIsManaged()
// 	}
// 	if signIn.GetDeviceDetail().GetOperatingSystem() != nil {
// 		deviceDetailInfo["operatingSystem"] = *signIn.GetDeviceDetail().GetOperatingSystem()
// 	}
// 	if signIn.GetDeviceDetail().GetTrustType() != nil {
// 		deviceDetailInfo["trustType"] = *signIn.GetDeviceDetail().GetTrustType()
// 	}
// 	return deviceDetailInfo
// }

// func (signIn *ADSignInReportInfo) SignInStatus() map[string]interface{} {
// 	if signIn.GetStatus() == nil {
// 		return nil
// 	}

// 	statusInfo := map[string]interface{}{}
// 	if signIn.GetStatus().GetErrorCode() != nil {
// 		statusInfo["errorCode"] = *signIn.GetStatus().GetErrorCode()
// 	}
// 	if signIn.GetStatus().GetFailureReason() != nil {
// 		statusInfo["failureReason"] = *signIn.GetStatus().GetFailureReason()
// 	}
// 	if signIn.GetStatus().GetAdditionalDetails() != nil {
// 		statusInfo["additionalDetails"] = *signIn.GetStatus().GetAdditionalDetails()
// 	}
// 	return statusInfo
// }

// func (signIn *ADSignInReportInfo) SignInLocation() map[string]interface{} {
// 	if signIn.GetLocation() == nil {
// 		return nil
// 	}

// 	locationInfo := map[string]interface{}{}
// 	if signIn.GetLocation().GetCity() != nil {
// 		locationInfo["city"] = *signIn.GetLocation().GetCity()
// 	}
// 	if signIn.GetLocation().GetCountryOrRegion() != nil {
// 		locationInfo["countryOrRegion"] = *signIn.GetLocation().GetCountryOrRegion()
// 	}
// 	if signIn.GetLocation().GetState() != nil {
// 		locationInfo["state"] = *signIn.GetLocation().GetState()
// 	}
// 	if signIn.GetLocation().GetGeoCoordinates() != nil {
// 		coordinateInfo := map[string]interface{}{}
// 		if signIn.GetLocation().GetGeoCoordinates().GetAltitude() != nil {
// 			coordinateInfo["altitude"] = *signIn.GetLocation().GetGeoCoordinates().GetAltitude()
// 		}
// 		if signIn.GetLocation().GetGeoCoordinates().GetLatitude() != nil {
// 			coordinateInfo["latitude"] = *signIn.GetLocation().GetGeoCoordinates().GetLatitude()
// 		}
// 		if signIn.GetLocation().GetGeoCoordinates().GetLongitude() != nil {
// 			coordinateInfo["longitude"] = *signIn.GetLocation().GetGeoCoordinates().GetLongitude()
// 		}
// 		locationInfo["geoCoordinates"] = coordinateInfo
// 	}
// 	return locationInfo
// }

// func (user *ADUserInfo) UserMemberOf() []map[string]interface{} {
// 	if user.GetMemberOf() == nil {
// 		return nil
// 	}

// 	members := []map[string]interface{}{}
// 	for _, i := range user.GetMemberOf() {
// 		member := map[string]interface{}{
// 			"@odata.type": i.GetType(),
// 			"id":          i.GetId(),
// 		}
// 		members = append(members, member)
// 	}
// 	return members
// }

// func (user *ADUserInfo) UserPasswordProfile() map[string]interface{} {
// 	if user.GetPasswordProfile() == nil {
// 		return nil
// 	}

// 	passwordProfileData := map[string]interface{}{}
// 	if user.GetPasswordProfile().GetForceChangePasswordNextSignIn() != nil {
// 		passwordProfileData["forceChangePasswordNextSignIn"] = *user.GetPasswordProfile().GetForceChangePasswordNextSignIn()
// 	}
// 	if user.GetPasswordProfile().GetForceChangePasswordNextSignInWithMfa() != nil {
// 		passwordProfileData["forceChangePasswordNextSignInWithMfa"] = *user.GetPasswordProfile().GetForceChangePasswordNextSignInWithMfa()
// 	}
// 	if user.GetPasswordProfile().GetPassword() != nil {
// 		passwordProfileData["password"] = *user.GetPasswordProfile().GetPassword()
// 	}

// 	return passwordProfileData
// }
