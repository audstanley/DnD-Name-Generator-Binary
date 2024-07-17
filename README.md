# DnD-Name-Generator-Binary
Generate DnD Names in the console

[Install Go](https://go.dev/doc/install)

### Clone:
```bash
git clone https://github.com/audstanley/DnD-Name-Generator-Binary
cd DnD-Name-Generator-Binary/
go build .
mv DnD-Name-Generator-Binary dnd-names
./dnd-names --number 10
```

### Output Example:
```
Zebulon Phena
Zavor Timoteo
Calaudra Lasla
Cantoria Halogil
Aegon Druse
Skawae Kaggen
Ruven Kaylina
Marinell Odvar
Jereni Chalono
Szanto Vanir
```

### Simple one name generator:

```bash
./dnd-names
```

### Output Example:
```
Nifai Bralas
```

Feel Free to fork, and add new names or new cli features. I've got cobra and viper setup for a config file - perhaps for last names from a yaml file. There are currently no categories/species, but plans for it. I just needed a very simple way to do this in the cli to get stared. Enjoy. No First and Last names either, but plans for it. I may have to devolop a last name feature instead of merging from a large list of generic fantasy names.

### Development:
```bash
# generate go file from names list
# first, add new names to generator/names.txt
# then...
./dnd-names generate
# then
go build -o dnd-names
./dnd-names --number 10
```

if you're forking and adding new names to the giant names list for a pull request - just make changes to the generator/names.txt file.
if you're working on name categories by spiecies (formally "races") using a yaml file, please take a look at the format in names.yaml

```bash
./dnd-names generate # when in the project directory ...
```

will eventually work through the generator/speciesTypes subfolders and generate first and last names based on the text files.
Those text files will be automatically be alphabetised, and converted into golang structs so that **eventually** you'll be able to use

```bash
# (not yet implemented)
./dnd-names --female --dragonborn -n 10
# print 10 female dragonborn names, first and last name.
# or...
./dnd-names --nonbinary --dragonborn -n 8
# print 8 nonbinary dragonborn names, first and last name.

```

See [cmd/generator/README.md](https://github.com/audstanley/DnD-Name-Generator-Binary/tree/main/cmd/generator) since it's in active development.
