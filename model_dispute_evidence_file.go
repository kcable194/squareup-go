/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A file to be uploaded as dispute evidence.
type DisputeEvidenceFile struct {
	// The file name including the file extension. For example: \"receipt.tiff\".
	Filename string `json:"filename,omitempty"`
	// Dispute evidence files must be application/pdf, image/heic, image/heif, image/jpeg, image/png, or image/tiff formats.
	Filetype string `json:"filetype,omitempty"`
}