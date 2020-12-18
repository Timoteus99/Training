import re
import time
import random
import requests
import datetime
import subprocess
from bs4 import BeautifulSoup

#class FBGroupScraper:
#    def _init_(self, group_id):
#        self.group_id = group_id
#        self.page_url = "https://facebook.com/groups/" + self.group_id
#        self.page_content = ""
#    
#    def get_page_content(self):
#        self.page_content = requests.get(self.page_url).text

#    def parse(self):
#        soup = BeautifulSoup(self.page_content, "html.parser")
#        feed_container = soup.find(id = "m_group_stories_container").find_all("p")
#        for i in feed_container:
#            print(i.text)

#group_id = "549562675144155"
#d = FBGroupScraper(group_id)
#d.get_page_content()
#d.parse()

def check_temperature():
    return subprocess.call(["/opt/vc/bin/vcgencmd", "measure_temp"])

def send_telegram_alarm(bot_message):
    #function to send bot message via telegram
    
    bot_token = '1461320034:AAHEGk-pBFw53gfJSW852VD4MUyG_6rdfpY'
    bot_chatID = '1490237517'
    send_text = 'https://api.telegram.org/bot' + bot_token + '/sendMessage?chat_id=' + bot_chatID + '&parse_mode=Markdown&text=' + bot_message

    response = requests.get(send_text)

    return response.json()


def get_weather_info(city_name):
    #function to return city weather info

    city = []

    city_string = city_name[0:2] + ': '

    hour = datetime.datetime.now().hour
    day = datetime.datetime.now().day
    month = datetime.datetime.now().month
    
    #Sevnica
    page_url = "https://vreme.siol.net/mesto/slovenija/" + city_name
    page_content = ""
    page_content = requests.get(page_url).text
    #print(page_content)
    time.sleep(3)
    soup = BeautifulSoup(page_content, "html.parser")
    feed_container = soup.find_all("p")

    for data in feed_container:
        string_data = str(data)
        #print(string_data)
        if "hour" in string_data and str(hour) in string_data[:20]:
            if str(day) + '.' + str(month) + '.' in string_data:
                index = feed_container.index(data) + 1
                #print(index)
                for _ in range(5):
                    line = str(feed_container[index]).strip('</p>')
                    digit = re.findall(r'\d+', line)
                    if len(digit) != 0:
                        city.append(digit)
                    index += 1
        if 'km/h' in string_data or 'mbar' in string_data:
            index = (feed_container.index(data))
            line = str(feed_container[index]).strip('</p>')
            #print(line)
            city.append([line])
            index += 1

    for element in city:
        for data in element:
            print(data)
            index = city.index(element)
            if index == 0 or index == 1:
                city_string += data + ', '
            elif index == 2 or index == 3:
                city_string += data + '°C, '
            elif index == 4:
                city_string += data + ' %, '
            else:
                city_string += data + ' %'


    return city_string


def GroupScraper(group_id):
    page_url = "https://mobile.facebook.com/groups/" + group_id
    page_content = ""
    page_content = requests.get(page_url).text
    #print(page_content)
    soup = BeautifulSoup(page_content, "html.parser")
    #print(len(soup))
    feed_container = soup.find(id="m_group_stories_container").find_all("p")
    #print(len(feed_container))
    #print(feed_container)
    #print(feed_container)
    posts = []
    for element in feed_container:
        posts.append(element.text)
    return posts


last_posts = []

while True:
    #print(check_temperature())
    try:
        new_posts = GroupScraper("549562675144155")
        for post in new_posts:
            if post not in last_posts:
                
                sevnica_string = get_weather_info("Sevnica")
                krsko_string = get_weather_info("Krsko")
                brezice_string = get_weather_info("Brezice")
                ljubljana_string =get_weather_info("Ljubljana")
                weather_forecast = sevnica_string + '\n' + krsko_string + '\n' + brezice_string + '\n' + ljubljana_string
                
                send_telegram_alarm("OBVESTILO: " + post + '\n' + 'VREME:' + '\n' + weather_forecast)
                #print("OBVESTILO: " + post + '\n' + 'VREME:' + '\n' + weather_forecast)
                

                if len(last_posts) < 30:
                    last_posts.append(post)
                else:
                    last_posts.append(post)
                    last_posts.remove(last_posts[0])

        time.sleep(180)
        
    except:
        #print("none")
        time.sleep(120)
        continue
    




