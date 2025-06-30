HTTP_TUI

A Terminal-based HTTP client built with Go, Bubble Tea, and Charm CLI tools.

Features

TUI built with Bubble Tea and Lip Gloss
Supports HTTP methods: GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS
Real-time response viewer with status code
Keyboard navigation across tabs and fields
Add custom headers easily
Installation

git clone https://github.com/jayanthsharabu/HTTP_TUI.git
cd HTTP_TUI
go mod tidy
go build -o HTTP_TUI
./HTTP_TUI
Overview

URL Tab – Enter API endpoint and select HTTP method
Headers Tab – Add custom request headers
Body Tab – Enter request body for POST/PUT/PATCH
Response Tab – View formatted response and headers
Keyboard Shortcuts

Navigation
Ctrl + → / Ctrl + L – Next tab
Ctrl + ← / Ctrl + H – Previous tab
Tab / Alt + → / Alt + L – Next field within tab
Alt + ← / Alt + H – Previous field within tab
↑ / ↓ or j / k – Move through HTTP methods
Enter – Select HTTP method
Actions
Ctrl + S – Send the HTTP request
Ctrl + A – Add header (in Headers tab)
Ctrl + X – Clear all headers
Ctrl + W – Save current request to .quest file
Ctrl + R – Load saved request
Shift + ← / → – Switch response sub-tabs (Headers/Body)
Esc – Cancel load dialog
/ – Search saved requests (in load dialog)
? – Toggle help menu
q or Ctrl + C – Quit application
Making Your First Request

Enter URL: Type your endpoint (e.g. https://jsonplaceholder.typicode.com/posts/1)
Select Method: Press Tab, then use ↑/↓ to choose (default: GET)
Add Headers (Optional): Go to Headers tab and input key-value pairs
Add Body (Optional): For POST/PUT/PATCH, go to Body tab and enter your JSON/XML
Send Request: Press Ctrl + S
View Response: Response will appear in the Response tab
Upcoming Features

WebSocket support
Built-in Mock Server (like Postman)
