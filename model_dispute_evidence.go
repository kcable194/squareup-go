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

type DisputeEvidence struct {
	// The Square-generated ID of the evidence.
	EvidenceId string `json:"evidence_id,omitempty"`
	// The Square-generated ID of the evidence.
	Id string `json:"id,omitempty"`
	// The ID of the dispute the evidence is associated with.
	DisputeId string `json:"dispute_id,omitempty"`
	EvidenceFile *DisputeEvidenceFile `json:"evidence_file,omitempty"`
	// Raw text
	EvidenceText string `json:"evidence_text,omitempty"`
	// The time when the evidence was uploaded, in RFC 3339 format.
	UploadedAt string `json:"uploaded_at,omitempty"`
	// The type of the evidence.
	EvidenceType string `json:"evidence_type,omitempty"`
}
