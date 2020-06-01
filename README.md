# asciigraph

A hard fork of [guptarohit/asciigraph](https://github.com/guptarohit/asciigraph) with some radical changes to render ASCII graphs on terminal UIs.

1. Width and heights provided are fully respected. In the original library, specifying the graphs width does not take into account the space needed for the left axis.
2. The visible unicode width of each rune drawn in the graph are taken into account while drawing against the width and height of the graph.
3. Any expensive calls to `fmt.Sprintf` are removed and replaced with `strconv.FormatFloat`.
4. A slice of runes is kept for drawing the graph vs a slice of strings.

```
 15.00 ┤       ╭────╮                                   ╭────╮                  
 12.86 ┤     ╭─╯    ╰─╮                               ╭─╯    ╰─╮                
 10.71 ┤    ╭╯        ╰╮                             ╭╯        ╰╮               
  8.57 ┤   ╭╯          ╰─╮                          ╭╯          ╰─╮             
  6.43 ┤  ╭╯             ╰╮                        ╭╯             ╰╮            
  4.29 ┤ ╭╯               ╰╮                      ╭╯               ╰╮           
  2.14 ┤╭╯                 ╰╮                    ╭╯                 │           
  0.00 ┼╯                   │                   ╭╯                  ╰╮          
 -2.14 ┤                    ╰╮                 ╭╯                    ╰╮         
 -4.29 ┤                     ╰╮               ╭╯                      ╰╮        
 -6.43 ┤                      ╰╮             ╭╯                        ╰╮       
 -8.57 ┤                       ╰─╮          ╭╯                          ╰─╮     
-10.71 ┤                         ╰╮        ╭╯                             ╰╮    
-12.86 ┤                          ╰─╮    ╭─╯                               ╰─╮  
-15.00 ┤                            ╰────╯                                   ╰─  
```

This fork was done in a rush while building [blanc](https://github.com/lithdew/blanc). If you have any questions, please open a Github issue!