package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configuration     Configuration
	configurationFile string
)

type Configuration struct {
	Name string   `mapstructure:"name"`
	Func []string `mapstructure:"func"`
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "run configured func",
	Run: func(cmd *cobra.Command, args []string) {
		for _, item := range configuration.Func {
			utility, ok := UtilityMap[item]
			if !ok {
				fmt.Printf("error: could not find %s in utility map\n", item)
				continue
			}
			utility()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if configurationFile != "" {
		viper.SetConfigFile(configurationFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("configuration")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("ERROR: could not read config: %s\n", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("ERROR: unable to decode into struct: %v\n", err)
	}
}

type Utility func()

var UtilityMap = map[string]Utility{
	"hello":   Hello,
	"goodbye": Goodbye,
}

func Hello() {
	fmt.Printf("Hello %s!\n", configuration.Name)
}

func Goodbye() {
	fmt.Printf("Goodbye %s!\n", configuration.Name)
}
