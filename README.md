# Password grabber for Mozilla Lockwise

This command-line allows you to decrypt passwords that Mozilla Lockwise stores.

# Installation

You just need to have Go compile to compile this command-line tool and GCC compiler because 
github.com/mattn/go-sqlite3 dependency need it.

# Arguments

##-path 

Path to Mozilla Firefox Profiles Folder. In windows it looks like 
`C:\Users\{Username}\AppData\Roaming\Mozilla\Firefox\Profiles` by default.

##-o

Output. You could enter the path to file where you want to store credentials data. 
By default program will print data to stdout.

##-pwd

With this flag you could specify the master password if it set.

