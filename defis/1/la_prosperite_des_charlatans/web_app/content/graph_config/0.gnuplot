#!/gnuplot

set term png
set output "../graph_render/0.png"

set xdata time
set timefmt '%Y-%m-%d %H:%M:%S'

plot "../graph_dataset/0.graph-dataset" using 1:4 title 'Carmack'  with lines , \
"../graph_dataset/0.graph-dataset" using 1:5 title 'GroupTest'  with lines, \
"../graph_dataset/0.graph-dataset" using 1:6 title 'Scableur'  with lines, \
"../graph_dataset/0.graph-dataset" using 1:7 title 'babar'  with lines, \
"../graph_dataset/0.graph-dataset" using 1:8 title 'nw-group' with lines
