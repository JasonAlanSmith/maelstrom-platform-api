-- Create the Maelstrom Platform API tablespaces.
-- PostgreSQL 16

CREATE TABLESPACE maelstrom
       LOCATION '/var/pg/data1';
       
CREATE TABLESPACE maelstrom_archive
       LOCATION '/var/pg/data2';
       
CREATE TABLESPACE maelstrom_audit
       LOCATION '/var/pg/data3';
       
CREATE TABLESPACE maelstrom_dw
       LOCATION '/var/pg/data1';
       
CREATE TABLESPACE maelstrom_log
       LOCATION '/var/pg/data2';

CREATE TABLESPACE maelstrom_staging
       LOCATION '/var/pg/data3';
