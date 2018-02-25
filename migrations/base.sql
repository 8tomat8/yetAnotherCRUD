CREATE database yetAnotherCRUD;

CREATE TABLE yetAnotherCRUD.users (
  ID INT primary key auto_increment,
  Username VARCHAR(255) NOT NULL,
  Password VARCHAR(32) NOT NULL,
  Firstname VARCHAR(255) NOT NULL,
  Lastname VARCHAR(255) NOT NULL,
  Sex ENUM('male', 'female') NOT NULL,
  Birthdate DATE NOT NULL
);
