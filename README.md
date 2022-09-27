# zh-word
单词打卡器

create table info(
id bigint  auto_increment,
name  varchar(20) not null comment '姓名',
today datetime default CURRENT_TIMESTAMP comment '时间',
word varchar(26)  not null comment '单词',
chinese varchar(2048) not null comment '中文',
primary key(id,name,word)
);

create table info(
name  varchar(20) not null comment '姓名',
today datetime default CURRENT_TIMESTAMP comment '时间',
word varchar(26)  not null comment '单词',
chinese varchar(2048) not null comment '中文',
primary key(name,word)
);

insert into info(name,word,chinese) values
('章皓', 'a', 'a'),
('刘俊新', 'a', 'a'),
('李丽', 'a', 'a'),
('李锦晖', 'a', 'a'),
('班风伦', 'a', 'a'),
('郭斌', 'a', 'a'),
('梁郑宇', 'a', 'a'),
('刘智裕', 'a', 'a')
;

insert into info(name,word,chinese, today) values
('章皓', 'zh', '测试', '2021-09-11');

(select count(*) num, today, name from info group by name order by num desc)  union all   (select count(*) tnum, today, name from info  where to_days(today) = to_days(now())    group by name order by tnum;);



select count(*) num, today, (select count(*) tnum from info  where to_days(today) = to_days(now()) group by name order by tnum)  as tnum,  name from info group by name order by num desc;

select a.name,
        a.num,
       b.tnum
from
        (
        select 
            name, today, count(*) as num
        from
            info
        group by name
) as a left join (
select count(*) tnum, today, name 
from info  where to_days(today) = to_days(now()) 
group by name order by tnum
)as b on a.name = b.name


set @@global.sql_mode ='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
如果第一条命令修改不成功的话，使用第二条命令
set @@global.sql_mode='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

flush privileges;
exit;

//结果
select a.name, a.today ,   a.num,        b.tnum from         (         select             name, today, count(*) as num         from             info         group by name ) as a left join ( select count(*) tnum, today, name  from info  where to_days(today) = to_days(now())  group by name order by tnum )as b on a.name = b.name;













