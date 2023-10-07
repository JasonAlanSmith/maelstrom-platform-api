-- Create the maelstrom database procedures and functions.
-- PostgreSQL 16

DROP FUNCTION maelstrom.select_issue_all;
CREATE OR REPLACE FUNCTION
maelstrom.select_issue_all()
RETURNS TABLE (
       iss_sysid		uuid,
       iss_identifier		varchar(16),
       iss_summary_brief	varchar(72),
       iss_summary_long		varchar(128)
)
LANGUAGE plpgsql
AS $$
BEGIN
	RETURN QUERY
	       SELECT  sysid,
	       	       identifier,
	               summary_brief,
	               summary_long
		    
	  	 FROM  issue;
END;
$$;


DROP FUNCTION maelstrom.select_issue_by_sysid;
CREATE OR REPLACE FUNCTION
maelstrom.select_issue_by_sysid(targ_sysid uuid)
RETURNS TABLE (
       iss_sysid		uuid,
       iss_identifier		varchar(16),
       iss_summary_brief	varchar(72),
       iss_summary_long		varchar(128)
)
LANGUAGE plpgsql
AS $$
BEGIN
	RETURN QUERY
	       SELECT  sysid,
		       identifier,
		       summary_brief,
		       summary_long
		       
	  	 FROM  issue
                       
	 	WHERE  sysid = targ_sysid;
END;
$$;
