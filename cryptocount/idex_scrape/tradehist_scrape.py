import requests
import bs4


link = "https://idex.market/eth/idex"

data = {"includeOutliers" : "false", "market" : "ETH_IDEX", "start": 1593704752, "end" : 1593705357688}

page = requests.post(link, data)

soup = bs4.BeautifulSoup(page.content, 'html.parser')

#title = soup.find("meta",  property="og:title")
#url = soup.find("meta",  property="og:url")

#print(title["content"] if title else "No meta title given")
#print(url["content"] if url else "No meta url given")

#for tag in soup.find_all("meta"):
#    if tag.get("property", None) == "og:title":
#        print(tag.get("content", None))
#    elif tag.get("property", None) == "og:url":
#        print(tag.get("content", None))

#f = open("idexhtml.txt", 'w')
#for element in list(soup.children):
#    try:
#        f.write(element.get_text())
#    except:
#        continue
#f.close()

f = open("idexhtml.txt", 'w')
lepse = soup.prettify()
f.write(lepse)
f.close()


#print(lepse)