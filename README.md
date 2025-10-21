# zap

This is a simple CLI project that allows users to perform a few basic Unix-based operations, including recursive copy.

## Features

- Recursive copy of files and directories
- Lightweight and easy to use

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/akhakpouri/bash-cli.git
   ```

2. Navigate to the project directory:
   ```bash
   cd bash-cli
   ```

3. Build the project:
   ```bash
   go build
   ```

4. Run the CLI:
   ```bash
   zap <command> {arguments}
   ```

## Usage

To perform a recursive copy, use the following command:
   ```bash
   zap copy -src <source-path> -dest <destination-path>
   ```
Replace `<source-path>` with the directory or file you want to copy and `<destination-path>` with the target location.

## Requirements
- Go programming language installed (version 1.18 or higher)

## License
This project is licensed under the terms of the [MIT License](LICENSE).

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request for any improvements or additional features.
This README provides a comprehensive overview of your CLI project and its functionality.