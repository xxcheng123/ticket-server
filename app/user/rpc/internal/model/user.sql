CREATE TABLE IF NOT EXISTS user(
    id int auto_increment,
    username varchar(15) not null comment '用户名',
    password char(32) not null comment '密码',
    balance decimal(10,2) not null default 0 comment '余额',
    last_change_time datetime not null default '1000-01-01 00:00:00' comment '最后修改密码时间',
    primary key (id),
    unique key (username)
)charset=utf8mb4,engine=innodb;