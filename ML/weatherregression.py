import pandas as pd  
import numpy as np  
import matplotlib.pyplot as plt  
import seaborn as seabornInstance 
from sklearn.model_selection import train_test_split 
from sklearn.linear_model import LinearRegression
from sklearn import metrics


datasets = pd.read_csv('/home/legion/Desktop/3759_5944_bundle_archive/Summary of Weather.csv') #import csv datasets using pandas

datasets.shape #checking the number of rows and columns in our datasets

datasets.describe() #see statistical details of datasets

#plot our data points on a 2-D graph
datasets.plot(x='MinTemp', y='MaxTemp', style='+')  
plt.title('MinTemp vs MaxTemp')  
plt.xlabel('MinTemp')  
plt.ylabel('MaxTemp')  
plt.show()

#Min temp stored in X axis as independet variable
X = datasets['MinTemp'].values.reshape(-1,1)
#Max temp stored in Y axis ad dependent variable
y = datasets['MaxTemp'].values.reshape(-1,1)

#split the data! 80% for training, 20% for testing
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=0)

#nrdim instanco regresorja
regressor = LinearRegression()  
regressor.fit(X_train, y_train) #training the algorithm


#to retrieve the intercept:
print(regressor.intercept_)
#for retrieving the slope:
print(regressor.coef_)

#make predictions and visualize in pandas vs actual values
y_pred = regressor.predict(X_test)

df = pd.DataFrame({'Actual': y_test.flatten(), 'Predicted': y_pred.flatten()})

print(df)

#visualize comparison result as bar graph
df1 = df.head(25)
df1.plot(kind='bar',figsize=(16,10))
plt.grid(which='major', linestyle='-', linewidth='0.5', color='green')
plt.grid(which='minor', linestyle=':', linewidth='0.5', color='black')
plt.show()

#show plotted line (optimized)
plt.scatter(X_test, y_test,  color='gray', marker='+')
plt.plot(X_test, y_pred, color='red', linewidth=1)
plt.show()

#calculate mean errors
print('Mean Absolute Error:', metrics.mean_absolute_error(y_test, y_pred))  
print('Mean Squared Error:', metrics.mean_squared_error(y_test, y_pred))  
print('Root Mean Squared Error:', np.sqrt(metrics.mean_squared_error(y_test, y_pred)))