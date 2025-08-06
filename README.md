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
The file reading stuff are done by the Os package, cause it interacts with the os which has the file systems n stuff

Also btw in go "" are use in string literals like yk "helo world" and '' are used in rune literals 
like '\n' n stuff

File handling in golang is pretty simple when  you understand it like in python when you open a file in write mode it automatically creates a file it its not there and overwrites it when used more than once write.
All those file management tasks like making the file it its new and overwriting it is done by one mode 
See in places like this you can see golangs main theme, it makes everything simple, like instead of puting all those features on 1 Golangs write mode which is walled WRONLY( its technially a constant flag thing in the os package) and it only writes stuff in a like, it doesnt make a new file or erase the old data and write a new one instead it has other flags for those too. 
I tbh think go is a better lanugage to learn as your 1st lanugage than python but this might be a pretty shalow take since ive only being playing around with this or less than a week.

Go has 4 types of numbers
ints floats uints(only postive ints) complex numbers(whatever that means)
yk in java theres int and long ryt, like a thresshold for ints same goes here theres int32 int64
float64..
this is good for you if your worried bout memeory 
00010 ten but more memrmory
010 also ten but less memory

const doesnt work with :=

why is go doesnt have OOP and not OOP heavy
the structs in Go is kinda like  js objects eveywhere thing but we also do the proper 
classes shit 
# Figureing out how to make the actual CLI tool

I wanna get my hands on the shell history
My code should go to the OS user folder in home and then check for the bash_history file
(what if the user put it in a different folder or renamed it? well fix that in future)
After that i should read the file 
When user types it should chose the matches of in the file and display like a autocorrect
listen to the user inputs, bash doesnt have event listeners, bash only knows what the user types after the \n through thats standed input
whats a Deamon, its a fancy way of says background process(i dont think there much more to it tbh)
its mostly a stuff like event listners yk somthing thats always runnig and that the user doesnt see or like have direct access to
im gonna write a that listen to what the user types and matches it with whats in the shellhistory
To make a daemon in Go we just have to make a go file and make infinete loop that runs for ever.
We dont print in that file (cause then the user will know the file is running
aka user ta horen thama run karanne file eka)
So basically any file that the user doesnt know that runs is a DAEMON.
We make a text file call a log file, then we put what ever we are supposed to print ( but cant cause the user is watching) to that file, this helps in debugging stuff. Thats a Deamon in a nutshell.
Making a fuzzy finder on my own to search and display the cmds
WTF is a fuzzy finder and finder and how it works 


# Learning DSA
A fuzzy finder uses a very complex algorithm to make it work with typos n stuff so imma start with the prefix finder 1st.
Its not fuzzy like if yo spell dev, it wil say developer but if you spell del it wont casue it just search for the words starting with that on memeory and it cant deal with typos
Okay lets learn bout trees then.
whats a tree it is a data srtucture that has values and those value has nother values in a hierachyal way thing
like the top most part or the like head of the tree is the root, thats where eveything branch off to.
Then trees just branch naturally 
the values of the trees are called nodes,and there are  parent nodes and child nodes and nodes that dont have any children are called nodes leaves. 
There is a tree type called trie also called a prefix trie and this is not a recursive one like most other trees.(technially you can make  it using recursion but no need to it will only slow it down )
I dont understand why so imma look that up 


maps vs list:
maps are key based so its import when you gotta indivaltly recognize the data
list they just have any index so so they dont have a indivital name or anything, so if you car bout that then use list
 
maps are fast when searching compared to list, O(1) when searching with a key and in a list its just O(n) for everything thats linear which is not good when the data is incerasing cause its linear(duh)
we use List when the order matter and if you need duplicated, cause in a map you cant have duplicates. 

