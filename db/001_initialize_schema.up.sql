CREATE TABLE `user` (
    `user_id` VARCHAR(36) NOT NULL,
    `first_name` VARCHAR(32) NOT NULL,
    `last_name` VARCHAR(32) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`)
);

CREATE TABLE `wallet` (
    `wallet_id` VARCHAR(36) NOT NULL,
    `name` VARCHAR(32) NOT NULL,
    `desc` VARCHAR(32),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`wallet_id`)
);
 
CREATE TABLE `user_wallet` (
    `user_id` VARCHAR(36) NOT NULL,
    `wallet_id` VARCHAR(36) NOT NULL,
    `balance` DECIMAL(12,2) NOT NULL,
    PRIMARY KEY (`user_id`, `wallet_id`)
);

CREATE TABLE `card` (
    `card_id` VARCHAR(36) NOT NULL,
    `name` VARCHAR(32) NOT NULL,
    `desc` VARCHAR(32),
    `daily_limit` DECIMAL(12,2) NOT NULL,
    `monthly_limit` DECIMAL(12,2) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`card_id`)
);

CREATE TABLE `wallet_card` (
    `user_id` VARCHAR(36) NOT NULL,
    `wallet_id` VARCHAR(36) NOT NULL,
    `balance` DECIMAL(12,2) NOT NULL,
    PRIMARY KEY (`user_id`, `wallet_id`)
);

CREATE TABLE `team` (
    `team_id` VARCHAR(36) NOT NULL,
    `name` VARCHAR(32) NOT NULL,
    `desc` VARCHAR(32),
    `wallet_id` VARCHAR(36) NOT NULL,
    `daily_limit` DECIMAL(12,2) NOT NULL,
    `monthly_limit` DECIMAL(12,2) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`team_id`)
);

CREATE TABLE `user_team` (
    `user_id` VARCHAR(36) NOT NULL,
    `team_id` VARCHAR(36) NOT NULL,
    PRIMARY KEY (`team_id`,`user_id`)
);

