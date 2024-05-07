-- See User and Host --
SELECT User, Host FROM mysql.user;

-- Create Database name "Balance" --
CREATE DATABASE Balance;

-- Use Balance Database --
USE Balance;

-- Create Table --
-- Fields Explanation --
-- Date = Date report -- 
-- Code = Stock code --
-- IS = Insurance --
-- CP = Corporate --
-- PF = Pension Fund --
-- IB = Bank --
-- ID = Individual --
-- MF = Mutual Fund --
-- SC = Securities  --
-- FD = Foundation --
-- OT = Other --
CREATE TABLE Stock(
	`Date` DATE,
	`Code` CHAR(4),
	`Local_IS` BIGINT UNSIGNED,
	`Local_CP` BIGINT UNSIGNED,
	`Local_PF` BIGINT UNSIGNED,
	`Local_IB` BIGINT UNSIGNED,
	`Local_ID` BIGINT UNSIGNED,
	`Local_MF` BIGINT UNSIGNED,
	`Local_SC` BIGINT UNSIGNED,
	`Local_FD` BIGINT UNSIGNED,
	`Local_OT` BIGINT UNSIGNED,
	`Foreign_IS` BIGINT UNSIGNED,
	`Foreign_CP` BIGINT UNSIGNED,
	`Foreign_PF` BIGINT UNSIGNED,
	`Foreign_IB` BIGINT UNSIGNED,
	`Foreign_ID` BIGINT UNSIGNED,
	`Foreign_MF` BIGINT UNSIGNED,
	`Foreign_SC` BIGINT UNSIGNED,
	`Foreign_FD` BIGINT UNSIGNED,
	`Foreign_OT` BIGINT UNSIGNED,
	PRIMARY KEY (`Date`, `Code`)
);

-- Testing query --
SELECT * FROM Stock WHERE `Code` = 'BBCA' ORDER BY `Date`;

-- Check the Stocks Table --
DESC Stock;

-- Show tables in Database --
SHOW TABLES;

-- Clear table data --
TRUNCATE Table Stock;