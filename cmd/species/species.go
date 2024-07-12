package species

import (
	"fmt"
	"os"

	"github.com/audstanley/DnD-Name-Generator-Binary/cmd"
	"github.com/spf13/cobra"
)

var fisrtFName bool
var fisrtMName bool
var firstNBName bool
var lastName bool
var specialName bool
var locationName bool

var aarakocra bool
var aasimar bool
var autognome bool
var bugbear bool
var centaur bool
var changeling bool
var dhampir bool
var dragonborn bool
var dwarf bool
var elf bool
var fairy bool
var firbolg bool
var genasi bool
var giff bool
var gith bool
var gnome bool
var goblin bool
var goliath bool
var halfElf bool
var halfOrc bool
var halfling bool
var harengon bool
var hexblood bool
var hobgoblin bool
var human bool
var kalashtar bool
var kender bool
var kenku bool
var kobold bool
var leonin bool
var lexodon bool
var lizardfolk bool
var minotaur bool
var orc bool
var owlin bool
var plasmoid bool
var reborn bool
var satyr bool
var shifter bool
var simicHybrid bool
var tabaxi bool
var thriKreen bool
var tiefling bool
var tortle bool
var triton bool
var vendalken bool

// Add SpeciesCommand to handle the logic
var SpeciesCommand = &cobra.Command{
	Use:   "species",
	Short: "Generate names for Dungeons and Dragons species based on the species name.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Generating names for Dungeons and Dragons species...")
		return nil
	},
}

func init() {
	cobra.OnInitialize(cmd.InitConfig)
	cmd.RootCmd.AddCommand(cmd.Generator)
	cmd.RootCmd.AddCommand(SpeciesCommand)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cmd.RootCmd.PersistentFlags().StringVar(&cmd.CfgFile, "config", "", "config file (default is $HOME/.dnd-names.yaml)")
	//RootCmd.PersistentFlags().StringArrayP("number", "n", []string{"1"}, "Number of names to generate")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cmd.RootCmd.PersistentFlags().BoolVar(&cmd.EnableDebug, "debug", os.Getenv("DEBUG") == "true", "Enable debug mode")
	cmd.RootCmd.Flags().IntVarP(&cmd.Count, "number", "n", 1, "Number of names to generate")
	//RootCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate names from the generator/names.txt file, outputs go code to cmd/names.go")
	SpeciesCommand.PersistentFlags().BoolVar(&aarakocra, "aarakocra", false, "generate aarakocra names")
	SpeciesCommand.PersistentFlags().BoolVar(&aasimar, "aasimar", false, "generate aasimar names")
	SpeciesCommand.PersistentFlags().BoolVar(&autognome, "autognome", false, "generate autognome names")
	SpeciesCommand.PersistentFlags().BoolVar(&bugbear, "bugbear", false, "generate bugbear names")
	SpeciesCommand.PersistentFlags().BoolVar(&centaur, "centaur", false, "generate centaur names")
	SpeciesCommand.PersistentFlags().BoolVar(&changeling, "changeling", false, "generate changeling names")
	SpeciesCommand.PersistentFlags().BoolVar(&dhampir, "dhampir", false, "generate dhampir names")
	SpeciesCommand.PersistentFlags().BoolVar(&dragonborn, "dragonborn", false, "generate dragonborn names")
	SpeciesCommand.PersistentFlags().BoolVar(&dwarf, "dwarf", false, "generate dwarf names")
	SpeciesCommand.PersistentFlags().BoolVar(&elf, "elf", false, "generate elf names")
	SpeciesCommand.PersistentFlags().BoolVar(&fairy, "fairy", false, "generate fairy names")
	SpeciesCommand.PersistentFlags().BoolVar(&firbolg, "firbolg", false, "generate firbolg names")
	SpeciesCommand.PersistentFlags().BoolVar(&genasi, "genasi", false, "generate genasi names")
	SpeciesCommand.PersistentFlags().BoolVar(&giff, "giff", false, "generate giff names")
	SpeciesCommand.PersistentFlags().BoolVar(&gith, "gith", false, "generate gith names")
	SpeciesCommand.PersistentFlags().BoolVar(&gnome, "gnome", false, "generate gnome names")
	SpeciesCommand.PersistentFlags().BoolVar(&goblin, "goblin", false, "generate goblin names")
	SpeciesCommand.PersistentFlags().BoolVar(&goliath, "goliath", false, "generate goliath names")
	SpeciesCommand.PersistentFlags().BoolVar(&halfElf, "half-elf", false, "generate half-elf names")
	SpeciesCommand.PersistentFlags().BoolVar(&halfOrc, "half-orc", false, "generate half-orc names")
	SpeciesCommand.PersistentFlags().BoolVar(&halfling, "halfling", false, "generate halfling names")
	SpeciesCommand.PersistentFlags().BoolVar(&harengon, "harengon", false, "generate harengon names")
	SpeciesCommand.PersistentFlags().BoolVar(&hexblood, "hexblood", false, "generate hexblood names")
	SpeciesCommand.PersistentFlags().BoolVar(&hobgoblin, "hobgoblin", false, "generate hobgoblin names")
	SpeciesCommand.PersistentFlags().BoolVar(&human, "human", false, "generate human names")
	SpeciesCommand.PersistentFlags().BoolVar(&kalashtar, "kalashtar", false, "generate kalashtar names")
	SpeciesCommand.PersistentFlags().BoolVar(&kender, "kender", false, "generate kender names")
	SpeciesCommand.PersistentFlags().BoolVar(&kenku, "kenku", false, "generate kenku names")
	SpeciesCommand.PersistentFlags().BoolVar(&kobold, "kobold", false, "generate kobold names")
	SpeciesCommand.PersistentFlags().BoolVar(&leonin, "leonin", false, "generate leonin names")
	SpeciesCommand.PersistentFlags().BoolVar(&lexodon, "lexodon", false, "generate lexodon names")
	SpeciesCommand.PersistentFlags().BoolVar(&lizardfolk, "lizardfolk", false, "generate lizardfolk names")
	SpeciesCommand.PersistentFlags().BoolVar(&minotaur, "minotaur", false, "generate minotaur names")
	SpeciesCommand.PersistentFlags().BoolVar(&orc, "orc", false, "generate orc names")
	SpeciesCommand.PersistentFlags().BoolVar(&owlin, "owlin", false, "generate owlin names")
	SpeciesCommand.PersistentFlags().BoolVar(&plasmoid, "plasmoid", false, "generate plasmoid names")
	SpeciesCommand.PersistentFlags().BoolVar(&reborn, "reborn", false, "generate reborn names")
	SpeciesCommand.PersistentFlags().BoolVar(&satyr, "satyr", false, "generate satyr names")
	SpeciesCommand.PersistentFlags().BoolVar(&shifter, "shifter", false, "generate shifter names")
	SpeciesCommand.PersistentFlags().BoolVar(&simicHybrid, "simic-hybrid", false, "generate simic-hybrid names")
	SpeciesCommand.PersistentFlags().BoolVar(&tabaxi, "tabaxi", false, "generate tabaxi names")
	SpeciesCommand.PersistentFlags().BoolVar(&thriKreen, "thri-kreen", false, "generate thri-kreen names")
	SpeciesCommand.PersistentFlags().BoolVar(&tiefling, "tiefling", false, "generate tiefling names")
	SpeciesCommand.PersistentFlags().BoolVar(&tortle, "tortle", false, "generate tortle names")
	SpeciesCommand.PersistentFlags().BoolVar(&triton, "triton", false, "generate triton names")
	SpeciesCommand.PersistentFlags().BoolVar(&vendalken, "vendalken", false, "generate vendalken names")
}
