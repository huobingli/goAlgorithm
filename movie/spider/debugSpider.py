from scrapy.cmdline import execute
import sys
import os
sys.path.append(os.path.dirname(os.path.abspath(__file__)))

def spider_run():
    execute(['scrapy', 'crawlall'])

def debug():
    execute(['scrapy', 'crawl', 'BonusFinancing'])

if __name__ == '__main__':
    debug()