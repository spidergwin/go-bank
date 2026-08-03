for Folder in $(eza -D); do
  (cd "$Folder" && mv * ../ && cd .. && rmdir "$Folder")
done
