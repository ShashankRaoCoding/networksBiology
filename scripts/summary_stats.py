import pandas
import igraph
import sys

data = pandas.read_csv(sys.stdin) 

vertices = data[["from", "to"]].stack().unique()

print("Vertices: ", len(vertices)) 
print("Edges: ", len(data)) 










































