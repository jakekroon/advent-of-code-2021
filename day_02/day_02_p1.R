input <- read.table('puzzle_input.txt')

sum(input[input$V1 == 'forward',]$V2) *
  (sum(input[input$V1 == 'down',]$V2) -
   sum(input[input$V1 == 'up',]$V2))
