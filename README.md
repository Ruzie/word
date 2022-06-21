## word

A tiny program which can scrap definition of various kind of words.

#### Screenshot
![image](https://user-images.githubusercontent.com/71683721/174860011-c2662129-fdc4-4065-821a-68b09a65d357.png)

## Installation

### Source

Install go on your system. Recommended (1.15 or above)
```sh
# Debian (or based Ubuntu etc)
sudo apt install golang

# Arch (or based Endeavouros etc)
sudo pacman -S golang
```

Compilation
```sh
git clone https://github.com/Ruzie/word.git
cd word && chmod +x build.sh
./build.sh
```

Pack and redistribute .tar.xz (or you can do gzip) file
```sh
tar cf word.tar out/word && xz -z /out/word
```
