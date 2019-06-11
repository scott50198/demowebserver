package dbhelper

const CREATE_USERINFO_TABLE = `create table userInfo(
	id int not null auto_increment,
    account varchar(30) not null unique,
    password varchar(70) not null,
    name varchar(20) not null,
    email varchar(30) not null unique,
    createTime timestamp not null,
    updateTime timestamp not null,
    primary key(id),
    index(updateTime));`

const INSERT_USERINFO = `INSERT INTO userInfo (account, password, name, email, createTime, updateTime) 
    SELECT ?, sha2(?, 256), ?, ?, now(), now()`

const CHECK_ACCOUNT_EXIST = `select * from userInfo
    where account= ?`

const CHECK_EMAIL_EXIST = `select * from userInfo
    where email= ?`

const GET_USER_ID_FROM_ACCOUNT = `select id from userInfo
    where account = ?`

const CHECK_ACCOUNT_AND_PASSWORD_VALIDATE = `select * from userInfo
    where account = ? and password = sha2(?, 256)`

const GET_USER_INFO_FROM_ACCOUNT = `select id, account, name, email from userInfo
    where account = ?`
