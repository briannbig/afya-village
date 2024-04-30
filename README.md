creating postgres database docker instance
```shell
docker run --name afya-village-datastore -e POSTGRES_PASSWORD=afy4 -e POSTGRES_DB=afya_village -e POSTGRES_USER=afya -p 5444:5432 -d postgres:alpine
```
