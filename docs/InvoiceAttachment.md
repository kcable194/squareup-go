# InvoiceAttachment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the attachment. | [optional] [default to null]
**Filename** | **string** | The file name of the attachment, which is displayed on the invoice. | [optional] [default to null]
**Description** | **string** | The description of the attachment, which is displayed on the invoice. This field maps to the seller-defined **Message** field. | [optional] [default to null]
**Filesize** | **int32** | The file size of the attachment in bytes. | [optional] [default to null]
**Hash** | **string** | The MD5 hash that was generated from the file contents. | [optional] [default to null]
**MimeType** | **string** | The mime type of the attachment. The following mime types are supported:  image/gif, image/jpeg, image/png, image/tiff, image/bmp, application/pdf. | [optional] [default to null]
**UploadedAt** | **string** | The timestamp when the attachment was uploaded, in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

