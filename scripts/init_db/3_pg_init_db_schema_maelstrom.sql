-- Create the maelstrom database schema.
-- PostgreSQL 16

-- Database roles are created at the database
-- cluster level.
CREATE ROLE maelstrom
WITH
	LOGIN
	SUPERUSER
	NOCREATEDB
	NOCREATEROLE
	NOINHERIT
	NOBYPASSRLS
	PASSWORD	NULL;

-- Database schemas are created at the
-- database level.
CREATE SCHEMA
IF NOT EXISTS maelstrom
AUTHORIZATION maelstrom;

-- Grant permissions to role maelstrom.
GRANT SELECT, INSERT, UPDATE, DELETE
ON ALL TABLES
IN SCHEMA maelstrom
TO maelstrom;

-- Extensions are created at the
-- database level.  This extension is
-- needed for uuid column support.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the Issue relation in the maelstrom
-- schema in the maelstrom database.
CREATE TABLE IF NOT EXISTS maelstrom.issue (
       sysid		uuid		DEFAULT uuid_generate_v4(),	-- System identifier
       identifier	varchar(16)	NOT NULL,			-- 
       summary_brief	varchar(50)	NOT NULL,			-- 
       summary_long	varchar(72)	NOT NULL			--
)
