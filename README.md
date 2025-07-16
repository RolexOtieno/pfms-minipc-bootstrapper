PFMS MINIPC BOOTSTRAPER
This is a lightweight backend service written in Go for securely bootstrapping PFMS MiniPCs in the field. This server validates MiniPC devices and delivers the correct software installer based on their request — without requiring user tokens, manual input, or interactive authentication.


Overview
The PFMS MiniPC Bootstrapper is designed to support automated setups in remote fuel station environments. It provides:
-  Self-validation of MiniPCs using `deviceId`
-  Secure distribution of install scripts (version-controlled)
-  Static file hosting for `.sh` installer scripts
-  Request logging for monitoring and auditing
-  No user interaction required (fully backend-driven)

How It Works
1. A MiniPC sends a POST request to `/init` with its `deviceId`, `os`, and requested `version`.
2. The backend checks whether the device is authorized.
3. If valid, it responds with a download URL pointing to the correct installer script.
4. The MiniPC then downloads and executes the script to complete installation.

Example of a Request
when its run it should show:
✅ Server running on http://localhost:8080...

Test using curl the powershell version

curl -Method POST http://localhost:8080/init `
  -Headers @{"Content-Type"="application/json"} `
  -Body '{"deviceId":"MINIPC_123456","os":"linux","version":"v1.0.0"}'




