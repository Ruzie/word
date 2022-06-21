if  [ ! go ]; then
    echo "Go isn't installed or not in global PATH"
    exit 1
fi

rm -rf out
mkdir out && go build -ldflags "-w" -o out/word word.go
