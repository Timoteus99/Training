#!/bin/bash

echo "Testing birthdayboy.py"

for file in birthdayboy-sample*
do
	echo "Dobljen rezultat: $(cat $file | python3 birthdayboy.py) / Pravilen rezultat: $(echo $file | awk -F'_' '{print $2}')"
done
