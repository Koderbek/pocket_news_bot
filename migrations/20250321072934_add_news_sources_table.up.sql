CREATE TABLE IF NOT EXISTS news_sources
(
    domain      VARCHAR(255) NOT NULL,
    category    VARCHAR(100) NOT NULL,
    active      CHAR(1)      NOT NULL CHECK (active IN ('Y', 'N')) DEFAULT 'Y',
    last_update TIMESTAMP    NOT NULL                              DEFAULT NOW(),

    CONSTRAINT unique_domain_category UNIQUE (domain, category)
);