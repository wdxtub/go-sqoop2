# go-sqoop2

go client for sqoop2 (1.99.7)

+ [Apache Doc](https://sqoop.apache.org/docs/1.99.7/dev/RESTAPI.html)
+ [1.99.7 API Change](https://maprdocs.mapr.com/home/Sqoop/Sqoop1.99.7DeprecatedAPI.html)

因为 API 文档不全，没有办法继续下去，这个项目似乎也很久不更新了，Sad

## API curl Test

+ `curl 127.0.0.1:12000/sqoop/version`
+ `curl 127.0.0.1:12000/sqoop/v1/connector/all`
+ `curl 127.0.0.1:12000/sqoop/v1/job/all`
+ `curl 127.0.0.1:12000/sqoop/v1/link/all`


curl -XPOST -d '{"links":[{"id":-1,"enabled":true,"creation-user":"yiplatform","update-user":"yiplatform","name":"test_link","creation-date":1526266566774,"update-date":1526266566774,"connector-name":"generic-jdbc-connector","link-config-values":{"configs":[{"id":1,"type":"LINK","validators":[],"inputs":[{"editable":"ANY","id":1,"name":"linkConfig.jdbcDriver","overrides":"","sensitive":false,"size":128,"type":"STRING","validators":[],"value":"com.mysql.jdbc.Driver"},{"editable":"ANY","id":2,"name":"linkConfig.connectionString","overrides":"","sensitive":false,"size":2000,"type":"STRING","validators":[],"value":"jdbc:mysql://127.0.0.1:3306/yiplatform"},{"editable":"ANY","id":3,"name":"linkConfig.username","overrides":"","sensitive":false,"size":40,"type":"STRING","validators":[],"value":"root"},{"editable":"ANY","id":4,"name":"linkConfig.password","overrides":"","sensitive":false,"size":40,"type":"STRING","validators":[],"value":"zyqqv587^^"},{"editable":"ANY","id":6,"name":"linkConfig.jdbcProperties","overrides":"","sensitive":false,"sensitive-pattern":"","type":"MAP","validators":[],"value":{"protocol":"tcp"}}],"name":"linkConfig"},{"id":2,"type":"LINK","validators":[],"inputs":[{"editable":"ANY","id":7,"name":"dialect.identifierEnclose","overrides":"","sensitive":false,"size":5,"type":"STRING","validators":[],"value":"++"}],"name":"dialect"}],"validators":[]}}]}' 127.0.0.1:12000/sqoop/v1/link
