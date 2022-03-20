# grep_in_go
Implementing grep command in go language.

Learning Doc - https://docs.google.com/document/d/1eHkmxWRycg5_U4NZ-HPVv5x4KKU-duRUXMNEveARr3A/edit

I/O:
1. Input - word, filename/filepath
2. Output - word matching in the given file

Example:
1. randomfile.txt -> a. pranavwadekar
                     b. Pranavisotwork
                     c. ppranavisgoodatpeople.
2. Input - ('pranav', 'randomfile.txt')
3. Output - pranavwadekar
            ppranavisgoodatpeople


For naive_main.go
Assumption: 
1. Word matching is case sensitive.
2. For naive_main input would be a file which would be in the current directory
3. Only .txt files will get consider

Test Cases:
1. throw err if file does not exists, does not have permission
2. print null if word does not matches
3. print all words from file which matches

I/O:
I - "some_text.txt" (Expecting in the current directory)
O - ["Ppranavisotwork" "pranavwadekar"]

For goroutine_single_file_main.go
Assumption: 
1. Word matching is case sensitive.
2. For this input would be a file which would be in the current directory
3. It will read the file divide the content and call different concurrent processes. 
4. This approach would be faster than the naive one.
5. Only .txt files will get consider

Test Cases:
1. throw err if file does not exists, does not have permission
2. print null if word does not matches
3. print all words from file which matches

I/O:
I - "some_text.txt" (Expecting in the current directory)
O - ["Ppranavisotwork" "pranavwadekar"]

For main.go

Assumption: 
1. Word matching is case sensitive.
2. For this input would be a absolute path of the directory or file.
3. It will recursively read all the files from the folder and subfolder as well and find the match.
4. Only .txt files will get consider

Test Cases:
1. throw err if file does not exists, does not have permission
2. print null if word does not matches
3. print all words from file which matches

I/O:
I - "/home/username/grep_in_go/"
O - ["pranavwadekar" "Ppranavisotwork" "pranavwadekarsampl1.1" "pranavwadekarsampl1.2"]

I have created a greptest.go file test cases evaluation/TDD.
In this I have created only 1 test for matching output of the Grep function, but we can add many of
test cases and scenarios.

We can Also add table driven tests.
Reference - https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package

The code coverage is upto 68.3 percent.
(TODO - Add test case for main function)

![alt text](https://github.com/pranavwadekar98/grep_in_go/blob/develop/images/code_coverage.png?raw=true)
