package resources

import (
	"errors"
	"fmt"
)

const suffixRelatedList = "RelatedList"

// Account struct
type Account struct {
	method   string
	metadata map[string]interface{}
}

func (a *Account) objectName() string {
	const obName = "Account"
	return obName
}

func (a *Account) BuildConnectionKey() string {
	objectName := "account"
	return objectName+"_"+a.method
}

// NewAccount writes that new Customer instance
func NewAccount(metadata map[string]interface{}) (*Account, error) {
	rawMethod, ok := metadata["method"]
	if !ok {
		return nil, errors.New("missing required parameters: method")
	}
	method, ok := rawMethod.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	return &Account{
		method:   method,
		metadata: metadata,
	}, nil
}

// getMetadata mold customer get metadata
func (a *Account) getMetadata() (map[string]interface{}, error) {
	return buildMetadata("account_bulk_get", a.method, a.objectName()+suffixRelatedList, "", nil,""), nil
}


// BuildMetadata
func (a *Account) BuildMetadata() (map[string]interface{}, error) {
	switch a.method {
	case "get":
		return a.getMetadata()
	}
	return nil, fmt.Errorf("invalid method: %s", a.method)
}

func buildMetadata(connectionKey, method, object, pathParam string, queryParams map[string]string, body string) map[string]interface{} {
	metadata := map[string]interface{}{
		"method":         method,
		"object":         object,
		"connection_key": connectionKey,
	}
	if len(pathParam) > 0 {
		metadata["path_param"] = pathParam
	}
	if queryParams != nil {
		metadata["query_params"] = queryParams
	}
	if body != "" {
		metadata["body"] = body
	}
	return metadata
}