package office365

import (
	"context"
	"encoding/json"
	"strings"

	betaErrors "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

type RequestError struct {
	Code    string
	Message string
}

func (m *RequestError) Error() string {
	errStr, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(errStr)
}

func getErrorObject(err error) *RequestError {
	if oDataError, ok := err.(*odataerrors.ODataError); ok {
		if terr := oDataError.GetError(); terr != nil {
			return &RequestError{
				Code:    *terr.GetCode(),
				Message: *terr.GetMessage(),
			}
		}
	}

	return nil
}

func getBetaErrorObject(err error) *RequestError {
	if oDataError, ok := err.(*betaErrors.ODataError); ok {
		if terr := oDataError.GetError(); terr != nil {
			return &RequestError{
				Code:    *terr.GetCode(),
				Message: *terr.GetMessage(),
			}
		}
	}

	return nil
}

func isIgnorableErrorPredicate(ignoreErrorCodes []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		if err != nil {
			if terr, ok := err.(*RequestError); ok {
				for _, item := range ignoreErrorCodes {
					if terr != nil && (terr.Code == item || strings.Contains(terr.Message, item)) {
						return true
					}
				}
			}
		}
		return false
	}
}
