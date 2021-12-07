crabs <- scan("puzzle_input.txt", sep=",", quiet=TRUE)

gauss <- function(num) {
  sum(num * (num + 1) / 2)
}

min(sapply(0:max(crabs), function(x) {
  return(sum(mapply(function(y, z) {
    return(sum(gauss(abs(y - z))))
  }, x, crabs)))
}))
