-- Миграция для создания таблицы пользователей (User)
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    name VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255),
    role INTEGER NOT NULL DEFAULT 1,
    registered_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    register_token VARCHAR(255) UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT FALSE
);

-- Миграция для хранения refresh токенов (и сессий пользоватетелей)
CREATE TABLE IF NOT EXISTS sessions (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    refresh_token VARCHAR(255),
    last_used TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ip_address VARCHAR(255) NOT NULL
);
