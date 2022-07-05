# To Do List API

This repository contains the "_To Do List API_" project, developed in Golang. It aims to put into practice the knowledge learned about the language.


### Author:

- _[Flávio Bianchetti](https://www.linkedin.com/in/flaviobianchetti/)_

---

### Were used in the construction of this page:

<section>
  <a href="https://go.dev/" target="_blank">
    <img
      align="center"
      height="30"
      src="https://img.shields.io/badge/GoLang-00AFF0?style=for-the-badge&logo=Go&logoColor=black"
      target="_blank"
    />
  </a>
</section>

---

### Installation

1. Open the terminal, in a directory of your choice, and clone the project:
```bash
  git clone git@github.com:flavio-bianchetti/go-todo-api.git
```

2. Enter the project directory:
```bash
  cd go-todo-api
```

3. The database used in this project is _[MySQL](https://dev.mysql.com/doc/)_. The file "_database.sql_" located in the root folder of this project contains the database creation script:
```SQL
CREATE DATABASE IF NOT EXISTS go_todo;

USE go_todo;

CREATE TABLE IF NOT EXISTS todo
(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(1000) NOT NULL
);

INSERT INTO todo (name, description)
VALUES ('Supermarket', 'milk, bread, egg, pasta, meat, cookies and fruit'),
    ('Meeting - work team', 'today, at 5 pm'),
    ('Football game', 'Friday, at 7 pm');

```

4. Configure the _[dotenv](https://www.npmjs.com/package/dotenv)_ file with your information, as shown below:
```javascript
  DBUSER=root
  DBPASSWORD=1234
  DBHOST=localhost
  DBPORT=3306
  DBDATABASE=go_todo
```

5. Launch the application:
```bash
  go run main.go
```

<!-- 3. Install dependencies:
```bash
  npm install
``` -->
<!-- 4. The API uses the MySQL database to store the information. If you do not have MySQL installed, consult the documentation available _[here](https://dev.mysql.com/doc/)_ or change the file "_src/config/config.js_" to database of your choice. -->

<!-- 5. Create the database:
```bash
  npx sequelize db:create
```
6. Create the tables:
```bash
  npx sequelize db:migrate
```
7. Feed the database:
```bash
  npx sequelize db:seed:all
```
8. Configure the _[dotenv](https://www.npmjs.com/package/dotenv)_ file with your information:
```javascript
  DB_HOSTNAME=host_name // Ex.: localhost
  DB_USER=db_username
  DB_PASSWORD=db_password
  DB_DATABASE=db_database
  DB_PORT=number // Ex.: 3000
  DB_DIALECT=dialect_name // Ex.: mysql
```
9. Launch the application:
```bash
  npm start
``` -->

### API usage:

<!-- 1. ...:
![flavio-bianchetti-blogs-api-project](https://docs.google.com/uc?id=)

2. ...:
![flavio-bianchetti-blogs-api-project](https://docs.google.com/uc?id=)

3. ...:
![flavio-bianchetti-blogs-api-project](https://docs.google.com/uc?id=) -->

### Available resources:
- **Check API health:**
  - **URL:** /ping
  - **Method:** GET
  - **response code:** 200
  - **response:** pong
- **Create an registry:**
  - **URL:** /
  - **Method:** POST
  - **response code:** 201 
  - **response:** object created - json
- **Read all registries:**
  - **URL:** /
  - **Method:** GET
  - **response code:** 200
  - **response:** all registries - json
- **Read one registry by id:**
  - **URL:** /:id
  - **Method** GET
  - **response code:** 200
  - **response:** selected registry - json
- **Delete registry by id:**
  - **URL:** /:id
  - **Method:** DELETE
  - **response code:** 200
  - **response:** none

### Author's considerations:

- There is no schedule of changes planned for this project.
- This project is under development.

---

by _[Flávio Bianchetti - 2022](https://github.com/flavio-bianchetti)_.

