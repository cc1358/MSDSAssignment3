import pandas as pd
import time

N = 100
start_time = time.time()

with open('housesOutputPy.txt', 'wt') as outfile:
    for i in range(N):
        houses = pd.read_csv("housesInput.csv")

    elapsed_time = time.time() - start_time

    outfile.write(houses.describe().to_string(header=True, index=True))
    outfile.write("\n")

# Print the runtime to the CLI
print(f"Script execution time: {elapsed_time} seconds")