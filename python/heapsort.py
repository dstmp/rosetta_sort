def heapsort(sortable):
    """Return the sorted list

    >>> heapsort(['zebra', 'aardvark', 'coyote', 'abalone', '123', '456'])
    ['123', '456', 'aardvark', 'abalone', 'coyote', 'zebra']

    Given a list including duplicates
    >>> heapsort(['12', '12', 'zebra', 'aardvark', 'cat', 'cat', 'abalone'])
    ['12', '12', 'aardvark', 'abalone', 'cat', 'cat', 'zebra']

    Given a list that starts out already sorted
    >>> heapsort(['aardvark', 'abalone', 'coyote', 'zebra'])
    ['aardvark', 'abalone', 'coyote', 'zebra']

    Given a list that starts out reversed
    >>> heapsort(['zebra', 'coyote', 'abalone', 'aardvark'])
    ['aardvark', 'abalone', 'coyote', 'zebra']

    Given a single element list
    >>> heapsort(['zebra'])
    ['zebra']

    Given an empty list
    >>> heapsort([])
    []
    """

    end = len(sortable) - 1
    __max_heapify(sortable, end)

    while end > 0:
      sortable[end], sortable[0] = sortable[0], sortable[end]
      end -= 1
      __sift_down(sortable, 0, end)

    return sortable

def __max_heapify(sortable, end):
  start = __parent_of(end)
  while start >= 0:
    __sift_down(sortable, start, end)
    start -= 1

def __sift_down(sortable, root, end):
  parent = root
  child = __max_child(sortable, parent, end)
  while child <= end:
    if sortable[parent] >= sortable[child]:
      return
    sortable[parent], sortable[child] = sortable[child], sortable[parent]
    parent = child
    child = __max_child(sortable, parent, end)

def __max_child(sortable, parent, end):
  left = parent * 2 + 1
  right = left + 1
  if right <= end and sortable[right] > sortable[left]:
    return right
  return left

def __parent_of(node):
  return (node - 1)/2

if __name__ == "__main__":
  import doctest
  doctest.testmod()
