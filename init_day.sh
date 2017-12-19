#/bin/sh

dir="day$1"

if [ ! -d $dir ]
then
    echo "Creating $dir and files"
    mkdir $dir
    cp day.go $dir/day$1.go
    cp day_test.go $dir/day$1_test.go
else
    echo "$dir exists, exiting"
fi


