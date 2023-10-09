
curl -d '{"sysid":"25b1e783-6a4f-4825-afb8-c667eaa195ce", "identifier":"ISS-1", "summary_brief":"Test issue #1", "summary_long":"This is a test issue, the first one.", "problem_description":"This is a problem description.", "work_around": "This is a work-around for the problem.", "steps_to_reproduce":"These are the steps to reproduce the problem: 1. Launch the application. 2. Press the button. 3. ISSUE: Here is the problem.", "kind":"b64f6f23-9094-49f6-8e25-cded27bfc351"}' -H "Content-Type: application/json" -X POST http://localhost:8080/issue

# curl -d '{"sysid":"38d946c4-a67a-4d05-9071-acde4d50665f", "identifier":"ISS-2", "summary_brief":"Test issue #2", "summary_long":"This is a test issue, the second one."}' -H "Content-Type: application/json" -X POST http://localhost:8080/issue
