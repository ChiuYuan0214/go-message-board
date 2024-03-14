create database if not exists go_project;

use go_project;

create table if not exists users (
    user_id int auto_increment primary key,
    username varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    phone varchar(255),
    job varchar(255),
    address varchar(255),
    is_active boolean default false,
    creation_time datetime default current_timestamp,
    update_time datetime default current_timestamp on update current_timestamp
);

create index users_creation_time_index on users(creation_time);

create table if not exists verification_codes (
    code_id int auto_increment primary key,
    user_id int not null,
    code varchar(255) not null,
    valid boolean default true,
    creation_time datetime default current_timestamp,
    expire_time datetime default current_timestamp,
    foreign key (user_id) references users(user_id)
);

create index verification_codes_creation_time_index on verification_codes(creation_time);

create table if not exists articles (
    article_id int auto_increment primary key,
    user_id int not null,
    title text not null,
    content text not null,
    top_comment_id int,
    edited boolean default false,
    view_count smallint default 0,
    publish_time datetime default current_timestamp,
    creation_time datetime default current_timestamp,
    update_time datetime default current_timestamp on update current_timestamp,
    foreign key (user_id) references users(user_id)
);

create index articles_user_id_index on articles(user_id);
create index articles_publish_time_index on articles(publish_time);

create table if not exists tags (
    tag_id int auto_increment primary key,
    name varchar(255) not null
);

create table if not exists article_tag_maps (
    article_id int,
    tag_id int,
    primary key (article_id, tag_id)
);

create table if not exists comments (
    comment_id int auto_increment primary key,
    user_id int not null,
    article_id int not null,
    title text not null,
    content text not null,
    edited boolean default false,
    creation_time datetime default current_timestamp,
    update_time datetime default current_timestamp on update current_timestamp,
    foreign key (user_id) references users(user_id),
    foreign key (article_id) references articles(article_id)
);

create index comments_user_id_index on comments(user_id);
create index comments_article_id_index on comments(article_id);
create index comments_creation_time_index on comments(creation_time);

create table if not exists votes (
    vote_id int auto_increment primary key,
    user_id int not null,
    source_id int not null,
    score tinyint default 0,
    vote_type varchar(255) default 'article',
    creation_time datetime default current_timestamp,
    update_time datetime default current_timestamp on update current_timestamp,
    foreign key (user_id) references users(user_id)
);

create index votes_user_id_index on votes(user_id);
create index votes_source_index on votes(source_id, vote_type);

create table if not exists images (
    user_id int primary key,
    file_name varchar(255) not null,
    descript varchar(255),
    creation_time datetime default current_timestamp,
    update_time datetime default current_timestamp on update current_timestamp
);

create table if not exists followers (
    user_id int,
    follower_id int,
    foreign key (user_id) references users(user_id),
    primary key (user_id, follower_id)
);

create table if not exists collections (
    user_id int,
    article_id int,
    foreign key (user_id) references users(user_id),
    foreign key (article_id) references articles(article_id),
    primary key (user_id, article_id)
);
