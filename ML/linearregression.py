import pandas as pd 
import numpy as np
import matplotlib.pyplot as plt
from sklearn import linear_model


df = pd.read_csv("dataset.csv")

#print(df)

plt.scatter(df.Area, df.Price)
#plt.show()

# y = m * x + b

reg = linear_model.LinearRegression()
reg.fit(df[['Area']], df.Price)

print(df[['Area']])


print(reg.predict([[3300]]))
print(reg.coef_) #koeficient m
print(reg.intercept_) #vrednost b