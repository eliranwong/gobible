# References

https://developer.fyne.io

fyne_demo source code:
https://github.com/fyne-io/fyne/tree/a83492533a2c9f56dc3e7e1d100ca9b4216edfb6/cmd/fyne_demo

# setup on macOS:

https://www.ookangzheng.com/fyne-first-try-on-mac/

# Notarizing macOS software before distribution

https://developer.apple.com/documentation/security/notarizing_macos_software_before_distribution

# create and sign a dmg file on macOS

run:

> chmod +x build_dmg.sh

> ./build_dmg.sh

Before running the script:

Go to https://developer.apple.com/account/resources/certificates/list
create a "Developer ID Application" certificate

Go to https://developer.apple.com/account/resources/identifiers/list
select "Register a new identifier" > "App IDs" > continue
save the name and ID of the app

To set up a password for the app
Go to https://appleid.apple.com/
Select "App-specific passwords"
generate
give it a name which is the same name as APP_ID
enter Apple ID password
write down the generated password, appeared on the screen
enter the pass below

edit lines 21-41 in build_dmg.sh accordingly
