// Code generated by generators/resource/main.go; DO NOT EDIT.

package efs

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_efs_access_point", accessPointResourceType)
}

// accessPointResourceType returns the Terraform awscc_efs_access_point resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EFS::AccessPoint resource type.
func accessPointResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_point_id": {
			// Property: AccessPointId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"access_point_tags": {
			// Property: AccessPointTags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"client_token": {
			// Property: ClientToken
			// CloudFormation resource type schema:
			// {
			//   "description": "(optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.",
			//   "type": "string"
			// }
			Description: "(optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // ClientToken is a force-new property.
			},
		},
		"file_system_id": {
			// Property: FileSystemId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the EFS file system that the access point provides access to.",
			//   "type": "string"
			// }
			Description: "The ID of the EFS file system that the access point provides access to.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // FileSystemId is a force-new property.
			},
		},
		"posix_user": {
			// Property: PosixUser
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Gid": {
			//       "description": "The POSIX group ID used for all file system operations using this access point.",
			//       "type": "string"
			//     },
			//     "SecondaryGids": {
			//       "description": "Secondary POSIX group IDs used for all file system operations using this access point.",
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "Uid": {
			//       "description": "The POSIX user ID used for all file system operations using this access point.",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Uid",
			//     "Gid"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"gid": {
						// Property: Gid
						Description: "The POSIX group ID used for all file system operations using this access point.",
						Type:        types.StringType,
						Required:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							tfsdk.RequiresReplace(), // Gid is a force-new property.
						},
					},
					"secondary_gids": {
						// Property: SecondaryGids
						Description: "Secondary POSIX group IDs used for all file system operations using this access point.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							tfsdk.RequiresReplace(), // SecondaryGids is a force-new property.
						},
					},
					"uid": {
						// Property: Uid
						Description: "The POSIX user ID used for all file system operations using this access point.",
						Type:        types.StringType,
						Required:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							tfsdk.RequiresReplace(), // Uid is a force-new property.
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // PosixUser is a force-new property.
			},
		},
		"root_directory": {
			// Property: RootDirectory
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "CreationInfo": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "OwnerGid": {
			//           "description": "Specifies the POSIX group ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).",
			//           "type": "string"
			//         },
			//         "OwnerUid": {
			//           "description": "Specifies the POSIX user ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).",
			//           "type": "string"
			//         },
			//         "Permissions": {
			//           "description": "Specifies the POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.",
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "OwnerUid",
			//         "OwnerGid",
			//         "Permissions"
			//       ],
			//       "type": "object"
			//     },
			//     "Path": {
			//       "description": "Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the CreationInfo.",
			//       "maxLength": 100,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"creation_info": {
						// Property: CreationInfo
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"owner_gid": {
									// Property: OwnerGid
									Description: "Specifies the POSIX group ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).",
									Type:        types.StringType,
									Required:    true,
								},
								"owner_uid": {
									// Property: OwnerUid
									Description: "Specifies the POSIX user ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).",
									Type:        types.StringType,
									Required:    true,
								},
								"permissions": {
									// Property: Permissions
									Description: "Specifies the POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							tfsdk.RequiresReplace(), // CreationInfo is a force-new property.
						},
					},
					"path": {
						// Property: Path
						Description: "Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the CreationInfo.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 100),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							tfsdk.RequiresReplace(), // Path is a force-new property.
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // RootDirectory is a force-new property.
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::EFS::AccessPoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EFS::AccessPoint").WithTerraformTypeName("awscc_efs_access_point")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_point_id":   "AccessPointId",
		"access_point_tags": "AccessPointTags",
		"arn":               "Arn",
		"client_token":      "ClientToken",
		"creation_info":     "CreationInfo",
		"file_system_id":    "FileSystemId",
		"gid":               "Gid",
		"key":               "Key",
		"owner_gid":         "OwnerGid",
		"owner_uid":         "OwnerUid",
		"path":              "Path",
		"permissions":       "Permissions",
		"posix_user":        "PosixUser",
		"root_directory":    "RootDirectory",
		"secondary_gids":    "SecondaryGids",
		"uid":               "Uid",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_efs_access_point", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}