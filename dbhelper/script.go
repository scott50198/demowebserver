package dbhelper

const CREATE_USERINFO_TABLE = `create table userInfo(
	id int not null auto_increment,
    account varchar(30) not null unique,
    password varchar(60) not null,
    name varchar(20) not null,
    email varchar(30) not null unique,
    createTime timestamp not null,
    updateTime timestamp not null,
    primary key(id),
    index(updateTime));`

const INSERT_USERINFO = `INSERT INTO userInfo (account, password, name, email, createTime, updateTime) 
    VALUES (?, ?, ?, ?, now(), now())`
