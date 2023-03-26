#find . -type f -name "*.sh" | cut -d "." -f2 | tr "/" "." | cut -d "." -f2
#find . -type f -name "*.sh" | cut -d "." -f2 | cut -d "/" -f3
find -type f -name "*.sh" -exec basename -s .sh {} \;