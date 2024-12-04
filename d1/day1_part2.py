import pandas as pd

input = pd.read_csv("day1_input.csv", sep=" +")
col1 = input['col1'].tolist()
col2 = input['col2'].tolist()

total = 0
for val in col1:
    filtered = [x for x in col2 if x == val]
    total += val * len(filtered)

print(total)