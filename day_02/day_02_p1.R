input <- read.table('puzzle_input.txt')

sum(input$V2[input$V1 == 'forward']) *
   (sum(input$V2[input$V1 == 'down']) -
    sum(input$V2[input$V1 == 'up']))
