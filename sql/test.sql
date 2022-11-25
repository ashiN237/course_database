USE test;

CREATE TABLE test_table(
    id int,
    name varchar(100),
    age int
);

INSERT INTO test_table (id, name, age) VALUES(1, "Taro", 20);

SELECT * FROM test_table;
