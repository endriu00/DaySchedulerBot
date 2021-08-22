-- Schema for telegram bot database

-- Telegram users table
CREATE TABLE `telegram_user` (
    `chatid` bigint(11) NOT NULL,
    `username` varchar(150) NOT NULL,
    PRIMARY KEY (`chatid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Telegram users registered to the bot';

-- Telegram user event, with event description and time decided by the user
CREATE TABLE `event` (
    `username` varchar(150) NOT NULL,
    `time` datetime NOT NULL,
    `description` varchar(1000) NOT NULL,
    PRIMARY KEY (`username`, `time`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Event for a telegram user at a given time';

-- Table for joining telegram_users and event
CREATE TABLE `agenda` (
    `chatid` bigint(11) NOT NULL,
    `username` varchar(150) NOT NULL,
    `time` datetime NOT NULL,
    PRIMARY KEY (`chatid`, `username`, `time`)
) ENGINE=InnoDB COMMENT='Agenda for joining tables telegram_users and event';