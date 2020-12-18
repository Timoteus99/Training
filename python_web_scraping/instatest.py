from instapy import InstaPy

#docu link: https://github.com/timgrossmann/InstaPy/blob/master/DOCUMENTATION.md

session = InstaPy(username="username", password="password")

session.login()

session.like_by_tags(["chess", "music"], amount = 30)