# golang-mysql
Basic Golang Mysql Connection and Query Template

# Steps

Run the go mod init command, giving it your new code’s module path.

    go mod init golang-mysql

Create Table and Insert Data

    DB Table create script
    
    -- testdb.TestUsers definition
    DROP TABLE IF EXISTS testdb.TestUsers;
    
    CREATE TABLE testdb.TestUsers (
      	id INT AUTO_INCREMENT NOT NULL,
    	name VARCHAR(50) NOT NULL,
    	email VARCHAR(100) NOT null,
    	PRIMARY KEY (`id`)
    )
    ENGINE=InnoDB
    DEFAULT CHARSET=utf8mb4
    COLLATE=utf8mb4_0900_ai_ci;
    
    INSERT INTO testdb.TestUsers
      (name, email)
    VALUES
      ('Blue Train', 'bluetrain@mail.com'),
      ('Giant Steps', 'john@mail.com')

Use the go get to add the github.com/go-sql-driver/mysql module as a dependency for your own module. Use a dot argument to mean “get dependencies for code in the current directory.”

    go get .

Run

    go run .

Terminal

    Connected!
    User found: [{1 Blue Train bluetrain@mail.com}]
    ID of added users: 3

  
