import pandas, sys, tqdm 

data = pandas.read_csv(sys.argv[1], sep="\t")

sampids = data["SAMPID"].to_list() 
indices = [0,1] 

for i in sys.stdin:
    # print(i) 
    columns = i.split("\t")
    for i in range(len(columns)):
        indices.append(i) if columns[i] in sampids else None 
    columns = "\t".join(columns[i] for i in indices) 
    print(columns) 
    break 

for i in tqdm.tqdm(sys.stdin):
    vals = i.split("\t")
    line = [vals[i] for i in indices]
    # for i in indices:
        # line.append(vals[i])
    print("\t".join(line)) 

# print(indices) 












































