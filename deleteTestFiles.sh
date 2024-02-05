## create files_to_delete by matching ls tmp/testfiles/ | grep for files ending in .bin
## then for loop through the files and deletes them
files_to_delete=$(ls ./tmp/testfiles/ | grep .bin)

for file in $files_to_delete; do
  rm tmp/testfiles/$file
  echo "Deleted $file"
done
echo "All files deleted"