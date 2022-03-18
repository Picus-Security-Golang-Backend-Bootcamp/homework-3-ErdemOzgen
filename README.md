# homework-3-ErdemOzgen
homework-3-ErdemOzgen created by GitHub Classroom

# Run Mysql Database in Docker
```bash
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -e MYSQL_USER=admin -e MYSQL_PASSWORD=root -d mysql
```
Create new schema with name "library" in the database with Mysql Workbench
