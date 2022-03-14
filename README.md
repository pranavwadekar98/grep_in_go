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

Assumption: 
1. Word matching is case sensitive.

Test Cases:
1. return err if file does not exists
2. return null if word does not matches
3. return all words from file which matches