DROP SCHEMA IF EXISTS testdb;
CREATE SCHEMA testdb;
USE testdb;

DROP TABLE IF EXISTS words;
CREATE TABLE words(
    id INT AUTO_INCREMENT, 
    word VARCHAR(100),
    PRIMARY KEY(id)
);

INSERT INTO words(word) VALUE("ねこ");
INSERT INTO words(word) VALUE("テスト");
INSERT INTO words(word) VALUE("test");