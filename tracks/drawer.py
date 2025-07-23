# This is a temporary file that I use to detect whether
# my program creates correct track map or no

import os

import matplotlib.pyplot as plt
import pandas as pd

script_dir = os.path.dirname(os.path.abspath(__file__))
file_path = os.path.join(script_dir, "generated_tracks", "bahrain_2024_racingline.txt")

df = pd.read_csv(file_path)
df_closed = pd.concat([df, df.iloc[[0]]], ignore_index=True)

plt.figure(figsize=(8, 6))
plt.plot(df_closed['z'], df_closed['x'], linewidth=1)
plt.xlabel("X position (m)")
plt.ylabel("Z position (m)")
plt.title("Track map (closed loop)")
plt.axis("equal")
plt.grid(True)
plt.show()
