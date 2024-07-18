package generator

import (
	"bufio"
	"fmt"

	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/goccy/go-yaml"
)

var EnableDebug bool

// Names represents the data used for generating the file
type Names struct {
	PackageName string
	Names       []string
}

type SpeciesData struct {
	// Other fields in the struct (if any)
	Species map[string]SpeciesInfo `yaml:"Species"` // Assuming data comes from a YAML file
}

type NamesTypes struct {
	FirstFNames  []string `yaml:"FirstFNames,omitempty"`  // Field for list
	FirstMNames  []string `yaml:"FirstMNames,omitempty"`  // Field for list
	FirstNBNames []string `yaml:"FirstNBNames,omitempty"` // Field for list
	Special      []string `yaml:"Special,omitempty"`      // Field for list
	Last         []string `yaml:"Last,omitempty"`         // Field for list
}

type SpeciesInfo struct {
	NameOfSpecies      string     `yaml:"NameOfSpecies"`           // Field for species name
	GlobalVariableName string     `yaml:"GlobalVariableName"`      // Field for global variable name
	NameOrder          []string   `yaml:"NameOrder,omitempty"`     // Field for order of names
	Variable           string     `yaml:"Variable,omitempty"`      // Field for your string value
	TextFileNames      []string   `yaml:"TextFileNames,omitempty"` // Field for list of text files
	Names              NamesTypes `yaml:"Names,omitempty"`         // Field for list of names
	// Other fields specific to species data (if any)
}

func printSpeciesData(speciesData []SpeciesInfo) {
	for i, species := range speciesData {
		fmt.Printf("%d: Name: %s, Variable: %s, Files: %v\n", i, species.NameOfSpecies, species.Variable, species.TextFileNames)
		fmt.Printf("  FirstFNames: %v\n", species.Names.FirstFNames)
		fmt.Printf("  FirstMNames: %v\n", species.Names.FirstMNames)
		fmt.Printf("  FirstNBNames: %v\n", species.Names.FirstNBNames)
		fmt.Printf("  Special: %v\n", species.Names.Special)
		fmt.Printf("  Last: %v\n", species.Names.Last)
	}
}

func readNamesFromFile() ([]string, error) {
	filename := "cmd/generator/names.txt"
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var names []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return names, nil
}

func generateNamesFile(names []string) error {
	// Trim whitespace from each name, just in case.
	for i, name := range names {
		names[i] = strings.TrimSpace(name)
	}

	var newNames []string
	currentGroup := []string{}
	for i, name := range names {
		isLastElement := len(names) == i+1
		isLastInGroup := len(currentGroup) == 4 || isLastElement // Check for last in group or last element overall

		currentGroup = append(currentGroup, fmt.Sprintf("\"%s\"", name))

		if isLastInGroup {
			newNames = append(newNames, strings.Join(currentGroup, ",")+",")
			currentGroup = []string{} // Clear the current group for the next
			if !isLastElement {       // Add newline and indentation if not the absolute last element
				newNames = append(newNames, "\n\t\t")
			}
		}
	}

	// Define the template for the generated file (same as before)
	tmpl := template.Must(template.New("names").Parse(`package {{.PackageName}}
// This file is generated by the generator tool; DO NOT EDIT

var NamesList = []string{
	{{- range .Names }}{{.}}{{end -}}} // List of names for DnD characters
`))

	// Define the data for the template
	data := Names{
		PackageName: "generator",
		Names:       newNames,
	}

	// Create the output file (same as before)
	f, err := os.OpenFile("cmd/generator/names.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Generate the file content and write it to the output file (same as before)
	return tmpl.Execute(f, data)
}

// Generate the names.go file by reading names from a file
func Generate(enableDebug bool) error {
	EnableDebug = enableDebug
	if EnableDebug {
		fmt.Println("Debug mode enabled from generator function")
	}
	names, err := readNamesFromFile()
	if err != nil {
		return err
	}

	err = generateNamesFile(names)
	if err != nil {
		return err
	} else {
		fmt.Println("Generated cmd/generator/names.go file successfully")
	}

	// Generate all the text files for the species
	fmt.Println("Generating species text files, if they don't already exist...")
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)
	os.Chdir("cmd/generator")
	speciesStructure, err := ProcessYAMLAndCreateFiles("names.yaml")
	if err != nil {
		os.Chdir(cwd)
		return err
	}
	speciesStructure, err = OrgnaizeSpeciesData(speciesStructure)
	if err != nil {
		fmt.Println("Error organizing species data:", err)
		return err
	}

	os.Chdir(cwd)

	GenerateSpecies(speciesStructure)

	return nil
}

// Generate cmd/species/species.go file species template files
func GenerateSpecies(speciesStructure []SpeciesInfo) error {
	if (speciesStructure == nil) || (len(speciesStructure) == 0) {
		return fmt.Errorf("speciesStructure is empty")
	} else {
		// Parse templates
		templSpeciesCommand, err := template.ParseFiles("cmd/species/species.tmpl")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			panic(err)
		}

		// Create a file for output
		f, err := os.Create("cmd/species/species-test.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
			panic(err)
		}
		defer f.Close()
		if EnableDebug {
			printSpeciesData(speciesStructure)
		}

		// Execute the template and write to the file
		err = templSpeciesCommand.Execute(f, speciesStructure)
		if err != nil {
			fmt.Println("Error executing template:", err)
			panic(err)
		}
		return nil
	}
}

// OrgnaizeSpeciesData function organizes an array of SpeciesInfo structs by
// looping throught each SpeciesInfo.TextFileNames and passing the files to
// AlphabetizeAndTrimFile
func OrgnaizeSpeciesData(speciesStructure []SpeciesInfo) ([]SpeciesInfo, error) {
	if EnableDebug {
		fmt.Println("\n\nOrganizing species data...")
	}

	for i, species := range speciesStructure {
		if EnableDebug {
			fmt.Printf("  Organizing species %s... %v\n", species.NameOfSpecies, species.TextFileNames)
		}
		for j, textFileName := range species.TextFileNames {
			sortedSpecies, err := AlphabetizeAndTrimFile(textFileName, species)
			if err != nil {
				return nil, fmt.Errorf("error alphabetizing and trimming file %s: %w", textFileName, err)
			}
			switch species.NameOrder[j] {
			case "FirstFNames":
				speciesStructure[i].Names.FirstFNames = sortedSpecies
			case "FirstMNames":
				speciesStructure[i].Names.FirstMNames = sortedSpecies
			case "FirstNBNames":
				speciesStructure[i].Names.FirstNBNames = sortedSpecies
			case "Special":
				speciesStructure[i].Names.Special = sortedSpecies
			case "Last":
				speciesStructure[i].Names.Last = sortedSpecies
			}
		}
	}
	return speciesStructure, nil
}

// ProcessSpeciesData function processes the provided YAML data and creates folders and text files
func ProcessSpeciesDataForFolderStructredTextData(speciesMap map[string]interface{}) (speciesStructure []SpeciesInfo, error error) {
	fmt.Println("Processing species data, Creating folders and text files, and Sorting SpeciesStructure by Variable...")
	if EnableDebug {
		fmt.Println("Species data:", speciesMap)
	}

	for species, v := range speciesMap {
		nameOfSpecies := v.(map[string]interface{})["NameOfSpecies"].(string)
		globalVariableName := strings.ReplaceAll(nameOfSpecies, "-", "")
		globalVariableName = strings.ReplaceAll(globalVariableName, " ", "")

		speciesStructure = append(speciesStructure, SpeciesInfo{
			NameOfSpecies:      nameOfSpecies,
			GlobalVariableName: globalVariableName,
			Variable:           v.(map[string]interface{})["Variable"].(string),
			TextFileNames:      []string{},
		})
		latestSpecies := speciesStructure[len(speciesStructure)-1]
		latestSpecies.NameOrder = latestSpecies.TextFileNames
		for _, textFileName := range v.(map[string]interface{})["TextFileNames"].([]interface{}) {
			filepath := latestSpecies.NameOfSpecies + "/" + latestSpecies.NameOfSpecies + "-" + textFileName.(string) + ".txt"
			latestSpecies.TextFileNames = append(latestSpecies.TextFileNames, filepath)
			latestSpecies.NameOrder = append(latestSpecies.NameOrder, textFileName.(string))
		}

		if EnableDebug {
			fmt.Println("species:", species,
				"globalVariableName:", globalVariableName,
				"variable:", v.(map[string]interface{})["Variable"],
				"textFileName:", latestSpecies.TextFileNames,
				"NameOrder:", latestSpecies.NameOrder)
		}

		speciesStructure[len(speciesStructure)-1] = latestSpecies

		// Create folder for the species
		if err := os.MkdirAll(species, os.ModePerm); err != nil {
			switch {
			case os.IsExist(err):
				if EnableDebug {
					fmt.Printf("\ncmd/generator/%s already exists. Skipping creation.\n", species)
				}
			default:
				return nil, fmt.Errorf("error creating folder %s: %w", species, err)
			}
		} else {
			if EnableDebug {
				fmt.Printf("\ncmd/generator/%s folder available\n", species)
			}
		}

		for _, textFileNames := range latestSpecies.TextFileNames {
			// Check if the file already exists
			_, err := os.Stat(textFileNames)
			if err != nil && !os.IsNotExist(err) {
				// Handle other errors
				return nil, fmt.Errorf("error checking file %s: %w", textFileNames, err)
			}

			// If file doesn't exist, create it
			if err == nil {
				if EnableDebug {
					fmt.Printf("  cmd/generator/%s already exists. Skipping creation.\n", textFileNames)
				}
				continue // Skip to next set of textFileNames
			}

			// Create empty text file
			err = os.WriteFile(textFileNames, []byte{}, 0644)
			if err != nil && !os.IsNotExist(err) { // Check for specific os.IsNotExist error
				fmt.Printf("Error creating file %s: %v\n", textFileNames, err)
			}
		}

		sort.Slice(speciesStructure, func(i, j int) bool {
			return speciesStructure[i].Variable < speciesStructure[j].Variable
		})
	}
	return speciesStructure, nil
}

// ProcessYAMLAndCreateFiles function opens, parses YAML data and creates folders/files
func ProcessYAMLAndCreateFiles(filePath string) (SpeciesStructure []SpeciesInfo, error error) {
	// Read YAML file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file %s: %w", filePath, err)
	}

	// Parse YAML data into a map
	var speciesData map[string]map[string]interface{}
	err = yaml.Unmarshal(data, &speciesData)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML data: %w", err)
	}

	speciesMap := speciesData["Species"] // Access the map associated with "Species" key

	// Now you can use speciesMap as a map[string]interface{}
	if speciesMap != nil {
		// Get keys from the speciesMap (using loop)
		speciesKeys := make([]string, 0, len(speciesMap))
		for key := range speciesMap {
			speciesKeys = append(speciesKeys, key)
		}
		if EnableDebug {
			fmt.Printf("speciesKeys: %v\n", speciesKeys)
		}
	} else {
		fmt.Println("Species data not found")
	}

	// Call ProcessSpeciesData to create folders and files
	return ProcessSpeciesDataForFolderStructredTextData(speciesMap)
}

// AlphabetizeAndTrimFile reads a text file, sorts lines alphabetically, trims whitespace, and rewrites the file
// and returns the sorted species data as a slice of strings
func AlphabetizeAndTrimFile(filepath string, species SpeciesInfo) ([]string, error) {
	if EnableDebug {
		fmt.Printf("  Alphabetizing and trimming file %s...\n", filepath)
	}

	// Open the file for reading
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Ensure file is closed

	// Use bufio.Scanner for efficient line-by-line processing
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines) // Split on newlines

	// Use a map to store unique lines
	uniqueLines := make(map[string]bool)
	for scanner.Scan() {
		// Trim leading and trailing whitespace from each line
		line := strings.TrimSpace(scanner.Text())
		if line != "" { // Add only non-empty lines
			uniqueLines[line] = true // Add to map if unique
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Convert map to a slice for sorting
	var lines []string
	for line := range uniqueLines {
		lines = append(lines, line)
	}

	// Sort the trimmed lines alphabetically
	sort.Strings(lines)

	// Join the sorted lines back into a string with newline characters
	sortedSpeciesData := strings.Join(lines, "\n")

	// Overwrite the file with the sorted and trimmed data
	err = os.WriteFile(filepath, []byte(sortedSpeciesData), 0644) // Adjust permissions as needed
	if err != nil {
		return nil, err
	}

	return lines, nil
}
