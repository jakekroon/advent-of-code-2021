library(dplyr)

input <- read.fwf('puzzle_input.txt', widths=rep(1, 12))

mostCommonBits <- input
for (ii in 1:NCOL(input)) {
  mostCommonBits <- mostCommonBits %>% filter(tail(names(sort(table(mostCommonBits[,ii]))), 1) == mostCommonBits[,ii])
}

leastCommonBits <- input
for (ii in 1:NCOL(input)) {
  leastCommonBits <- leastCommonBits %>% filter(head(names(sort(table(leastCommonBits[,ii]))), 1) == leastCommonBits[,ii])
}

strtoi(paste(mostCommonBits, collapse=''), base=2) * strtoi(paste(leastCommonBits, collapse=''), base=2)
