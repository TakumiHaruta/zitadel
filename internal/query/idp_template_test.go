package query

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/zitadel/zitadel/internal/database"
	"github.com/zitadel/zitadel/internal/domain"
	errs "github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/repository/idp"
)

var (
	idpTemplateQuery = `SELECT projections.idp_templates4.id,` +
		` projections.idp_templates4.resource_owner,` +
		` projections.idp_templates4.creation_date,` +
		` projections.idp_templates4.change_date,` +
		` projections.idp_templates4.sequence,` +
		` projections.idp_templates4.state,` +
		` projections.idp_templates4.name,` +
		` projections.idp_templates4.type,` +
		` projections.idp_templates4.owner_type,` +
		` projections.idp_templates4.is_creation_allowed,` +
		` projections.idp_templates4.is_linking_allowed,` +
		` projections.idp_templates4.is_auto_creation,` +
		` projections.idp_templates4.is_auto_update,` +
		// oauth
		` projections.idp_templates4_oauth2.idp_id,` +
		` projections.idp_templates4_oauth2.client_id,` +
		` projections.idp_templates4_oauth2.client_secret,` +
		` projections.idp_templates4_oauth2.authorization_endpoint,` +
		` projections.idp_templates4_oauth2.token_endpoint,` +
		` projections.idp_templates4_oauth2.user_endpoint,` +
		` projections.idp_templates4_oauth2.scopes,` +
		` projections.idp_templates4_oauth2.id_attribute,` +
		// oidc
		` projections.idp_templates4_oidc.idp_id,` +
		` projections.idp_templates4_oidc.issuer,` +
		` projections.idp_templates4_oidc.client_id,` +
		` projections.idp_templates4_oidc.client_secret,` +
		` projections.idp_templates4_oidc.scopes,` +
		` projections.idp_templates4_oidc.id_token_mapping,` +
		// jwt
		` projections.idp_templates4_jwt.idp_id,` +
		` projections.idp_templates4_jwt.issuer,` +
		` projections.idp_templates4_jwt.jwt_endpoint,` +
		` projections.idp_templates4_jwt.keys_endpoint,` +
		` projections.idp_templates4_jwt.header_name,` +
		// azure
		` projections.idp_templates4_azure.idp_id,` +
		` projections.idp_templates4_azure.client_id,` +
		` projections.idp_templates4_azure.client_secret,` +
		` projections.idp_templates4_azure.scopes,` +
		` projections.idp_templates4_azure.tenant,` +
		` projections.idp_templates4_azure.is_email_verified,` +
		// github
		` projections.idp_templates4_github.idp_id,` +
		` projections.idp_templates4_github.client_id,` +
		` projections.idp_templates4_github.client_secret,` +
		` projections.idp_templates4_github.scopes,` +
		// github enterprise
		` projections.idp_templates4_github_enterprise.idp_id,` +
		` projections.idp_templates4_github_enterprise.client_id,` +
		` projections.idp_templates4_github_enterprise.client_secret,` +
		` projections.idp_templates4_github_enterprise.authorization_endpoint,` +
		` projections.idp_templates4_github_enterprise.token_endpoint,` +
		` projections.idp_templates4_github_enterprise.user_endpoint,` +
		` projections.idp_templates4_github_enterprise.scopes,` +
		// gitlab
		` projections.idp_templates4_gitlab.idp_id,` +
		` projections.idp_templates4_gitlab.client_id,` +
		` projections.idp_templates4_gitlab.client_secret,` +
		` projections.idp_templates4_gitlab.scopes,` +
		// gitlab self hosted
		` projections.idp_templates4_gitlab_self_hosted.idp_id,` +
		` projections.idp_templates4_gitlab_self_hosted.issuer,` +
		` projections.idp_templates4_gitlab_self_hosted.client_id,` +
		` projections.idp_templates4_gitlab_self_hosted.client_secret,` +
		` projections.idp_templates4_gitlab_self_hosted.scopes,` +
		// google
		` projections.idp_templates4_google.idp_id,` +
		` projections.idp_templates4_google.client_id,` +
		` projections.idp_templates4_google.client_secret,` +
		` projections.idp_templates4_google.scopes,` +
		// ldap
		` projections.idp_templates4_ldap.idp_id,` +
		` projections.idp_templates4_ldap.host,` +
		` projections.idp_templates4_ldap.port,` +
		` projections.idp_templates4_ldap.tls,` +
		` projections.idp_templates4_ldap.base_dn,` +
		` projections.idp_templates4_ldap.user_object_class,` +
		` projections.idp_templates4_ldap.user_unique_attribute,` +
		` projections.idp_templates4_ldap.admin,` +
		` projections.idp_templates4_ldap.password,` +
		` projections.idp_templates4_ldap.id_attribute,` +
		` projections.idp_templates4_ldap.first_name_attribute,` +
		` projections.idp_templates4_ldap.last_name_attribute,` +
		` projections.idp_templates4_ldap.display_name_attribute,` +
		` projections.idp_templates4_ldap.nick_name_attribute,` +
		` projections.idp_templates4_ldap.preferred_username_attribute,` +
		` projections.idp_templates4_ldap.email_attribute,` +
		` projections.idp_templates4_ldap.email_verified,` +
		` projections.idp_templates4_ldap.phone_attribute,` +
		` projections.idp_templates4_ldap.phone_verified_attribute,` +
		` projections.idp_templates4_ldap.preferred_language_attribute,` +
		` projections.idp_templates4_ldap.avatar_url_attribute,` +
		` projections.idp_templates4_ldap.profile_attribute` +
		` FROM projections.idp_templates4` +
		` LEFT JOIN projections.idp_templates4_oauth2 ON projections.idp_templates4.id = projections.idp_templates4_oauth2.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_oauth2.instance_id` +
		` LEFT JOIN projections.idp_templates4_oidc ON projections.idp_templates4.id = projections.idp_templates4_oidc.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_oidc.instance_id` +
		` LEFT JOIN projections.idp_templates4_jwt ON projections.idp_templates4.id = projections.idp_templates4_jwt.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_jwt.instance_id` +
		` LEFT JOIN projections.idp_templates4_azure ON projections.idp_templates4.id = projections.idp_templates4_azure.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_azure.instance_id` +
		` LEFT JOIN projections.idp_templates4_github ON projections.idp_templates4.id = projections.idp_templates4_github.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_github.instance_id` +
		` LEFT JOIN projections.idp_templates4_github_enterprise ON projections.idp_templates4.id = projections.idp_templates4_github_enterprise.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_github_enterprise.instance_id` +
		` LEFT JOIN projections.idp_templates4_gitlab ON projections.idp_templates4.id = projections.idp_templates4_gitlab.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_gitlab.instance_id` +
		` LEFT JOIN projections.idp_templates4_gitlab_self_hosted ON projections.idp_templates4.id = projections.idp_templates4_gitlab_self_hosted.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_gitlab_self_hosted.instance_id` +
		` LEFT JOIN projections.idp_templates4_google ON projections.idp_templates4.id = projections.idp_templates4_google.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_google.instance_id` +
		` LEFT JOIN projections.idp_templates4_ldap ON projections.idp_templates4.id = projections.idp_templates4_ldap.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_ldap.instance_id` +
		` AS OF SYSTEM TIME '-1 ms'`
	idpTemplateCols = []string{
		"id",
		"resource_owner",
		"creation_date",
		"change_date",
		"sequence",
		"state",
		"name",
		"type",
		"owner_type",
		"is_creation_allowed",
		"is_linking_allowed",
		"is_auto_creation",
		"is_auto_update",
		// oauth config
		"idp_id",
		"client_id",
		"client_secret",
		"authorization_endpoint",
		"token_endpoint",
		"user_endpoint",
		"scopes",
		"id_attribute",
		// oidc config
		"id_id",
		"issuer",
		"client_id",
		"client_secret",
		"scopes",
		"id_token_mapping",
		// jwt
		"idp_id",
		"issuer",
		"jwt_endpoint",
		"keys_endpoint",
		"header_name",
		// azure
		"idp_id",
		"client_id",
		"client_secret",
		"scopes",
		"tenant",
		"is_email_verified",
		// github config
		"idp_id",
		"client_id",
		"client_secret",
		"scopes",
		// github enterprise config
		"idp_id",
		"client_id",
		"client_secret",
		"authorization_endpoint",
		"token_endpoint",
		"user_endpoint",
		"scopes",
		// gitlab config
		"idp_id",
		"client_id",
		"client_secret",
		"scopes",
		// gitlab self hosted config
		"idp_id",
		"issuer",
		"client_id",
		"client_secret",
		"scopes",
		// google config
		"idp_id",
		"client_id",
		"client_secret",
		"scopes",
		// ldap config
		"idp_id",
		"host",
		"port",
		"tls",
		"base_dn",
		"user_object_class",
		"user_unique_attribute",
		"admin",
		"password",
		"id_attribute",
		"first_name_attribute",
		"last_name_attribute",
		"display_name_attribute",
		"nick_name_attribute",
		"preferred_username_attribute",
		"email_attribute",
		"email_verified",
		"phone_attribute",
		"phone_verified_attribute",
		"preferred_language_attribute",
		"avatar_url_attribute",
		"profile_attribute",
	}
	idpTemplatesQuery = `SELECT projections.idp_templates4.id,` +
		` projections.idp_templates4.resource_owner,` +
		` projections.idp_templates4.creation_date,` +
		` projections.idp_templates4.change_date,` +
		` projections.idp_templates4.sequence,` +
		` projections.idp_templates4.state,` +
		` projections.idp_templates4.name,` +
		` projections.idp_templates4.type,` +
		` projections.idp_templates4.owner_type,` +
		` projections.idp_templates4.is_creation_allowed,` +
		` projections.idp_templates4.is_linking_allowed,` +
		` projections.idp_templates4.is_auto_creation,` +
		` projections.idp_templates4.is_auto_update,` +
		// oauth
		` projections.idp_templates4_oauth2.idp_id,` +
		` projections.idp_templates4_oauth2.client_id,` +
		` projections.idp_templates4_oauth2.client_secret,` +
		` projections.idp_templates4_oauth2.authorization_endpoint,` +
		` projections.idp_templates4_oauth2.token_endpoint,` +
		` projections.idp_templates4_oauth2.user_endpoint,` +
		` projections.idp_templates4_oauth2.scopes,` +
		` projections.idp_templates4_oauth2.id_attribute,` +
		// oidc
		` projections.idp_templates4_oidc.idp_id,` +
		` projections.idp_templates4_oidc.issuer,` +
		` projections.idp_templates4_oidc.client_id,` +
		` projections.idp_templates4_oidc.client_secret,` +
		` projections.idp_templates4_oidc.scopes,` +
		` projections.idp_templates4_oidc.id_token_mapping,` +
		// jwt
		` projections.idp_templates4_jwt.idp_id,` +
		` projections.idp_templates4_jwt.issuer,` +
		` projections.idp_templates4_jwt.jwt_endpoint,` +
		` projections.idp_templates4_jwt.keys_endpoint,` +
		` projections.idp_templates4_jwt.header_name,` +
		// azure
		` projections.idp_templates4_azure.idp_id,` +
		` projections.idp_templates4_azure.client_id,` +
		` projections.idp_templates4_azure.client_secret,` +
		` projections.idp_templates4_azure.scopes,` +
		` projections.idp_templates4_azure.tenant,` +
		` projections.idp_templates4_azure.is_email_verified,` +
		// github
		` projections.idp_templates4_github.idp_id,` +
		` projections.idp_templates4_github.client_id,` +
		` projections.idp_templates4_github.client_secret,` +
		` projections.idp_templates4_github.scopes,` +
		// github enterprise
		` projections.idp_templates4_github_enterprise.idp_id,` +
		` projections.idp_templates4_github_enterprise.client_id,` +
		` projections.idp_templates4_github_enterprise.client_secret,` +
		` projections.idp_templates4_github_enterprise.authorization_endpoint,` +
		` projections.idp_templates4_github_enterprise.token_endpoint,` +
		` projections.idp_templates4_github_enterprise.user_endpoint,` +
		` projections.idp_templates4_github_enterprise.scopes,` +
		// gitlab
		` projections.idp_templates4_gitlab.idp_id,` +
		` projections.idp_templates4_gitlab.client_id,` +
		` projections.idp_templates4_gitlab.client_secret,` +
		` projections.idp_templates4_gitlab.scopes,` +
		// gitlab self hosted
		` projections.idp_templates4_gitlab_self_hosted.idp_id,` +
		` projections.idp_templates4_gitlab_self_hosted.issuer,` +
		` projections.idp_templates4_gitlab_self_hosted.client_id,` +
		` projections.idp_templates4_gitlab_self_hosted.client_secret,` +
		` projections.idp_templates4_gitlab_self_hosted.scopes,` +
		// google
		` projections.idp_templates4_google.idp_id,` +
		` projections.idp_templates4_google.client_id,` +
		` projections.idp_templates4_google.client_secret,` +
		` projections.idp_templates4_google.scopes,` +
		// ldap
		` projections.idp_templates4_ldap.idp_id,` +
		` projections.idp_templates4_ldap.host,` +
		` projections.idp_templates4_ldap.port,` +
		` projections.idp_templates4_ldap.tls,` +
		` projections.idp_templates4_ldap.base_dn,` +
		` projections.idp_templates4_ldap.user_object_class,` +
		` projections.idp_templates4_ldap.user_unique_attribute,` +
		` projections.idp_templates4_ldap.admin,` +
		` projections.idp_templates4_ldap.password,` +
		` projections.idp_templates4_ldap.id_attribute,` +
		` projections.idp_templates4_ldap.first_name_attribute,` +
		` projections.idp_templates4_ldap.last_name_attribute,` +
		` projections.idp_templates4_ldap.display_name_attribute,` +
		` projections.idp_templates4_ldap.nick_name_attribute,` +
		` projections.idp_templates4_ldap.preferred_username_attribute,` +
		` projections.idp_templates4_ldap.email_attribute,` +
		` projections.idp_templates4_ldap.email_verified,` +
		` projections.idp_templates4_ldap.phone_attribute,` +
		` projections.idp_templates4_ldap.phone_verified_attribute,` +
		` projections.idp_templates4_ldap.preferred_language_attribute,` +
		` projections.idp_templates4_ldap.avatar_url_attribute,` +
		` projections.idp_templates4_ldap.profile_attribute,` +
		` COUNT(*) OVER ()` +
		` FROM projections.idp_templates4` +
		` LEFT JOIN projections.idp_templates4_oauth2 ON projections.idp_templates4.id = projections.idp_templates4_oauth2.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_oauth2.instance_id` +
		` LEFT JOIN projections.idp_templates4_oidc ON projections.idp_templates4.id = projections.idp_templates4_oidc.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_oidc.instance_id` +
		` LEFT JOIN projections.idp_templates4_jwt ON projections.idp_templates4.id = projections.idp_templates4_jwt.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_jwt.instance_id` +
		` LEFT JOIN projections.idp_templates4_azure ON projections.idp_templates4.id = projections.idp_templates4_azure.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_azure.instance_id` +
		` LEFT JOIN projections.idp_templates4_github ON projections.idp_templates4.id = projections.idp_templates4_github.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_github.instance_id` +
		` LEFT JOIN projections.idp_templates4_github_enterprise ON projections.idp_templates4.id = projections.idp_templates4_github_enterprise.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_github_enterprise.instance_id` +
		` LEFT JOIN projections.idp_templates4_gitlab ON projections.idp_templates4.id = projections.idp_templates4_gitlab.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_gitlab.instance_id` +
		` LEFT JOIN projections.idp_templates4_gitlab_self_hosted ON projections.idp_templates4.id = projections.idp_templates4_gitlab_self_hosted.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_gitlab_self_hosted.instance_id` +
		` LEFT JOIN projections.idp_templates4_google ON projections.idp_templates4.id = projections.idp_templates4_google.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_google.instance_id` +
		` LEFT JOIN projections.idp_templates4_ldap ON projections.idp_templates4.id = projections.idp_templates4_ldap.idp_id AND projections.idp_templates4.instance_id = projections.idp_templates4_ldap.instance_id` +
		` AS OF SYSTEM TIME '-1 ms'`
	idpTemplatesCols = []string{
		"id",
		"resource_owner",
		"creation_date",
		"change_date",
		"sequence",
		"state",
		"name",
		"type",
		"owner_type",
		"is_creation_allowed",
		"is_linking_allowed",
		"is_auto_creation",
		"is_auto_update",
		// oauth config
		"idp_id",
		"client_id",
		"client_secret",
		"authorization_endpoint",
		"token_endpoint",
		"user_endpoint",
		"scopes",
		"id_attribute",
		// oidc config
		"id_id",
		"issuer",
		"client_id",
		"client_secret",
		"scopes",
		"id_token_mapping",
		// jwt
		"idp_id",
		"issuer",
		"jwt_endpoint",
		"keys_endpoint",
		"header_name",
		// azure
		"idp_id",
		"client_id",
		"client_secret",
		"scopes",
		"tenant",
		"is_email_verified",
		// github config
		"idp_id",
		"client_id",
		"client_secret",
		"scopes",
		// github enterprise config
		"idp_id",
		"client_id",
		"client_secret",
		"authorization_endpoint",
		"token_endpoint",
		"user_endpoint",
		"scopes",
		// gitlab config
		"idp_id",
		"client_id",
		"client_secret",
		"scopes",
		// gitlab self hosted config
		"idp_id",
		"issuer",
		"client_id",
		"client_secret",
		"scopes",
		// google config
		"idp_id",
		"client_id",
		"client_secret",
		"scopes",
		// ldap config
		"idp_id",
		"host",
		"port",
		"tls",
		"base_dn",
		"user_object_class",
		"user_unique_attribute",
		"admin",
		"password",
		"id_attribute",
		"first_name_attribute",
		"last_name_attribute",
		"display_name_attribute",
		"nick_name_attribute",
		"preferred_username_attribute",
		"email_attribute",
		"email_verified",
		"phone_attribute",
		"phone_verified_attribute",
		"preferred_language_attribute",
		"avatar_url_attribute",
		"profile_attribute",
		"count",
	}
)

func Test_IDPTemplateTemplatesPrepares(t *testing.T) {
	type want struct {
		sqlExpectations sqlExpectation
		err             checkErr
	}
	tests := []struct {
		name    string
		prepare interface{}
		want    want
		object  interface{}
	}{
		{
			name:    "prepareIDPTemplateByIDQuery no result",
			prepare: prepareIDPTemplateByIDQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(idpTemplateQuery),
					nil,
					nil,
				),
				err: func(err error) (error, bool) {
					if !errs.IsNotFound(err) {
						return fmt.Errorf("err should be zitadel.NotFoundError got: %w", err), false
					}
					return nil, true
				},
			},
			object: (*IDPTemplate)(nil),
		},
		{
			name:    "prepareIDPTemplateByIDQuery oauth idp",
			prepare: prepareIDPTemplateByIDQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(idpTemplateQuery),
					idpTemplateCols,
					[]driver.Value{
						"idp-id",
						"ro",
						testNow,
						testNow,
						uint64(20211109),
						domain.IDPConfigStateActive,
						"idp-name",
						domain.IDPTypeOAuth,
						domain.IdentityProviderTypeOrg,
						true,
						true,
						true,
						true,
						// oauth
						"idp-id",
						"client_id",
						nil,
						"authorization",
						"token",
						"user",
						database.StringArray{"profile"},
						"id-attribute",
						// oidc
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// jwt
						nil,
						nil,
						nil,
						nil,
						nil,
						// azure
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// github
						nil,
						nil,
						nil,
						nil,
						// github enterprise
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// gitlab
						nil,
						nil,
						nil,
						nil,
						// gitlab self hosted
						nil,
						nil,
						nil,
						nil,
						nil,
						// google
						nil,
						nil,
						nil,
						nil,
						// ldap config
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
					},
				),
			},
			object: &IDPTemplate{
				CreationDate:      testNow,
				ChangeDate:        testNow,
				Sequence:          20211109,
				ResourceOwner:     "ro",
				ID:                "idp-id",
				State:             domain.IDPStateActive,
				Name:              "idp-name",
				Type:              domain.IDPTypeOAuth,
				OwnerType:         domain.IdentityProviderTypeOrg,
				IsCreationAllowed: true,
				IsLinkingAllowed:  true,
				IsAutoCreation:    true,
				IsAutoUpdate:      true,
				OAuthIDPTemplate: &OAuthIDPTemplate{
					IDPID:                 "idp-id",
					ClientID:              "client_id",
					ClientSecret:          nil,
					AuthorizationEndpoint: "authorization",
					TokenEndpoint:         "token",
					UserEndpoint:          "user",
					Scopes:                []string{"profile"},
					IDAttribute:           "id-attribute",
				},
			},
		},
		{
			name:    "prepareIDPTemplateByIDQuery oidc idp",
			prepare: prepareIDPTemplateByIDQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(idpTemplateQuery),
					idpTemplateCols,
					[]driver.Value{
						"idp-id",
						"ro",
						testNow,
						testNow,
						uint64(20211109),
						domain.IDPConfigStateActive,
						"idp-name",
						domain.IDPTypeOIDC,
						domain.IdentityProviderTypeOrg,
						true,
						true,
						true,
						true,
						// oauth
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// oidc
						"idp-id",
						"issuer",
						"client_id",
						nil,
						database.StringArray{"profile"},
						true,
						// jwt
						nil,
						nil,
						nil,
						nil,
						nil,
						// azure
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// github
						nil,
						nil,
						nil,
						nil,
						// github enterprise
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// gitlab
						nil,
						nil,
						nil,
						nil,
						// gitlab self hosted
						nil,
						nil,
						nil,
						nil,
						nil,
						// google
						nil,
						nil,
						nil,
						nil,
						// ldap config
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
					},
				),
			},
			object: &IDPTemplate{
				CreationDate:      testNow,
				ChangeDate:        testNow,
				Sequence:          20211109,
				ResourceOwner:     "ro",
				ID:                "idp-id",
				State:             domain.IDPStateActive,
				Name:              "idp-name",
				Type:              domain.IDPTypeOIDC,
				OwnerType:         domain.IdentityProviderTypeOrg,
				IsCreationAllowed: true,
				IsLinkingAllowed:  true,
				IsAutoCreation:    true,
				IsAutoUpdate:      true,
				OIDCIDPTemplate: &OIDCIDPTemplate{
					IDPID:            "idp-id",
					Issuer:           "issuer",
					ClientID:         "client_id",
					ClientSecret:     nil,
					Scopes:           []string{"profile"},
					IsIDTokenMapping: true,
				},
			},
		},
		{
			name:    "prepareIDPTemplateByIDQuery jwt idp",
			prepare: prepareIDPTemplateByIDQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(idpTemplateQuery),
					idpTemplateCols,
					[]driver.Value{
						"idp-id",
						"ro",
						testNow,
						testNow,
						uint64(20211109),
						domain.IDPConfigStateActive,
						"idp-name",
						domain.IDPTypeJWT,
						domain.IdentityProviderTypeOrg,
						true,
						true,
						true,
						true,
						// oauth
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// oidc
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// jwt
						"idp-id",
						"issuer",
						"jwt",
						"keys",
						"header",
						// azure
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// github
						nil,
						nil,
						nil,
						nil,
						// github enterprise
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// gitlab
						nil,
						nil,
						nil,
						nil,
						// gitlab self hosted
						nil,
						nil,
						nil,
						nil,
						nil,
						// google
						nil,
						nil,
						nil,
						nil,
						// ldap config
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
					},
				),
			},
			object: &IDPTemplate{
				CreationDate:      testNow,
				ChangeDate:        testNow,
				Sequence:          20211109,
				ResourceOwner:     "ro",
				ID:                "idp-id",
				State:             domain.IDPStateActive,
				Name:              "idp-name",
				Type:              domain.IDPTypeJWT,
				OwnerType:         domain.IdentityProviderTypeOrg,
				IsCreationAllowed: true,
				IsLinkingAllowed:  true,
				IsAutoCreation:    true,
				IsAutoUpdate:      true,
				JWTIDPTemplate: &JWTIDPTemplate{
					IDPID:        "idp-id",
					Issuer:       "issuer",
					Endpoint:     "jwt",
					KeysEndpoint: "keys",
					HeaderName:   "header",
				},
			},
		},
		{
			name:    "prepareIDPTemplateByIDQuery github idp",
			prepare: prepareIDPTemplateByIDQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(idpTemplateQuery),
					idpTemplateCols,
					[]driver.Value{
						"idp-id",
						"ro",
						testNow,
						testNow,
						uint64(20211109),
						domain.IDPConfigStateActive,
						"idp-name",
						domain.IDPTypeGitHub,
						domain.IdentityProviderTypeOrg,
						true,
						true,
						true,
						true,
						// oauth
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// oidc
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// jwt
						nil,
						nil,
						nil,
						nil,
						nil,
						// azure
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// github
						"idp-id",
						"client_id",
						nil,
						database.StringArray{"profile"},
						// github enterprise
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// gitlab
						nil,
						nil,
						nil,
						nil,
						// gitlab self hosted
						nil,
						nil,
						nil,
						nil,
						nil,
						// google
						nil,
						nil,
						nil,
						nil,
						// ldap config
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
					},
				),
			},
			object: &IDPTemplate{
				CreationDate:      testNow,
				ChangeDate:        testNow,
				Sequence:          20211109,
				ResourceOwner:     "ro",
				ID:                "idp-id",
				State:             domain.IDPStateActive,
				Name:              "idp-name",
				Type:              domain.IDPTypeGitHub,
				OwnerType:         domain.IdentityProviderTypeOrg,
				IsCreationAllowed: true,
				IsLinkingAllowed:  true,
				IsAutoCreation:    true,
				IsAutoUpdate:      true,
				GitHubIDPTemplate: &GitHubIDPTemplate{
					IDPID:        "idp-id",
					ClientID:     "client_id",
					ClientSecret: nil,
					Scopes:       []string{"profile"},
				},
			},
		},
		{
			name:    "prepareIDPTemplateByIDQuery gitlab idp",
			prepare: prepareIDPTemplateByIDQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(idpTemplateQuery),
					idpTemplateCols,
					[]driver.Value{
						"idp-id",
						"ro",
						testNow,
						testNow,
						uint64(20211109),
						domain.IDPConfigStateActive,
						"idp-name",
						domain.IDPTypeGitLab,
						domain.IdentityProviderTypeOrg,
						true,
						true,
						true,
						true,
						// oauth
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// oidc
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// jwt
						nil,
						nil,
						nil,
						nil,
						nil,
						// azure
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// github
						nil,
						nil,
						nil,
						nil,
						// github enterprise
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// gitlab
						"idp-id",
						"client_id",
						nil,
						database.StringArray{"profile"},
						// gitlab self hosted
						nil,
						nil,
						nil,
						nil,
						nil,
						// google
						nil,
						nil,
						nil,
						nil,
						// ldap config
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
					},
				),
			},
			object: &IDPTemplate{
				CreationDate:      testNow,
				ChangeDate:        testNow,
				Sequence:          20211109,
				ResourceOwner:     "ro",
				ID:                "idp-id",
				State:             domain.IDPStateActive,
				Name:              "idp-name",
				Type:              domain.IDPTypeGitLab,
				OwnerType:         domain.IdentityProviderTypeOrg,
				IsCreationAllowed: true,
				IsLinkingAllowed:  true,
				IsAutoCreation:    true,
				IsAutoUpdate:      true,
				GitLabIDPTemplate: &GitLabIDPTemplate{
					IDPID:        "idp-id",
					ClientID:     "client_id",
					ClientSecret: nil,
					Scopes:       []string{"profile"},
				},
			},
		},
		{
			name:    "prepareIDPTemplateByIDQuery gitlab self hosted idp",
			prepare: prepareIDPTemplateByIDQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(idpTemplateQuery),
					idpTemplateCols,
					[]driver.Value{
						"idp-id",
						"ro",
						testNow,
						testNow,
						uint64(20211109),
						domain.IDPConfigStateActive,
						"idp-name",
						domain.IDPTypeGitLabSelfHosted,
						domain.IdentityProviderTypeOrg,
						true,
						true,
						true,
						true,
						// oauth
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// oidc
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// jwt
						nil,
						nil,
						nil,
						nil,
						nil,
						// azure
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// github
						nil,
						nil,
						nil,
						nil,
						// github enterprise
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// gitlab
						nil,
						nil,
						nil,
						nil,
						// gitlab self hosted
						"idp-id",
						"issuer",
						"client_id",
						nil,
						database.StringArray{"profile"},
						// google
						nil,
						nil,
						nil,
						nil,
						// ldap config
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
					},
				),
			},
			object: &IDPTemplate{
				CreationDate:      testNow,
				ChangeDate:        testNow,
				Sequence:          20211109,
				ResourceOwner:     "ro",
				ID:                "idp-id",
				State:             domain.IDPStateActive,
				Name:              "idp-name",
				Type:              domain.IDPTypeGitLabSelfHosted,
				OwnerType:         domain.IdentityProviderTypeOrg,
				IsCreationAllowed: true,
				IsLinkingAllowed:  true,
				IsAutoCreation:    true,
				IsAutoUpdate:      true,
				GitLabSelfHostedIDPTemplate: &GitLabSelfHostedIDPTemplate{
					IDPID:        "idp-id",
					Issuer:       "issuer",
					ClientID:     "client_id",
					ClientSecret: nil,
					Scopes:       []string{"profile"},
				},
			},
		},
		{
			name:    "prepareIDPTemplateByIDQuery google idp",
			prepare: prepareIDPTemplateByIDQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(idpTemplateQuery),
					idpTemplateCols,
					[]driver.Value{
						"idp-id",
						"ro",
						testNow,
						testNow,
						uint64(20211109),
						domain.IDPConfigStateActive,
						"idp-name",
						domain.IDPTypeGoogle,
						domain.IdentityProviderTypeOrg,
						true,
						true,
						true,
						true,
						// oauth
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// oidc
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// jwt
						nil,
						nil,
						nil,
						nil,
						nil,
						// azure
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// github
						nil,
						nil,
						nil,
						nil,
						// github enterprise
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// gitlab
						nil,
						nil,
						nil,
						nil,
						// gitlab self hosted
						nil,
						nil,
						nil,
						nil,
						nil,
						// google
						"idp-id",
						"client_id",
						nil,
						database.StringArray{"profile"},
						// ldap config
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
					},
				),
			},
			object: &IDPTemplate{
				CreationDate:      testNow,
				ChangeDate:        testNow,
				Sequence:          20211109,
				ResourceOwner:     "ro",
				ID:                "idp-id",
				State:             domain.IDPStateActive,
				Name:              "idp-name",
				Type:              domain.IDPTypeGoogle,
				OwnerType:         domain.IdentityProviderTypeOrg,
				IsCreationAllowed: true,
				IsLinkingAllowed:  true,
				IsAutoCreation:    true,
				IsAutoUpdate:      true,
				GoogleIDPTemplate: &GoogleIDPTemplate{
					IDPID:        "idp-id",
					ClientID:     "client_id",
					ClientSecret: nil,
					Scopes:       []string{"profile"},
				},
			},
		},
		{
			name:    "prepareIDPTemplateByIDQuery ldap idp",
			prepare: prepareIDPTemplateByIDQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(idpTemplateQuery),
					idpTemplateCols,
					[]driver.Value{
						"idp-id",
						"ro",
						testNow,
						testNow,
						uint64(20211109),
						domain.IDPConfigStateActive,
						"idp-name",
						domain.IDPTypeLDAP,
						domain.IdentityProviderTypeOrg,
						true,
						true,
						true,
						true,
						// oauth
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// oidc
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// jwt
						nil,
						nil,
						nil,
						nil,
						nil,
						// azure
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// github
						nil,
						nil,
						nil,
						nil,
						// github enterprise
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// gitlab
						nil,
						nil,
						nil,
						nil,
						// gitlab self hosted
						nil,
						nil,
						nil,
						nil,
						nil,
						// google
						nil,
						nil,
						nil,
						nil,
						// ldap config
						"idp-id",
						"host",
						"port",
						true,
						"base",
						"user",
						"uid",
						"admin",
						nil,
						"id",
						"first",
						"last",
						"display",
						"nickname",
						"username",
						"email",
						"emailVerified",
						"phone",
						"phoneVerified",
						"lang",
						"avatar",
						"profile",
					},
				),
			},
			object: &IDPTemplate{
				CreationDate:      testNow,
				ChangeDate:        testNow,
				Sequence:          20211109,
				ResourceOwner:     "ro",
				ID:                "idp-id",
				State:             domain.IDPStateActive,
				Name:              "idp-name",
				Type:              domain.IDPTypeLDAP,
				OwnerType:         domain.IdentityProviderTypeOrg,
				IsCreationAllowed: true,
				IsLinkingAllowed:  true,
				IsAutoCreation:    true,
				IsAutoUpdate:      true,
				LDAPIDPTemplate: &LDAPIDPTemplate{
					IDPID:               "idp-id",
					Host:                "host",
					Port:                "port",
					TLS:                 true,
					BaseDN:              "base",
					UserObjectClass:     "user",
					UserUniqueAttribute: "uid",
					Admin:               "admin",
					LDAPAttributes: idp.LDAPAttributes{
						IDAttribute:                "id",
						FirstNameAttribute:         "first",
						LastNameAttribute:          "last",
						DisplayNameAttribute:       "display",
						NickNameAttribute:          "nickname",
						PreferredUsernameAttribute: "username",
						EmailAttribute:             "email",
						EmailVerifiedAttribute:     "emailVerified",
						PhoneAttribute:             "phone",
						PhoneVerifiedAttribute:     "phoneVerified",
						PreferredLanguageAttribute: "lang",
						AvatarURLAttribute:         "avatar",
						ProfileAttribute:           "profile",
					},
				},
			},
		},
		{
			name:    "prepareIDPTemplateByIDQuery no config",
			prepare: prepareIDPTemplateByIDQuery,
			want: want{
				sqlExpectations: mockQuery(
					regexp.QuoteMeta(idpTemplateQuery),
					idpTemplateCols,
					[]driver.Value{
						"idp-id",
						"ro",
						testNow,
						testNow,
						uint64(20211109),
						domain.IDPConfigStateActive,
						"idp-name",
						domain.IDPTypeLDAP,
						domain.IdentityProviderTypeOrg,
						true,
						true,
						true,
						true,
						// oauth
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// oidc
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// jwt
						nil,
						nil,
						nil,
						nil,
						nil,
						// azure
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// github
						nil,
						nil,
						nil,
						nil,
						// github enterprise
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						// gitlab
						nil,
						nil,
						nil,
						nil,
						// gitlab self hosted
						nil,
						nil,
						nil,
						nil,
						nil,
						// google config
						nil,
						nil,
						nil,
						nil,
						// ldap config
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
					},
				),
			},
			object: &IDPTemplate{
				CreationDate:      testNow,
				ChangeDate:        testNow,
				Sequence:          20211109,
				ResourceOwner:     "ro",
				ID:                "idp-id",
				State:             domain.IDPStateActive,
				Name:              "idp-name",
				Type:              domain.IDPTypeLDAP,
				OwnerType:         domain.IdentityProviderTypeOrg,
				IsCreationAllowed: true,
				IsLinkingAllowed:  true,
				IsAutoCreation:    true,
				IsAutoUpdate:      true,
			},
		},
		{
			name:    "prepareIDPTemplateByIDQuery sql err",
			prepare: prepareIDPTemplateByIDQuery,
			want: want{
				sqlExpectations: mockQueryErr(
					regexp.QuoteMeta(idpTemplateQuery),
					sql.ErrConnDone,
				),
				err: func(err error) (error, bool) {
					if !errors.Is(err, sql.ErrConnDone) {
						return fmt.Errorf("err should be sql.ErrConnDone got: %w", err), false
					}
					return nil, true
				},
			},
			object: nil,
		},
		{
			name:    "prepareIDPTemplatesQuery no result",
			prepare: prepareIDPTemplatesQuery,
			want: want{
				sqlExpectations: mockQueries(
					regexp.QuoteMeta(idpTemplatesQuery),
					nil,
					nil,
				),
				err: func(err error) (error, bool) {
					if !errs.IsNotFound(err) {
						return fmt.Errorf("err should be zitadel.NotFoundError got: %w", err), false
					}
					return nil, true
				},
			},
			object: &IDPTemplates{Templates: []*IDPTemplate{}},
		},
		{
			name:    "prepareIDPTemplatesQuery ldap idp",
			prepare: prepareIDPTemplatesQuery,
			want: want{
				sqlExpectations: mockQueries(
					regexp.QuoteMeta(idpTemplatesQuery),
					idpTemplatesCols,
					[][]driver.Value{
						{
							"idp-id",
							"ro",
							testNow,
							testNow,
							uint64(20211109),
							domain.IDPConfigStateActive,
							"idp-name",
							domain.IDPTypeLDAP,
							domain.IdentityProviderTypeOrg,
							true,
							true,
							true,
							true,
							// oauth
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// oidc
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// jwt
							nil,
							nil,
							nil,
							nil,
							nil,
							// azure
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// github
							nil,
							nil,
							nil,
							nil,
							// github enterprise
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// gitlab
							nil,
							nil,
							nil,
							nil,
							// gitlab self hosted
							nil,
							nil,
							nil,
							nil,
							nil,
							// google config
							nil,
							nil,
							nil,
							nil,
							// ldap config
							"idp-id",
							"host",
							"port",
							true,
							"base",
							"user",
							"uid",
							"admin",
							nil,
							"id",
							"first",
							"last",
							"display",
							"nickname",
							"username",
							"email",
							"emailVerified",
							"phone",
							"phoneVerified",
							"lang",
							"avatar",
							"profile",
						},
					},
				),
			},
			object: &IDPTemplates{
				SearchResponse: SearchResponse{
					Count: 1,
				},
				Templates: []*IDPTemplate{
					{
						CreationDate:      testNow,
						ChangeDate:        testNow,
						Sequence:          20211109,
						ResourceOwner:     "ro",
						ID:                "idp-id",
						State:             domain.IDPStateActive,
						Name:              "idp-name",
						Type:              domain.IDPTypeLDAP,
						OwnerType:         domain.IdentityProviderTypeOrg,
						IsCreationAllowed: true,
						IsLinkingAllowed:  true,
						IsAutoCreation:    true,
						IsAutoUpdate:      true,
						LDAPIDPTemplate: &LDAPIDPTemplate{
							IDPID:               "idp-id",
							Host:                "host",
							Port:                "port",
							TLS:                 true,
							BaseDN:              "base",
							UserObjectClass:     "user",
							UserUniqueAttribute: "uid",
							Admin:               "admin",
							LDAPAttributes: idp.LDAPAttributes{
								IDAttribute:                "id",
								FirstNameAttribute:         "first",
								LastNameAttribute:          "last",
								DisplayNameAttribute:       "display",
								NickNameAttribute:          "nickname",
								PreferredUsernameAttribute: "username",
								EmailAttribute:             "email",
								EmailVerifiedAttribute:     "emailVerified",
								PhoneAttribute:             "phone",
								PhoneVerifiedAttribute:     "phoneVerified",
								PreferredLanguageAttribute: "lang",
								AvatarURLAttribute:         "avatar",
								ProfileAttribute:           "profile",
							},
						},
					},
				},
			},
		},
		{
			name:    "prepareIDPTemplatesQuery no config",
			prepare: prepareIDPTemplatesQuery,
			want: want{
				sqlExpectations: mockQueries(
					regexp.QuoteMeta(idpTemplatesQuery),
					idpTemplatesCols,
					[][]driver.Value{
						{
							"idp-id",
							"ro",
							testNow,
							testNow,
							uint64(20211109),
							domain.IDPConfigStateActive,
							"idp-name",
							domain.IDPTypeLDAP,
							domain.IdentityProviderTypeOrg,
							true,
							true,
							true,
							true,
							// oauth
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// oidc
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// jwt
							nil,
							nil,
							nil,
							nil,
							nil,
							// azure
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// github
							nil,
							nil,
							nil,
							nil,
							// github enterprise
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// gitlab
							nil,
							nil,
							nil,
							nil,
							// gitlab self hosted
							nil,
							nil,
							nil,
							nil,
							nil,
							// google config
							nil,
							nil,
							nil,
							nil,
							// ldap config
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
						},
					},
				),
			},
			object: &IDPTemplates{
				SearchResponse: SearchResponse{
					Count: 1,
				},
				Templates: []*IDPTemplate{
					{
						CreationDate:      testNow,
						ChangeDate:        testNow,
						Sequence:          20211109,
						ResourceOwner:     "ro",
						ID:                "idp-id",
						State:             domain.IDPStateActive,
						Name:              "idp-name",
						Type:              domain.IDPTypeLDAP,
						OwnerType:         domain.IdentityProviderTypeOrg,
						IsCreationAllowed: true,
						IsLinkingAllowed:  true,
						IsAutoCreation:    true,
						IsAutoUpdate:      true,
					},
				},
			},
		},
		{
			name:    "prepareIDPTemplatesQuery all config types",
			prepare: prepareIDPTemplatesQuery,
			want: want{
				sqlExpectations: mockQueries(
					regexp.QuoteMeta(idpTemplatesQuery),
					idpTemplatesCols,
					[][]driver.Value{
						{
							"idp-id-ldap",
							"ro",
							testNow,
							testNow,
							uint64(20211109),
							domain.IDPConfigStateActive,
							"idp-name",
							domain.IDPTypeLDAP,
							domain.IdentityProviderTypeOrg,
							true,
							true,
							true,
							true,
							// oauth
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// oidc
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// jwt
							nil,
							nil,
							nil,
							nil,
							nil,
							// azure
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// github
							nil,
							nil,
							nil,
							nil,
							// github enterprise
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// gitlab
							nil,
							nil,
							nil,
							nil,
							// gitlab self hosted
							nil,
							nil,
							nil,
							nil,
							nil,
							// google config
							nil,
							nil,
							nil,
							nil,
							// ldap config
							"idp-id-ldap",
							"host",
							"port",
							true,
							"base",
							"user",
							"uid",
							"admin",
							nil,
							"id",
							"first",
							"last",
							"display",
							"nickname",
							"username",
							"email",
							"emailVerified",
							"phone",
							"phoneVerified",
							"lang",
							"avatar",
							"profile",
						},
						{
							"idp-id-google",
							"ro",
							testNow,
							testNow,
							uint64(20211109),
							domain.IDPConfigStateActive,
							"idp-name",
							domain.IDPTypeGoogle,
							domain.IdentityProviderTypeOrg,
							true,
							true,
							true,
							true,
							// oauth
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// oidc
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// jwt
							nil,
							nil,
							nil,
							nil,
							nil,
							// azure
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// github
							nil,
							nil,
							nil,
							nil,
							// github enterprise
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// gitlab
							nil,
							nil,
							nil,
							nil,
							// gitlab self hosted
							nil,
							nil,
							nil,
							nil,
							nil,
							// google
							"idp-id-google",
							"client_id",
							nil,
							database.StringArray{"profile"},
							// ldap config
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
						},
						{
							"idp-id-oauth",
							"ro",
							testNow,
							testNow,
							uint64(20211109),
							domain.IDPConfigStateActive,
							"idp-name",
							domain.IDPTypeOAuth,
							domain.IdentityProviderTypeOrg,
							true,
							true,
							true,
							true,
							// oauth
							"idp-id-oauth",
							"client_id",
							nil,
							"authorization",
							"token",
							"user",
							database.StringArray{"profile"},
							"id-attribute",
							// oidc
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// jwt
							nil,
							nil,
							nil,
							nil,
							nil,
							// azure
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// github
							nil,
							nil,
							nil,
							nil,
							// github enterprise
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// gitlab
							nil,
							nil,
							nil,
							nil,
							// gitlab self hosted
							nil,
							nil,
							nil,
							nil,
							nil,
							// google
							nil,
							nil,
							nil,
							nil,
							// ldap config
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
						},
						{
							"idp-id-oidc",
							"ro",
							testNow,
							testNow,
							uint64(20211109),
							domain.IDPConfigStateActive,
							"idp-name",
							domain.IDPTypeOIDC,
							domain.IdentityProviderTypeOrg,
							true,
							true,
							true,
							true,
							// oauth
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// oidc
							"idp-id-oidc",
							"issuer",
							"client_id",
							nil,
							database.StringArray{"profile"},
							true,
							// jwt
							nil,
							nil,
							nil,
							nil,
							nil,
							// azure
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// github
							nil,
							nil,
							nil,
							nil,
							// github enterprise
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// gitlab
							nil,
							nil,
							nil,
							nil,
							// gitlab self hosted
							nil,
							nil,
							nil,
							nil,
							nil,
							// google
							nil,
							nil,
							nil,
							nil,
							// ldap config
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
						},
						{
							"idp-id-jwt",
							"ro",
							testNow,
							testNow,
							uint64(20211109),
							domain.IDPConfigStateActive,
							"idp-name",
							domain.IDPTypeJWT,
							domain.IdentityProviderTypeOrg,
							true,
							true,
							true,
							true,
							// oauth
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// oidc
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// jwt
							"idp-id-jwt",
							"issuer",
							"jwt",
							"keys",
							"header",
							// azure
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// github
							nil,
							nil,
							nil,
							nil,
							// github enterprise
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							// gitlab
							nil,
							nil,
							nil,
							nil,
							// gitlab self hosted
							nil,
							nil,
							nil,
							nil,
							nil,
							// google
							nil,
							nil,
							nil,
							nil,
							// ldap config
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
							nil,
						},
					},
				),
			},
			object: &IDPTemplates{
				SearchResponse: SearchResponse{
					Count: 5,
				},
				Templates: []*IDPTemplate{
					{
						CreationDate:      testNow,
						ChangeDate:        testNow,
						Sequence:          20211109,
						ResourceOwner:     "ro",
						ID:                "idp-id-ldap",
						State:             domain.IDPStateActive,
						Name:              "idp-name",
						Type:              domain.IDPTypeLDAP,
						OwnerType:         domain.IdentityProviderTypeOrg,
						IsCreationAllowed: true,
						IsLinkingAllowed:  true,
						IsAutoCreation:    true,
						IsAutoUpdate:      true,
						LDAPIDPTemplate: &LDAPIDPTemplate{
							IDPID:               "idp-id-ldap",
							Host:                "host",
							Port:                "port",
							TLS:                 true,
							BaseDN:              "base",
							UserObjectClass:     "user",
							UserUniqueAttribute: "uid",
							Admin:               "admin",
							LDAPAttributes: idp.LDAPAttributes{
								IDAttribute:                "id",
								FirstNameAttribute:         "first",
								LastNameAttribute:          "last",
								DisplayNameAttribute:       "display",
								NickNameAttribute:          "nickname",
								PreferredUsernameAttribute: "username",
								EmailAttribute:             "email",
								EmailVerifiedAttribute:     "emailVerified",
								PhoneAttribute:             "phone",
								PhoneVerifiedAttribute:     "phoneVerified",
								PreferredLanguageAttribute: "lang",
								AvatarURLAttribute:         "avatar",
								ProfileAttribute:           "profile",
							},
						},
					},
					{
						CreationDate:      testNow,
						ChangeDate:        testNow,
						Sequence:          20211109,
						ResourceOwner:     "ro",
						ID:                "idp-id-google",
						State:             domain.IDPStateActive,
						Name:              "idp-name",
						Type:              domain.IDPTypeGoogle,
						OwnerType:         domain.IdentityProviderTypeOrg,
						IsCreationAllowed: true,
						IsLinkingAllowed:  true,
						IsAutoCreation:    true,
						IsAutoUpdate:      true,
						GoogleIDPTemplate: &GoogleIDPTemplate{
							IDPID:        "idp-id-google",
							ClientID:     "client_id",
							ClientSecret: nil,
							Scopes:       []string{"profile"},
						},
					},

					{
						CreationDate:      testNow,
						ChangeDate:        testNow,
						Sequence:          20211109,
						ResourceOwner:     "ro",
						ID:                "idp-id-oauth",
						State:             domain.IDPStateActive,
						Name:              "idp-name",
						Type:              domain.IDPTypeOAuth,
						OwnerType:         domain.IdentityProviderTypeOrg,
						IsCreationAllowed: true,
						IsLinkingAllowed:  true,
						IsAutoCreation:    true,
						IsAutoUpdate:      true,
						OAuthIDPTemplate: &OAuthIDPTemplate{
							IDPID:                 "idp-id-oauth",
							ClientID:              "client_id",
							ClientSecret:          nil,
							AuthorizationEndpoint: "authorization",
							TokenEndpoint:         "token",
							UserEndpoint:          "user",
							Scopes:                []string{"profile"},
							IDAttribute:           "id-attribute",
						},
					},
					{
						CreationDate:      testNow,
						ChangeDate:        testNow,
						Sequence:          20211109,
						ResourceOwner:     "ro",
						ID:                "idp-id-oidc",
						State:             domain.IDPStateActive,
						Name:              "idp-name",
						Type:              domain.IDPTypeOIDC,
						OwnerType:         domain.IdentityProviderTypeOrg,
						IsCreationAllowed: true,
						IsLinkingAllowed:  true,
						IsAutoCreation:    true,
						IsAutoUpdate:      true,
						OIDCIDPTemplate: &OIDCIDPTemplate{
							IDPID:            "idp-id-oidc",
							Issuer:           "issuer",
							ClientID:         "client_id",
							ClientSecret:     nil,
							Scopes:           []string{"profile"},
							IsIDTokenMapping: true,
						},
					},
					{
						CreationDate:      testNow,
						ChangeDate:        testNow,
						Sequence:          20211109,
						ResourceOwner:     "ro",
						ID:                "idp-id-jwt",
						State:             domain.IDPStateActive,
						Name:              "idp-name",
						Type:              domain.IDPTypeJWT,
						OwnerType:         domain.IdentityProviderTypeOrg,
						IsCreationAllowed: true,
						IsLinkingAllowed:  true,
						IsAutoCreation:    true,
						IsAutoUpdate:      true,
						JWTIDPTemplate: &JWTIDPTemplate{
							IDPID:        "idp-id-jwt",
							Issuer:       "issuer",
							Endpoint:     "jwt",
							KeysEndpoint: "keys",
							HeaderName:   "header",
						},
					},
				},
			},
		},
		{
			name:    "prepareIDPTemplatesQuery sql err",
			prepare: prepareIDPTemplatesQuery,
			want: want{
				sqlExpectations: mockQueryErr(
					regexp.QuoteMeta(idpTemplatesQuery),
					sql.ErrConnDone,
				),
				err: func(err error) (error, bool) {
					if !errors.Is(err, sql.ErrConnDone) {
						return fmt.Errorf("err should be sql.ErrConnDone got: %w", err), false
					}
					return nil, true
				},
			},
			object: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertPrepare(t, tt.prepare, tt.object, tt.want.sqlExpectations, tt.want.err, defaultPrepareArgs...)
		})
	}
}
