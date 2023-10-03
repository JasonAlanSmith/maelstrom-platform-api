
curl -d '{"sysid":1, "identifier":"ISS-1", "summary_brief":"Test issue #1", "summary_long":"This is a test issue, the first one."}' -H "Content-Type: application/json" -X POST http://localhost:8080/issue

curl -d '{"sysid":2, "identifier":"ISS-2", "summary_brief":"Test issue #2", "summary_long":"This is a test issue, the second one."}' -H "Content-Type: application/json" -X POST http://localhost:8080/issue
