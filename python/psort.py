from __future__ import print_function
import argparse
import string
import sys
from heapsort import heapsort
from quicksort import quicksort

def main():
  """Return the sorted list, using the correct algorithm
  """
  parser = argparse.ArgumentParser(description='Sort an array.')
  parser.add_argument('--algorithm', dest='algorithm', default='heapsort',
                      help='quicksort or heapsort (default: heapsort)')
  args = parser.parse_args()

  sortable = sys.stdin.readlines()
  if len(sortable) == 0:
    print('You need to pass in something to sort!', file=sys.stderr)
    return

  if args.algorithm == 'heapsort':
    heapsort(sortable)
  elif args.algorithm == 'quicksort':
    quicksort(sortable)
  else:
    err = 'The only valid algorithm options are heapsort or quicksort.'
    print(err, file=sys.stderr)
    return

  out = ''.join(sortable)
  print(out.rstrip('\n'))

if __name__ == "__main__":
  main()
