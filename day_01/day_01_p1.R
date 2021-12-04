input <- read.table('puzzle_input.txt')$V1
sum(input[-1] - input[-length(input)] > 0)
