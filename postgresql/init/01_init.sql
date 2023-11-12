--ユーザーの作成
CREATE USER docker;
--DBの作成
CREATE DATABASE asset_sim;
--ユーザーにDBの権限をまとめて付与
GRANT ALL PRIVILEGES ON DATABASE asset_sim TO docker;
--ユーザーを切り替え
\c docker
--テーブルを作成
CREATE TABLE book (
  id integer,
  name varchar(30)
);

CREATE TABLE nikkei (
  date varchar(100),
  value varchar(100)
);


--テーブルにデータを挿入
INSERT INTO book VALUES (1,'The Very Hungry Caterpillar');
