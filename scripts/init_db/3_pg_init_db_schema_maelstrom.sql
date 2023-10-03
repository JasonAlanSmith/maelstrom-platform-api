-- Create the maelstrom database relations.
-- PostgreSQL 16

CREATE TABLE IF NOT EXISTS maelstrom.issue (
       sysid		uuid		DEFAULT uuid_generate_v4(),	-- System identifier
       identifier	varchar(16)	NOT NULL,			-- 
       summary_brief	varchar(50)	NOT NULL,			-- 
       summary_long	varchar(72)	NOT NULL,			-- 
)
