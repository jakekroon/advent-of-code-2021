input <- read.fwf('puzzle_input.txt', widths=rep(1, 12))

gammaRate <- paste(sapply(input, function(x) {
  tail(names(sort(table(x))), 1)
}), collapse = '')

episolonRate <- paste(sapply(input, function(x) {
  head(names(sort(table(x))), 1)
}), collapse = '')

strtoi(gammaRate, base=2) * strtoi(episolonRate, base=2)
