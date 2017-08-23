sudo yum update -y
sudo yum install mysql-server -yi
sudo service mysqld starti
mysql --host=day-planner-db1.c124lns4pyxg.us-east-1.rds.amazonaws.com --user=day_planner_db1 --password=$DB_PASSWORD day_planner_db1
