--ユーザーの作成
CREATE USER postgres;
--DBの作成
CREATE DATABASE asset_sim;
--ユーザーにDBの権限をまとめて付与
GRANT ALL PRIVILEGES ON DATABASE asset_sim TO postgres;
--ユーザーを切り替え
\c postgres
--テーブルを作成
CREATE TABLE book (
  id integer,
  name varchar(30)
);
--テーブルにデータを挿入
INSERT INTO book VALUES (1,'The Very Hungry Caterpillar');
