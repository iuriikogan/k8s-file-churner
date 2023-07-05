## create files_to_delete by matching ls testfiles/ | grep for files ending in .txt
## then for loop through the files and deletes them
files_to_delete=$(ls testfiles/ | grep .txt)

for file in $files_to_delete; do
  rm testfiles/$file
  echo "Deleted $file"
done
echo "All files deleted"