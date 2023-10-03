-- Create schemas for the Maelstrom API databases.
-- PostgreSQL 16

CREATE SCHEMA IF NOT EXISTS maelstrom AUTHORIZATION maelstromapi
CREATE SCHEMA IF NOT EXISTS maelstrom_archive AUTHORIZATION maelstromapi
CREATE SCHEMA IF NOT EXISTS maelstrom_audit AUTHORIZATION maelstromapi
CREATE SCHEMA IF NOT EXISTS maelstrom_dw AUTHORIZATION maelstromapi
CREATE SCHEMA IF NOT EXISTS maelstrom_log AUTHORIZATION maelstromapi
CREATE SCHEMA IF NOT EXISTS maelstrom_staging AUTHORIZATION maelstromapi
