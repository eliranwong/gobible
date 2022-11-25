# Merge Fonts

Use fontforge

To install fontforge, e.g. on mac:
> brew install fontforge
> brew install --cask fontforge

> cd fonts
> chmod +x merge_fonts.sh

e.g.
> ./merge_fonts.sh font1.ttf font2.ttf

The code in 'merge_fonts.sh' is pretty straight forward. The “$1” argument is for your primary font while the “$2” argument is for your secondary font. Fonts are first scaled to a uniform size and then merged to generate a new font. If you do not scale them to a common size, you may get uneven text rendering from the final merged font.

Alternately, use font-family-merger

Read https://github.com/dfrankland/font-family-merger

# Bundle Fonts

Read https://developer.fyne.io/extend/bundle

Example: https://github.com/lusingander/fyne-font-example