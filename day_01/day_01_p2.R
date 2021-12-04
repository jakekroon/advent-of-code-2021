input <- read.table('puzzle_input.txt')$V1
windows <- rollapply(input, 3, sum)
sum(windows[-1] - windows[-NROW(windows)] > 0)
