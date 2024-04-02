// PLAY LIST : https://www.youtube.com/playlist?list=PL0wd2rfixHH23Ej7qVSrbgr-yhI6iSC05

package adress
// https://docs.gofiber.io/


```
mkdir go-react-blog
cd go-react-blog/
mkdir server
mkdir client
cd server/
go-react-blog/server$ go mod init github.com/blalyasar/go-react-blog
go-react-blog/server$ go get github.com/gofiber/fiber/v2
go-react-blog/server$ go run server.go 
go-react-blog/server$ go get -u gorm.io/gorm
go-react-blog/server$ go get -u gorm.io/driver/mysql
```

Second Video:
// https://www.youtube.com/watch?v=PPsTNQpT3e0&list=PL0wd2rfixHH23Ej7qVSrbgr-yhI6iSC05&index=2

```
MariaDB [(none)]> CREATE DATABASE fiberblog;

MariaDB [(none)]> show databases;
+--------------------+
| Database           |
+--------------------+
| fiberblog          |

```

After Created
```
MariaDB [fiberblog]> show tables;
+---------------------+
| Tables_in_fiberblog |
+---------------------+
| blogs               |

MariaDB [fiberblog]> explain blogs;
+-------+---------------------+------+-----+---------+----------------+
| Field | Type                | Null | Key | Default | Extra          |
+-------+---------------------+------+-----+---------+----------------+
| id    | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
| title | varchar(255)        | NO   |     | NULL    |                |
| post  | longtext            | YES  |     | NULL    |                |
+-------+---------------------+------+-----+---------+----------------+
```

3.VIDEO: 
https://www.youtube.com/watch?v=2fonkDhlPIs&list=PL0wd2rfixHH23Ej7qVSrbgr-yhI6iSC05&index=3



 
After  video3 
$ curl -X POST  http://localhost:8000

{"msg":"Blog Create","statusText":"Ok"}

$ curl localhost:8000

{"msg":"Blog List","statusText":"Ok"}

$ curl -X PUT  http://localhost:8000

{"msg":"Blog Update","statusText":"Ok"} 
 
$ curl -X DELETE  http://localhost:8000
{"msg":"Blog Delete from given id","statusText":"Ok"} 

4.VIDEO
BEFORE 4.VIDEO CREATE MYSQL DATA IN TABLES....
```

MariaDB [fiberblog]> 

MariaDB [fiberblog]> INSERT INTO `blogs` (`id`,`title`,`post`) VALUES (NULL, 'FIRST BLOG',  'FIRST POST');
Query OK, 1 row affected (0.002 sec)

MariaDB [fiberblog]> INSERT INTO `blogs` (`id`,`title`,`post`) VALUES (NULL, 'SECOND BLOG',  'SECOND POST');

Query OK, 1 row affected (0.001 sec)

MariaDB [fiberblog]> SELECT * from blogs;

+----+-------------+-------------+
| id | title       | post        |
+----+-------------+-------------+
|  1 | FIRST BLOG  | FIRST POST  |
|  2 | SECOND BLOG | SECOND POST |
+----+-------------+-------------+
2 rows in set (0.000 sec)
```

After blog.go update list data.... and dont forget data is list..

curl -X GET http://localhost:8000

{"blog_records":[{"id":1,"title":"FIRST BLOG","post":"FIRST POST"},{"id":2,"title":"SECOND BLOG","post":"SECOND POST"}],"msg":"Blog List","statusText":"Ok"}

5.VIDEO:
go get github.com/joho/godotenv


Add .env file in server directory for db info

```
db_user=""
db_password="...."
db_name="fiberblog"
```

6.VIDEO:
```
$curl -X POST --header "Content-Type: application/json"  --data '{"title":"newtitle","post":"new post"}'   http://localhost:8000
{"data":{"id":6,"title":"newtitle","post":"new post"},"msg":"Record is saved successfuly.","statusText":"Ok"

$ curl -X GET http://localhost:8000
{"blog_records":[{"id":1,"title":"FIRST BLOG","post":"FIRST POST"},{"id":2,"title":"SECOND BLOG","post":"SECOND POST"},{"id":3,"title":"","post":""},{"id":4,"title":"","post":""}],"msg":"Blog List","statusText":"Ok"}
:~$ 
```



7.VIDEO
```
put examples

curl -X GET http://localhost:8000
{"blog_records":[{"id":1,"title":"FIRST BLOG","post":"FIRST POST"},{"id":2,"title":"SECOND BLOG","post":"SECOND POST"},{"id":7,"title":"newtitle","post":"new post"},{"id":8,"title":"newtitle","post":"new post"},{"id":9,"title":"newtitle updated444","post":"new post updated 444"},{"id":10,"title":"newtitle updated444","post":"new post updated 444"}],"msg":"Blog List","statusText":"Ok"}


curl -X PUT --header "Content-Type: application/json"  --data '{"title":"newtitle updated444 TO 555","post":"new post updated 444 TO 555"}'   http://localhost:8000/9
{"data":{"id":9,"title":"newtitle updated444 TO 555","post":"new post updated 444 TO 555"},"msg":"record updated successfully","statusText":"Ok"}


curl -X GET http://localhost:8000
{"blog_records":[{"id":1,"title":"FIRST BLOG","post":"FIRST POST"},{"id":2,"title":"SECOND BLOG","post":"SECOND POST"},{"id":7,"title":"newtitle","post":"new post"},{"id":8,"title":"newtitle","post":"new post"},{"id":9,"title":"newtitle updated444 TO 555","post":"new post updated 444 TO 555"},{"id":10,"title":"newtitle updated444","post":"new post updated 444"}],"msg":"Blog List","statusText":"Ok"}

```