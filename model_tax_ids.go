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

// Identifiers for the location used by various governments for tax purposes.
type TaxIds struct {
	// The EU VAT number for this location. For example, `IE3426675K`. If the EU VAT number is present, it is well-formed and has been validated with VIES, the VAT Information Exchange System.
	EuVat string `json:"eu_vat,omitempty"`
	// The SIRET (Système d'Identification du Répertoire des Entreprises et de leurs Etablissements) number is a 14-digit code issued by the French INSEE. For example, `39922799000021`.
	FrSiret string `json:"fr_siret,omitempty"`
	// The French government uses the NAF (Nomenclature des Activités Françaises) to display and track economic statistical data. This is also called the APE (Activite Principale de l’Entreprise) code. For example, `6910Z`.
	FrNaf string `json:"fr_naf,omitempty"`
	// The NIF (Numero de Identificacion Fiscal) number is a nine-character tax identifier used in Spain. If it is present, it has been validated. For example, `73628495A`.
	EsNif string `json:"es_nif,omitempty"`
	// The QII (Qualified Invoice Issuer) number is a 14-character tax identifier used in Japan. For example, `T1234567890123`.
	JpQii string `json:"jp_qii,omitempty"`
}
