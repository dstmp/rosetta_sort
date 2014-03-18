package quicksort

import("sort")

func QuickSort(sortable sort.Interface){
  quickSort(sortable, 0, sortable.Len() - 1)
}

func quickSort(sortable sort.Interface, first, last int) {
  if first >= last {
    return
  }

  pivot := medianOfThree(sortable, first, last)
  sortable.Swap(pivot, last)

  left := first
  right := last - 1
  for left <= right {
    for left <= last && sortable.Less(left, last) {
      left++
    }
    for right >= first && sortable.Less(last, right) {
      right--
    }
    if left <= right {
      sortable.Swap(left, right)
      left++
      right--
    }
  }
  sortable.Swap(left, last)

  quickSort(sortable, first, left - 1)
  quickSort(sortable, left + 1, last)
}

func medianOfThree(sortable sort.Interface, left, right int) int {
  center := (left + right)/2
  small := left
  large := right
  if sortable.Less(large, small) {
    small, large = large, small
  }
  if sortable.Less(large, center) {
    return large
  }
  if sortable.Less(center, small) {
    return small
  }
  return center
}
