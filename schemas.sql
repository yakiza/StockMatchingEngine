DROP TABLE tradebook;
DROP TABLE tickers;
DROP TABLE trades;
DROP TABLE orders;
DROP TABLE users;


CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    firstname TEXT NOT NULL,
	lastname TEXT NOT NULL,
    tickers INTEGER,
    trades INTEGER
);

CREATE TABLE IF NOT EXISTS orders
(
    id SERIAL PRIMARY KEY,
    userID INTEGER NOT NULL references users(id),
	tickerID TEXT NOT NULL,
	price NUMERIC(10,2) NOT NULL ,
	quantity INTEGER NOT NULL,
	command CHARACTER(5)
);

CREATE TABLE IF NOT EXISTS trades
(
    id SERIAL PRIMARY KEY,
    buyerid INTEGER NOT NULL,
    sellerid INTEGER NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    quantity INTEGER NOT NULL,
    ticker  CHARACTER(5)
);

CREATE TABLE IF NOT EXISTS tickers
(
    id CHARACTER(5) PRIMARY KEY,
    userid INTEGER NOT NULL references users(id),
	quantity INTEGER NOT NULL    
);

CREATE TABLE tradebook
(
    userid INTEGER NOT NULL references users(id),
    tradeid INTEGER NOT NULL references orders(id)
);