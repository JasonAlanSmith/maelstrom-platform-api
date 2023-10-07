
curl -d '{"sysid":"25b1e783-6a4f-4825-afb8-c667eaa195ce", "identifier":"ISS-2", "summary_brief":"Test issue #2", "summary_long":"This is a test issue, the second one."}' -H "Content-Type: application/json" -X POST http://localhost:8080/issue

curl -d '{"sysid":"38d946c4-a67a-4d05-9071-acde4d50665f", "identifier":"ISS-3", "summary_brief":"Test issue #3", "summary_long":"This is a test issue, the third one."}' -H "Content-Type: application/json" -X POST http://localhost:8080/issue
