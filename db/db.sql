create database julo;
use julo;

create table user (
userid int auto_increment not null,
username varchar(20) not null,
password varchar(100) not null,
token varchar(50),
generated_at datetime,
expired_at datetime,
primary key (userid)
);

create table wallet (
walletid int auto_increment not null,
userid int not null,
balance float,
enabled_at datetime,
disabled_at datetime,
status varchar(10),
primary key (walletid),
foreign key (userid) references user(userid)
);

insert into user(username, password) values("user1","pass1");