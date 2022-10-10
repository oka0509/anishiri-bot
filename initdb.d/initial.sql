DROP SCHEMA IF EXISTS testdb;
CREATE SCHEMA testdb;
USE testdb;

DROP TABLE IF EXISTS words;
CREATE TABLE words(
    id INT AUTO_INCREMENT, 
    word VARCHAR(100),
    PRIMARY KEY(id)
);

INSERT INTO words(word) VALUE("あそびあそばせ");
INSERT INTO words(word) VALUE("いぬやしき");
INSERT INTO words(word) VALUE("うららめいろちょう");
INSERT INTO words(word) VALUE("えいぞうけんにはてをだすな");
INSERT INTO words(word) VALUE("おじゃまじょどれみ");
INSERT INTO words(word) VALUE("かけぐるい");
INSERT INTO words(word) VALUE("きののたび");
INSERT INTO words(word) VALUE("くらなど");
INSERT INTO words(word) VALUE("けろろぐんそう");
INSERT INTO words(word) VALUE("このすば");
INSERT INTO words(word) VALUE("さえないかのじょのそだてかた");
INSERT INTO words(word) VALUE("しゃーろっと");
INSERT INTO words(word) VALUE("すずめのとじまり");
INSERT INTO words(word) VALUE("せいとかいのいちぞん");
INSERT INTO words(word) VALUE("そらよりもとおいばしょ");
INSERT INTO words(word) VALUE("たまごっち");
INSERT INTO words(word) VALUE("ちゅうにびょうでもこいがしたい");
INSERT INTO words(word) VALUE("つきがきれい");
INSERT INTO words(word) VALUE("てんしのすりーぴーす");
INSERT INTO words(word) VALUE("とあるかがくのレールガン");
INSERT INTO words(word) VALUE("なるにあこくものがたり");
INSERT INTO words(word) VALUE("にせこい");
INSERT INTO words(word) VALUE("ぬらりひょんのまご");
INSERT INTO words(word) VALUE("ねとげのすすめ");
INSERT INTO words(word) VALUE("のーげーむ・のーらいふ");

