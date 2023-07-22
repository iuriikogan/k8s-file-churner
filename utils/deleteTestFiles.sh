## create files_to_delete by matching ls app/testfiles/ | grep for files ending in .bin
## then for loop through the files and deletes them
files_to_delete=$(ls app/testfiles/ | grep .bin)

for file in $files_to_delete; do
  rm app/testfiles/$file
  echo "Deleted $file"
done
echo "All files deleted"