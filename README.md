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

Feel Free to fork, and add new names or new cli features. I've got cobra and viper setup for a config file - perhaps for last names from a yaml file. There are currently no categories, and no plan for that yet. I just needed a very simple way to do this in the cli. Enjoy. No First and Last names Either. I may have to devolop a last name feature instead of merging from a large list of generic fantasy names.