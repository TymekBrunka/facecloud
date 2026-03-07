### users

| name            | type         | additional             |
|-----------------|--------------|------------------------|
| id              | serial       | primary key            |
| name            | varchar(70)  | not null               |
| surname         | varchar(40)  | not null               |
| email           | varchar(100) | not null unique        |
| password        | char(65)     | not null               |
| birth_date      | date         | not null               |
| roles           | short array  | not null               |
| /               | /            | /                      |
| profile_picture | varchar(200) |                        |
| bio             | varchar(250) |                        |
| last_login      | timestamp    |                        |
| /               | /            | /                      |
| online          | boolean      | not null default false |
| activated       | boolean      | not null default false |

### groups

| name       | type        | additional                         |
|------------|-------------|------------------------------------|
| id         | serial      | primary key                        |
| name       | varchar(70) | not null                           |
| /          | /           | /                                  |
| created_at | timestamp   | not null default current_timestamp |


### group_members

| name      | type      | additional                              |
|-----------|-----------|-----------------------------------------|
| id        | serial    | primary key                             |
| group_id  | int       | references groups(id) on delete cascade |
| user_id   | int       | references  uses(id) on delete cascade  |
| /         | /         | /                                       |
| joined_at | timestamp | not null default current_timestamp      |

### messages

| name              | type        | additional                                      |
|-------------------|-------------|-------------------------------------------------|
| id                | serial      | primary key                                     |
| sender_id         | int         | not null references users(id) on delete cascade |
| receiver_id       | int         | references users(id) on delete set null         |
| group_id          | int         | references groups(id) on delete set null        |
| content           | text        | not null                                        |
| created_at        | timestamp   | not null default current_timestamp              |
| group_name_backup | varchar(70) |                                                 |

```sql
constraint chk_sent_somwehere CHECK (  
    sender_id NOT NULL OR
    receiver_id NOT NULL  
)
```
