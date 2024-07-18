package species

import (
	"fmt"
	"os"

	"github.com/audstanley/DnD-Name-Generator-Binary/cmd"
	"github.com/spf13/cobra"
)

var firstFName bool
var firstMName bool
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

// this function will get the species from a list of species from the
// cmd/species/speciesNames.go file and return a random first and last name
// based on the species
func getSpeciesForSpeciesCommand(species string) {
	fmt.Println("Generating names for " + species)
}

// Add SpeciesCommand to handle the logic
var SpeciesCommand = &cobra.Command{
	Use:   "species",
	Short: "Generate names for Dungeons and Dragons species based on the species name.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("CHECKING FOR GENDER...")
		aarakocra, _ = cmd.Flags().GetBool("aarakocra")
		if aarakocra {
			fmt.Println("Generating aarakocra names")
		}
		aasimar, _ = cmd.Flags().GetBool("aasimar")
		if aasimar {
			fmt.Println("Generating aasimar names")
		}
		autognome, _ = cmd.Flags().GetBool("autognome")
		if autognome {
			fmt.Println("Generating autognome names")
		}
		bugbear, _ = cmd.Flags().GetBool("bugbear")
		if bugbear {
			fmt.Println("Generating bugbear names")
		}
		centaur, _ = cmd.Flags().GetBool("centaur")
		if centaur {
			fmt.Println("Generating centaur names")
		}
		changeling, _ = cmd.Flags().GetBool("changeling")
		if changeling {
			fmt.Println("Generating changeling names")
		}
		dhampir, _ = cmd.Flags().GetBool("dhampir")
		if dhampir {
			fmt.Println("Generating dhampir names")
		}
		dragonborn, _ = cmd.Flags().GetBool("dragonborn")
		if dragonborn {
			fmt.Println("Generating dragonborn names")
		}
		dwarf, _ = cmd.Flags().GetBool("dwarf")
		if dwarf {
			fmt.Println("Generating dwarf names")
		}
		elf, _ = cmd.Flags().GetBool("elf")
		if elf {
			fmt.Println("Generating elf names")
		}
		fairy, _ = cmd.Flags().GetBool("fairy")
		if fairy {
			fmt.Println("Generating fairy names")
		}
		firbolg, _ = cmd.Flags().GetBool("firbolg")
		if firbolg {
			fmt.Println("Generating firbolg names")
		}
		genasi, _ = cmd.Flags().GetBool("genasi")
		if genasi {
			fmt.Println("Generating genasi names")
		}
		giff, _ = cmd.Flags().GetBool("giff")
		if giff {
			fmt.Println("Generating giff names")
		}
		gith, _ = cmd.Flags().GetBool("gith")
		if gith {
			fmt.Println("Generating gith names")
		}
		gnome, _ = cmd.Flags().GetBool("gnome")
		if gnome {
			fmt.Println("Generating gnome names")
		}
		goblin, _ = cmd.Flags().GetBool("goblin")
		if goblin {
			fmt.Println("Generating goblin names")
		}
		goliath, _ = cmd.Flags().GetBool("goliath")
		if goliath {
			fmt.Println("Generating goliath names")
		}
		halfElf, _ = cmd.Flags().GetBool("halfElf")
		if halfElf {
			fmt.Println("Generating halfElf names")
		}
		halfOrc, _ = cmd.Flags().GetBool("halfOrc")
		if halfOrc {
			fmt.Println("Generating halfOrc names")
		}
		halfling, _ = cmd.Flags().GetBool("halfling")
		if halfling {
			fmt.Println("Generating halfling names")
		}
		harengon, _ = cmd.Flags().GetBool("harengon")
		if harengon {
			fmt.Println("Generating harengon names")
		}
		hexblood, _ = cmd.Flags().GetBool("hexblood")
		if hexblood {
			fmt.Println("Generating hexblood names")
		}
		hobgoblin, _ = cmd.Flags().GetBool("hobgoblin")
		if hobgoblin {
			fmt.Println("Generating hobgoblin names")
		}
		human, _ = cmd.Flags().GetBool("human")
		if human {
			fmt.Println("Generating human names")
		}
		kalashtar, _ = cmd.Flags().GetBool("kalashtar")
		if kalashtar {
			fmt.Println("Generating kalashtar names")
		}
		kender, _ = cmd.Flags().GetBool("kender")
		if kender {
			fmt.Println("Generating kender names")
		}
		kenku, _ = cmd.Flags().GetBool("kenku")
		if kenku {
			fmt.Println("Generating kenku names")
		}
		kobold, _ = cmd.Flags().GetBool("kobold")
		if kobold {
			fmt.Println("Generating kobold names")
		}
		leonin, _ = cmd.Flags().GetBool("leonin")
		if leonin {
			fmt.Println("Generating leonin names")
		}
		lexodon, _ = cmd.Flags().GetBool("lexodon")
		if lexodon {
			fmt.Println("Generating lexodon names")
		}
		lizardfolk, _ = cmd.Flags().GetBool("lizardfolk")
		if lizardfolk {
			fmt.Println("Generating lizardfolk names")
		}
		minotaur, _ = cmd.Flags().GetBool("minotaur")
		if minotaur {
			fmt.Println("Generating minotaur names")
		}
		orc, _ = cmd.Flags().GetBool("orc")
		if orc {
			fmt.Println("Generating orc names")
		}
		owlin, _ = cmd.Flags().GetBool("owlin")
		if owlin {
			fmt.Println("Generating owlin names")
		}
		plasmoid, _ = cmd.Flags().GetBool("plasmoid")
		if plasmoid {
			fmt.Println("Generating plasmoid names")
		}
		reborn, _ = cmd.Flags().GetBool("reborn")
		if reborn {
			fmt.Println("Generating reborn names")
		}
		satyr, _ = cmd.Flags().GetBool("satyr")
		if satyr {
			fmt.Println("Generating satyr names")
		}
		shifter, _ = cmd.Flags().GetBool("shifter")
		if shifter {
			fmt.Println("Generating shifter names")
		}
		simicHybrid, _ = cmd.Flags().GetBool("simicHybrid")
		if simicHybrid {
			fmt.Println("Generating simicHybrid names")
		}
		tabaxi, _ = cmd.Flags().GetBool("tabaxi")
		if tabaxi {
			fmt.Println("Generating tabaxi names")
		}
		thriKreen, _ = cmd.Flags().GetBool("thriKreen")
		if thriKreen {
			fmt.Println("Generating thriKreen names")
		}
		tiefling, _ = cmd.Flags().GetBool("tiefling")
		if tiefling {
			fmt.Println("Generating tiefling names")
		}
		tortle, _ = cmd.Flags().GetBool("tortle")
		if tortle {
			fmt.Println("Generating tortle names")
		}
		triton, _ = cmd.Flags().GetBool("triton")
		if triton {
			fmt.Println("Generating triton names")
		}
		vendalken, _ = cmd.Flags().GetBool("vendalken")
		if vendalken {
			fmt.Println("Generating vendalken names")
		}

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
	SpeciesCommand.PersistentFlags().BoolVar(&firstFName, "first", false, "generate a first name, make sure to specify the species")

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
