import requests
from bs4 import BeautifulSoup

link = "https://en.wikipedia.org/wiki/List_of_current_members_of_the_United_States_Senate"

html = requests.get(link).text

# If you do not want to use requests then you can use the following code below
# with urllib (the snippet above). It should not cause any issue."""
soup = BeautifulSoup(html, "lxml")
res = soup.findAll("span", {"class": "fn"})
for r in res:
    print("Name: " + r.find('a').text)
    table_body=soup.find('senators')
    rows = table_body.find_all('tr')
    for row in rows:
        cols=row.find_all('td')
        cols=[x.text.strip() for x in cols]
        print(cols)