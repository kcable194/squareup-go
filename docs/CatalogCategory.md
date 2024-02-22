# CatalogCategory

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The category name. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points. | [optional] [default to null]
**ImageIds** | **[]string** | The IDs of images associated with this &#x60;CatalogCategory&#x60; instance. Currently these images are not displayed by Square, but are free to be displayed in 3rd party applications. | [optional] [default to null]
**CategoryType** | **string** | The type of the category. | [optional] [default to null]
**ParentCategory** | [***CatalogObjectCategory**](CatalogObjectCategory.md) |  | [optional] [default to null]
**IsTopLevel** | **bool** | Indicates whether a category is a top level category, which does not have any parent_category. | [optional] [default to null]
**Channels** | **[]string** | A list of IDs representing channels, such as a Square Online site, where the category can be made visible. | [optional] [default to null]
**AvailabilityPeriodIds** | **[]string** | The IDs of the &#x60;CatalogAvailabilityPeriod&#x60; objects associated with the category. | [optional] [default to null]
**OnlineVisibility** | **bool** | Indicates whether the category is visible (&#x60;true&#x60;) or hidden (&#x60;false&#x60;) on all of the seller&#x27;s Square Online sites. | [optional] [default to null]
**RootCategory** | **string** | The top-level category in a category hierarchy. | [optional] [default to null]
**EcomSeoData** | [***CatalogEcomSeoData**](CatalogEcomSeoData.md) |  | [optional] [default to null]
**PathToRoot** | [**[]CategoryPathToRootNode**](CategoryPathToRootNode.md) | The path from the category to its root category. The first node of the path is the parent of the category and the last is the root category. The path is empty if the category is a root category. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

