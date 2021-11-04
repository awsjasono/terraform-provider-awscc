// Code generated by generators/resource/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_redshift_scheduled_action", scheduledActionResourceType)
}

// scheduledActionResourceType returns the Terraform awscc_redshift_scheduled_action resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Redshift::ScheduledAction resource type.
func scheduledActionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"enable": {
			// Property: Enable
			// CloudFormation resource type schema:
			// {
			//   "description": "If true, the schedule is enabled. If false, the scheduled action does not trigger.",
			//   "type": "boolean"
			// }
			Description: "If true, the schedule is enabled. If false, the scheduled action does not trigger.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"end_time": {
			// Property: EndTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger.",
			//   "type": "string"
			// }
			Description: "The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger.",
			Type:        types.StringType,
			Optional:    true,
		},
		"iam_role": {
			// Property: IamRole
			// CloudFormation resource type schema:
			// {
			//   "description": "The IAM role to assume to run the target action.",
			//   "type": "string"
			// }
			Description: "The IAM role to assume to run the target action.",
			Type:        types.StringType,
			Optional:    true,
		},
		"next_invocations": {
			// Property: NextInvocations
			// CloudFormation resource type schema:
			// {
			//   "description": "List of times when the scheduled action will run.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "List of times when the scheduled action will run.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"schedule": {
			// Property: Schedule
			// CloudFormation resource type schema:
			// {
			//   "description": "The schedule in `at( )` or `cron( )` format.",
			//   "type": "string"
			// }
			Description: "The schedule in `at( )` or `cron( )` format.",
			Type:        types.StringType,
			Optional:    true,
		},
		"scheduled_action_description": {
			// Property: ScheduledActionDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the scheduled action.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The description of the scheduled action.",
			Type:        types.StringType,
			Optional:    true,
		},
		"scheduled_action_name": {
			// Property: ScheduledActionName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the scheduled action. The name must be unique within an account.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the scheduled action. The name must be unique within an account.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"start_time": {
			// Property: StartTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger.",
			//   "type": "string"
			// }
			Description: "The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger.",
			Type:        types.StringType,
			Optional:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of the scheduled action.",
			//   "enum": [
			//     "ACTIVE",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Description: "The state of the scheduled action.",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_action": {
			// Property: TargetAction
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A JSON format string of the Amazon Redshift API operation with input parameters.",
			//   "properties": {
			//     "PauseCluster": {
			//       "additionalProperties": false,
			//       "description": "Describes a pause cluster operation. For example, a scheduled action to run the `PauseCluster` API operation.",
			//       "properties": {
			//         "ClusterIdentifier": {
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "ClusterIdentifier"
			//       ],
			//       "type": "object"
			//     },
			//     "ResizeCluster": {
			//       "additionalProperties": false,
			//       "description": "Describes a resize cluster operation. For example, a scheduled action to run the `ResizeCluster` API operation.",
			//       "properties": {
			//         "Classic": {
			//           "type": "boolean"
			//         },
			//         "ClusterIdentifier": {
			//           "type": "string"
			//         },
			//         "ClusterType": {
			//           "type": "string"
			//         },
			//         "NodeType": {
			//           "type": "string"
			//         },
			//         "NumberOfNodes": {
			//           "type": "integer"
			//         }
			//       },
			//       "required": [
			//         "ClusterIdentifier"
			//       ],
			//       "type": "object"
			//     },
			//     "ResumeCluster": {
			//       "additionalProperties": false,
			//       "description": "Describes a resume cluster operation. For example, a scheduled action to run the `ResumeCluster` API operation.",
			//       "properties": {
			//         "ClusterIdentifier": {
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "ClusterIdentifier"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A JSON format string of the Amazon Redshift API operation with input parameters.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"pause_cluster": {
						// Property: PauseCluster
						Description: "Describes a pause cluster operation. For example, a scheduled action to run the `PauseCluster` API operation.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"cluster_identifier": {
									// Property: ClusterIdentifier
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
					"resize_cluster": {
						// Property: ResizeCluster
						Description: "Describes a resize cluster operation. For example, a scheduled action to run the `ResizeCluster` API operation.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"classic": {
									// Property: Classic
									Type:     types.BoolType,
									Optional: true,
								},
								"cluster_identifier": {
									// Property: ClusterIdentifier
									Type:     types.StringType,
									Required: true,
								},
								"cluster_type": {
									// Property: ClusterType
									Type:     types.StringType,
									Optional: true,
								},
								"node_type": {
									// Property: NodeType
									Type:     types.StringType,
									Optional: true,
								},
								"number_of_nodes": {
									// Property: NumberOfNodes
									Type:     types.NumberType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"resume_cluster": {
						// Property: ResumeCluster
						Description: "Describes a resume cluster operation. For example, a scheduled action to run the `ResumeCluster` API operation.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"cluster_identifier": {
									// Property: ClusterIdentifier
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The `AWS::Redshift::ScheduledAction` resource creates an Amazon Redshift Scheduled Action.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::ScheduledAction").WithTerraformTypeName("awscc_redshift_scheduled_action")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"classic":                      "Classic",
		"cluster_identifier":           "ClusterIdentifier",
		"cluster_type":                 "ClusterType",
		"enable":                       "Enable",
		"end_time":                     "EndTime",
		"iam_role":                     "IamRole",
		"next_invocations":             "NextInvocations",
		"node_type":                    "NodeType",
		"number_of_nodes":              "NumberOfNodes",
		"pause_cluster":                "PauseCluster",
		"resize_cluster":               "ResizeCluster",
		"resume_cluster":               "ResumeCluster",
		"schedule":                     "Schedule",
		"scheduled_action_description": "ScheduledActionDescription",
		"scheduled_action_name":        "ScheduledActionName",
		"start_time":                   "StartTime",
		"state":                        "State",
		"target_action":                "TargetAction",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
