package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/audstanley/DnD-Name-Generator-Binary/cmd/generator"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var EnableDebug bool
var CfgFile string

// Config stores configuration values
type Config struct {
	Names []string `mapstructure:"names"`
}

var Count int

// Add Generator command to handle the logic
var Generator = &cobra.Command{
	Use:   "generator",
	Short: "Generate cmd/generator/names.go file from generator/names.txt file (for development purposes)",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Generating names.go file...")
		err := generator.Generate() // Generate the names.go file
		if err != nil {
			return err
		}
		return nil
	},
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "DnD Name Generator",
	Short: "Generate names for Dungeons and Dragons characters.",
	Long:  `DnD Name Generator is a CLI application that generates names for Dungeons and Dragons characters.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if EnableDebug { // Enable debug mode if `--debug=true` or `DEBUG=true`.
			fmt.Println("Debug mode enabled")
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(generator.NamesList) == 0 {
			return fmt.Errorf("names array is empty")
		}
		// Seed the random number generator
		// Use a new source for randomness (recommended)
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)

		for i := 0; i < Count; i++ {
			// Pick two random indices
			index1 := random.Intn(len(generator.NamesList))
			index2 := random.Intn(len(generator.NamesList))

			// Ensure indices are unique (avoid duplicates)
			for index1 == index2 {
				index2 = random.Intn(len(generator.NamesList))
			}

			fmt.Printf("%s %s\n", generator.NamesList[index1], generator.NamesList[index2])
		}
		return nil
	},
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra-example" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".dnd-names")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
