package rolemanagementpolicyassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ScopedRoleManagementPolicyAssignmentId{}

// ScopedRoleManagementPolicyAssignmentId is a struct representing the Resource ID for a Scoped Role Management Policy Assignment
type ScopedRoleManagementPolicyAssignmentId struct {
	Scope                              string
	RoleManagementPolicyAssignmentName string
}

// NewScopedRoleManagementPolicyAssignmentID returns a new ScopedRoleManagementPolicyAssignmentId struct
func NewScopedRoleManagementPolicyAssignmentID(scope string, roleManagementPolicyAssignmentName string) ScopedRoleManagementPolicyAssignmentId {
	return ScopedRoleManagementPolicyAssignmentId{
		Scope:                              scope,
		RoleManagementPolicyAssignmentName: roleManagementPolicyAssignmentName,
	}
}

// ParseScopedRoleManagementPolicyAssignmentID parses 'input' into a ScopedRoleManagementPolicyAssignmentId
func ParseScopedRoleManagementPolicyAssignmentID(input string) (*ScopedRoleManagementPolicyAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedRoleManagementPolicyAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedRoleManagementPolicyAssignmentId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, fmt.Errorf("the segment 'scope' was not found in the resource id %q", input)
	}

	if id.RoleManagementPolicyAssignmentName, ok = parsed.Parsed["roleManagementPolicyAssignmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'roleManagementPolicyAssignmentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseScopedRoleManagementPolicyAssignmentIDInsensitively parses 'input' case-insensitively into a ScopedRoleManagementPolicyAssignmentId
// note: this method should only be used for API response data and not user input
func ParseScopedRoleManagementPolicyAssignmentIDInsensitively(input string) (*ScopedRoleManagementPolicyAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedRoleManagementPolicyAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedRoleManagementPolicyAssignmentId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, fmt.Errorf("the segment 'scope' was not found in the resource id %q", input)
	}

	if id.RoleManagementPolicyAssignmentName, ok = parsed.Parsed["roleManagementPolicyAssignmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'roleManagementPolicyAssignmentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateScopedRoleManagementPolicyAssignmentID checks that 'input' can be parsed as a Scoped Role Management Policy Assignment ID
func ValidateScopedRoleManagementPolicyAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedRoleManagementPolicyAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Role Management Policy Assignment ID
func (id ScopedRoleManagementPolicyAssignmentId) ID() string {
	fmtString := "/%s/providers/Microsoft.Authorization/roleManagementPolicyAssignments/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.RoleManagementPolicyAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Role Management Policy Assignment ID
func (id ScopedRoleManagementPolicyAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticRoleManagementPolicyAssignments", "roleManagementPolicyAssignments", "roleManagementPolicyAssignments"),
		resourceids.UserSpecifiedSegment("roleManagementPolicyAssignmentName", "roleManagementPolicyAssignmentValue"),
	}
}

// String returns a human-readable description of this Scoped Role Management Policy Assignment ID
func (id ScopedRoleManagementPolicyAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Role Management Policy Assignment Name: %q", id.RoleManagementPolicyAssignmentName),
	}
	return fmt.Sprintf("Scoped Role Management Policy Assignment (%s)", strings.Join(components, "\n"))
}
