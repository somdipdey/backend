package cmd

import (
	"fmt"
	coreCmd "github.com/bitclout/core/cmd"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

type Config struct {
	// Core
	TXIndex                bool
	APIPort                uint16

	// Onboarding
	StarterBitcloutSeed    string
	StarterBitcloutNanos   uint64
	StarterPrefixNanosMap  map[string]uint64
	TwilioAccountSID       string
	TwilioAuthToken        string
	TwilioVerifyServiceID  string
	CompProfileCreation    bool
	MinSatoshisForProfile  uint64

	// Global State
	GlobalStateRemoteNode   string
	GlobalStateRemoteSecret string

	// Web Security
	AccessControlAllowOrigins []string
	SecureHeaderDevelopment   bool
	SecureHeaderAllowHosts    []string
	AdminPublicKeys           []string

	// Analytics + Profiling
	AmplitudeKey           string
	AmplitudeDomain        string
	DatadogProfiler        bool

	// User Interface
	SupportEmail           string
	ShowProcessingSpinners bool

	// Images
	GCPCredentialsPath     string
	GCPBucketName          string
}

func LoadConfig(coreConfig *coreCmd.Config) *Config {
	config := Config{}

	// Core
	config.TXIndex = viper.GetBool("txindex")
	config.APIPort = uint16(viper.GetUint64("api-port"))
	if config.APIPort <= 0 {
		// TODO: pull this out of core. we shouldn't need core's config here
		config.APIPort = coreConfig.Params.DefaultJSONPort
	}

	// Onboarding
	config.StarterBitcloutSeed = viper.GetString("starter-bitclout-seed")
	config.StarterBitcloutNanos = viper.GetUint64("starter-bitclout-nanos")
	starterPrefixNanosMap := viper.GetString("starter-prefix-nanos-map")
	if len(starterPrefixNanosMap) > 0 {
		config.StarterPrefixNanosMap = make(map[string]uint64)
		for _, pair := range strings.Split(starterPrefixNanosMap, ",") {
			entry := strings.Split(pair, "=")
			nanos, err := strconv.Atoi(entry[1])
			if err != nil {
				fmt.Printf("invalid nanos: %s", entry[1])
			}
			config.StarterPrefixNanosMap[entry[0]] = uint64(nanos)
		}
	}
	config.TwilioAccountSID = viper.GetString("twilio-account-sid")
	config.TwilioAuthToken = viper.GetString("twilio-auth-token")
	config.TwilioVerifyServiceID = viper.GetString("twilio-verify-service-id")
	config.CompProfileCreation = viper.GetBool("comp-profile-creation")
	config.MinSatoshisForProfile = viper.GetUint64("min-satoshis-for-profile")

	// Global State
	config.GlobalStateRemoteNode = viper.GetString("global-state-remote-node")
	config.GlobalStateRemoteSecret = viper.GetString("global-state-remote-secret")

	// Web Security
	config.AccessControlAllowOrigins = viper.GetStringSlice("access-control-allow-origins")
	config.SecureHeaderDevelopment = viper.GetBool("secure-header-development")
	config.SecureHeaderAllowHosts =  viper.GetStringSlice("secure-header-allow-hosts")
	config.AdminPublicKeys = viper.GetStringSlice("admin-public-keys")

	// Analytics + Profiling
	config.AmplitudeKey = viper.GetString("amplitude-key")
	config.AmplitudeDomain = viper.GetString("amplitude-domain")

	// User Interface
	config.SupportEmail = viper.GetString("support-email")
	config.ShowProcessingSpinners = viper.GetBool("show-processing-spinners")

	// Images
	config.GCPCredentialsPath = viper.GetString("gcp-credentials-path")
	config.GCPBucketName = viper.GetString("gcp-bucket-name")

	return &config
}
