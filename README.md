# GoBible

Go Bible App project here is to upgrade our existing bible project, written in python:

Unique Bible App https://github.com/eliranwong/UniqueBible

# Screenshot

<img width="1440" alt="gobible_screenshot" src="https://user-images.githubusercontent.com/25262722/205455063-247152e2-9cdc-4624-8738-c678a5e25f7e.png">

# in progress

... gobible development is in progress ...

Current version supports the following features:
* Bible reading
* Bible searching
* Bible comparison on chapters
* Bible comparison on verses
* Bible cross-reference
* Opening multiple references in one go
* Loading content in multiple tabs and windows
* Light and dark themes

You may also try our full-featured bible app written in python:

<b>Unique Bible App https://github.com/eliranwong/UniqueBible</b>

# Install on Windows

... pending ...

# Install on macOS

1) download https://github.com/eliranwong/gobible/archive/refs/heads/main.zip

2) double-click the downloaded zip file to unzip it

3) double-click "GoBible.app" in the unzipped folder

To by-pass security check, run:

> xattr -d com.apple.quarantine GoBible.app

Read more at: https://hiringengineersbook.com/post/disable-quarantine/

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

# Add More Bibles

Download more at https://github.com/otseng/UniqueBible_Bibles

unzip files and place in folder "data/bibles"

restart Go Bible App

