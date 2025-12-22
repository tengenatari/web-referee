Invoke-RestMethod -Uri "http://127.0.0.1:8000/api/v1/user?id=123&name=John%20Doe&tigrId=456&rating=85" `
-Method Post `
-ContentType "application/json" `
-Headers @{"Accept" = "application/json"}