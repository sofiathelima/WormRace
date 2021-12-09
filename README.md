# WormRace
 Simulating motor effects of genetic mutations

 My program simulates any number of "competitors" in the worm race on a track length of 100 cm, which can also be adjusted.

My program takes input of two types: (1) txt files containing data of 1 worm, and (2) xlsx files containing data of samples from groups of genotypes.

Dependencies: excelize package on github.com to access xlsx files

To run my program, use the following example command line with -ftype flag:
-ftype txt DATA1 DATA2 DATA3 ...
-ftype xlsx project.xlsx

The expected output will be: printed updates in the console, and an animation gif of the simulated race.

Changes to my project from what I originally proposed: I was unable to compute statistical significance between groups due to time constraints. I spent more time than expected trying to orient myself to GitHub version control practice, which was an added component to my project.