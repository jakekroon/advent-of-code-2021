input <- read.table('puzzle_input.txt')
input$V2[input$V1 == "up"] <- input$V2[input$V1 == "up"] * -1

depth <- 0
aim <- 0

for (ii in seq_along(input$V2)) { 
  if (input$V1[ii] == "forward") {
    depth <- depth + (aim * input$V2[ii])
  } else {
    aim <- aim + input$V2[ii]
  }
}

print(depth * sum(input$V2[input$V1 == 'forward']))
