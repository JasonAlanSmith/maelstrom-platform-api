
curl -d '[{"op": "replace", "path": "/summary_long", "value": "This is yet another issue."}]' -X PATCH http://localhost:8080/issue/2 -H "Content-Type: application/json"
