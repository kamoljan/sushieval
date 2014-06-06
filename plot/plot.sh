#!/bin/sh
SITE=$1

echo 'SITE ' $SITE

echo '\
set terminal png font "/Library/Fonts/Times\ New\ Roman.ttf" 14; \
set output "'$SITE'.jpg"; \
set datafile separator ","; \
set xlabel "Files"; \
set ylabel "Duplicates"; \
plot "'$SITE'.plot" notitle' | gnuplot -persist
