#!/gnuplot

set term png
set output "./web_app/content/graph_render/6.png"
set title "Test 5: change of two formulas F1->F2-> change F2"

set xdata time
set timefmt '%Y-%m-%d %H:%M:%S'

plot './web_app/content/graph_dataset/6.graph-dataset' using 1:4 title 'Carmack' with linespoints, \
  './web_app/content/graph_dataset/6.graph-dataset' using 1:5 title 'GroupTest' with linespoints, \
  './web_app/content/graph_dataset/6.graph-dataset' using 1:6 title 'Scableur' with linespoints, \
  './web_app/content/graph_dataset/6.graph-dataset' using 1:7 title 'babar' with linespoints, \
  './web_app/content/graph_dataset/6.graph-dataset' using 1:8 title 'nw-group' with linespoints
