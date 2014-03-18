package heapsort

import("sort")

func HeapSort(sortable sort.Interface) {
  end := sortable.Len() - 1

  maxHeapify(sortable, end)

  for end > 0 {
    sortable.Swap(end, 0)
    end--
    siftDown(sortable, 0, end)
  }
}

func maxHeapify(sortable sort.Interface, end int) {
  start := parentOf(end)
  for start >= 0 {
    siftDown(sortable, start, end)
    start--
  }
}

func siftDown(sortable sort.Interface, root, end int) {
  parent := root
  child := maxChild(sortable, parent, end)
  for child <= end {
    if !sortable.Less(parent, child) {
      return
    }
    sortable.Swap(parent, child)
    parent = child
    child = maxChild(sortable, parent, end)
  }
}

func maxChild(sortable sort.Interface, parent, end int) int {
  left := parent * 2 + 1
  right := left + 1
  if right <= end && sortable.Less(left, right) {
    return right
  }
  return left
}

func parentOf(node int) int {
  return (node - 1)/2
}
