use AI_Models_Manager;
alter table users add nick_name varchar(255) not null ;
alter table users modify column nick_name varchar(255) comment '昵称';
