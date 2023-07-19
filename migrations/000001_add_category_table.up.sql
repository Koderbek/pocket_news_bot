CREATE TABLE IF NOT EXISTS category
(
    id   SMALLSERIAL PRIMARY KEY,
    code VARCHAR(100) UNIQUE,
    name VARCHAR(100)
);

INSERT INTO category (code, name)
VALUES ('general', 'Главное'),
       ('world', 'Мир'),
       ('nation', 'Россия'),
       ('business', 'Бизнес'),
       ('technology', 'Технологии'),
       ('entertainment', 'Развлечения'),
       ('sports', 'Спорт'),
       ('science', 'Наука'),
       ('health', 'Здоровье');