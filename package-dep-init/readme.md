run the go file and check out in the logs, the order in which the init methods of each file are called

HINT: the go tool sorts the go files before invoking the compiler. Hence the init methods will be called based on the
dependency go file's name. Change the file names of the dependencies to understand package initialization.