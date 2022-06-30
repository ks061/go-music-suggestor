/**
Final Project in Go!

Name: Kartikeya Sharma
Professor Wittie
CSCI 308
May 2022
*/

package main

import "fmt"
// import "reflect"
import "unsafe"

/**
Topic 1: Primitives

IMPORTANT NOTE: I am using poetic license NOT to make all the functions and variable names
musically themed. The primitive-types information would not be very apparent to readers. I'm just trying to communicate the basics of the types in this
language. After this question, I'll move towards that if that works. Thank you for understanding.
*/
func primitives() {
    fmt.Println("Topic 1: Primitives")

    var b bool
    var u8 uint8
    var u16 uint16
    var u32 uint32
    var u64 uint64
    var i8 int8
    var i16 int16
    var i32 int32
    var i64 int64
    var f32 float32
    var f64 float64
    var c64 complex64
    var c128 complex128
    var s string // immutable primitive type

    var by byte // uint8
    var r rune // int32
    var i int // int64 or 64 bits (signed)
    var ui uint // uint64 or 64 bits (unsigned)
    var uip uintptr // unsigned integer that can hold a pointer value (32 bits on a 32-bit system, and 64 bits on a 64-bit system)
    var e error // essentially a string

    // Citation used: [1]
    fmt.Println("Primitive type sizes below are in bytes.")
    fmt.Printf("Size of bool: %d\n", unsafe.Sizeof(b))  // prints 1
    fmt.Printf("Size of uint8: %d\n", unsafe.Sizeof(u8)) // prints 1
    fmt.Printf("Size of uint16: %d\n", unsafe.Sizeof(u16)) // prints 2
    fmt.Printf("Size of uint32 : %d\n", unsafe.Sizeof(u32)) // prints 4
    fmt.Printf("Size of uint64: %d\n", unsafe.Sizeof(u64)) // prints 8 
    fmt.Printf("Size of int8: %d\n", unsafe.Sizeof(i8)) // prints 1
    fmt.Printf("Size of int16: %d\n", unsafe.Sizeof(i16)) // prints 2
    fmt.Printf("Size of int32: %d\n", unsafe.Sizeof(i32)) // prints 4
    fmt.Printf("Size of int64: %d\n", unsafe.Sizeof(i64)) // prints 8
    fmt.Printf("Size of float32: %d\n", unsafe.Sizeof(f32)) // prints 4
    fmt.Printf("Size of float64: %d\n", unsafe.Sizeof(f64)) // prints 8
    fmt.Printf("Size of complex64: %d\n", unsafe.Sizeof(c64)) // prints 8
    fmt.Printf("Size of complex128: %d\n", unsafe.Sizeof(c128)) // prints 16
    fmt.Printf("Size of string: %d\n", unsafe.Sizeof(s)) // prints 16

    fmt.Printf("Size of byte: %d\n", unsafe.Sizeof(by))  // prints 1
    fmt.Printf("Size of rune: %d\n", unsafe.Sizeof(r)) // prints 4
    fmt.Printf("Size of int: %d\n", unsafe.Sizeof(i)) // prints 8
    fmt.Printf("Size of uint: %d\n", unsafe.Sizeof(ui)) // prints 8
    fmt.Printf("Size of uintptr: %d\n", unsafe.Sizeof(uip)) // prints 8
    fmt.Printf("Size of error: %d\n", unsafe.Sizeof(e)) // prints 16

    /**
    // Citation used: [1]
    // Note that 8-bit unsigned integers use all 8 bits to represent
    // a number. 2^8 integers from and including 0 through 2^8-1.
    // Because unsigned numbers use one of these bits for signage,
    // the values range from -2^7 through 2^7-1 for an 8-bit signed integer.
    bool has the possible values of false and true.
    uint8 has any non-negative integer from 0 through 255. Examples: 0, 126
    uint16 has any non-negative integer from 0 through 65535. Examples: 0, 126, 60000
    uint32 has any non-negative integer from 0 through 4294967295. Examples: 0, 126, 4731828
    uint64 has any non-negative integer from 0 through 18446744073709551615. Examples: 0, 126, 198219181
    int8 has any integer from -128 to 127. Examples: -91, 0, 092
    int16 has any integer from -32768 to 32767.
    int32 can hold any integer from -2147483648 to 2147483647.
    int64 can hold any integer from -9223372036854775808 to 9223372036854775807.
    // The section below uses citation [2] but also my knowledge from ECEG 240
    float32 can hold any floating-point number that can be either exactly expressed by or expressed with some loss of precision by an IEEE-754 number having 1 bit for the sign, 8 bits for the exponent, and 23 bits for the mantissa. This is also known as a single-precision floating-point number. Examples: 1.2, -0.14121213
    float64 can hold any floating-point number that can be either exactly expressed by or expressed with some loss of precision by an IEEE-754 number having 1 bit for the sign, 11 bits for the exponent, and 52 bits for the mantissa. This is also known as a double-precision floating-point number. Examples: 1.412, -0.141213
    complex64 holds a 32-bit floating-point value for the real part and a 32-bit floating-point value for the imaginary part. Examples: complex(6,-20)
    complex128 holds a 64-bit floating-point value for the real part and a 64-bit floating-point value for the imaginary part. Examples: complex(5,-17)
    string holds 128 bits of data, i.e., two 64-bit values. One is an 8-bit integer representing the length of the string, and the other is an unsigned integer pointer storing a pointer to the start of the string in memory. Examples: \"hello world!\"
    byte behaves like uint8, rune behaves like int32, int behaves like int64, uint behaves like uint64, uintptr behaves like uint64 on a 64-bit machine, and uint32 on a 32-bit machine and error behaves like a string.
    */

    var num_songs int
    fmt.Println("Every int appears to be initialized as a 64-bit integer. It does not grow automatically, as it fails if trying to set it to 2^63. The analogous is true for uint.")
    num_songs = 4294967295 // 2^32 - 1
    fmt.Println("setting num_songs to 2^32-1")
    fmt.Println("value of num_songs: ",num_songs) // prints 4294967295
    fmt.Println("int size: ",unsafe.Sizeof(num_songs)) // prints 8 (for 8 bytes)
    num_songs = 4294967296 // 2^32 (overflowing 32-bit int)
    fmt.Println("setting num_songs to 2^32")
    fmt.Println("value of num_songs: ",num_songs) // prints 4294967296
    fmt.Println("int size: ",unsafe.Sizeof(num_songs)) // prints 8
    num_songs = 9223372036854775807 // 2^63-1
    fmt.Println("setting num_songs to 2^63-1")
    fmt.Println("value of num_songs: ",num_songs) // prints 9223372036854775807
    fmt.Println("int size: ",unsafe.Sizeof(num_songs)) // prints 8
/**
    The code below does not work; it terminates with the following runtime error: "cannot use 9223372036854775808 (untyped int constant) as int value in assignment (overflows)"
    Go does not, therefore, allow for overflowing or dynamic resizing.
 
   num_songs = 9223372036854775808 // 2^63 (overflowing 64-bit int)
    fmt.Println("setting num_songs to 2^63")
    fmt.Println("value of num_songs: ",num_songs)
    fmt.Println("int size: ",unsafe.Sizeof(num_songs))
*/
    fmt.Println("\n")
}

/**
Topic 2: Background

The barebones development of the Go programming language began on September 21, 2007, and was publicly released on November 10, 2009. Robert Griesemer, Rob Pike, and Ken Thompson developed it. Go is a compiled language. It has the safety of typing, unlike Python, and the efficiency of avoiding repetitive pieces of code, i.e., without "declarations" or "header files" [2]. Go is used heavily within Google, specifically for "site reliability engineering" and "large-scale data processing" [2].
*/

/**
Topic 3: Default Params [3]

func default_params(num_songs int=5) int {
    return num_songs+num_songs
}

OR

func default_params(num_songs=5) int {
    return num_songs+num_songs
}

yields the following error:

"syntax error: unexpected =, expecting comma or )"

Go does NOT allow for default parameters.
*/

/**
Topic 4: Dynamic vs. Static Typing [4]

Go is statically typed. Once a particular type, always a particular type (that cannot be changed). I call the function below and pass in a float instead of the expected int. I compile the code and then run it to see where the type checking is done (where the type mismatch is caught -- we already know that Go cares because it's strongly typed). It gives "cannot use f (variable of type float64) as type int in argument to num_song_incrementer," so it infers that f is of type float64 during compilation and does the type-checking then.

func num_song_incrementer(num_songs int) int {
    return num_songs+1
}
*/

/**
Topic 5: Implicitly or explicitly typed

Variables can be implicitly declared or explicitly declared.

num_minutes_listened := 2 // implicit

OR 

var num_minutes_listened int // explicit
num_minutes_listened = 2

Note that num_minutes_listened := 2 will default to int (int64) rather than the minimally-sized int of int8 necessary to store the value 2.
The benefit of explicitly declaring types in Go specifies a particular variable size.
*/

/**
Topic 6: Static vs. Dynamic Scope

Go is statically, or lexically, scoped. The code below modifies the global copy of the variable num_minutes_listened rather than the main function call's copy of the variable num_minutes_listened when inc_num_mins_listened() is called.
*/

func inc_num_mins_listened() {
    num_minutes_listened = num_minutes_listened+1
}

/**
Topic 7: Overloading

No overloading of methods OR operators. [5] The decision was to keep things cleaner over the convenience of having the same name for multiple types. Each method call is made intentionally with the parameter inputs in mind
and cannot be confused with any other type. This is, from my view, useful in a language where things are implicitly typed, and programmers, consequently, could make unintended calls to a method with a parameter type different than
what they had anticipated the type to be inferred as. e.g. num_minutes_listened := 5. Many lines down, they call popsicle(num_minutes_listened int) instead of popsicle(num_minutes_listened float) because they had forgotten to put a 5.0 instead of a 5. Specifically, num_minutes_ads_listened := 5.0 implicitly types num_minutes_ads_listened to be a float64, whereas num_minutes_ads_listened := 5 implicitly types num_minutes_ads_listened to be an int.
*/


/**
Topic 8: Strict vs. Non-Strict Evaluation

Strictly evaluated by default. The parameters are evaluated before any of the statements inside of the method/function are executed. The tests below AFTER the comment, namely num_song_eval and all_songs_eval called by main, prove this with the output for those method calls being the following:

num_song_eval w/ input  1
num_song_eval w/ input  2
all_songs_eval

One COULD theoretically do non-strict evaluation, like in Python, by passing in a reference to the function and the parameters to pass into that function as distinct parameters, allowing the method
to delay, using control structures like if-else, evaluation of the parameters. That seems like a workaround and less of an example of non-strict evaluation. See the brief example below in this comment:

func wittie(fav_song string) {

}

func meng(fav_song string) {

}

func choose_profs_fav_song(choose_wittie bool, wittie, meng func(fav_song string)) {
    if choose_wittie {
        wittie(fav_song)
    } else {
        meng(fav_song)
    }
}
*/

func num_song_eval(num_minutes_listened int) int {
    fmt.Println("num_song_eval w/ input ", num_minutes_listened)
    return num_minutes_listened+1
}

func all_songs_eval(num_minutes_ads_listened int, level_of_music_interest int) int {
    fmt.Println("all_songs_eval")
    return num_minutes_ads_listened*level_of_music_interest
}

/**
Topic 9: Parameter Passing/Calling Mechanisms

Pass by value. However, you can still share a reference between two functions with a pointer.
*/

func set_one_song (num_minutes_listened int) {
    num_minutes_listened = 1
}

/**
Topic 10: The string type

Strings, in Go, are primitive and immutable. They are a struct with an 8-bit integer length field and an 8-bit unsigned integer memory address pointer pointing to the start of the string. += and + can be used for concatenation [6]. Strings can also be lexicographically compared with >, <, <=, and >=. They can be declared as variables or constants (note: immutable does not mean constant; variables of an immutable type have to recreate a new structure to hold the new value and then change the variable to point to that new structure). \0 is illegal, so \000 (3 octal digits) or \x00 (two hex digits) must be used [7]. There are several string methods in the built-in string library, like in Java. Some examples include the following: Clone (new copy of string), Compare (lexicographical comparison of strings w/ -1, 0, and 1 return values), Contains (search for substring), Count ("number of non-overlapping instances of [a substring]"), Cut (splice around a substring), HasPrefix and HasSuffix (checks if the substring is a prefix or suffix of a string), Index (first index of substring), Join (concatenates two strings with joining substring in the middle), LastIndex (index of the last instance of a substring in a string), Replace (replace substring a with substring b in a string), Split (return a list of substrings generated when a string is spliced on a particular inputted separator), and more. (Note: I paraphrased ALL of those after reading and understanding them.) 
*/

/**
Topic 11: Multidimensional Arrays

You can make two or more dimensional arrays. The function below demonstrates the generation of 3D lists in Go.
*/

type fav_song struct{}

func get_fav_songs() [][][]*fav_song {
	var fav_songs [][][]*fav_song
	fav_songs = make([][][]*fav_song, 2)
	for i := range fav_songs { // indices are more clear as letters than full music-themed names
		fav_songs[i] = make([][]*fav_song, 2)
		for j := range fav_songs[i] {
			fav_songs[i][j] = make([]*fav_song, 2)
		}
	}
	return fav_songs
}

/**
Topic 12: Dangling Else

Go does a beautiful job eliminating the dangling else problem. The else must follow on the same line after the closing } of
the if block; otherwise, the following error is thrown: "syntax error: unexpected else, expecting }." Having no else is OK, but having an else on its line or, for that matter, without a } closing a prior if clause results in the error mentioned above. 

num_minutes_listened := 3
if num_minutes_listened == 3 {
    fmt.Println("joseph")
} if num_minutes_listened == 4 {
    fmt.Println("wittie")
} else {
    fmt.Println("kartikeya")
} // joseph and then kartikeya prints because the else is unambiguously tied to the if num_minutes_listened == 4 block due to the else following that if block's closing } curly brace
*/

/**
Topic 13: Garbage Collection

Go has a garbage collector. Despite concurrent threads having access to the same memory, it safely frees memory, so it has to make sure that all threads are not or will not use a piece of memory before freeing it [9]. It can be invoked by first importing runtime and then calling runtime.GC() [10]. It currently uses a "mark-and-sweep" collection algorithm [9]. It maintains bits to mark a reference as 0 once unreachable. Then a collector individually frees memory (including memory that depends on the current freed memory to, in turn, also be accessed).
*/

/**
Topic 14: Short Circuit Evaluation

It does use short circuit. If Go is short circuit, it will be able to evaluate to true for the if statement in both_songs and that in at_least_one_song after evaluating the first operand in the if-statement. It did not short circuit if it assessed the entire conditional statement before branching within the code. The former, i.e., Go is short circuit, holds per the test performed and shown below.
*/

func ret_true_song_listened() bool {
    fmt.Println("true")
    return true
}

func ret_false_song_listened() bool {
    fmt.Println("false")
    return false
}

func both_songs() {
    if ret_false_song_listened() && ret_false_song_listened() {
    } else {
    }
}

func at_least_one_song() {
    if ret_true_song_listened() || ret_true_song_listened() {
    } else {
    }
}

/**
Topic 15: Exceptions

There are no exceptions because Go says that try-catch-finally "results in convoluted code" [11]. (This is an aside from me writing this at 2:58 AM... I turned it in on time... yup...: I feel like Go's programmers read a bunch of Dr. Seuss as children. "A person's a person, no matter how small!" It's kind of like how they think that some problems are ignorable, so we'll treat them special. Go says... treat every error like it's essential... stop blowing it off! I'm not sure why I thought of this odd analogy, but it made sense. Maybe you found it semi-humorous while grading the projects as well.) While you can't use control structures to raise and catch custom exceptions, you can log errors by calling log.Fatal while passing in the error itself [12]. e.g.:

file, err := os.Open("music.txt")
if err != nil {
    log.Fatal(err)
}
*/

/**
Topic 16: Shallow v. Deep Copy

Default is shallow copy. See copy_music_listener function below. To force a deep copy, you can create a new struct and manually assign each field to each string value OR use a deep copy library like https://github.com/jinzhu/copier. Note that strings, but rather structs within a struct, will need to be deep copied because strings are immutable and primitive. They will point to a new value rather than an old value when the value is changed (aliasing side effects will not be present with strings).
*/

type MusicListener struct {
    FavSong string
    FavSinger string
    FavBand string
}

func copy_music_listener() {
    ml1 := new(MusicListener)
    ml1.FavSong = "Merry Dangling Else"
    ml1.FavSinger = "Dan Hustavan"
    ml1.FavBand = "Maroon 5"

    ml2 := ml1
    fmt.Printf("ml1 is pointing to %p\n", ml1)
    fmt.Printf("ml2 is pointing to %p", ml2) // both return same thing (are pointing to the same object)
}

/**
Topic 17: Math operators on numbers & bits

In addition to + (adding), - (subtracting), * (multiplying), / (dividing), and % (remainder), there are also bitwise logical operators, such as & (bitwise AND), | (bitwise OR), ^ (bitwise XOR), and &^ (setting a bit to 0), << and >> (left and right bit-shifting arithmetic operators to multiply and divide by 2 respectively) [13].
*/

/**
Topic 18: Regular expression support

Regex is compiled and then used with methods in the built-in regexp library, analogous to the regexp compilation and method usage of the re library in Python. Below is an example:

import "regexp"
var rexp, _ = regexp.Compile("N*") // trying to find all the bands starting with N, such as Nirvana
rexp.MatchString("Nirvana") // yields true

It does do matching the whole string and search (finding matches inside). It can run a greedy search. For example, it has Longest(), which finds the "leftmost-longest match" [14]. It does not appear to find non-overlapping matches without the regex having a divider within it, so to speak, such as "SomeRegex(divider)?OtherRegex(divider)?MoreRegex." with FindSubmatch() [14]. However, it does capture with FindSubmatch(), as Python does with grouping. Moreover, it can do replace with Replace().
*/

/**
Topic 19: Pointers/references

Objects can be accessed through pointers with dereferencing, as shown by the following code. Go has no pointer arithmetic [15]. 

*/

func num_songs_changer() {
    var ptr_num_songs *int
    num_songs := 2
    ptr_num_songs = &num_songs
    *ptr_num_songs = 71
    fmt.Println("value of num_songs changed via dereferencing a pointer pointing to its address should now be 71:", num_songs)
}

var num_minutes_listened int8

func main() {
    fmt.Println()
    primitives()
    // f := 4.3
    // num_song_incrementer(f) // returns error if not commented
    var num_minutes_listened int8
    num_minutes_listened = 5
    inc_num_mins_listened()
    fmt.Println("Topic 6: Static vs. Dynamic Scope")
    fmt.Println(num_minutes_listened) // if 5, then statically scoped (global num_minutes_listened was incremented from -1 to 0 with main's num_minutes_listened untouched)
                   // else if 6, inc_num_mins_listened incremented main's num_minutes_listened because main was the scope immediately prior to the function call to inc_num_mins_listened.
                   // 5 prints --> statically scoped
    fmt.Println("\n")
    fmt.Println("Topic 8: Strict vs. Non-Strict Evaluation")
    all_songs_eval(num_song_eval(1), num_song_eval(2))
    fmt.Println("\n")
    num_minutes_ads_listened := 5
    set_one_song(num_minutes_ads_listened)
    fmt.Println("Topic 9: Parameter Passing/Calling Mechanism")
    fmt.Println("num_minutes_ads_listened prints", num_minutes_ads_listened, "(5 meaning num_minutes_ads_listened did not change but rather the value of 5 was copied and pass into the set_one_song call, and 1 meaning that set_one_song had access to and changed main's num_minutes_ads_listened)")
    fmt.Println("\n")
    fmt.Println("Topic 11: Multidimensional Arrays")
    l := get_fav_songs()
    fmt.Println("A 3D 2x2x2 array:", l)
    fmt.Println("\n")
    fmt.Println("Topic 14: Short Circuit Evaluation")
    both_songs()
    at_least_one_song()
    fmt.Println("\n")
    fmt.Println("Topic 16: Copy Test")
    copy_music_listener()
    fmt.Println("\n")
    fmt.Println("Topic 19: Pointer Test")
    num_songs_changer()
    fmt.Println("\n")
}


/**
Works Cited, referenced in the text above with [citation #]

1. https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go
2. https://go.dev/doc/faq
3. https://www.golang-book.com/books/intro/7
4. http://www.club.cc.cmu.edu/~cmccabe/blog_golang_type_system.html
5. https://go.dev/doc/faq#overloading
6. https://go101.org/article/string.html
7. https://go.dev/ref/spec#String_literals
8. https://pkg.go.dev/strings
9. https://go.dev/doc/faq#garbage_collection
10. https://golang.org/pkg/runtime/#GC
11. https://go.dev/doc/faq#exceptions
12. https://go.dev/blog/error-handling-and-go
13. https://go.dev/ref/spec#Arithmetic_operators
14. https://pkg.go.dev/regexp
15. https://go.dev/tour/moretypes/1
*/
