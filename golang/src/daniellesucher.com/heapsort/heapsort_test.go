package heapsort

import(
  "sort"
  "testing"
)

func TestRandom(t *testing.T) {
  unsorted := []string{"zebra", "aardvark", "coyote", "abalone", "123", "456"}
  sorted := []string{"123", "456", "aardvark", "abalone", "coyote", "zebra"}

  sortable := sort.StringSlice(unsorted)
  HeapSort(sortable)

  if !sort.IsSorted(sortable) {
    t.Errorf("Expected %v, but received %v", sorted, sortable)
  }
}

func TestDuplicates(t *testing.T) {
  unsorted := []string{"12", "12", "zebra", "aardvark", "cat", "cat", "abalone"}
  sorted := []string{"12", "12", "aardvark", "abalone", "cat", "cat", "zebra"}

  sortable := sort.StringSlice(unsorted)
  HeapSort(sortable)

  if !sort.IsSorted(sortable) {
    t.Errorf("Expected %v, but received %v", sorted, sortable)
  }
}

func TestAlreadySorted(t *testing.T) {
  unsorted := []string{"aardvark", "abalone", "coyote", "zebra"}
  sorted := []string{"aardvark", "abalone", "coyote", "zebra"}

  sortable := sort.StringSlice(unsorted)
  HeapSort(sortable)

  if !sort.IsSorted(sortable) {
    t.Errorf("Expected %v, but received %v", sorted, sortable)
  }
}

func TestReversed(t *testing.T) {
  reversed := []string{"zebra", "coyote", "abalone", "aardvark"}
  sorted := []string{"aardvark", "abalone", "coyote", "zebra"}

  sortable := sort.StringSlice(reversed)
  HeapSort(sortable)

  if !sort.IsSorted(sortable) {
    t.Errorf("Expected %v, but received %v", sorted, sortable)
  }
}

func TestInts(t *testing.T) {
  unsorted := []int{6, 3, 9, 1, 2}
  sorted := []int{1, 2, 3, 6, 9}

  sortable := sort.IntSlice(unsorted)
  HeapSort(sortable)

  if !sort.IsSorted(sortable) {
    t.Errorf("Expected %v, but received %v", sorted, sortable)
  }
}

func TestSingleElement(t *testing.T) {
  single := []string{"zebra"}
  sortable := sort.StringSlice(single)
  HeapSort(sortable)

  if !sort.IsSorted(sortable) {
    t.Errorf("Expected %v, but received %v", single, sortable)
  }
}

func TestEmpty(t *testing.T) {
  empty := []string{}
  sortable := sort.StringSlice(empty)
  HeapSort(sortable)

  if !sort.IsSorted(sortable) {
    t.Errorf("Expected %v, but received %v", empty, sortable)
  }
}
