/**
 * Final Project in Go!

 * songs.go can be run in the command line, assuming that go is installed, by running `go run songs.go`

 * The program prompts the user for how they are feeling (from a selection of the following emotions: happy,
 * despair, peaceful, and angry). It then selects music that emulates how the user is feeling.
 
 * However, to do this, it must import and identify which songs are associated with each emotion.
 * The preprocessing of the data takes considerable time and would benefit from parallelism, particularly when
 * preprocessing steps are independent. The Go programming feature of its built-in sync library's WaitGroup is used.
 * Processes are grouped, executed in parallel, with each one indicating when it is done right before returning with the keyword defer.
 
 * Importing the emotion-suggesting words and song titles are done in parallel. Then, the population of each song list, i.e., happy songs, songs of despair,
 * peaceful songs, and angry songs, are also done concurrently. Now, the user can be prompted, and the program is ready to provide song recommendations that
 * emulate how the user is feeling, that is, concerning the four emotion options currently programmed in this software: happy, despair, peaceful, and angry.

 * Name: Kartikeya Sharma
 * Professor Wittie
 * CSCI 308
 * May 2022
 */

package main

// Import statements
import (
  "fmt"
  "math/rand"
  "strings"
  "sync"
  "time"
)

// List of song titles
var songTitles [20]string

/**
 * Loads all the song titles into the songTitles array
 */
func loadSongs(wg *sync.WaitGroup) {
    defer wg.Done()
    
    songTitles[0] = "Easy On Me"
    songTitles[1] = "Neon Lights"
    songTitles[2] = "The Adults Are Talking"
    songTitles[3] = "Sweet Talk"
    songTitles[4] = "Ruin My Life"
    songTitles[5] = "Shotgun"
    songTitles[6] = "As It Was"
    songTitles[7] = "First Class"
    songTitles[8] = "Neon Moon"
    songTitles[9] = "Drunk Driving"
    songTitles[10] = "Paralyzed"
    songTitles[11] = "Lullaby"
    songTitles[12] = "New Light"
    songTitles[13] = "Stronger"
    songTitles[14] = "Already Gone"
    songTitles[15] = "Night Changes"
    songTitles[16] = "Rich & Sad"
    songTitles[17] = "Happy"
    songTitles[18] = "Bad Girlfriend"
    songTitles[19] = "Burning"
}

/**
 * Struct to hold lists of words that describe happy songs, songs of despair,
 * peaceful songs and songs with anger, respectively.
 */
type WordBank struct { // stores emotion-word associations
    happy []string
    despair []string
    peaceful []string
    angry []string
}

// Global pointer to the WordBank struct
var wb WordBank

/**
 * Initializes the aforementioned WordBank struct
 */
func fillWordBank(wg *sync.WaitGroup) {
    defer wg.Done()
    
    wb.happy = []string{"easy", "neon", "lights", "sweet", "first class", "new", "stronger", "happy"}
    wb.despair = []string{"ruin", "shotgun", "drunk", "paralyzed", "gone", "sad"}
    wb.peaceful = []string{"night", "lullaby"}
    wb.angry = []string{"bad", "burning"}
}

/**
 * Struct to hold songs that describe happy songs, songs of despair,
 * peaceful songs and songs with anger, respectively.
 */
type SongBank struct {
    happy []string
    despair []string
    peaceful []string
    angry []string
}

// Global pointer to the SongBank struct
var sb SongBank

/**
 * Helper function that searches each song title for the happy words, and
 * if a happy word is found in a particular song title,
 * add it to the list of happy songs in the SongBank struct
 */
func fillSongBankWithHappySongs(wg *sync.WaitGroup) {
    defer wg.Done()
    
    sb.happy = []string{}
    for songNum := 0; songNum < len(songTitles); songNum++ {
        for wordNum := 0; wordNum < len(wb.happy); wordNum++ {
            if strings.Contains(strings.ToLower(songTitles[songNum]),
                                strings.ToLower(wb.happy[wordNum])) {
                sb.happy = append(sb.happy, songTitles[songNum])
                break
            }
        }
    }
}

/**
 * Helper function that searches each song title for the despair words, and
 * if a word of despair is found in a particular song title,
 * add it to the list of despair songs in the SongBank struct
 */
func fillSongBankWithDespairSongs(wg *sync.WaitGroup) {
    defer wg.Done()
    
    sb.despair = []string{}
    for songNum := 0; songNum < len(songTitles); songNum++ {
        for wordNum := 0; wordNum < len(wb.despair); wordNum++ {
            if strings.Contains(strings.ToLower(songTitles[songNum]),
                                strings.ToLower(wb.despair[wordNum])) {
                sb.despair = append(sb.despair, songTitles[songNum])
                break
            }
        }
    }
}

/**
 * Helper function that searches each song title for the peaceful words, and
 * if a peaceful word is found in a particular song title,
 * add it to the list of peaceful songs in the SongBank struct
 */
func fillSongBankWithPeacefulSongs(wg *sync.WaitGroup) {
    defer wg.Done()
    
    sb.peaceful = []string{}
    for songNum := 0; songNum < len(songTitles); songNum++ {
        for wordNum := 0; wordNum < len(wb.peaceful); wordNum++ {
            if strings.Contains(strings.ToLower(songTitles[songNum]),
                                strings.ToLower(wb.peaceful[wordNum])) {
                sb.peaceful = append(sb.peaceful, songTitles[songNum])
                break
            }
        }
   }
}

/**
 * Helper function that searches each song title for the angry words, and
 * if an angry word is found in a particular song title,
 * add it to the list of angry songs in the SongBank struct
 */
func fillSongBankWithAngrySongs(wg *sync.WaitGroup) {
    defer wg.Done()
    
    sb.angry = []string{}
    for songNum := 0; songNum < len(songTitles); songNum++ {
        for wordNum := 0; wordNum < len(wb.angry); wordNum++ {
            if strings.Contains(strings.ToLower(songTitles[songNum]),
                                strings.ToLower(wb.angry[wordNum])) {
                sb.angry = append(sb.angry, songTitles[songNum])
                break
            }
        }
    }
}

/**
 * Calls helper functions to fill song titles for the happy, despair, peaceful, and angry songs based on
 * if the respective types of words are found in those song titles. The critical aspect to note is the concurrency:
 * since the determination and setting of each type of song are independent of one another, their function calls
 * are put into a Go WaitGroup to be then executed in parallel.
 */
func fillSongBank() {
    var wgSongBank sync.WaitGroup
    wgSongBank.Add(4)
    go fillSongBankWithHappySongs(&wgSongBank)
    go fillSongBankWithDespairSongs(&wgSongBank)
    go fillSongBankWithPeacefulSongs(&wgSongBank)
    go fillSongBankWithAngrySongs(&wgSongBank)
    
    wgSongBank.Wait()
    
   fmt.Println("All my happy songs are: ")
   for i := 0; i < len(sb.happy); i++ {
        fmt.Println(i, ". ", sb.happy[i])
    }
 
   fmt.Println("All my despair songs are: ")
   for i := 0; i < len(sb.despair); i++ {
        fmt.Println(i, ". ", sb.despair[i])
    }   

   fmt.Println("All my peaceful songs are: ")
   for i := 0; i < len(sb.peaceful); i++ {
        fmt.Println(i, ". ", sb.peaceful[i])
    }   

   fmt.Println("All my angry songs are ")
   for i := 0; i < len(sb.angry); i++ {
        fmt.Println(i, ". ", sb.angry[i])
    }   
}

/**
 * A happy song is suggested from the SongBank struct at random.
 */
func suggest_happy() {
    fmt.Println("Suggesting a happy song...")
    s := rand.NewSource(time.Now().UnixNano())   
    r := rand.New(s)
    selected_index := r.Intn(len(sb.happy))
    fmt.Println("Now suggesting: ", sb.happy[selected_index])
}

/**
 * A song of despair is suggested from the SongBank struct at random.
 */
func suggest_despair() {
    fmt.Println("Suggesting a song of despair...")
    s := rand.NewSource(time.Now().UnixNano())   
    r := rand.New(s)
    selected_index := r.Intn(len(sb.despair))
    fmt.Println("Now suggesting: ", sb.despair[selected_index])
}

/**
 * A peaceful song is suggested from the SongBank struct at random.
 */
func suggest_peaceful() {
    fmt.Println("Suggesting a peaceful song...")
    s := rand.NewSource(time.Now().UnixNano())   
    r := rand.New(s)
    selected_index := r.Intn(len(sb.peaceful))
    fmt.Println("Now suggesting: ", sb.peaceful[selected_index])
}

/**
 * An angry song is suggested from the SongBank struct at random.
 */
func suggest_angry() {
    fmt.Println("Suggesting an angry song...")
    s := rand.NewSource(time.Now().UnixNano())   
    r := rand.New(s)
    selected_index := r.Intn(len(sb.angry))
    fmt.Println("Now suggesting: ", sb.angry[selected_index])
}

/**
 * Loading the song titles and the list of words related to each emotion are two independent processes
 * that are placed into the same Go WaitGroup and executed concurrently. Once done, fillSongBank is called to
 * use these two lists and find songs that belong to each emotion category. Now, the program can suggest
 * songs that the end-user can relate to based on how they feel; to do this, it first prompts the user about their feelings.
 * It then suggests a song of the inputted emotion accordingly.
 */
func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go loadSongs(&wg)
    go fillWordBank(&wg)
    
    fmt.Println("DATA LOADING")
    wg.Wait()
    fmt.Println("DATA LOADED")
    
    fillSongBank()
    
    const prompt = "What mood are you in? Choose one of the following: happy, despair, peaceful, angry. "
    var input string
    try_again := true
    for try_again == true {
        fmt.Print(prompt)
        fmt.Scanln(&input)
        switch (input) {
            case "happy":
                suggest_happy()
                try_again = false
            case "despair":
                suggest_despair()
                try_again = false      
            case "peaceful":
                suggest_peaceful()
                try_again = false
            case "angry":
                suggest_angry()
                try_again = false
            default:
                try_again = true
        }
    }
}
