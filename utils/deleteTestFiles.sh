## create files_to_delete by matching ls data/ | grep for files ending in .txt
## then for loop through the files and deletes them
files_to_delete=$(ls data/ | grep .txt)

for file in $files_to_delete; do
  rm data/$file
  echo "Deleted $file"
done
echo "All files deleted"