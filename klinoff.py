# pyinstaller --onefile --icon=klinoff.ico --add-data "klinoff.ogg;." --hidden-import ipaddress klinoff.py
import os
import random
import pygame
import time

# Initialize pygame
pygame.init()

# Get the current directory where the script is located
current_directory = os.path.dirname(os.path.abspath(__file__))

# Specify the path to the sound file and the file name
sound_file = os.path.join(current_directory, 'klinoff.mp3')

# Load the sound file
start_sound = pygame.mixer.Sound(sound_file)
level_of_pain = 0
file_name = "klinoff.hns"

while True:
    try:
        level_of_pain = int(input("Enter the level of pain (a number): "))
        break  # Exit the loop if a valid number is entered
    except ValueError:
        print("Not a valid number. Please try again.")

# Now you have the valid level_of_pain as an integer
print("Level of pain:", level_of_pain)

the_true_klinoff = "klinoff "

# Create a list with "level of pain"-amount of "klinoff" strings
strings = ["klanoff "] * level_of_pain

# Define the list of fake words
fake_words = [the_true_klinoff, "kalnoff ", "kannöff "]

def Start():
    # clear console
    os.system('cls' if os.name == 'nt' else 'clear')
    
    # Play the start sound
    start_sound.play()

    # Shuffle the list of fake words randomly
    random.shuffle(fake_words)

    # Insert the shuffled fake words at random positions
    for fake_word in fake_words:
        index_to_insert = random.randint(0, len(strings))
        strings.insert(index_to_insert, fake_word)

    # Join the list of strings into a single string
    result_string = "".join(strings)

    # Write the result string to a file
    with open(file_name, "w") as file:
        file.write(result_string)

    Countdown()

def Countdown():
    # clear console
    os.system('cls' if os.name == 'nt' else 'clear')
    print("Find the true klinoff!")
    # open file in notepad or gedit
    if os.name == 'nt':
        os.system("start notepad.exe " + file_name)
    else:
        os.system("gedit " + file_name)
        
    # when sound is stopped, close the program
    while pygame.mixer.get_busy():
        time.sleep(0.1)

    # close file in notepad if on Windows, or close the specific gedit instance on Linux
    if os.name == 'nt':
        os.system("taskkill /im notepad.exe /f")
    else:
        # Get the process ID of the gedit instance associated with file_name
        gedit_pid = os.popen(f"pgrep -f 'gedit {file_name}'").read()
        if gedit_pid:
            # Kill the gedit instance using its PID
            os.system(f"kill {gedit_pid}")

    # read file file_name and read it line by line, and find if the_true_klinoff is in
    with open(file_name, "r") as file:
        for line in file:
            if the_true_klinoff in line:
                print("You are not a true klinoff!")
                break
            else:
                print("You are a true klinoff!")
    
    os.remove(file_name)
    input("Press enter to close the klinoff")
    
if __name__ == "__main__":
    Start()
