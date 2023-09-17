package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
)

var the_true_klinoff = "klinoff "
var file_name string = "klinoff.hns"
var current_directory, _ = os.Getwd()
var sound_file string
var level_of_pain int
var difficulty int = 0
var isTrueKlinoff bool = false
var asciiArt string

func main() {
	asciiArt = `
	Beta Release 0.0.2 / Minor bugs may occur / Please report any bugs to the developer at Klinoff@klinoff.dev 
	If you experience any bugs, restart the program and try again.
	if you are a true klinoff you will find the true klinoff in the file
		_    _           _____               _  __ _  _                 __   __ 
		| |  | |  ___    / ____|             | |/ /| |(_)               / _| / _|
		| |__| | ( _ )  | (___   __      __  | ' / | | _  _ __    ___  | |_ | |_ 
		|  __  | / _ \/\ \___ \  \ \ /\ / /  |  <  | || || '_ \  / _ \ |  _||  _|
		| |  | || (_>  < ____) |  \ V  V /_  | . \ | || || | | || (_) || |  | |  
		|_|  |_| \___/\/|_____/    \_/\_/(_) |_|\_\|_||_||_| |_| \___/ |_|  |_|  
	`
	fmt.Println(asciiArt)
	playSound("sound/klinoff_introduction.mp3")
	levelOfPainMaker()
	// set difficulty between 1 and 3
	clearConsole()
	fmt.Println("Choose the difficulty:")
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")

	for {
		fmt.Println("Enter the difficulty (1-3): ")
		_, err := fmt.Scanf("%d", &difficulty)
		if err == nil && difficulty >= 1 && difficulty <= 3 {
			break
		}
		fmt.Print("Not a valid number. Please try again. ")
	}

	if difficulty == 1 {
		sound_file = "sound/easy_mode.ogg"
	} else if difficulty == 2 {
		sound_file = "sound/medium_mode.mp3"
	} else if difficulty == 3 {
		sound_file = "sound/hard_mode.mp3"
	}

	startFileMaker()

	colors := []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m"} // ANSI color escape codes

	if !isTrueKlinoff {
		// ASCII art
		asciiArt = `
			@@@@@@@@@@@@@@@@@%*@@@@@@%..&@@@@*.,&@@(..@@@&@@@@@@@*....,@@@@#.     .%#/(#.*...*%@@@@@@@@@@@@@@@@@
			@@@@@@@@@@@@@@@@@@@@@@@( .(@@@@@@@@@@@@@@@&&&@@@@@@@@@@@@@@@@@, /#####(  &@@@@@@@@@@@@@@@@@@@@@@@@@@
			@@@@@@@@@@@@@@@@@@@@@@....,,.. &@@@@@@%, .,*, .#@@@@@@@@@@@@@%.,%/..../* /@@@@@@@@@@@@@@@@@@@@@@@@@@
			@@@@@@@@@@@@@@@@@@@@%//*//(/,/*@@@@@@. /###/*(%. #@%/,        .(##*.,*/ ,@@@@@@@@@@@@@@@@@@@@@@@@@@@
			@@@@@@@@@@@@&..*///////((###/(&@@@@@% .##(....(#(//###################%%/,   /&@@@@@@@@@@@@@@@@@@@@@
			@@@@@@@@@@@(     ,///(#%%#(/%%%@@@@@@&.  ,#############################%%%%%%#*  /@@@@@@@@@@@@@@@@@@
			@@@@@@@#..  ..*(((((%&&#(*%%#(/@@@@@@@%, .##################################%%%%%(. ,%@@@@@@@@@@@@@@
			@@@@*.  ,....*(,..(&&(//(%%(/..(/%@@#  /########################################%%%%( .%@@@@@@@@@@@@
			@@@,.,. ..,*(&/../.,(*(%%##/,/#.*/((.###############################################%%#. %@@@@@@@@@@
			@@@,,.,/      .&/,#.,&%#%#/*#,.  ,,*//*(##############################################%%#  (@@@@@@@@
			@@@,,*&        .@./.,##(.,/%   .  .,*/((//((##########################################(*,*(. %@@@@@@
			@@&,*,(        *(,(./#/#./.*   .  .**/(#%%%##################((((((#########################, (@@@@@
			@@@#/*,,/&#((*(,,*.,(##.(.,       *(##%%%%%%###########(/(#%%%%%%%%%%#/(#####################* (@@@@
			@@@@@/************###/((#####*   ./#%#%%###/########((%%%%%%%%%%%%%%%&&&%#/(########//(///(###* %@@@
			@@@@@(//.,,/#/**###(,,(##%%%%((  *(##%%#(&@@@%/###/(%%%%%%%%%%%%%%%%%%%%&&%%((##(*#@@@@@@@@&/##,,@@@
			@@@@@/.*.... .%..#/.,#/(/((##((. *##%%%(**@@@@@/(/#%%%%#,..../%%%%%#/,*#%%&&%(((*%@@@@@%,./@@/#* (@@
			@@@@@/,*.    ,/***.,/*.#*****/((((/#%%%*.,&@@@@//(%%%%%*.....,#%%%#,....%%%%%#(/*%@@@@@&/*#@@/#(.,@@
			@@@@@. #,////*.#,,*,,/&&&&&%%%##(#(,*/#@@@@@@@((/(%%%%%%(*,,*#%%%%%(,.,(%%%%%%(//*(@@@@@@@@%*##(.,@@
			@@@@@@..     ..  .*######%%%%%%%%#((/**/(((/*(###/(##%%%%%%%%%%%%%%%%%%%%%%%%#(#(((/**//**/####* (@@
			@@@@@@. ..     ..*/((((((((######(((###############(/(####%%%%%%%%%%%%%%%###(/################/.*@@@
			@@@@@@             ..,,***///(#%##(###################(//((/(((((((((((///((#################/.,&@@@
			@@@@@@/     ....,.*/(#####%%%%%%%#/(###################(///////////////####################(* *&@@@@
			@@@@@@@@.      . ,**///(((((#####/ ./(#####################/*******//###################(/. .&@@@@@@
			&@@@@@@@@&. .,.   .,****///*/(##%&&(  ,//(###########################################(/,  #&@@@@@@@@
			&&&&&&&&&&&#   ,*((((#((((##(((&&&&&&&(.  ,//(((###############################(((/*. .(&&&&&&&&&&&&
			&&&&&&&&&&&&&. .*////*//(((((%&&&&&&&&&&&%/. .,*///(((((((((((((((((((((((((//*,. .*#&&&&&&&&&&&&&&&
			&&&&&&&&&&&&&&&&/*..,,**/&&&&&&&&&&&&&&&&&&&&&#/.    .,**///////////**,..    ./%&&&&&&&&&&&&&&&&&&&&
			&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&%#((/**/(((#&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&                                 
		`
		//wait for second
		time.Sleep(100 * time.Millisecond)

		go func() {
			playSound("sound/klinoff_outro.mp3")
		}()
	} else {
		// ASCII art
		asciiArt = `
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@& *##(* @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#,((///(#.@@@@@@@@@@@@%*,(@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ /(///(( @@@@@@@@@@ ((#(((/,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@&. ,((((((((((((((*.  (/////##.@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@% .((###################(#(((#((. @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ .#(###(##################(#######, &@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@% /(######(##(#(((###########((((#####((..@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ ,/,*(##############(((##(###########,,,((#. @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@( ((((##(#((##############################((((( %@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@, ((###(#((##((#(//(######(//((####(####((((####( *@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@% (/#@@@@@&,(#(*################/((##(*@@@@@@@*((## *@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ ./@@@@@@@@@&/####################*#(#@@@@@@@@@&/##/ @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ ,&@.  @@@@@((###*...,###(....####(//&@@@@@( ,@@*### &@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ .,@@@@@@@@@/####/...,####....####(/#*&@@@@@@@@*/##( %@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@( //**%&&/*/(/####################/(#(//**/**//(#((, @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@* /((((((#((((/################/(((#####(((#((###* @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@& */(#########((///////////(((##############(((..@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@*. ,(((((####(#((((((((((((############(##(/..@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@....,,,. ,//(((##########(#(((##((##((##((//. ...(@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@,    .  ,,,,,,,.  ,///((((((((((((#(((((((*.  ....,..     *&@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@@@@@@/.          ..,,,,,,,                           ...,,,.           *@@@@@@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@@@# ...........     ....,,,.                             .,........ . .....   (@@@@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@@#............... . .  ...,..                            ...,, ................   @@@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@@@   ................... ....,,.                          .  ....................... *@@@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@@@ &#/%*. ............. .. .,....,.                       ...... ,.........., .... ....&&/@@@@@@@@@@@@@@@
		@@@@@@@@@@@@@@@/.@.@%..*&%............... ,,.....,.   . ..               ....... /###*@&,.. ........@*&/,*%@@@@@@@@@@@@@
		@@@@@@@@@@@@@& &&#,(@(@@..,.............. ....,.....  .....             ......,*(/*//(@&@%........,(#,,&&&&,@@@@@@@@@@@@
		@@@@@@@@@@@@@   ..,..(@@..*/............... ......... ....... .        ,,,.., .,,,,***//((......#*@*&@*.....,@@@@@@@@@@@
		@@@@@@@@@@@@,    .......,&@# .......,... ........,...  .......        ........,,.,,,...... . .../............,@@@@@@@@@@
		@@@@@@@@@@@@       ......... .. ............@@/&,...,..**************.....,............... ..., ........(@*.. @@@@@@@@@@
		@@@@@@@@@@@*..&&&    /@(... .%%%&&...&@,...%@%(&@* ...*%/#&##(##%%%/*...,#@*.,@@&(@@@(@&./,,... ...... &%@%(@#&@@@@@@@@@
		@@@@@@@@@@@       .. .(*... .. .@@,(@&.....@@..,@@* .. ,*//*/******,...,&@@%.#@%..&@#.%@*(. ..    ... */ /***,#@@@@@@@@@
		@@@@@@@@@@*        . .(#*.., . .*@@@,.%@# *@&(%@@&. .,, ....        .. ..(@@*@@.,,@@,..@@ . ./    ........,... @@@@@@@@@
		@@@@@@@@@@ (/&@% ..  ......  .../@%...*,. ......,,,..., .           ......,,,,,,,,,,,*@@,. ./,   ...............@@@@@@@@
		@@@@@@@@@  *(/. %@@.... ...  ..%@/....... .,,,,,,.(&# ,. ..        ......,,,,,,,,,,,@@(,,, ..      .......,.....%@@@@@@@
		@@@@@@@@(      (#@(...... .   ,*..../@% .. %@%,,,,,.. ., ..        .....,,,,,,,,,,,,,,,,,,..   ./@@@@%,@@&/,*@.&.@@@@@@@
		@@@@@@@@     */  @%.*,  ....   .....%@@@,  @@*,,,,,..  . .        ./*.../&&(,,,*/,,,,,,.....,.(@,,.%@/(@#...,@%% @@@@@@@
		@@@@@@@& . .((,,@@@,..(%*,*   *.....&@#,/..@@,.%@@&@@. ...         (@(/@@,...&@(.,.%@@&@@#&.,#@&..,*,.......,@(/.@@@@@@@
		@@@@@@@@ ./  .       /@@...    . ...&&, ,/(@&.@@@&&&@& ..        , .....#@@#.,#@%(&@,..%*..*  .,........ .*&..*,*@@@@@@@
		@@@@@@@@,,.        .*(#. #@(   .*   //,  */(& %&(..%*.* .        .//.%(&@@*(#&@@*,@@&%@%.    , ..,***(..........@@@@@@@@
		@@@@@@@@%  ...,*.&@/ .&@* ,/%@#            .. . ,**.    .       .     ..................,.**./@@,&@@@/.%@,..... @@@@@@@@
		@@@@@@@@@ .,@@@@ #@%..*&@,%@@@#.     . ...              .       .......,....................(&&,(@@,..,#%&@&%@(*@@@@@@@@
		@@@@@@@@@%./% .@%.&@/.. .........      ... ...........  .        .,,,,,,,,,,.%@@@@%..    .........,/(*..@&,.,@*&@@@@@@@@
		@@@@@@@@@@,%@/@&(.............*@&      %/ ..,... .....  .      .(,,,.//*.,,,,@@.,@%..      ..............%@@@/,@@@@@@@@@
		@@@@@@@@@@@. /................@@@#..  &. *@&.,. #@@@@&,..       @(.#@@,,,,,.%@@@@&.&*.. .......#@@@@/.......*/@@@@@@@@@@
		@@@@@@@@@@@&.    .....        /%##@&/**(/.*&&@.@@/ . (...       @@@@,......,@@..%@*........(%#&/. .@@*.......#@@@@@@@@@@
		@@@@@@@@@@@@#    ..   ..  .. #%,        .%&&@#.&@%..&# .      ..@@,..@@@.,.&@@@&/(%...........#&&#  (......./@@@@@@@@@@@
		@@@@@@@@@@@@@# . #&/.&&...  .  ....     .*  ...,,//,(* .       &*..,,,,,,,...../@/..................*   ,..&@@@@@@@@@@@@
		`
		//wait for second
		time.Sleep(100 * time.Millisecond)

		go func() {
			playSound("sound/klinoff_happy_ending.mp3")
		}()
	}

	for i := 0; i < 1000; i++ {
		randomColor := colors[rand.Intn(len(colors))] // Select a random color from the list
		resetColor := "\033[0m"                       // Reset color to default

		if isTrueKlinoff {
			fmt.Printf("%sYou are The true Klinoff%s\n", randomColor, resetColor)
		} else {
			fmt.Printf("%sYou are not a true klinoff!%s\n", randomColor, resetColor)
		}

		// Print ASCII art below
		fmt.Fprintln(os.Stderr, asciiArt)
		if isTrueKlinoff {
			time.Sleep(385 * time.Millisecond)
		} else {
			time.Sleep(430 * time.Millisecond)
		}
		clearConsole()
	}
}

func levelOfPainMaker() {
	for {
		fmt.Print("Enter the level of pain: ")
		_, err := fmt.Scanf("%d", &level_of_pain)
		if err == nil {
			break
		}
		fmt.Println("Not a valid number. Please try again.")
	}

	fmt.Println("Level of pain:", level_of_pain)
}

func startFileMaker() {
	stringSlice := make([]string, level_of_pain)
	for i := range stringSlice {
		stringSlice[i] = "klanoff "
	}

	fakeWords := []string{the_true_klinoff, "kalnoff ", "kannöff "}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(fakeWords), func(i, j int) {
		fakeWords[i], fakeWords[j] = fakeWords[j], fakeWords[i]
	})

	for _, fakeWord := range fakeWords {
		indexToInsert := rand.Intn(len(stringSlice) + 1)
		stringSlice = append(stringSlice[:indexToInsert], append([]string{fakeWord}, stringSlice[indexToInsert:]...)...)
	}

	resultString := strings.Join(stringSlice, "")

	file, _ := os.Create(file_name)
	file.WriteString(resultString)
	file.Close()

	countdown()
}

func playSound(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening sound file:", err)
		return
	}
	defer f.Close()

	var streamer beep.StreamSeekCloser
	var format beep.Format

	if strings.HasSuffix(file, ".mp3") {
		streamer, format, err = mp3.Decode(f)
		if err != nil {
			fmt.Println("Error decoding MP3 sound file:", err)
			return
		}
	} else if strings.HasSuffix(file, ".ogg") {
		streamer, format, err = vorbis.Decode(f)
		if err != nil {
			fmt.Println("Error decoding OGG sound file:", err)
			return
		}
	} else if strings.HasSuffix(file, ".wav") {
		fmt.Println("Unsupported sound file format")
		return
	} else {
		fmt.Println("Unsupported sound file format")
		return
	}

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		fmt.Println("Error initializing speaker:", err)
		return
	}

	done := make(chan struct{})
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(done)
	})))
	<-done
}

func countdown() {
	clearConsole()
	fmt.Println("Find the true klinoff!")

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("notepad.exe", file_name)
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("xdg-open", file_name)
	}

	if cmd != nil {
		cmd.Start()
		playSound(current_directory + string(os.PathSeparator) + sound_file)
		// shut down the editor
		forceShutEditor()
		scanFile(file_name, the_true_klinoff)

		// After waiting, the text editor should have released the file
		os.Remove(current_directory + string(os.PathSeparator) + file_name)

	}
}

func clearConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(asciiArt)
}

func scanFile(file_name string, the_true_klinoff string) {
	file, _ := os.Open(file_name)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), the_true_klinoff) {
			fmt.Println("You are not a true klinoff!")
			isTrueKlinoff = false
			return
		}
	}

	// if the scanner never found the true klinoff
	if !scanner.Scan() {
		fmt.Println("You are The true Klinoff")
		isTrueKlinoff = true
	}
}

func forceShutEditor() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("taskkill", "/im", "notepad.exe", "/f")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
