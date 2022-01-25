CREATE TABLE IF NOT EXISTS rwm_config (
    `config` TEXT
);

INSERT INTO `rwm_config` (`config`) values ('');

CREATE TABLE IF NOT EXISTS rwm_connection (
    `id` TEXT key,
    `connection` TEXT
);