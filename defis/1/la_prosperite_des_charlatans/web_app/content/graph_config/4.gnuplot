#!/gnuplot

set term png
set output "./web_app/content/graph_render/4.png"
set title "Test 3: user.txt change a formula"

set xdata time
set timefmt '%Y-%m-%d %H:%M:%S'

plot './web_app/content/graph_dataset/4.graph-dataset' using 1:4 title 'Carmack' with linespoints, \
  './web_app/content/graph_dataset/4.graph-dataset' using 1:5 title 'GroupTest' with linespoints, \
  './web_app/content/graph_dataset/4.graph-dataset' using 1:6 title 'Scableur' with linespoints, \
  './web_app/content/graph_dataset/4.graph-dataset' using 1:7 title 'babar' with linespoints, \
  './web_app/content/graph_dataset/4.graph-dataset' using 1:8 title 'nw-group' with linespoints
