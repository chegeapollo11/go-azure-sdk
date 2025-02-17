package diagnostic

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DiagnosticId{}

// DiagnosticId is a struct representing the Resource ID for a Diagnostic
type DiagnosticId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	DiagnosticId      string
}

// NewDiagnosticID returns a new DiagnosticId struct
func NewDiagnosticID(subscriptionId string, resourceGroupName string, serviceName string, diagnosticId string) DiagnosticId {
	return DiagnosticId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		DiagnosticId:      diagnosticId,
	}
}

// ParseDiagnosticID parses 'input' into a DiagnosticId
func ParseDiagnosticID(input string) (*DiagnosticId, error) {
	parser := resourceids.NewParserFromResourceIdType(DiagnosticId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DiagnosticId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.DiagnosticId, ok = parsed.Parsed["diagnosticId"]; !ok {
		return nil, fmt.Errorf("the segment 'diagnosticId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDiagnosticIDInsensitively parses 'input' case-insensitively into a DiagnosticId
// note: this method should only be used for API response data and not user input
func ParseDiagnosticIDInsensitively(input string) (*DiagnosticId, error) {
	parser := resourceids.NewParserFromResourceIdType(DiagnosticId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DiagnosticId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.DiagnosticId, ok = parsed.Parsed["diagnosticId"]; !ok {
		return nil, fmt.Errorf("the segment 'diagnosticId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDiagnosticID checks that 'input' can be parsed as a Diagnostic ID
func ValidateDiagnosticID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDiagnosticID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Diagnostic ID
func (id DiagnosticId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/diagnostics/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.DiagnosticId)
}

// Segments returns a slice of Resource ID Segments which comprise this Diagnostic ID
func (id DiagnosticId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticDiagnostics", "diagnostics", "diagnostics"),
		resourceids.UserSpecifiedSegment("diagnosticId", "diagnosticIdValue"),
	}
}

// String returns a human-readable description of this Diagnostic ID
func (id DiagnosticId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Diagnostic: %q", id.DiagnosticId),
	}
	return fmt.Sprintf("Diagnostic (%s)", strings.Join(components, "\n"))
}
