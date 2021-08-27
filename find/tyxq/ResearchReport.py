import sys
from sprequrl import *

def openurl(_page_id):
    basic_url = "https://www.fupanyoudao.com/v1/api/report?page_num=%d&page_size=10&type=1&userid=24214" % 1
    ret = requestURL(basic_url)
    # print(ret)
    print(ret.text)

if __name__ == "__main__":
    openurl()