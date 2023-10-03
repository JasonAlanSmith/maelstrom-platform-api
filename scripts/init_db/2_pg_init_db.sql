-- Create the Maelstrom Platform API databases.
-- PostgreSQL 16

CREATE DATABASE maelstrom
       OWNER		maelstromapi
       TABLESPACE	maelstrom;
       
CREATE DATABASE maelstrom_archive
       OWNER		maelstromapi
       TABLESPACE	maelstrom_archive;
       
CREATE DATABASE maelstrom_audit
       OWNER		maelstromapi
       TABLESPACE	maelstrom_audit;
       
CREATE DATABASE maelstrom_dw
       OWNER		maelstromapi
       TABLESPACE	maelstrom_dw;
       
CREATE DATABASE maelstrom_log
       OWNER		maelstromapi
       TABLESPACE	maelstrom_log;

CREATE DATABASE maelstrom_staging
       OWNER		maelstromapi
       TABLESPACE	maelstrom_staging;
