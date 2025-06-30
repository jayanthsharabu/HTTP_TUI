# HTTP TUI 🚀

A modern terminal-based HTTP client built with Go, featuring a beautiful TUI powered by Bubble Tea and Charm CLI tools.

## ✨ Features

- 🎨 **Beautiful TUI** - Built with Bubble Tea and Lip Gloss
- 🔥 **Full HTTP Support** - GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS
- ⚡ **Real-time Response** - Instant response viewer with status codes
- ⌨️ **Keyboard Navigation** - Seamless tab and field navigation
- 📝 **Custom Headers** - Easy header management
- 💾 **Save & Load** - Store requests as .quest files

## 🚀 Quick Start

```bash
git clone https://github.com/jayanthsharabu/HTTP_TUI.git
cd HTTP_TUI
go mod tidy
go build -o HTTP_TUI
./HTTP_TUI
```

## 📖 How to Use

### Tabs Overview
| Tab | Purpose |
|-----|---------|
| **URL** | Enter endpoint and select HTTP method |
| **Headers** | Add custom request headers |
| **Body** | Enter request body (POST/PUT/PATCH) |
| **Response** | View formatted response and headers |

### Making Your First Request

1. **Enter URL**: Type your endpoint
2. **Select Method**: Use `Tab` then `↑/↓` to choose method
3. **Add Headers** _(optional)_: Switch to Headers tab, add key-value pairs
4. **Add Body** _(optional)_: For POST/PUT/PATCH, go to Body tab
5. **Send Request**: Press `Ctrl + S`
6. **View Response**: Check the Response tab for results

## ⌨️ Keyboard Shortcuts

### Navigation
| Key | Action |
|-----|--------|
| `Ctrl + →` / `Ctrl + L` | Next tab |
| `Ctrl + ←` / `Ctrl + H` | Previous tab |
| `Tab` / `Alt + →` / `Alt + L` | Next field |
| `Alt + ←` / `Alt + H` | Previous field |
| `↑/↓` or `j/k` | Move through HTTP methods |
| `Enter` | Select HTTP method |

### Actions
| Key | Action |
|-----|--------|
| `Ctrl + S` | Send HTTP request |
| `Ctrl + A` | Add header (Headers tab) |
| `Ctrl + X` | Clear all headers |
| `Ctrl + W` | Save request to .quest file |
| `Ctrl + R` | Load saved request |
| `Shift + ←/→` | Switch response sub-tabs |
| `Esc` | Cancel load dialog |
| `/` | Search saved requests |
| `?` | Toggle help menu |
| `q` / `Ctrl + C` | Quit application |

## 🔮 Coming Soon

- 🌐 **WebSocket Support** - Real-time communication
- 🖥️ **Built-in Mock Server** - Local testing environment

---

