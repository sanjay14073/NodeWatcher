# NodeWatcher

**NodeWatcher** is a lightweight file-watching tool that automatically restarts a Node.js application whenever changes are detected in the files you're working on. It is built with Go and comes with a convenient Node.js wrapper script.

## Features

- Watches for file changes in your project directory.
- Automatically restarts Node.js applications on changes.
- Allows manual restarts by entering `rs` in the terminal.

## Prerequisites

- [Node.js](https://nodejs.org/) and npm installed.

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd nodewatcher
   ```

2. Replace the binary path in `index.js`:
   Open `index.js` and update the `binary` variable with the full path to the provided `nodewatcher` binary. Example:

   ```javascript
   const binary = "E:/path/to/nodewatcher.exe"; // Update this path to the actual binary location
   ```

3. Set up the wrapper script globally:
   ```bash
   npm link
   ```

   This will make the `nodewatcher` command available globally.

4. Verify the setup:
   ```bash
   nodewatcher --help
   ```

## Usage

Run the `nodewatcher` command followed by the file or script you want to monitor:

```bash
nodewatcher <file-to-watch>
```

### Example:
To watch a `server.js` file:
```bash
nodewatcher server.js
```

## Manual Restart

While `nodewatcher` is running, you can manually restart the application by typing:
```
rs
```

## Contributing

Feel free to fork the repository and submit pull requests to improve or extend functionality.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

