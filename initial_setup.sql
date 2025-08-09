CREATE SCHEMA bot;
CREATE TABLE bot.guild_members (
    id BIGSERIAL PRIMARY KEY,
    user_id VARCHAR(21) NOT NULL,
    guild_id VARCHAR(21) NOT NULL,
    xp BIGINT DEFAULT 0 NOT NULL,
    level INT DEFAULT 0 NOT NULL,
    last_message_timestamp TIMESTAMPTZ,
	last_message TEXT,
    CONSTRAINT unique_member UNIQUE (user_id, guild_id)
);
