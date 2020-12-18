from booking_scraper import bkscraper

result = bkscraper.get_result(city="Ljubljana", limit=40)


f = open('bookingscrape.txt', 'w')

f.write(str(result))

f.close()

sum_prices = 0

sum_offers = 0

for element in result:
    for lst in element:
        for price in lst:
            sum_prices += int(lst['price'])
            sum_offers += 1

print('Average price: ' + str(sum_prices / sum_offers))