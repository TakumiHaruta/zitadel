package command

import (
	"context"
	"strings"

	"github.com/zitadel/zitadel/internal/api/authz"
	"github.com/zitadel/zitadel/internal/command/preparation"
	"github.com/zitadel/zitadel/internal/crypto"
	"github.com/zitadel/zitadel/internal/domain"
	caos_errs "github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/repository/instance"
)

func (c *Commands) AddInstanceGenericOAuthProvider(ctx context.Context, provider GenericOAuthProvider) (string, *domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	id, err := c.idGenerator.Next()
	if err != nil {
		return "", nil, err
	}
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.prepareAddInstanceOAuthProvider(instanceAgg, id, provider))
	if err != nil {
		return "", nil, err
	}
	pushedEvents, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return "", nil, err
	}
	return id, pushedEventsToObjectDetails(pushedEvents), nil
}

func (c *Commands) UpdateInstanceGenericOAuthProvider(ctx context.Context, id string, provider GenericOAuthProvider) (*domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.prepareUpdateInstanceOAuthProvider(instanceAgg, id, provider))
	if err != nil {
		return nil, err
	}
	if len(cmds) == 0 {
		// no change, so return directly
		return &domain.ObjectDetails{}, nil
	}
	pushedEvents, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return nil, err
	}
	return pushedEventsToObjectDetails(pushedEvents), nil
}

func (c *Commands) AddInstanceGitHubProvider(ctx context.Context, provider GitHubProvider) (string, *domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	id, err := c.idGenerator.Next()
	if err != nil {
		return "", nil, err
	}
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.prepareAddInstanceGitHubProvider(instanceAgg, id, provider))
	if err != nil {
		return "", nil, err
	}
	pushedEvents, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return "", nil, err
	}
	return id, pushedEventsToObjectDetails(pushedEvents), nil
}

func (c *Commands) UpdateInstanceGitHubProvider(ctx context.Context, id string, provider GitHubProvider) (*domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.prepareUpdateInstanceGitHubProvider(instanceAgg, id, provider))
	if err != nil {
		return nil, err
	}
	if len(cmds) == 0 {
		// no change, so return directly
		return &domain.ObjectDetails{}, nil
	}
	pushedEvents, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return nil, err
	}
	return pushedEventsToObjectDetails(pushedEvents), nil
}

func (c *Commands) AddInstanceGitHubEnterpriseProvider(ctx context.Context, provider GitHubEnterpriseProvider) (string, *domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	id, err := c.idGenerator.Next()
	if err != nil {
		return "", nil, err
	}
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.prepareAddInstanceGitHubEnterpriseProvider(instanceAgg, id, provider))
	if err != nil {
		return "", nil, err
	}
	pushedEvents, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return "", nil, err
	}
	return id, pushedEventsToObjectDetails(pushedEvents), nil
}

func (c *Commands) UpdateInstanceGitHubEnterpriseProvider(ctx context.Context, id string, provider GitHubEnterpriseProvider) (*domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.prepareUpdateInstanceGitHubEnterpriseProvider(instanceAgg, id, provider))
	if err != nil {
		return nil, err
	}
	if len(cmds) == 0 {
		// no change, so return directly
		return &domain.ObjectDetails{}, nil
	}
	pushedEvents, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return nil, err
	}
	return pushedEventsToObjectDetails(pushedEvents), nil
}

func (c *Commands) AddInstanceGoogleProvider(ctx context.Context, provider GoogleProvider) (string, *domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	id, err := c.idGenerator.Next()
	if err != nil {
		return "", nil, err
	}
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.prepareAddInstanceGoogleProvider(instanceAgg, id, provider))
	if err != nil {
		return "", nil, err
	}
	pushedEvents, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return "", nil, err
	}
	return id, pushedEventsToObjectDetails(pushedEvents), nil
}

func (c *Commands) UpdateInstanceGoogleProvider(ctx context.Context, id string, provider GoogleProvider) (*domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.prepareUpdateInstanceGoogleProvider(instanceAgg, id, provider))
	if err != nil {
		return nil, err
	}
	if len(cmds) == 0 {
		// no change, so return directly
		return &domain.ObjectDetails{}, nil
	}
	pushedEvents, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return nil, err
	}
	return pushedEventsToObjectDetails(pushedEvents), nil
}

func (c *Commands) AddInstanceLDAPProvider(ctx context.Context, provider LDAPProvider) (string, *domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	id, err := c.idGenerator.Next()
	if err != nil {
		return "", nil, err
	}
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.prepareAddInstanceLDAPProvider(instanceAgg, id, provider))
	if err != nil {
		return "", nil, err
	}
	pushedEvents, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return "", nil, err
	}
	return id, pushedEventsToObjectDetails(pushedEvents), nil
}

func (c *Commands) UpdateInstanceLDAPProvider(ctx context.Context, id string, provider LDAPProvider) (*domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.prepareUpdateInstanceLDAPProvider(instanceAgg, id, provider))
	if err != nil {
		return nil, err
	}
	if len(cmds) == 0 {
		// no change, so return directly
		return &domain.ObjectDetails{}, nil
	}
	pushedEvents, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return nil, err
	}
	return pushedEventsToObjectDetails(pushedEvents), nil
}

func (c *Commands) DeleteInstanceProvider(ctx context.Context, id string) (*domain.ObjectDetails, error) {
	instanceAgg := instance.NewAggregate(authz.GetInstance(ctx).InstanceID())
	cmds, err := preparation.PrepareCommands(ctx, c.eventstore.Filter, c.prepareDeleteInstanceProvider(instanceAgg, id))
	if err != nil {
		return nil, err
	}
	pushedEvents, err := c.eventstore.Push(ctx, cmds...)
	if err != nil {
		return nil, err
	}
	return pushedEventsToObjectDetails(pushedEvents), nil
}

func (c *Commands) prepareAddInstanceOAuthProvider(a *instance.Aggregate, id string, provider GenericOAuthProvider) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		if provider.Name = strings.TrimSpace(provider.Name); provider.Name == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-D32ef", "Errors.Invalid.Argument")
		}
		if provider.ClientID = strings.TrimSpace(provider.ClientID); provider.ClientID == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-Dbgzf", "Errors.Invalid.Argument")
		}
		if provider.ClientSecret = strings.TrimSpace(provider.ClientSecret); provider.ClientSecret == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-DF4ga", "Errors.Invalid.Argument")
		}
		if provider.AuthorizationEndpoint = strings.TrimSpace(provider.AuthorizationEndpoint); provider.AuthorizationEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-B23bs", "Errors.Invalid.Argument")
		}
		if provider.TokenEndpoint = strings.TrimSpace(provider.TokenEndpoint); provider.TokenEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-D2gj8", "Errors.Invalid.Argument")
		}
		if provider.UserEndpoint = strings.TrimSpace(provider.UserEndpoint); provider.UserEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-Fb8jk", "Errors.Invalid.Argument")
		}
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			writeModel := NewOAuthInstanceIDPWriteModel(a.InstanceID, id)
			events, err := filter(ctx, writeModel.Query())
			if err != nil {
				return nil, err
			}
			writeModel.AppendEvents(events...)
			if err = writeModel.Reduce(); err != nil {
				return nil, err
			}
			secret, err := crypto.Encrypt([]byte(provider.ClientSecret), c.idpConfigEncryption)
			if err != nil {
				return nil, err
			}
			return []eventstore.Command{
				instance.NewOAuthIDPAddedEvent(
					ctx,
					&a.Aggregate,
					id,
					provider.Name,
					provider.ClientID,
					secret,
					provider.AuthorizationEndpoint,
					provider.TokenEndpoint,
					provider.UserEndpoint,
					provider.Scopes,
					provider.IDPOptions,
				),
			}, nil
		}, nil
	}
}

func (c *Commands) prepareUpdateInstanceOAuthProvider(a *instance.Aggregate, id string, provider GenericOAuthProvider) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		if provider.Name = strings.TrimSpace(provider.Name); provider.Name == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-D32ef", "Errors.Invalid.Argument")
		}
		if provider.ClientID = strings.TrimSpace(provider.ClientID); provider.ClientID == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-Dbgzf", "Errors.Invalid.Argument")
		}
		if provider.AuthorizationEndpoint = strings.TrimSpace(provider.AuthorizationEndpoint); provider.AuthorizationEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-B23bs", "Errors.Invalid.Argument")
		}
		if provider.TokenEndpoint = strings.TrimSpace(provider.TokenEndpoint); provider.TokenEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-D2gj8", "Errors.Invalid.Argument")
		}
		if provider.UserEndpoint = strings.TrimSpace(provider.UserEndpoint); provider.UserEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-Fb8jk", "Errors.Invalid.Argument")
		}
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			writeModel := NewOAuthInstanceIDPWriteModel(a.InstanceID, id)
			events, err := filter(ctx, writeModel.Query())
			if err != nil {
				return nil, err
			}
			writeModel.AppendEvents(events...)
			if err = writeModel.Reduce(); err != nil {
				return nil, err
			}
			if !writeModel.State.Exists() {
				return nil, caos_errs.ThrowNotFound(nil, "INST-D3r1s", "Errors.Instance.IDPConfig.NotExisting")
			}
			event, err := writeModel.NewChangedEvent(
				ctx,
				&a.Aggregate,
				id,
				provider.Name,
				provider.ClientID,
				provider.ClientSecret,
				c.idpConfigEncryption,
				provider.AuthorizationEndpoint,
				provider.TokenEndpoint,
				provider.UserEndpoint,
				provider.Scopes,
				provider.IDPOptions,
			)
			if err != nil {
				return nil, err
			}
			if event == nil {
				return nil, nil
			}
			return []eventstore.Command{event}, nil
		}, nil
	}
}

func (c *Commands) prepareAddInstanceGitHubProvider(a *instance.Aggregate, id string, provider GitHubProvider) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		if provider.ClientID = strings.TrimSpace(provider.ClientID); provider.ClientID == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-Jdsgf", "Errors.Invalid.Argument")
		}
		if provider.ClientSecret = strings.TrimSpace(provider.ClientSecret); provider.ClientSecret == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-dsgz3", "Errors.Invalid.Argument")
		}
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			writeModel := NewGitHubInstanceIDPWriteModel(a.InstanceID, id)
			events, err := filter(ctx, writeModel.Query())
			if err != nil {
				return nil, err
			}
			writeModel.AppendEvents(events...)
			if err = writeModel.Reduce(); err != nil {
				return nil, err
			}
			secret, err := crypto.Encrypt([]byte(provider.ClientSecret), c.idpConfigEncryption)
			if err != nil {
				return nil, err
			}
			return []eventstore.Command{
				instance.NewGitHubIDPAddedEvent(
					ctx,
					&a.Aggregate,
					id,
					provider.Name,
					provider.ClientID,
					secret,
					provider.Scopes,
					provider.IDPOptions,
				),
			}, nil
		}, nil
	}
}

func (c *Commands) prepareUpdateInstanceGitHubProvider(a *instance.Aggregate, id string, provider GitHubProvider) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		if id = strings.TrimSpace(id); id == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-sdf4h", "Errors.Invalid.Argument")
		}
		if provider.ClientID = strings.TrimSpace(provider.ClientID); provider.ClientID == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-fdh5z", "Errors.Invalid.Argument")
		}
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			writeModel := NewGitHubInstanceIDPWriteModel(a.InstanceID, id)
			events, err := filter(ctx, writeModel.Query())
			if err != nil {
				return nil, err
			}
			writeModel.AppendEvents(events...)
			if err = writeModel.Reduce(); err != nil {
				return nil, err
			}
			if !writeModel.State.Exists() {
				return nil, caos_errs.ThrowNotFound(nil, "INST-Dr1gs", "Errors.Instance.IDPConfig.NotExisting")
			}
			event, err := writeModel.NewChangedEvent(
				ctx,
				&a.Aggregate,
				id,
				provider.Name,
				provider.ClientID,
				provider.ClientSecret,
				c.idpConfigEncryption,
				provider.Scopes,
				provider.IDPOptions,
			)
			if err != nil {
				return nil, err
			}
			if event == nil {
				return nil, nil
			}
			return []eventstore.Command{event}, nil
		}, nil
	}
}

func (c *Commands) prepareAddInstanceGitHubEnterpriseProvider(a *instance.Aggregate, id string, provider GitHubEnterpriseProvider) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		if provider.Name = strings.TrimSpace(provider.Name); provider.Name == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-Dg4td", "Errors.Invalid.Argument")
		}
		if provider.ClientID = strings.TrimSpace(provider.ClientID); provider.ClientID == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-dgj53", "Errors.Invalid.Argument")
		}
		if provider.ClientSecret = strings.TrimSpace(provider.ClientSecret); provider.ClientSecret == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-Ghjjs", "Errors.Invalid.Argument")
		}
		if provider.AuthorizationEndpoint = strings.TrimSpace(provider.AuthorizationEndpoint); provider.AuthorizationEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-sani2", "Errors.Invalid.Argument")
		}
		if provider.TokenEndpoint = strings.TrimSpace(provider.TokenEndpoint); provider.TokenEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-agj42", "Errors.Invalid.Argument")
		}
		if provider.UserEndpoint = strings.TrimSpace(provider.UserEndpoint); provider.UserEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-sd5hn", "Errors.Invalid.Argument")
		}
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			writeModel := NewGitHubEnterpriseInstanceIDPWriteModel(a.InstanceID, id)
			events, err := filter(ctx, writeModel.Query())
			if err != nil {
				return nil, err
			}
			writeModel.AppendEvents(events...)
			if err = writeModel.Reduce(); err != nil {
				return nil, err
			}
			secret, err := crypto.Encrypt([]byte(provider.ClientSecret), c.idpConfigEncryption)
			if err != nil {
				return nil, err
			}
			return []eventstore.Command{
				instance.NewGitHubEnterpriseIDPAddedEvent(
					ctx,
					&a.Aggregate,
					id,
					provider.Name,
					provider.ClientID,
					secret,
					provider.AuthorizationEndpoint,
					provider.TokenEndpoint,
					provider.UserEndpoint,
					provider.Scopes,
					provider.IDPOptions,
				),
			}, nil
		}, nil
	}
}

func (c *Commands) prepareUpdateInstanceGitHubEnterpriseProvider(a *instance.Aggregate, id string, provider GitHubEnterpriseProvider) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		if id = strings.TrimSpace(id); id == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-sdfh3", "Errors.Invalid.Argument")
		}
		if provider.Name = strings.TrimSpace(provider.Name); provider.Name == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-shj42", "Errors.Invalid.Argument")
		}
		if provider.ClientID = strings.TrimSpace(provider.ClientID); provider.ClientID == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-sdh73", "Errors.Invalid.Argument")
		}
		if provider.AuthorizationEndpoint = strings.TrimSpace(provider.AuthorizationEndpoint); provider.AuthorizationEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-acx2w", "Errors.Invalid.Argument")
		}
		if provider.TokenEndpoint = strings.TrimSpace(provider.TokenEndpoint); provider.TokenEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-dgj6q", "Errors.Invalid.Argument")
		}
		if provider.UserEndpoint = strings.TrimSpace(provider.UserEndpoint); provider.UserEndpoint == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-ybj62", "Errors.Invalid.Argument")
		}
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			writeModel := NewGitHubEnterpriseInstanceIDPWriteModel(a.InstanceID, id)
			events, err := filter(ctx, writeModel.Query())
			if err != nil {
				return nil, err
			}
			writeModel.AppendEvents(events...)
			if err = writeModel.Reduce(); err != nil {
				return nil, err
			}
			if !writeModel.State.Exists() {
				return nil, caos_errs.ThrowNotFound(nil, "INST-GBr42", "Errors.Instance.IDPConfig.NotExisting")
			}
			event, err := writeModel.NewChangedEvent(
				ctx,
				&a.Aggregate,
				id,
				provider.Name,
				provider.ClientID,
				provider.ClientSecret,
				c.idpConfigEncryption,
				provider.AuthorizationEndpoint,
				provider.TokenEndpoint,
				provider.UserEndpoint,
				provider.Scopes,
				provider.IDPOptions,
			)
			if err != nil {
				return nil, err
			}
			if event == nil {
				return nil, nil
			}
			return []eventstore.Command{event}, nil
		}, nil
	}
}

func (c *Commands) prepareAddInstanceGoogleProvider(a *instance.Aggregate, id string, provider GoogleProvider) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		if provider.ClientID = strings.TrimSpace(provider.ClientID); provider.ClientID == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-D3fvs", "Errors.Invalid.Argument")
		}
		if provider.ClientSecret = strings.TrimSpace(provider.ClientSecret); provider.ClientSecret == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-W2vqs", "Errors.Invalid.Argument")
		}
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			writeModel := NewGoogleInstanceIDPWriteModel(a.InstanceID, id)
			events, err := filter(ctx, writeModel.Query())
			if err != nil {
				return nil, err
			}
			writeModel.AppendEvents(events...)
			if err = writeModel.Reduce(); err != nil {
				return nil, err
			}
			secret, err := crypto.Encrypt([]byte(provider.ClientSecret), c.idpConfigEncryption)
			if err != nil {
				return nil, err
			}
			return []eventstore.Command{
				instance.NewGoogleIDPAddedEvent(
					ctx,
					&a.Aggregate,
					id,
					provider.Name,
					provider.ClientID,
					secret,
					provider.Scopes,
					provider.IDPOptions,
				),
			}, nil
		}, nil
	}
}

func (c *Commands) prepareUpdateInstanceGoogleProvider(a *instance.Aggregate, id string, provider GoogleProvider) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		if id = strings.TrimSpace(id); id == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-S32t1", "Errors.Invalid.Argument")
		}
		if provider.ClientID = strings.TrimSpace(provider.ClientID); provider.ClientID == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-ds432", "Errors.Invalid.Argument")
		}
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			writeModel := NewGoogleInstanceIDPWriteModel(a.InstanceID, id)
			events, err := filter(ctx, writeModel.Query())
			if err != nil {
				return nil, err
			}
			writeModel.AppendEvents(events...)
			if err = writeModel.Reduce(); err != nil {
				return nil, err
			}
			if !writeModel.State.Exists() {
				return nil, caos_errs.ThrowNotFound(nil, "INST-D3r1s", "Errors.Instance.IDPConfig.NotExisting")
			}
			event, err := writeModel.NewChangedEvent(
				ctx,
				&a.Aggregate,
				id,
				provider.Name,
				provider.ClientID,
				provider.ClientSecret,
				c.idpConfigEncryption,
				provider.Scopes,
				provider.IDPOptions,
			)
			if err != nil {
				return nil, err
			}
			if event == nil {
				return nil, nil
			}
			return []eventstore.Command{event}, nil
		}, nil
	}
}

func (c *Commands) prepareAddInstanceLDAPProvider(a *instance.Aggregate, id string, provider LDAPProvider) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		if provider.Name = strings.TrimSpace(provider.Name); provider.Name == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-SAfdd", "Errors.Invalid.Argument")
		}
		if provider.Host = strings.TrimSpace(provider.Host); provider.Host == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-SDVg2", "Errors.Invalid.Argument")
		}
		if provider.BaseDN = strings.TrimSpace(provider.BaseDN); provider.BaseDN == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-sv31s", "Errors.Invalid.Argument")
		}
		if provider.UserObjectClass = strings.TrimSpace(provider.UserObjectClass); provider.UserObjectClass == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-sdgf4", "Errors.Invalid.Argument")
		}
		if provider.UserUniqueAttribute = strings.TrimSpace(provider.UserUniqueAttribute); provider.UserUniqueAttribute == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-AEG2w", "Errors.Invalid.Argument")
		}
		if provider.Admin = strings.TrimSpace(provider.Admin); provider.Admin == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-SAD5n", "Errors.Invalid.Argument")
		}
		if provider.Password = strings.TrimSpace(provider.Password); provider.Password == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-sdf5h", "Errors.Invalid.Argument")
		}
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			writeModel := NewLDAPInstanceIDPWriteModel(a.InstanceID, id)
			events, err := filter(ctx, writeModel.Query())
			if err != nil {
				return nil, err
			}
			writeModel.AppendEvents(events...)
			if err = writeModel.Reduce(); err != nil {
				return nil, err
			}
			secret, err := crypto.Encrypt([]byte(provider.Password), c.idpConfigEncryption)
			if err != nil {
				return nil, err
			}
			return []eventstore.Command{
				instance.NewLDAPIDPAddedEvent(
					ctx,
					&a.Aggregate,
					id,
					provider.Name,
					provider.Host,
					provider.Port,
					provider.TLS,
					provider.BaseDN,
					provider.UserObjectClass,
					provider.UserUniqueAttribute,
					provider.Admin,
					secret,
					provider.LDAPAttributes,
					provider.IDPOptions,
				),
			}, nil
		}, nil
	}
}

func (c *Commands) prepareUpdateInstanceLDAPProvider(a *instance.Aggregate, id string, provider LDAPProvider) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		if id = strings.TrimSpace(id); id == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-Dgdbs", "Errors.Invalid.Argument")
		}
		if provider.Name = strings.TrimSpace(provider.Name); provider.Name == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-Sffgd", "Errors.Invalid.Argument")
		}
		if provider.Host = strings.TrimSpace(provider.Host); provider.Host == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-Dz62d", "Errors.Invalid.Argument")
		}
		if provider.BaseDN = strings.TrimSpace(provider.BaseDN); provider.BaseDN == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-vb3ss", "Errors.Invalid.Argument")
		}
		if provider.UserObjectClass = strings.TrimSpace(provider.UserObjectClass); provider.UserObjectClass == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-hbere", "Errors.Invalid.Argument")
		}
		if provider.UserUniqueAttribute = strings.TrimSpace(provider.UserUniqueAttribute); provider.UserUniqueAttribute == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-ASFt6", "Errors.Invalid.Argument")
		}
		if provider.Admin = strings.TrimSpace(provider.Admin); provider.Admin == "" {
			return nil, caos_errs.ThrowInvalidArgument(nil, "INST-DG45z", "Errors.Invalid.Argument")
		}
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			writeModel := NewLDAPInstanceIDPWriteModel(a.InstanceID, id)
			events, err := filter(ctx, writeModel.Query())
			if err != nil {
				return nil, err
			}
			writeModel.AppendEvents(events...)
			if err = writeModel.Reduce(); err != nil {
				return nil, err
			}
			if !writeModel.State.Exists() {
				return nil, caos_errs.ThrowNotFound(nil, "INST-ASF3F", "Errors.Instance.IDPConfig.NotExisting")
			}
			event, err := writeModel.NewChangedEvent(
				ctx,
				&a.Aggregate,
				id,
				writeModel.Name,
				provider.Name,
				provider.Host,
				provider.Port,
				provider.TLS,
				provider.BaseDN,
				provider.UserObjectClass,
				provider.UserUniqueAttribute,
				provider.Admin,
				provider.Password,
				c.idpConfigEncryption,
				provider.LDAPAttributes,
				provider.IDPOptions,
			)
			if err != nil {
				return nil, err
			}
			if event == nil {
				return nil, nil
			}
			return []eventstore.Command{event}, nil
		}, nil
	}
}

func (c *Commands) prepareDeleteInstanceProvider(a *instance.Aggregate, id string) preparation.Validation {
	return func() (preparation.CreateCommands, error) {
		return func(ctx context.Context, filter preparation.FilterToQueryReducer) ([]eventstore.Command, error) {
			writeModel := NewInstanceIDPRemoveWriteModel(a.InstanceID, id)
			events, err := filter(ctx, writeModel.Query())
			if err != nil {
				return nil, err
			}
			writeModel.AppendEvents(events...)
			if err = writeModel.Reduce(); err != nil {
				return nil, err
			}
			if !writeModel.State.Exists() {
				return nil, caos_errs.ThrowNotFound(nil, "INST-Se3tg", "Errors.Instance.IDPConfig.NotExisting")
			}
			return []eventstore.Command{instance.NewIDPRemovedEvent(ctx, &a.Aggregate, id, writeModel.name)}, nil
		}, nil
	}
}
