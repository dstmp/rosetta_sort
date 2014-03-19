#!/bin/sh

# user time, system time, realtime, % CPU usage, maximum resident set, command being timed
FORMAT='%Uu %Ss %er %MkB %Pcpu %C'

/usr/bin/time -f "$FORMAT" python ./python/psort.py --algorithm heapsort < $@ > /dev/null
/usr/bin/time -f "$FORMAT" python ./python/psort.py --algorithm quicksort < $@ > /dev/null
/usr/bin/time -f "$FORMAT" ./golang/gsort -algorithm=heapsort < $@ > /dev/null
/usr/bin/time -f "$FORMAT" ./golang/gsort -algorithm=quicksort < $@ > /dev/null
