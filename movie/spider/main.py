import requests
from bs4 import BeautifulSoup

link = "http://www.xxi5.cn/list/1-0-0-0-0-0.html"
base_http_header = "http://www.xxi5.cn"

html = requests.get(link).text

# If you do not want to use requests then you can use the following code below
# with urllib (the snippet above). It should not cause any issue."""
soup = BeautifulSoup(html, "lxml")
res = soup.findAll("td", {"class": "dt prel nobr"})
for r in res:
    for subr in r.find_all('a'):
        print("Down: " + subr["href"])
    # table_body=soup.find('senators')
    # # rows = table_body.find_all('tr')
    # for row in rows:
    #     cols=row.find_all('td')
    #     cols=[x.text.strip() for x in cols]
    #     print(cols)
# 加入redis
