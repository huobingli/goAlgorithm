import redis

class CodeGenerate(object):
    def __init__(self):
        self.redis = redis.Redis('47.114.171.118', port=6999, password='111111', decode_responses=True)
        print(self.redis)

    def generate(self):
        # for code in self.codes:
        #     print("添加code为：" + code + "的代码到爬取队列")
        #     # self.redis.lpush('Bonus:start_urls', 'http://basic.10jqka.com.cn/' + code + '/bonus.html')
        #     self.redis.lpush('move_id', '')
        # print("所有code已经加入队列")

       ret = self.redis.set("hot_news_0", "world_craft")
       ret = self.redis.set("hot_news_1", "test")
       ret = self.redis.set("hot_news_2", "add")
       ret = self.redis.set("hot_news_3", "delete")
       ret = self.redis.set("hot_news_4", "modify")
       ret = self.redis.set("hot_news_5", "update")
       print(ret)

    def get_key(self, key):
        ret = self.redis.get(key)
        return ret

def add_redis(self, movie_id, movie_url):
    self.redis.lpush(movie_id, movie_url)

def gen_run():
    c = CodeGenerate()
    # c.generate()
    ret = c.get_key("hot_news_0")
    print(ret)
    print(type(ret))

    # c.add_redis("test", "test")

if __name__ == '__main__':
    gen_run()
