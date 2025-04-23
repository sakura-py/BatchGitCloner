# BatchGitCloner Documentation

- [English Documentation](README.md)
- [中文文档](README-ZH.md)

## Project Overview
Are you tired of dealing with the hassle of cloning multiple repositories for microservices projects? BatchGitCloner is a Go language tool designed to clone multiple Git repositories in bulk. By reading a JSON configuration file, it can clone multiple specified branches of repositories to a unified local directory, streamlining the management of multiple repositories.

## Features
- **Batch Operation**: Supports configuring multiple repository cloning tasks at once, improving efficiency.
- **Flexible Configuration**: You can specify repository URL, target branch, and local storage path through a JSON file.
- **Smart Handling**: Automatically creates directory structures, handles existing directories, and outputs detailed execution logs.

## Usage

### Environment Preparation
Ensure that you have installed the Go language environment (Go 1.20 or higher) and the Git tool.

### Configure JSON File
Create a JSON file (e.g., `repos.json`) and fill in the repository information in the following format:
```
{
"basePath": "D:\\Code\\go",
"repos": [
{
"url": "https://github.com/sakura-py/BatchGitCloner.git",
"branch": "master",
"path": ""
},
{
"url": "https://github.com/sakura-py/BatchGitCloner.git",
"branch": "master",
"path": ""
}
]
}
```
- basePath: A unified project directory path, where all repositories will be cloned.
- repos: A list of repositories, containing detailed information for multiple repositories.
- url: The URL of the Git repository (must include the .git suffix).
- branch: The name of the branch to be cloned from the repository.
- path: The sub-path of the repository in the local environment (optional). If empty, the repository name will be used as the subdirectory.

### Run the Program
In the command line, run the following command to start the BatchGitCloner tool:
```
go run main.go
```
or

```
go build main.go
```
The program will prompt you to enter the JSON file path. After entering the path of the `repos.json` file you created, the program will clone the repositories according to the configuration.

## Code Structure Explanation

### Main Structures
- RepoInfo: Used to store information about a single repository, including the repository URL, branch name, and local path.
- Config: Stores the entire configuration information, including the unified project directory and the list of repositories.

### Core Logic
- **Read Configuration**: Obtains the JSON file path through standard input, reads and parses the file content into the Config structure.
- **Directory Preparation**: Checks and creates the unified project directory to ensure the validity of the cloning target path.
- **Clone Operation**: Iterates through the repository list, generates the complete clone path based on each repository's configuration, executes the Git clone command, and outputs the command execution results in real-time.

## Frequently Asked Questions

Q1：How to handle errors during the cloning process?
A1：The program handles each repository's cloning operation separately. Even if cloning a particular repository fails, it will not affect the cloning of other repositories. Error information is recorded in the logs, making it easier to troubleshoot issues.

Q2：Can the directory structure be modified?
A2：Yes, you can customize the local storage path for each repository by modifying the path field in the JSON configuration file.

Q3：How to update cloned repositories?
A3：Currently, the tool is primarily designed for cloning operations. If you need to update cloned repositories, you can manually execute the Git pull command in the corresponding repository directory or extend the functionality of this tool to support repository update operations.
