// Code generated by internal/rbac/yamldata. DO NOT EDIT.
package types

// NamespaceAction represents the action permitted in a namespace.
type NamespaceAction string

func (a NamespaceAction) String() string {
	return string(a)
}

const BatchChangesReadAction NamespaceAction = "READ"
const BatchChangesWriteAction NamespaceAction = "WRITE"
const OwnershipAssignAction NamespaceAction = "ASSIGN"
