START TRANSACTION;

DROP TABLE IF EXISTS "user";

CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TYPE gender_type AS ENUM ('male', 'female', 'unknown');
CREATE TYPE status_type AS ENUM ('active', 'blocked', 'deleted', 'pending');

CREATE TABLE "user"
(
    id           BIGSERIAL PRIMARY KEY,
    email        VARCHAR(64) UNIQUE       NOT NULL CHECK ( email <> '' ),
    password     VARCHAR(255)             NOT NULL CHECK ( octet_length(password) <> 0 ),
    first_name   VARCHAR(64)              NOT NULL DEFAULT '',
    last_name    VARCHAR(64)              NOT NULL DEFAULT '',
    about        VARCHAR(4096)            NOT NULL DEFAULT '',
    phone_number VARCHAR(64)              NOT NULL DEFAULT '',
    gender       gender_type              NOT NULL DEFAULT 'unknown',
    status       status_type              NOT NULL DEFAULT 'active',
    last_ip      VARCHAR(32)              NOT NULL DEFAULT '',
    last_device  VARCHAR(64)              NOT NULL DEFAULT '',
    avatar_url   VARCHAR(255)             NOT NULL DEFAULT '',
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMIT;