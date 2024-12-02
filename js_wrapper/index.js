#!/usr/bin/env node

import { execSync } from "child_process";
import { argv} from "process";

try {
    const args = argv.slice(2);
    // Check if the user provided the required arguments
    if (args.length === 0) {
        console.error("Usage: nodewatcher <file-to-watch>");
        process.exit(1);
    }
    // Specify the full path to the Go binary
    const binary = "E:/go_automation/nodewatcher/js_wrapper/nodewatcher.exe";
    const command = `${binary} ${args.join(" ")}`;
    console.log(command)
    execSync(command, { stdio: "inherit" });
} catch (error) {
    console.error("Error running nodewatcher:", error.message);
    process.exit(1); // Exit with an error code
}