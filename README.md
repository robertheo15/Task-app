## Task Tracker App
https://roadmap.sh/projects/task-tracker

Task tracker is a project used to track and manage your tasks. In this task, you will build a simple command 
line interface (CLI) to track what you need to do, what you have done, and what you are currently working on. 
This project will help you practice your programming skills, including working with the filesystem, handling user inputs, and building a simple CLI application.

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/robertheo15/Task-app.git
cd task-app
```


### Run the following command to build and run the project:
```
# to build the project
go build -o task-tracker

# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```
