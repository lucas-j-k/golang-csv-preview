# CSV Preview

Preview CSV files in the command line

## Commands

### Preview all columns
```preview --file ./myfile.csv --row-limit 10```

### Preview all values in specified column(s)
```preview columnOne,columnTwo --file ./myfile.csv --row-limit 10```

### Search for value in specific column
```search columnOne searchTerm --file ./myfile.csv --row-limit 10```
