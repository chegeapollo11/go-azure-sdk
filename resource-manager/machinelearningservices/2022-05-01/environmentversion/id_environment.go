package environmentversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = EnvironmentId{}

// EnvironmentId is a struct representing the Resource ID for a Environment
type EnvironmentId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	EnvironmentName   string
}

// NewEnvironmentID returns a new EnvironmentId struct
func NewEnvironmentID(subscriptionId string, resourceGroupName string, workspaceName string, environmentName string) EnvironmentId {
	return EnvironmentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		EnvironmentName:   environmentName,
	}
}

// ParseEnvironmentID parses 'input' into a EnvironmentId
func ParseEnvironmentID(input string) (*EnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(EnvironmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EnvironmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EnvironmentName, ok = parsed.Parsed["environmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'environmentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseEnvironmentIDInsensitively parses 'input' case-insensitively into a EnvironmentId
// note: this method should only be used for API response data and not user input
func ParseEnvironmentIDInsensitively(input string) (*EnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(EnvironmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EnvironmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EnvironmentName, ok = parsed.Parsed["environmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'environmentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateEnvironmentID checks that 'input' can be parsed as a Environment ID
func ValidateEnvironmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEnvironmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Environment ID
func (id EnvironmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/environments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.EnvironmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Environment ID
func (id EnvironmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticEnvironments", "environments", "environments"),
		resourceids.UserSpecifiedSegment("environmentName", "environmentValue"),
	}
}

// String returns a human-readable description of this Environment ID
func (id EnvironmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Environment Name: %q", id.EnvironmentName),
	}
	return fmt.Sprintf("Environment (%s)", strings.Join(components, "\n"))
}
