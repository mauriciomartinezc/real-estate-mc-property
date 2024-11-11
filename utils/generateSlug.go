package utils

import (
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"strings"
	"unicode"
)

// GenerateSlug creates a slug from a SimpleProperty, utilizing fixed attributes.
func GenerateSlug(property *domain.SimpleProperty) string {
	var slugParts []string

	// Append necessary attributes to form the slug in the required order
	if property.PropertyType.Name != "" {
		slugParts = append(slugParts, property.PropertyType.Name)
	}
	for _, managementType := range property.ManagementTypes {
		if managementType.Name != "" {
			slugParts = append(slugParts, managementType.Name)
		}
	}
	if property.City.Name != "" {
		slugParts = append(slugParts, property.City.Name)
	}
	if property.Neighborhood.Name != "" {
		slugParts = append(slugParts, property.Neighborhood.Name)
	}

	// Append unique ID (Hex) as the final part to ensure uniqueness
	slugParts = append(slugParts, property.ID.Hex())

	// Join parts with hyphens and sanitize the final slug
	slug := strings.Join(slugParts, "-")
	return sanitizeSlug(slug)
}

// sanitizeSlug cleans the slug, converting it to lowercase and replacing invalid characters with hyphens.
func sanitizeSlug(slug string) string {
	var sanitized strings.Builder
	for _, r := range slug {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' {
			sanitized.WriteRune(unicode.ToLower(r))
		} else {
			// Replace any non-alphanumeric character with a hyphen
			sanitized.WriteRune('-')
		}
	}
	return sanitized.String()
}
