CREATE TABLE IF NOT EXISTS cameras (
    id serial primary key,
    latitude float8 not null,
    longitude float8 not null,
    address varchar(255) not null,
    status varchar(255) not null,
    created_at timestamp not null
);

CREATE TABLE IF NOT EXISTS incident_types (
    id serial primary key,
    incident_type varchar(255) not null,
    severity varchar(255) not null,
    response_action varchar(255),
    created_at timestamp not null
);

CREATE TABLE IF NOT EXISTS situations (
    id serial primary key,
    camera_id int,
    incident_type_id int,
    description varchar(255),
    start_time timestamp,
    end_time timestamp,
    created_at timestamp not null,
    FOREIGN KEY (camera_id) REFERENCES cameras(id) on delete cascade,
    FOREIGN KEY (incident_type_id) REFERENCES incident_types(id) on delete cascade
);

CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    username varchar(255) not null,
    email varchar(255) not null,
    phone_number varchar(255),
    password varchar(255) not null,
    role varchar(255) not null,
    last_login timestamp,
    created_at timestamp not null
);

CREATE TABLE IF NOT EXISTS notifications (
    id serial primary key,
    user_id int,
    situation_id int,
    message varchar(255),
    status varchar(255) not null,
    created_at timestamp not null,
    sent_at timestamp not null,
    FOREIGN KEY (user_id) REFERENCES users(id) on delete cascade,
    FOREIGN KEY (situation_id) REFERENCES situations(id) on delete cascade
);

CREATE TABLE IF NOT EXISTS user_situation (
    id serial primary key,
    user_id int,
    situation_id int,
    created_at timestamp not null,
    assigned_at timestamp not null,
    FOREIGN KEY (user_id) REFERENCES users(id) on delete cascade,
    FOREIGN KEY (situation_id) REFERENCES situations(id) on delete cascade
);
