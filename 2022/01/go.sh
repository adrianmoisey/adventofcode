#!/bin/bash

count=0

totals=()

while read a; do
	if [ "$a" = "" ] ; then
		totals+=($count)
		count=0
	else
		count=$(( $count + $a ))
	fi

done < input.txt

IFS=$'\n' totals=( $(sort <<<"${totals[*]}") )

top_3=$((${totals[-1]} + ${totals[-2]} + ${totals[-3]}))
echo $top_3
