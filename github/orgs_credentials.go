package github

import (
	"context"
	"fmt"
)

// SSOAuthorization SSO authorization credential
type SSOAuthorization struct {
	Login                         *string    `json:"login,omitempty"`
	CredentialID                  *int64     `json:"credential_id,omitempty"`
	CredentialType                *string    `json:"credential_type,omitempty"`
	CredentialAuthorizedAt        *Timestamp `json:"credential_authorized_at,omitempty"`
	CredentialAccessedAt          *Timestamp `json:"credential_accessed_at,omitempty"`
	AuthorizedCredentialID        *int64     `json:"authorized_credential_id,omitempty"`
	TokenLastEight                *string    `json:"token_last_eight,omitempty"`
	Scopes                        []string   `json:"scopes,omitempty"`
	AuthorizedCredentialNote      *string    `json:"authorized_credential_note,omitempty"`
	AuthorizedCredentialExpiresAt *Timestamp `json:"authorized_credential_expires_at,omitempty"`
	Fingerprint                   *string    `json:"fingerprint,omitempty"`
	AuthorizedCredentialTitle     *string    `json:"authorized_credential_title,omitempty"`
}

// ListSSOAuthorizations lists SSO authoirized credentials for the organization
//
// GitHub API docs: https://docs.github.com/en/rest/reference/orgs#list-saml-sso-authorizations-for-an-organization
func (s *OrganizationsService) ListSSOAuthorizations(ctx context.Context, org string, opts *ListCursorOptions) ([]*SSOAuthorization, *Response, error) {
	u := fmt.Sprintf("orgs/%v/credential-authorizations", org)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	authorizations := []*SSOAuthorization{}
	resp, err := s.client.Do(ctx, req, &authorizations)
	if err != nil {
		return nil, resp, err
	}

	return authorizations, resp, nil
}
