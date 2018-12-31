package onepassword

import (
	"fmt"

	"github.com/sganon/env-secrets/common"
)

const (
	opBinary = "op"
)

// OP wraps utility methods to interact with op CLI
type OP struct {
	Tags       []string
	secretsIDs []ItemID
	values     map[string]string
}

// GenerateEnv implements common.EnvGenerator
func (op OP) GenerateEnv() string {
	var str string
	for key, value := range op.values {
		str += fmt.Sprintf("export %s=%s\n", key, value)
	}
	return str
}

// SetSecretsIDs get items IDs which match given tags
func (op *OP) SetSecretsIDs() error {
	var items []Item
	err := common.ExecCLI(opBinary, []string{"list", "items"}, &items)
	if err != nil {
		return fmt.Errorf("env-secrets error: an error occurred getting secrets IDs: %v", err)
	}
	for _, item := range items {
		for _, itemTag := range item.Overview.Tags {
			for _, targetTag := range op.Tags {
				if itemTag == targetTag {
					op.secretsIDs = append(op.secretsIDs, item.ID)
				}
			}
		}
	}
	return nil
}

// SetValues fetch and set values via secretsIDs
func (op *OP) SetValues() error {
	if op.values == nil {
		op.values = make(map[string]string)
	}
	for _, id := range op.secretsIDs {
		var item Item
		err := common.ExecCLI(opBinary, []string{"get", "item", string(id)}, &item)
		if err != nil {
			return fmt.Errorf("env-secrets error: error getting secret value: %v", err)
		}
		op.values[item.Overview.Title] = item.Details.Password

	}
	return nil
}
