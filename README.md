# AdventOfCode-2024
Solutions to Advent of code 2024: https://adventofcode.com/2024/

## How to Run the code

### 1. Clone the Repository
Clone the repository and navigate into the project directory:
```bash
git clone <repository-url>
cd <project-directory>
```

### 2. Install dependencies
Run `go mod tidy` from the root of the project.

### 3. Fetch your input data
To fetch the input data for Advent of Code, follow these detailed steps:

#### Grab Your Session Cookie
The script requires your session cookie from the Advent of Code website to authenticate and fetch the input data. Here’s how you can retrieve your session cookie:

1. Open your browser and navigate to the [Advent of Code website](https://adventofcode.com).
2. Log in using your preferred account (GitHub, Google, etc.).
3. Open the browser’s developer tools:
    - On Windows/Linux: Press `Ctrl+Shift+I`.
    - On Mac: Press `Cmd+Option+I`.
4. Navigate to the **Application** tab in the developer tools.
5. In the left sidebar, find **Storage > Cookies**, and select the `https://adventofcode.com` domain.
6. Look for the cookie named `session`. Copy the value of this cookie.

#### Create a `.env` File
You need to store your session cookie securely in a `.env` file at the root of the project. Follow these steps:

1. Create a new file named `.env` in the root of your project.
2. Add the following line to the file: `SESSION_ID=<your-session-cookie>`

#### Run the script

From the root of the project run `./scripts/fetch_input.sh <day> <output-file>`

- `day` should be the day you want to grap input for
- `output-file` should be the full path to the file you want to dump the fetched input to. Note: the file should be named `input.txt`
