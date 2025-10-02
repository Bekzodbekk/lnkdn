CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,                              -- Post ID (avtomatik)
    user_id INT NOT NULL,                               -- Post egasi (foydalanuvchi ID)
    content TEXT NOT NULL,                              -- Post matni
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,     -- Yaratilgan vaqt
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,     -- Yangilangan vaqt
    deleted_at BIGINT DEFAULT 0
);
