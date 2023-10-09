# ChangeLog

2023-09-27  Jason Alan Smith  <smith.jason.alan.me@gmail.com>

	Initial API
	
	* main.go: Implement simple API with single endpoint (/ping).

2023-09-30  Jason Alan Smith <smith.jason.alan.me@gmail.com>

	Add endpoint to create an issue
	
	* database/database.go: Connect to a PostgreSQL database.
	
	* main.go: Add createIssue and Issue struct. Add call to
	connect to database and route to main.

2023-10-01  Jason Alan Smith <smith.jason.alan.me@gmail.com>

	Add endpoint to get all issues
	
	* main.go: Add getIssues. Add route to main.

	Add endpoint to update an issue.
	
	* main.go: Add updateIssue. Add route to main.
	
2023-10-02  Jason Alan Smith <smith.jason.alan.me@gmail.com>

	Add endpoint to patch an issue
	
	* main.go: Add patchIssue and mergeIssue. These are two
	methods of patching a resource. Add route for patchIssue
	to main.
	
	Add endpoint to delete an issue
	
	* main.go: Add deleteIssue. Add route to main.

2023-10-03  Jason Alan Smith <smith.jason.alan.me@gmail.com>

	Add endpoint to get an issue
	
	* main.go: Add getIssueById. Add route to main.

2023-10-04  Jason Alan Smith <smith.jason.alan.me@gmail.com>

	Fix bug where http PUT route was using POST
	
	* main.go: Change POST to PUT in main.

	Refactor route handler names to match http methods
	
	* main.go: Rename functions to include http method name.
	Update main to refer to new function names.

2023-10-05  Jason Alan Smith <smith.jason.alan.me@gmail.com>

	Add logging using logger
	
	main.go: Update all functions to use a logger replacing
	print statements.

2023-10-07  Jason Alan Smith <smith.jason.alan.me@gmail.com>

	Add database infrastructure
	
	* database.go: Add support to read .env variables for
	connecting to the six databases. Set the database role
	passwords based on environment variables.
	
	* main.go: Use the new connections setup in database.go.

2023-10-08  Jason Alan Smith <smith.jason.alan.me@gmail.com>

	Build out issue relation
	
	* main.go: Build out Issue struct to match database
	columns. Add validation required to Issue struct fields.
	Add validation check and new function Unmarshal to
	return 400 Bad Request if not all fields provided in
	JSON request.
	
	* scripts/init_db/3_pg_init_db_schema_maelstrom.sql: Build
	out schema script for issue relation.
	
	* scripts/init_db/9_pg_init_db_maelstrom_procedures.sql: Build
	out function definitions to match issue relation definition.
