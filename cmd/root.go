package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Config stores configuration values
type Config struct {
	Names []string `mapstructure:"names"`
}

var count int

// Define the flag for number of times to print
//rootCmd.Flags().IntVarP(&count, "number", "n", 1, "Number of times to print DnD Names from a huge list")

// Add a command to handle the logic
var Number = &cobra.Command{
	Use:   "names",
	Short: "Pick random elements from the names array",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(NamesList) == 0 {
			return fmt.Errorf("names array is empty")
		}

		// Seed the random number generator
		// Use a new source for randomness (recommended)
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)

		for i := 0; i < count; i++ {
			// Pick two random indices
			index1 := random.Intn(len(NamesList))
			index2 := random.Intn(len(NamesList))

			// Ensure indices are unique (avoid duplicates)
			for index1 == index2 {
				index2 = random.Intn(len(NamesList))
			}

			fmt.Printf("%s %s\n", NamesList[index1], NamesList[index2])
		}
		return nil
	},
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "DnD Name Generator",
	Short: "Generate names for Dungeons and Dragons characters.",
	Long:  `DnD Name Generator is a CLI application that generates names for Dungeons and Dragons characters.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println(NamesList)
		if len(NamesList) == 0 {
			return fmt.Errorf("names array is empty")
		}

		// Seed the random number generator
		// Use a new source for randomness (recommended)
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)

		for i := 0; i < count; i++ {
			// Pick two random indices
			index1 := random.Intn(len(NamesList))
			index2 := random.Intn(len(NamesList))

			// Ensure indices are unique (avoid duplicates)
			for index1 == index2 {
				index2 = random.Intn(len(NamesList))
			}

			fmt.Printf("%s %s\n", NamesList[index1], NamesList[index2])
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println("Executing RootCmd")
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(Number)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dnd-names.yaml)")
	//RootCmd.PersistentFlags().StringArrayP("number", "n", []string{"1"}, "Number of names to generate")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootCmd.Flags().IntVarP(&count, "number", "n", 1, "Number of names to generate")
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
