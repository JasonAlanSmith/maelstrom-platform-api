-- Create the maelstrom_log database relations.
-- PostgreSQL 16

-- Database roles are created at the database
-- cluster level.
CREATE ROLE maelstrom_log
WITH
	LOGIN
	NOSUPERUSER
	NOCREATEDB
	NOCREATEROLE
	NOINHERIT
	NOBYPASSRLS
	PASSWORD	NULL;

-- Database schemas are created at the
-- database level.
CREATE SCHEMA
IF NOT EXISTS maelstrom_log
AUTHORIZATION maelstrom_log;

-- Grant permissions to role maelstrom_log.
GRANT SELECT, INSERT, UPDATE, DELETE
ON ALL TABLES
IN SCHEMA maelstrom_log
TO maelstrom_log;

-- Create the TBD relation in the maelstrom_log
-- schema in the maelstrom_log database.
