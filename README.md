# GoBible

Go Bible App project here is to upgrade our existing bible project, written in python:

Unique Bible App https://github.com/eliranwong/UniqueBible

# Screenshot

<img width="1440" alt="gobible_screenshot" src="https://user-images.githubusercontent.com/25262722/205455063-247152e2-9cdc-4624-8738-c678a5e25f7e.png">

# in progress

... gobible development is in progress ...

You may also try our full-featured bible app written in python:

<b>Unique Bible App https://github.com/eliranwong/UniqueBible</b>

Current GoBible version supports the following features:
* Bible reading
* Bible searching
* Bible comparison on chapters
* Bible comparison on verses
* Bible cross-reference
* Opening multiple references in one go
* Loading content in multiple tabs and windows
* Light and dark themes

# Install GUI on Windows

... pending ...

# Install GUI on macOS

1) download and unzip https://github.com/eliranwong/gobible/archive/refs/heads/main.zip

2) double-click "GoBible.dmg"

3) drag the "GoBible.app" file to Applications directory

4) drag the "gobible_data" folder to the same Applications directory

Remove "... Apple cannot check ..." message:

1) In the first run, you may see the following message:

<img width="922" alt="check1" src="https://user-images.githubusercontent.com/25262722/205467739-f4e5b28f-225c-497e-9ac4-b192a2e67f2b.png">

2) Go to "System Settings" > "Privacy & Security" and select "Open Anyway"

<img width="717" alt="check2" src="https://user-images.githubusercontent.com/25262722/205467742-2a9b77a6-4fc4-4ef3-a5b9-8a22dacd6996.png">

# Install GUI on Linux 

For example, debian-based Linux users:

1) Download to home directory<br>
> cd ~<br>
> git clone https://github.com/eliranwong/gobible.git<br>
> cd gobible<br>
2) Unpack to root directory<br>
> sudo apt install xz-utils<br>
> sudo tar -xvf GoBible.tar.xz -C "/"<br>
3) Run Go Bible App<br>
> gobible<br>
To add path, in case step 3 does not work:<br>
> echo "PATH=$PATH:/usr/local/bin" >> ~/.bashrc

For other Linux distros, read about installing *.tar.xz file at:
https://www.cyberciti.biz/faq/how-to-extract-tar-xz-files-in-linux-and-unzip-all-files/

# Install Terminal Mode

In addition to GUI mode, you can run GoBible in terminal mode:

<img width="746" alt="terminal_mode" src="https://user-images.githubusercontent.com/25262722/205730113-e90ab5e7-c7c7-43fd-898c-3a8ad37015a5.png">

Search example: search for all verses containing both "God" and "love":

https://user-images.githubusercontent.com/25262722/206020927-f2bcf406-af35-46fa-96e1-6f5d741cb78a.mov

To install, download to home directory<br>
> cd ~<br>
> git clone https://github.com/eliranwong/gobible.git<br>
> cd gobible<br>
> chmod +x gobible
> ./gobible

# For go developers

Prerequisites: Go 1.18 or later

> git clone https://github.com/eliranwong/gobible.git

> cd gobible

To run GUI mode:

> go run gobible.go fyne

To run Terminal mode:

> go run gobible.go

To build a binary file:

> go biuld gobible.go

# Data

We separate "gobible_data" folder from executable file on purpose, to let users manage or add data easily.

# Add More Bibles

Download more at https://github.com/otseng/UniqueBible_Bibles

unzip files and place in folder "gobible_data/bibles"

restart Go Bible App

