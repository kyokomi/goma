create database goma_test;
create user 'admin@localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON goma_test.* TO 'admin'@'localhost';
FLUSH PRIVILEGES;
