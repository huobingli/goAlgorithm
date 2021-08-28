import pymysql
from timeit import default_timer

from setting import *

def get_mysql_conn():
    conn = pymysql.connect(host=host, port=port, db=db, user=user, password=password)
    return conn

# ---- 使用 with 的方式来优化代码
class UsingMysql(object):

    def __init__(self, commit=True, log_time=True, log_label='总用时'):
        """

        :param commit: 是否在最后提交事务(设置为False的时候方便单元测试)
        :param log_time:  是否打印程序运行总时间
        :param log_label:  自定义log的文字
        """
        self._log_time = log_time
        self._commit = commit
        self._log_label = log_label

    def __enter__(self):

        # 如果需要记录时间
        if self._log_time is True:
            self._start = default_timer()

        # 在进入的时候自动获取连接和cursor
        conn = get_mysql_conn()
        cursor = conn.cursor(pymysql.cursors.DictCursor)
        conn.autocommit = False

        self._conn = conn
        self._cursor = cursor
        return self

    def __exit__(self, *exc_info):
        # 提交事务
        if self._commit:
            self._conn.commit()
        # 在退出的时候自动关闭连接和cursor
        self._cursor.close()
        self._conn.close()

        if self._log_time is True:
            diff = default_timer() - self._start
            print('-- %s: %.6f 秒' % (self._log_label, diff))

    @property
    def cursor(self):
        return self._cursor

def insert_datas(datas):     
    with UsingMysql(log_time=True) as um:
        for data in datas:
            sql = "insert into ARK_INNOVATION_ETF(ark_date, ark_id, ark_stock_name, ark_company, ark_shares, ark_market_value, ark_weight)  \
            values(%s, %s, %s, %s, %s, %s, %s)"
            params = ('%s' % data[6], '%d' % data[0], "%s" % data[2], "%s" % data[1], "%s" % data[3], "%s" % data[4], "%f" % data[5])
            um.cursor.execute(sql, params)

def insert_data(database, data):
    with UsingMysql(log_time=True) as um:
        sql = "insert into %s(ark_date, ark_id, ark_stock_name, ark_company, ark_shares, ark_market_value, ark_weight)  \
        values(%s, %s, %s, %s, %s, %s, %s)" % database
        params = ('%s' % data[6], '%d' % data[0], "%s" % data[2], "%s" % data[1], "%s" % data[3], "%s" % data[4], "%f" % data[5])
        um.cursor.execute(sql, params)

if __name__ == '__main__':
    insert_datas()