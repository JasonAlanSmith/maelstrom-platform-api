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
