# Tax Calculator

## A Golang application that takes a list of input prices from either a file or the command line and then returns output files or command line output of the prices inclusive of tax


### Criteria

1. Struct of tax rate, input prices and prices+tax
2. Struct constructor
3. Struct method to execute calculation logic
4. Struct method to load data from a file with os, bufio, strconv (filemanager package)
5. Struct method to write data to  JSON files with encoding/json for each tax rate (filemanager package)
6. Convert strings to floats (conversion package)
7. Use a FileManager struct and a CommandLineManager struct with the use of an IOManager interface
8. Use struct tags and exclude file manager from the json file
9. Error handling
