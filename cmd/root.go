package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v3"
)

// AppName - application name
const AppName string = "vamp-cloud-cli"

// Version - version of the cli
const Version string = "v1.0.8"

// ApiVersion - supported version of the api
const ApiVersion string = "v1"

// AddAppName - Application name can change over time so it is made parameteric
func AddAppName(str string) string {
	return strings.Replace(str, "$AppName", AppName, -1)
}

// Config - vamp cloud cli configuration
var Config models.VampCloudCliConfiguration

// Common code parameters
var cfgFile string
var configPath string
var configFileType string
var description string
var applicationName string

var apiKey string
var apiUrl string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   AddAppName("$AppName"),
	Short: AddAppName("A command line client for $AppName"),
	Long: AddAppName(`$AppName is a setup tool for vamp.
It is required to have a default config.
Envrionment variables can be used to override the values in the config.
Environment variables:
	VAMP_CLOUD_API_URL
	VAMP_CLOUD_API_KEY`),
}

// RootCmd - returns root command for integration tests
func RootCmd() *cobra.Command {
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	logging.Init(os.Stdout, os.Stderr)

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vamp/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVarP(&logging.Verbose, "verbose", "v", false, "Verbose")

	rootCmd.PersistentFlags().StringVarP(&apiKey, "key", "k", "", "Vamp Cloud api key")
	rootCmd.PersistentFlags().StringVarP(&apiUrl, "address", "a", "", "Vamp Cloud api address")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".$AppName" (without extension).
		path := filepath.FromSlash(home + AddAppName("/.$AppName"))

		logging.Info("Looking for config", logging.NewPair("path", path))

		if _, pathErr := os.Stat(path); os.IsNotExist(pathErr) {
			// path does not exist
			err = os.MkdirAll(path, 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		viper.AddConfigPath(path)
		viper.SetConfigName("config")
	}

	setupConfigurationDefaults()
	setupConfigurationEnvrionmentVariables()

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	viper.ReadInConfig()

	// unmarshal config
	c := viper.AllSettings()

	bs, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}

	unmarshallError := yaml.Unmarshal(bs, &Config)
	if unmarshallError != nil {
		panic(unmarshallError)
	}

	setupConfigurationOverrides(&Config)

	jsonConfig, _ := json.Marshal(Config)

	logging.Info(fmt.Sprintf("Vamp Cloud cli configuration %v", utils.PrettyJson(string(jsonConfig))))
}

func setupConfigurationEnvrionmentVariables() {
	viper.BindEnv("vamp-cloud-api-url", "VAMP_CLOUD_API_URL")
	viper.BindEnv("vamp-cloud-api-key", "VAMP_CLOUD_API_KEY")
}

func setupConfigurationOverrides(config *models.VampCloudCliConfiguration) {
	if apiUrl != "" {
		config.VampCloudApiURL = apiUrl
	}
	if apiKey != "" {
		config.APIKey = apiKey
	}
}

func checkValues(config models.VampCloudCliConfiguration) error {
	if config.VampCloudApiURL == "" {
		return errors.New("configuration is invalid: VAMP_CLOUD_API_URL is required")
	}
	if config.APIKey == "" {
		return errors.New("configuration is invalid: VAMP_CLOUD_API_KEY is required")
	}

	return nil
}

func setupConfigurationDefaults() {
	viper.SetDefault("vamp-cloud-addr", "https://vamp.cloud/api/public")
}

func handleErrorOnName(err error) error {

	if outputType == "name" {

		var check *usecase.ResourceNotFoundError
		if !errors.As(err, &check) {
			return nil
		}

		return err
	} else {
		return err
	}

}
