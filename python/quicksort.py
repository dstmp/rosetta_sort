def quicksort(sortable):
    """Return the sorted list

    >>> quicksort(['zebra', 'aardvark', 'coyote', 'abalone', '123', '456'])
    ['123', '456', 'aardvark', 'abalone', 'coyote', 'zebra']

    Given a list including duplicates
    >>> quicksort(['12', '12', 'zebra', 'aardvark', 'cat', 'cat', 'abalone'])
    ['12', '12', 'aardvark', 'abalone', 'cat', 'cat', 'zebra']

    Given a list that starts out already sorted
    >>> quicksort(['aardvark', 'abalone', 'coyote', 'zebra'])
    ['aardvark', 'abalone', 'coyote', 'zebra']

    Given a list that starts out reversed
    >>> quicksort(['zebra', 'coyote', 'abalone', 'aardvark'])
    ['aardvark', 'abalone', 'coyote', 'zebra']

    Given a single element list
    >>> quicksort(['zebra'])
    ['zebra']

    Given an empty list
    >>> quicksort([])
    []
    """

    __quicksort(sortable, 0, len(sortable) - 1)
    return sortable

def __quicksort(sortable, first, last):
  if first >= last:
    return

  pivot = __median_of_three(sortable, first, last)
  sortable[pivot], sortable[last] = sortable[last], sortable[pivot]

  left = first
  right = last - 1
  while left <= right:
    while left <= last and sortable[left] < sortable[last]:
      left += 1
    while right >= first and sortable[last] < sortable[right]:
      right -= 1
    if left <= right:
      sortable[left], sortable[right] = sortable[right], sortable[left]
      left += 1
      right -= 1
  sortable[left], sortable[last] = sortable[last], sortable[left]

  __quicksort(sortable, first, left - 1)
  __quicksort(sortable, left + 1, last)

def __median_of_three(sortable, left, right):
  center = (left + right)/2
  small = left
  large = right
  if sortable[large] < sortable[small]:
    small, large = large, small
  if sortable[large] < sortable[center]:
    return large
  if sortable[center] < sortable[small]:
    return small
  return center

if __name__ == "__main__":
  import doctest
  doctest.testmod()
