# grep_in_go
Implementing grep command in go language.

I/O:
1. Input - word, filename/filepath
2. Output - word matching in the given file

Example:
1. randomfile.txt -> a. pranavwadekar
                     b. Pranavisokokatwork
                     c. ppranavisnotgoodatpeople.
2. Input - ('pranav', 'randomfile.txt')
3. Output - pranavwadekar
            ppranav is not good at people


For naive_main.go
Assumption: 
1. Word matching is case sensitive.
2. For naive_main input would be a file which would be in the current directory

Test Cases:
1. throw err if file does not exists, does not have permission
2. print null if word does not matches
3. print all words from file which matches

For goroutine_single_file_main.go
Assumption: 
1. Word matching is case sensitive.
2. For this input would be a file which would be in the current directory
3. It will read the file divide the content and call different concurrent processes. 
4. This approach would be faster than the naive one.

Test Cases:
1. throw err if file does not exists, does not have permission
2. print null if word does not matches
3. print all words from file which matches

For main.go

Assumption: 
1. Word matching is case sensitive.
2. For this input would be a absolute path of the directory or file.
3. It will recursively read all the files from the folder and subfolder as well and find the match.

Test Cases:
1. throw err if file does not exists, does not have permission
2. print null if word does not matches
3. print all words from file which matches

I have created a greptest.go file test cases evaluation/TDD.
In this I have created only 1 test for matching output of the Grep function, but we can add many of
test cases and scenarios.

We can Also add table driven tests.
Reference - https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package