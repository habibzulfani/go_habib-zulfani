#!/bin/sh
# using arguments
main_folder="$1 at $(date +'%a %b %d %H:%M:%S %Z %Y')"
facebook_url="https://www.facebook.com/$2"
linkedin_url="https://www.linkedin.com/in/$3"

# creating main_folder 
mkdir "$main_folder"
cd "$main_folder"

# creating sub_folder
mkdir -p about_me/personal
mkdir -p about_me/professional
mkdir -p my_friends
mkdir -p my_system_info

# inputing data into files
echo "$facebook_url" > about_me/personal/facebook.txt
echo "$linkedin_url" > about_me/professional/linkedin.txt

curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt  -o my_friends/list_of_my_friends.txt

username=$(uname -n)
os=$(uname -a)
echo "My username : $username" > my_system_info/about_this_laptop.txt
echo "With host: $os" >> my_system_info/about_this_laptop.txt

ping=$(ping -c 3 google.com)
echo "$ping" > my_system_info/internet_connection.txt

# showing folder tree
tree 
