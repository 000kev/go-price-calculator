# Tax Calculator

## A Golang application that takes a list of input prices from either a file or the command line and then returns output files or command line output of the prices inclusive of tax

## Example Usage

Given input prices {10, 20, 30, 40, 50} and tax rates {0, 0.25, 0.5} => the application will produce new adjusted prices for each tax rate:
0 : {10, 20, 30, 40, 50}
0.25: {12.5, 25, 37.5, 50, 62.5}
0.5: {15, 30, 45, 60, 75}

These results are then saved to a JSON file and outputted to the command line.


### Criteria

1. Struct of tax rate, input prices and prices+tax √
2. Struct constructor √
3. Struct method to execute calculation logic √
4. Struct method to load data from a file with os, bufio, strconv (filemanager package) √
5. Struct method to write data to  JSON files with encoding/json for each tax rate (filemanager package) √
6. Convert strings to floats (conversion package) √
7. Use a FileManager struct and a CommandLineManager struct with the use of an IOManager interface √
8. Use struct tags √
9. Error handling
