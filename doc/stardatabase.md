# server address 192.168.2.62
# Mysql 
##  service mysql restart (/usr/local/mysql/bin/mysqld --basedir=/usr/local/mysql --datadir=/usr/local/mysql/data --plugin-dir=/usr/local/mysql/lib/plugin --user=mysql --log-error=serverless-node2.err --pid-file=/usr/local/mysql/data/serverless-node2.pid --port=3306)
##  pass : 123456
# Etcd 
## etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 &
# Redis 
## redis-server /jinxin/redis/redis-6.2.6/redis.conf & （redis-server *:9876）
## pass: pml@123:/:：
