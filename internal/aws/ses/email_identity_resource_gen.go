// Code generated by generators/resource/main.go; DO NOT EDIT.

package ses

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ses_email_identity", emailIdentityResourceType)
}

// emailIdentityResourceType returns the Terraform awscc_ses_email_identity resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SES::EmailIdentity resource type.
func emailIdentityResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"configuration_set_attributes": {
			// Property: ConfigurationSetAttributes
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Used to associate a configuration set with an email identity.",
			//   "properties": {
			//     "ConfigurationSetName": {
			//       "description": "The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Used to associate a configuration set with an email identity.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"configuration_set_name": {
						// Property: ConfigurationSetName
						Description: "The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"dkim_attributes": {
			// Property: DkimAttributes
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Used to enable or disable DKIM authentication for an email identity.",
			//   "properties": {
			//     "SigningEnabled": {
			//       "description": "Sets the DKIM signing configuration for the identity. When you set this value true, then the messages that are sent from the identity are signed using DKIM. If you set this value to false, your messages are sent without DKIM signing.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Used to enable or disable DKIM authentication for an email identity.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"signing_enabled": {
						// Property: SigningEnabled
						Description: "Sets the DKIM signing configuration for the identity. When you set this value true, then the messages that are sent from the identity are signed using DKIM. If you set this value to false, your messages are sent without DKIM signing.",
						Type:        types.BoolType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"dkim_dns_token_name_1": {
			// Property: DkimDNSTokenName1
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"dkim_dns_token_name_2": {
			// Property: DkimDNSTokenName2
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"dkim_dns_token_name_3": {
			// Property: DkimDNSTokenName3
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"dkim_dns_token_value_1": {
			// Property: DkimDNSTokenValue1
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"dkim_dns_token_value_2": {
			// Property: DkimDNSTokenValue2
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"dkim_dns_token_value_3": {
			// Property: DkimDNSTokenValue3
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"dkim_signing_attributes": {
			// Property: DkimSigningAttributes
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for Easy DKIM.",
			//   "properties": {
			//     "DomainSigningPrivateKey": {
			//       "description": "[Bring Your Own DKIM] A private key that's used to generate a DKIM signature. The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.",
			//       "type": "string"
			//     },
			//     "DomainSigningSelector": {
			//       "description": "[Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.",
			//       "type": "string"
			//     },
			//     "NextSigningKeyLength": {
			//       "description": "[Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day.",
			//       "pattern": "RSA_1024_BIT|RSA_2048_BIT",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for Easy DKIM.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"domain_signing_private_key": {
						// Property: DomainSigningPrivateKey
						Description: "[Bring Your Own DKIM] A private key that's used to generate a DKIM signature. The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.",
						Type:        types.StringType,
						Optional:    true,
					},
					"domain_signing_selector": {
						// Property: DomainSigningSelector
						Description: "[Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.",
						Type:        types.StringType,
						Optional:    true,
					},
					"next_signing_key_length": {
						// Property: NextSigningKeyLength
						Description: "[Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("RSA_1024_BIT|RSA_2048_BIT"), ""),
						},
					},
				},
			),
			Optional: true,
			// DkimSigningAttributes is a write-only property.
		},
		"email_identity": {
			// Property: EmailIdentity
			// CloudFormation resource type schema:
			// {
			//   "description": "The email address or domain to verify.",
			//   "type": "string"
			// }
			Description: "The email address or domain to verify.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"feedback_attributes": {
			// Property: FeedbackAttributes
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Used to enable or disable feedback forwarding for an identity.",
			//   "properties": {
			//     "EmailForwardingEnabled": {
			//       "description": "If the value is true, you receive email notifications when bounce or complaint events occur",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Used to enable or disable feedback forwarding for an identity.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"email_forwarding_enabled": {
						// Property: EmailForwardingEnabled
						Description: "If the value is true, you receive email notifications when bounce or complaint events occur",
						Type:        types.BoolType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"mail_from_attributes": {
			// Property: MailFromAttributes
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Used to enable or disable the custom Mail-From domain configuration for an email identity.",
			//   "properties": {
			//     "BehaviorOnMxFailure": {
			//       "description": "The action to take if the required MX record isn't found when you send an email. When you set this value to UseDefaultValue , the mail is sent using amazonses.com as the MAIL FROM domain. When you set this value to RejectMessage , the Amazon SES API v2 returns a MailFromDomainNotVerified error, and doesn't attempt to deliver the email.",
			//       "pattern": "USE_DEFAULT_VALUE|REJECT_MESSAGE",
			//       "type": "string"
			//     },
			//     "MailFromDomain": {
			//       "description": "The custom MAIL FROM domain that you want the verified identity to use",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Used to enable or disable the custom Mail-From domain configuration for an email identity.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"behavior_on_mx_failure": {
						// Property: BehaviorOnMxFailure
						Description: "The action to take if the required MX record isn't found when you send an email. When you set this value to UseDefaultValue , the mail is sent using amazonses.com as the MAIL FROM domain. When you set this value to RejectMessage , the Amazon SES API v2 returns a MailFromDomainNotVerified error, and doesn't attempt to deliver the email.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("USE_DEFAULT_VALUE|REJECT_MESSAGE"), ""),
						},
					},
					"mail_from_domain": {
						// Property: MailFromDomain
						Description: "The custom MAIL FROM domain that you want the verified identity to use",
						Type:        types.StringType,
						Optional:    true,
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
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::SES::EmailIdentity",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::EmailIdentity").WithTerraformTypeName("awscc_ses_email_identity")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"behavior_on_mx_failure":       "BehaviorOnMxFailure",
		"configuration_set_attributes": "ConfigurationSetAttributes",
		"configuration_set_name":       "ConfigurationSetName",
		"dkim_attributes":              "DkimAttributes",
		"dkim_dns_token_name_1":        "DkimDNSTokenName1",
		"dkim_dns_token_name_2":        "DkimDNSTokenName2",
		"dkim_dns_token_name_3":        "DkimDNSTokenName3",
		"dkim_dns_token_value_1":       "DkimDNSTokenValue1",
		"dkim_dns_token_value_2":       "DkimDNSTokenValue2",
		"dkim_dns_token_value_3":       "DkimDNSTokenValue3",
		"dkim_signing_attributes":      "DkimSigningAttributes",
		"domain_signing_private_key":   "DomainSigningPrivateKey",
		"domain_signing_selector":      "DomainSigningSelector",
		"email_forwarding_enabled":     "EmailForwardingEnabled",
		"email_identity":               "EmailIdentity",
		"feedback_attributes":          "FeedbackAttributes",
		"mail_from_attributes":         "MailFromAttributes",
		"mail_from_domain":             "MailFromDomain",
		"next_signing_key_length":      "NextSigningKeyLength",
		"signing_enabled":              "SigningEnabled",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/DkimSigningAttributes",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
