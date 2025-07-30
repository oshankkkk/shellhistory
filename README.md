# Understanding and using Golang
went to google and search golang and then when to the Docs
1st tut about installing go, did it through dnf ( linux pacman)
Then print helo world using Go
In Go theres a built in package.json thing yk like pom.xml in java to manage the dependancies.
Its called go.mod, and we need to set it up before we code
To do that we just put go init package/module or dir/file its the same thing
This makees a go.mod file thats links to your module/file so it manages its versions n shit
After that we have to write code for that we just write helo world in a main function with the fmt package
 **"fmt" not 'fmt'**
It lowkey like java with the system.out stuff
Then we run the file with go run . ( the dot means all the file like in git commit . )
In go there is this thing called export names in functions
If a function name is capitalisee, you can call it from another package in the same level in the filesystem
In go when you intilaze a variable you do it with 

var myvar string
myvar="oshan the great"
OR
myvar:="oshan the great"
myvar automatically set the type accoding to the 1st value

Also in go you should specifiy the return type and the arguments type as well

func FunctionName(args...type ) ReturnType {
    // function body
}

yes, the return type is at the end of the variable
Also in go you should specifiy the main function ( old python habbits lol)  
In Golang package is required, go has this package mindset thing and i think its cruical for ist compiler to have that so ye.


# Figureing out how to make the actual CLI tool

I wanna get my hands on the shell history
My code should go to the OS user folder in home and then check for the bash_history file
(what if the user put it in a different folder or renamed it? well fix that in future)
After that i should read the file 
When user types it should chose the matches of in the file and display like a autocorrect

