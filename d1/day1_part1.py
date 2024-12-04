import pandas as pd

input = pd.read_csv("day1_input.csv", sep=" +")
col1 = input['col1'].tolist()
col2 = input['col2'].tolist()

col1.sort()
col2.sort()

total = 0
for i in range(len(col1)):
    total += abs(col2[i] - col1[i])

print(total)