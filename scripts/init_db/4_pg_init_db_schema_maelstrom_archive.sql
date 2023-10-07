-- Create the maelstrom_archive database relations.
-- PostgreSQL 16

-- Database roles are created at the database
-- cluster level.
CREATE ROLE maelstrom_archive
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
IF NOT EXISTS maelstrom_archive
AUTHORIZATION maelstrom_archive;

-- Grant permissions to role maelstrom_archive.
GRANT SELECT, INSERT, UPDATE, DELETE
ON ALL TABLES
IN SCHEMA maelstrom_archive
TO maelstrom_archive;

-- Create the TBD relation in the maelstrom_archive
-- schema in the maelstrom_archive database.
