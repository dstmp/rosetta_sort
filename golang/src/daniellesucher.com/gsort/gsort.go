package main

import(
  "daniellesucher.com/heapsort"
  "daniellesucher.com/quicksort"
  "errors"
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "sort"
  "strings"
)

func main() {
  algorithmPtr := flag.String("algorithm", "heapsort", "heapsort or quicksort")
  flag.Parse()

  bytes, err := ioutil.ReadAll(os.Stdin)

  if err != nil {
    fmt.Fprint(os.Stderr, err)
    return
  }

  if len(bytes) == 0 {
    err = errors.New("You need to pass in something to sort!")
    fmt.Fprint(os.Stderr, err)
    return
  }

  lines := strings.Split(string(bytes), "\n")
  sortable := sort.StringSlice(lines)

  if *algorithmPtr == "heapsort" {
    heapsort.HeapSort(sortable)
  } else if *algorithmPtr == "quicksort" {
    quicksort.QuickSort(sortable)
  } else {
    err = errors.New("The only valid algorithm options are heapsort or quicksort.")
    fmt.Fprint(os.Stderr, err)
    return
  }

  out := strings.Join(sortable, "\n")
  fmt.Fprint(os.Stdout, out)
}
