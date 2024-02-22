/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package squareup

// An instance of a custom attribute. Custom attributes can be defined and added to `ITEM` and `ITEM_VARIATION` type catalog objects. [Read more about custom attributes](https://developer.squareup.com/docs/catalog-api/add-custom-attributes).
type CatalogCustomAttributeValue struct {
	// The name of the custom attribute.
	Name string `json:"name,omitempty"`
	// The string value of the custom attribute.  Populated if `type` = `STRING`.
	StringValue string `json:"string_value,omitempty"`
	// The id of the [CatalogCustomAttributeDefinition](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogCustomAttributeDefinition) this value belongs to.
	CustomAttributeDefinitionId string `json:"custom_attribute_definition_id,omitempty"`
	// A copy of type from the associated `CatalogCustomAttributeDefinition`.
	Type_ string `json:"type,omitempty"`
	// Populated if `type` = `NUMBER`. Contains a string representation of a decimal number, using a `.` as the decimal separator.
	NumberValue string `json:"number_value,omitempty"`
	// A `true` or `false` value. Populated if `type` = `BOOLEAN`.
	BooleanValue bool `json:"boolean_value,omitempty"`
	// One or more choices from `allowed_selections`. Populated if `type` = `SELECTION`.
	SelectionUidValues []string `json:"selection_uid_values,omitempty"`
	// If the associated `CatalogCustomAttributeDefinition` object is defined by another application, this key is prefixed by the defining application ID. For example, if the CatalogCustomAttributeDefinition has a key attribute of \"cocoa_brand\" and the defining application ID is \"abcd1234\", this key is \"abcd1234:cocoa_brand\" when the application making the request is different from the application defining the custom attribute definition. Otherwise, the key is simply \"cocoa_brand\".
	Key string `json:"key,omitempty"`
}
