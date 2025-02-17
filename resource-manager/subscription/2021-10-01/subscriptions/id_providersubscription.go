package subscriptions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ProviderSubscriptionId{}

// ProviderSubscriptionId is a struct representing the Resource ID for a Provider Subscription
type ProviderSubscriptionId struct {
	SubscriptionId string
}

// NewProviderSubscriptionID returns a new ProviderSubscriptionId struct
func NewProviderSubscriptionID(subscriptionId string) ProviderSubscriptionId {
	return ProviderSubscriptionId{
		SubscriptionId: subscriptionId,
	}
}

// ParseProviderSubscriptionID parses 'input' into a ProviderSubscriptionId
func ParseProviderSubscriptionID(input string) (*ProviderSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderSubscriptionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseProviderSubscriptionIDInsensitively parses 'input' case-insensitively into a ProviderSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseProviderSubscriptionIDInsensitively(input string) (*ProviderSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderSubscriptionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateProviderSubscriptionID checks that 'input' can be parsed as a Provider Subscription ID
func ValidateProviderSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Subscription ID
func (id ProviderSubscriptionId) ID() string {
	fmtString := "/providers/Microsoft.Subscription/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Subscription ID
func (id ProviderSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSubscription", "Microsoft.Subscription", "Microsoft.Subscription"),
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
	}
}

// String returns a human-readable description of this Provider Subscription ID
func (id ProviderSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("Provider Subscription (%s)", strings.Join(components, "\n"))
}
