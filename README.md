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

# Install on Windows

... pending ...

# Install on macOS

1) download and unzip https://github.com/eliranwong/gobible/archive/refs/heads/main.zip

2) double-click "GoBible.dmg"

3) drag the "GoBible.app" file to Applications directory

4) drag the "data" folder to the same Applications directory

Remove "... Apple cannot check ..." message:

1) In the first run, you may see the following message:

<img width="922" alt="check1" src="https://user-images.githubusercontent.com/25262722/205467739-f4e5b28f-225c-497e-9ac4-b192a2e67f2b.png">

2) Go to "System Settings" > "Privacy & Security" and select "Open Anyway"

<img width="717" alt="check2" src="https://user-images.githubusercontent.com/25262722/205467742-2a9b77a6-4fc4-4ef3-a5b9-8a22dacd6996.png">

# Install on Linux

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

# For go developers

> git clone https://github.com/eliranwong/gobible.git

> cd gobible

> go run gobible.go

# Data

We separate "data" folder from executable file on purpose, to allow users manage or add data easily.

# Add More Bibles

Download more at https://github.com/otseng/UniqueBible_Bibles

unzip files and place in folder "data/bibles"

restart Go Bible App

