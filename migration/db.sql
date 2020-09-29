
DROP TABLE trades CASCADE;
DROP TABLE orders CASCADE;
DROP TABLE tradebook CASCADE; 
DROP TABLE tickers CASCADE;
DROP TABLE users CASCADE;

CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    firstname TEXT NOT NULL,
	lastname TEXT NOT NULL
);

CREATE TABLE orders
(
    id SERIAL PRIMARY KEY,
    userID INTEGER NOT NULL references users(id),
	tickerID TEXT NOT NULL,
	price NUMERIC(10,2) NOT NULL ,
	quantity INTEGER NOT NULL,
	command CHARACTER(5)
);

CREATE TABLE trades
(
    id SERIAL PRIMARY KEY,
    buyerid INTEGER NOT NULL,
    sellerid INTEGER NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    quantity INTEGER NOT NULL,
    ticker  CHARACTER(5)
);

CREATE TABLE tickers
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

