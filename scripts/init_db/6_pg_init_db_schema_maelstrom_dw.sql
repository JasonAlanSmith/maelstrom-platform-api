-- Create the maelstrom_dw database relations.
-- PostgreSQL 16

-- Database roles are created at the database
-- cluster level.
CREATE ROLE maelstrom_dw
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
IF NOT EXISTS maelstrom_dw
AUTHORIZATION maelstrom_dw;

-- Grant permissions to role maelstrom_dw.
GRANT SELECT, INSERT, UPDATE, DELETE
ON ALL TABLES
IN SCHEMA maelstrom_dw
TO maelstrom_dw;

-- Create the TBD relation in the maelstrom_dw
-- schema in the maelstrom_dw database.
