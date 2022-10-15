# split_pdpqt
Split pdbqt file which contains many models into single file.


## Requirements
- [Go](https://go.dev/)


## Build
```bash
git clone https://github.com/Metaphorme/split_pdpqt.git     # Get split_pdpqt
cd split_pdpqt
go build main.go                                            # Build
```


## Examples
```bash
./main -inputFile BBEDRL.xaa.pdbqt -outputDir output/
./main -inputDir input/ -outputDir output/
./main -inputFile BBEDRL.xaa.pdbqt -inputDir input/ -outputDir output/
```


## Commands
```bash
Usage of ./main:
  -inputDir string
    	The Directory of pdbqt file which contains many models. (default "NoInputDir")
  -inputFile string
    	The pdbqt file which contains many models. (default "NoInputFile")
  -outputDir string
    	The output directory to save output models. (default "./output/")
```


## License
[MIT License](https://github.com/Metaphorme/split_pdpqt/blob/main/LICENSE)
