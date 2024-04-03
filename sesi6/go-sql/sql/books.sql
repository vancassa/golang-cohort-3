CREATE TABLE `db-go-sql`.`books` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(45) NOT NULL,
  `author` VARCHAR(45) NOT NULL,
  `stock` INT NOT NULL,
  PRIMARY KEY (`id`));

-- CREATE TABLE books (
--   id SERIAL PRIMARY KEY,
--   title VARCHAR(45) NOT NULL,
--   author VARCHAR(45) NOT NULL,
--   stock INT NOT NULL
-- );
