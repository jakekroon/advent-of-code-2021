crabs <- scan("puzzle_input.txt", sep=",", quiet=TRUE)

sum(abs(crabs - median(crabs)))
