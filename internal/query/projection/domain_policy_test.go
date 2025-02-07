package projection

import (
	"testing"

	"github.com/zitadel/zitadel/internal/domain"
	"github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/eventstore/handler"
	"github.com/zitadel/zitadel/internal/eventstore/repository"
	"github.com/zitadel/zitadel/internal/repository/instance"
	"github.com/zitadel/zitadel/internal/repository/org"
)

func TestDomainPolicyProjection_reduces(t *testing.T) {
	type args struct {
		event func(t *testing.T) eventstore.Event
	}
	tests := []struct {
		name   string
		args   args
		reduce func(event eventstore.Event) (*handler.Statement, error)
		want   wantReduce
	}{
		{
			name: "org reduceAdded",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(org.DomainPolicyAddedEventType),
					org.AggregateType,
					[]byte(`{
						"userLoginMustBeDomain": true,
						"validateOrgDomains": true,
						"smtpSenderAddressMatchesInstanceDomain": true
}`),
				), org.DomainPolicyAddedEventMapper),
			},
			reduce: (&domainPolicyProjection{}).reduceAdded,
			want: wantReduce{
				aggregateType:    eventstore.AggregateType("org"),
				sequence:         15,
				previousSequence: 10,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "INSERT INTO projections.domain_policies2 (creation_date, change_date, sequence, id, state, user_login_must_be_domain, validate_org_domains, smtp_sender_address_matches_instance_domain, is_default, resource_owner, instance_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
							expectedArgs: []interface{}{
								anyArg{},
								anyArg{},
								uint64(15),
								"agg-id",
								domain.PolicyStateActive,
								true,
								true,
								true,
								false,
								"ro-id",
								"instance-id",
							},
						},
					},
				},
			},
		},
		{
			name:   "org reduceChanged",
			reduce: (&domainPolicyProjection{}).reduceChanged,
			args: args{
				event: getEvent(testEvent(
					repository.EventType(org.DomainPolicyChangedEventType),
					org.AggregateType,
					[]byte(`{
						"userLoginMustBeDomain": true,
						"validateOrgDomains": true,
						"smtpSenderAddressMatchesInstanceDomain": true
		}`),
				), org.DomainPolicyChangedEventMapper),
			},
			want: wantReduce{
				aggregateType:    eventstore.AggregateType("org"),
				sequence:         15,
				previousSequence: 10,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "UPDATE projections.domain_policies2 SET (change_date, sequence, user_login_must_be_domain, validate_org_domains, smtp_sender_address_matches_instance_domain) = ($1, $2, $3, $4, $5) WHERE (id = $6) AND (instance_id = $7)",
							expectedArgs: []interface{}{
								anyArg{},
								uint64(15),
								true,
								true,
								true,
								"agg-id",
								"instance-id",
							},
						},
					},
				},
			},
		},
		{
			name:   "org reduceRemoved",
			reduce: (&domainPolicyProjection{}).reduceRemoved,
			args: args{
				event: getEvent(testEvent(
					repository.EventType(org.DomainPolicyRemovedEventType),
					org.AggregateType,
					nil,
				), org.DomainPolicyRemovedEventMapper),
			},
			want: wantReduce{
				aggregateType:    eventstore.AggregateType("org"),
				sequence:         15,
				previousSequence: 10,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "DELETE FROM projections.domain_policies2 WHERE (id = $1) AND (instance_id = $2)",
							expectedArgs: []interface{}{
								"agg-id",
								"instance-id",
							},
						},
					},
				},
			},
		},
		{
			name: "instance reduceInstanceRemoved",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(instance.InstanceRemovedEventType),
					instance.AggregateType,
					nil,
				), instance.InstanceRemovedEventMapper),
			},
			reduce: reduceInstanceRemovedHelper(DomainPolicyInstanceIDCol),
			want: wantReduce{
				aggregateType:    eventstore.AggregateType("instance"),
				sequence:         15,
				previousSequence: 10,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "DELETE FROM projections.domain_policies2 WHERE (instance_id = $1)",
							expectedArgs: []interface{}{
								"agg-id",
							},
						},
					},
				},
			},
		},
		{
			name:   "instance reduceAdded",
			reduce: (&domainPolicyProjection{}).reduceAdded,
			args: args{
				event: getEvent(testEvent(
					repository.EventType(instance.DomainPolicyAddedEventType),
					instance.AggregateType,
					[]byte(`{
						"userLoginMustBeDomain": true,
						"validateOrgDomains": true,
						"smtpSenderAddressMatchesInstanceDomain": true
					}`),
				), instance.DomainPolicyAddedEventMapper),
			},
			want: wantReduce{
				aggregateType:    eventstore.AggregateType("instance"),
				sequence:         15,
				previousSequence: 10,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "INSERT INTO projections.domain_policies2 (creation_date, change_date, sequence, id, state, user_login_must_be_domain, validate_org_domains, smtp_sender_address_matches_instance_domain, is_default, resource_owner, instance_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
							expectedArgs: []interface{}{
								anyArg{},
								anyArg{},
								uint64(15),
								"agg-id",
								domain.PolicyStateActive,
								true,
								true,
								true,
								true,
								"ro-id",
								"instance-id",
							},
						},
					},
				},
			},
		},
		{
			name:   "instance reduceChanged",
			reduce: (&domainPolicyProjection{}).reduceChanged,
			args: args{
				event: getEvent(testEvent(
					repository.EventType(instance.DomainPolicyChangedEventType),
					instance.AggregateType,
					[]byte(`{
						"userLoginMustBeDomain": true,
						"validateOrgDomains": true,
						"smtpSenderAddressMatchesInstanceDomain": true
					}`),
				), instance.DomainPolicyChangedEventMapper),
			},
			want: wantReduce{
				aggregateType:    eventstore.AggregateType("instance"),
				sequence:         15,
				previousSequence: 10,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "UPDATE projections.domain_policies2 SET (change_date, sequence, user_login_must_be_domain, validate_org_domains, smtp_sender_address_matches_instance_domain) = ($1, $2, $3, $4, $5) WHERE (id = $6) AND (instance_id = $7)",
							expectedArgs: []interface{}{
								anyArg{},
								uint64(15),
								true,
								true,
								true,
								"agg-id",
								"instance-id",
							},
						},
					},
				},
			},
		},
		{
			name:   "org.reduceOwnerRemoved",
			reduce: (&domainPolicyProjection{}).reduceOwnerRemoved,
			args: args{
				event: getEvent(testEvent(
					repository.EventType(org.OrgRemovedEventType),
					org.AggregateType,
					nil,
				), org.OrgRemovedEventMapper),
			},
			want: wantReduce{
				aggregateType:    eventstore.AggregateType("org"),
				sequence:         15,
				previousSequence: 10,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "UPDATE projections.domain_policies2 SET (change_date, sequence, owner_removed) = ($1, $2, $3) WHERE (instance_id = $4) AND (resource_owner = $5)",
							expectedArgs: []interface{}{
								anyArg{},
								uint64(15),
								true,
								"instance-id",
								"agg-id",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := baseEvent(t)
			got, err := tt.reduce(event)
			if _, ok := err.(errors.InvalidArgument); !ok {
				t.Errorf("no wrong event mapping: %v, got: %v", err, got)
			}

			event = tt.args.event(t)
			got, err = tt.reduce(event)
			assertReduce(t, got, err, DomainPolicyTable, tt.want)
		})
	}
}
