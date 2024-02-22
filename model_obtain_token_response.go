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

type ObtainTokenResponse struct {
	// A valid OAuth access token.  Provide the access token in a header with every request to Connect API endpoints. For more information, see [OAuth API: Walkthrough](https://developer.squareup.com/docs/oauth-api/walkthrough).
	AccessToken string `json:"access_token,omitempty"`
	// This value is always _bearer_.
	TokenType string `json:"token_type,omitempty"`
	// The date when the `access_token` expires, in [ISO 8601](http://www.iso.org/iso/home/standards/iso8601.htm) format.
	ExpiresAt string `json:"expires_at,omitempty"`
	// The ID of the authorizing merchant's business.
	MerchantId string `json:"merchant_id,omitempty"`
	// __LEGACY FIELD__. The ID of a subscription plan the merchant signed up for. The ID is only present if the merchant signed up for a subscription plan during authorization.
	SubscriptionId string `json:"subscription_id,omitempty"`
	// __LEGACY FIELD__. The ID of the subscription plan the merchant signed up for. The ID is only present if the merchant signed up for a subscription plan during authorization.
	PlanId string `json:"plan_id,omitempty"`
	// The OpenID token belonging to this person. This token is only present if the OPENID scope is included in the authorization request.
	IdToken string `json:"id_token,omitempty"`
	// A refresh token.  For more information, see [Refresh, Revoke, and Limit the Scope of OAuth Tokens](https://developer.squareup.com/docs/oauth-api/refresh-revoke-limit-scope).
	RefreshToken string `json:"refresh_token,omitempty"`
	// A Boolean indicating that the access token is a short-lived access token. The short-lived access token returned in the response expires in 24 hours.
	ShortLived bool `json:"short_lived,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The date when the `refresh_token` expires, in [ISO 8601](http://www.iso.org/iso/home/standards/iso8601.htm) format.
	RefreshTokenExpiresAt string `json:"refresh_token_expires_at,omitempty"`
}
