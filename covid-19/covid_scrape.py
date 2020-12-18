import pandas

url = 'https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/06-29-2020.csv'
df = pandas.read_csv(url, error_bad_lines=False)

print(df)