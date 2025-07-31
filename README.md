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

Error handling in go is very simple, the functions that can fail yk like readers n stuff returns
2 values.the actual value and the Error
If the functions works fine the error value is nill(null in go) so you can just put a If statement
to check if the error is null if yes no error very good, else error have very very bad.
WTF, ERROR HANDLING COULD BE THIS SIMPLE BUT OTHERS JUST FUCKING HAD TO USE TRY BLOCKS AND CATCHES AND THOSE PASSING THE EXCEPTION TO A HIGHER FUCTION THING JAVA DOES

In go we get the user input through os module, like system in in java
i think i should prolly check why these are so similar

Anyway we read theterminal input thorugh os.sdin method thing,like and a reader is returned by bufio which is the input out module i think   
in java we make a scanner with system in ryt, same here we make a reader( or whatever you wanna call it) that accespts whats returned to it by the bufio.Reader(os.sdin)

File reading
there are 2 ways of reading files in Golang and both are done by the os package 

we can read the whole file and pass it to the memory, and we print it to the terminal 
The other way is to only read line by line yk how we used to do it, with os package we use Open method and then we to open a file the file data object is returned and we can print it line by line using a for loop

to read a file in go we use the os module ii 

Also btw in go "" are use in string literals like yk "helo world" and '' are used in rune literals like '\n' n stuff 
# Figureing out how to make the actual CLI tool

I wanna get my hands on the shell history
My code should go to the OS user folder in home and then check for the bash_history file
(what if the user put it in a different folder or renamed it? well fix that in future)
After that i should read the file 
When user types it should chose the matches of in the file and display like a autocorrect

