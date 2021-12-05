is_bingo <- function(checks, dimension)
  colSums(marginSums(checks, c(1, 3)) == dimension) |
  colSums(marginSums(checks, c(2, 3)) == dimension)

play_bingo <- function(nums, bingoBoards) {
  # transform the bingo board data into 5x5 matrices
  dim(bingoBoards) <- c(5, 5, length(bingoBoards) / 5^2)
  # create a accompanying 5x5 check matrices
  checks <- array(FALSE, dim(bingoBoards))
  # init a dummy board
  bingo <- logical()
  
  for (num in nums) {
    checks[bingoBoards == num] <- TRUE
    prev_bingo <- bingo
    bingo <- is_bingo(checks, 5)
    
    if (all(bingo)) {
      bingo <- !prev_bingo
      break
    }
  }
  
  num * sum(bingoBoards[, , bingo][!checks[, , bingo]])
}

nums <- scan("puzzle_input.txt", nlines = 1, sep = ",", quiet = TRUE)
bingoBoards <- scan("puzzle_input.txt", skip=1, quiet=TRUE)

play_bingo(nums, bingoBoards)
