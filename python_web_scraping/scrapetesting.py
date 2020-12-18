import requests
import random
import bs4


link = "http://quotes.toscrape.com"

link_extensions = ["/tag/love",
                   "/tag/inspirational",
                   "/tag/life",
                   "/tag/humor",
                   "/tag/books",
                   "/tag/reading",
                   "/tag/friendship",
                   "/tag/friends",
                   "/tag/truth",
                   "/tag/simile"]

next_page = "/page/2"

quote_list = []

for extension in link_extensions:
    page = requests.get(link + extension)

    #print(page.content)

    soup = bs4.BeautifulSoup(page.content, 'html.parser')

    #print(soup.prettify())
    #print(list(soup.children))

    #seznam = list(soup.children)
    #print(seznam)
    #print(seznam[2])

    html = list(soup.children)[2]

    body = list(html.children)[3]

    p = list(body.children)[1]

    raw_quote = (p.get_text())

    raw_quote = raw_quote.split('\n')

    i = 0
    while i < len(raw_quote):
        if len(raw_quote[i]) != 0 and raw_quote[i][0] == '“':
            quote_list.append(raw_quote[i] +  '\n' + raw_quote[i + 1])
        i += 1

    try:
        page = requests.get(link + extension + next_page)

        soup = bs4.BeautifulSoup(page.content, 'html.parser')

        html = list(soup.children)[2]

        body = list(html.children)[3]

        p = list(body.children)[1]

        raw_quote = (p.get_text())

        raw_quote = raw_quote.split('\n')

        i = 0
        while i < len(raw_quote):
            if len(raw_quote[i]) != 0 and raw_quote[i][0] == '“':
                quote_list.append(raw_quote[i] +  '\n' + raw_quote[i + 1])
            i += 1
            
    except:
        continue


random_quote = random.choice(quote_list)

print(random_quote)