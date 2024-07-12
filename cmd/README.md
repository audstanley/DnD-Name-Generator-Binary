### RootCmd

if you run

```bash
./dnd-names -n 10
```

You'll get a list or first and last names that come from [cmd/generator/names.go](https://github.com/audstanley/DnD-Name-Generator-Binary/tree/main/generator) which is generated from a long list of text: 
[cmd/generator/names.txt](https://github.com/audstanley/DnD-Name-Generator-Binary/blob/main/cmd/generator/names.txt)

### List of Subcommands:

```bash
./dnd-names generator
./dnd-names species
```

#### Generator Command

Will generate files, folders, based on [cmd/generator/names.yaml](https://github.com/audstanley/DnD-Name-Generator-Binary/blob/main/cmd/generator/names.yaml)
That same file will also generate variables for the file [cmd/generator/names.go](https://github.com/audstanley/DnD-Name-Generator-Binary/blob/main/cmd/generator/names.go).
The large list of names was essentially phase 1 of this while project. Just generate a bunch of names. Generator also does some extra work for phase 2, will be
to generate (cmd/species/species.go)[https://github.com/audstanley/DnD-Name-Generator-Binary/blob/main/cmd/species/species.go]. The current state of the species.go file
is that it's handwritten, but for ease of use, a new species could be added to cmd/generator/names.yaml, run the generator, modify the species text files.
The generator can be run again, and the flag associated with that species will be generated as code. The generator does all the work, and just the text files
need to be edited to keep things looking great. Generator also trims whitespace, before and after each line in all the text files, and alphebatizes everyting.

So if new names get added to the textfiles:

```bash
./dnd-names generator
# will clean up all of the species text files before pushing back up to github.
```


#### Species Command

Species Command will be generated from the Generator Command. This will allow for an unlimited number of species to be added over time.
The hard part will be to get contributors to add names.

```bash
./dnd-names species --dragonborn --location --number 10
# should eventually produce something like John Doe of Baulder's Gate - with 9 additional generated names.
./dnd-names species --tabaxi --special --location --number 1
# will eventually produce something like Jane Doe, Clan of Riverrunners, of Neverwinter
```

#### Future Goals after completing:
* Ai tools for creating character importabe [FoundryVTT](https://foundryvtt.com/) stat sheets based on character names using the [5e-statblock-importer](https://foundryvtt.com/packages/5e-statblock-importer).
* Going to see if I can figure out [this kind of model](https://cros.land/ai-powered-dnd-5e-monster-statblock-generator/), and create an API alongside llama to create perfect stat sheets.