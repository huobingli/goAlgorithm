import requests


'''
request请求 
返回 请求数据
'''
def requestURL(url):
    headers = { 
        'Accept-Encoding': 'gzip, deflate, br',
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat',
        'Referer': 'https://servicewechat.com/wxe04e1890bc944202/33/page-frame.html',
        'Host': 'www.fupanyoudao.com'
    }

    ret = requests.get(url, headers=headers)
    return ret
'''
request请求 
带cookie 
返回 请求数据
'''
def reqeustWithCookie(url, cookie):
    pass
