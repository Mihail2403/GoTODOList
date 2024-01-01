CREATE TABLE IF NOT EXISTS users  (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash CHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS todo_list (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS users_lists (
    id SERIAL PRIMARY KEY,
    user_id INT,
    list_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_list(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS todo_items  (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL, 
    description VARCHAR(255),
    done BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS lists_item  (
    id SERIAL PRIMARY KEY,
    item_id INT,
    list_id INT,
    FOREIGN KEY (item_id) REFERENCES todo_items(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_list(id) ON DELETE CASCADE
);