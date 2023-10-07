-- Create the maelstrom_audit database relations.
-- PostgreSQL 16

-- Database roles are created at the database
-- cluster level.
CREATE ROLE maelstrom_audit
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
IF NOT EXISTS maelstrom_audit
AUTHORIZATION maelstrom_audit;

-- Grant permissions to role maelstrom_audit.
GRANT SELECT, INSERT, UPDATE, DELETE
ON ALL TABLES
IN SCHEMA maelstrom_audit
TO maelstrom_audit;

-- Create the TBD relation in the maelstrom_audit
-- schema in the maelstrom_audit database.
