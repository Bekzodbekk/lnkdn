CREATE TABLE IF NOT EXISTS likes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    post_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    UNIQUE(user_id, post_id) -- har bir user faqat bitta marta like qila olishi uchun
);
