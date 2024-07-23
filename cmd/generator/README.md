# Generator Function

 - [x] Generate folders based on [Species keys](https://github.com/audstanley/DnD-Name-Generator-Binary/blob/main/cmd/generator/names.yaml)
 - [x] Create text file if they DNE ([example](https://github.com/audstanley/DnD-Name-Generator-Binary/blob/main/cmd/generator/Aarakocra/Aarakocra-MaleFirst.txt))
 - [x] Once the structure for all species and genders are created, parse all text files according to the species in the yaml file.
 - [x] Some species might have 'special names' - I'm thinking Tabaxi Clans - there might be others - another example could be other species from specific areas 'Alice Doe of Baulder's Gate' (of Baulders Gate) could be a special name. These names would appear after last names.

This has been updated - special names will be used for tabaxi and tiefling. I'm sure there are other species that might have special names, but "of Baulder's Gate" will come from a location.yaml file. This way DMs can customize NPC names based on locations in their own world!

 - [ ] I need to scrape - or scour the internet for more names to add to the various species text files.
 - [ ] Contributing Guidelines - I'm thinking if someone wants to add a new species to the list, they need to provide some minimum number of names for the species in the pull request. 300 minimum per gender before merging with master branch. It's not enough to add a species and leave out data that is usable.
 - [x] Write code to generate GoLang Species DataTypes, per Species based on the Species yaml keys in [generator/names.yaml](https://github.com/audstanley/DnD-Name-Generator-Binary/blob/main/cmd/generator/names.yaml)
 - [ ] I may want to have an generator/output.yaml  -for after parsing all the text files - for debugging when creating GoLang Structs (generated code)
 - [ ] figure out how to create a template that can be used in the cmd package that will automatically add new species as flags for cobra/viper (this will be challenging). I might just have to do this manually, though it would be really nice to have this be an automated, generated task. I need to read more docs on Go Templates. This would need to modify the cmd/root.go - or generate the init function, which might need to move to another seperate go file (cmd/init.go) - this might not be so tough.
 - [ ] The output will be an joined string array ["First - Gendered", "Last", "Special"]. Think about how I could merge species, for example:

 ```bash
 # proof of concept
 dnd-names species dragonborn male --number 20;
```

- [ ] Special Names should be specified in the CLI. It's not a priority right now. if --male --dragonborn is picked, Asume --last --dragonborn.
- [ ] Last, but not least, perhaps i can look into creating a web assembly module for node or vinalla JS integration. o.O