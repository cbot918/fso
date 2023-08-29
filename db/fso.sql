CREATE TABLE users(
  id        int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  email     varchar(255) NOT NULL,
  name      varchar(255) NOT NULL,
  password  varchar(255) NOT NULL
);

CREATE TABLE chars(
  id        int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  level     varchar(50) NOT NULL,
  userid    int NOT NULL,
  name      varchar(255) NOT NULL
);

ALTER TABLE `chars` ADD FOREIGN KEY (`userid`) REFERENCES `users` (`id`);