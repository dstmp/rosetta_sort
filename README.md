##Description

Command line sort utility that uses either heapsort or quicksort, in both Python and Go.

##Requirements
* Python 2.7
* Go 1.2.1

##Installation

1. Clone this repository!
1. `cd rosetta_sort/golang`
1. `export GOPATH=$PWD`
1. `go build src/daniellesucher.com/gsort/gsort.go`

##Running the tests

Go unit tests, run from the ./golang subdirectory -

1. `export GOPATH=$PWD`
1. `go test daniellesucher.com/heapsort/`
1. `go test daniellesucher.com/quicksort/`

Python unit tests, run from the ./python subdirectory -

1. `python heapsort.py`
1. `python quicksort.py`


##Usage

From the main directory of this repo:

Python:

`python ./python/psort.py --algorithm heapsort < input.txt`
`python ./python/psort.py --algorithm quicksort < input.txt`

Go:

`./golang/gsort -algorithm=heapsort < input.txt`
`./golang/gsort -algorithm=quicksort < input.txt`

##Validation with Gensort/Valsort

1. Install gensort from http://www.ordinal.com/gensort.html
1. Generate an input file with `gensort -a 100000000 input.txt`
1. Pipe the input file into psort or gsort, and pipe the output to output.txt. For example: `./golang/gsort -algorithm=heapsort < input.txt > output.txt`
1. Validate output.txt with `valsort output.txt`


##Assumptions

1. Input will be encoded in ASCII with lines separated by \n.
1. Output is desired in ascending ASCIIbetical order.
