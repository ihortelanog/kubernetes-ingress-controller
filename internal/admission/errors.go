package admission

const (
	ErrTextAdminAPIUnavailable                = "could not talk to Kong admin API"
	ErrTextConsumerCredentialSecretNotFound   = "consumer referenced non-existent credentials secret"
	ErrTextConsumerCredentialValidationFailed = "consumer credential failed validation"
	ErrTextConsumerExists                     = "consumer already exists"
	ErrTextConsumerUnretrievable              = "failed to fetch consumer from kong"
	ErrTextConsumerGroupUnsupported           = "consumer group support requires Kong Enterprise"
	ErrTextConsumerGroupUnlicensed            = "consumer group support requires a valid Kong Enterprise license"
	ErrTextConsumerGroupUnexpected            = "unexpected error during checking support for consumer group"
	ErrTextConsumerUsernameEmpty              = "username cannot be empty"
	ErrTextFailedToRetrieveSecret             = "could not retrieve secrets from the kubernetes API" //nolint:revive,gosec
	ErrTextPluginConfigInvalid                = "could not parse plugin configuration"
	ErrTextPluginConfigValidationFailed       = "unable to validate plugin schema"
	ErrTextPluginConfigViolatesSchema         = "plugin failed schema validation: %s"
	ErrTextPluginNameEmpty                    = "plugin name cannot be empty"
	ErrTextPluginSecretConfigUnretrievable    = "could not load secret plugin configuration"
	ErrTextPluginUsesBothConfigTypes          = "plugin cannot use both Config and ConfigFrom"
)

const (
	ErrTextCantRetrieveGatewayClass    = "gatewayclass for this gateway could not be retrieved"
	ErrTextInvalidGatewayConfiguration = "gateway metadata and/or spec are invalid"
)
