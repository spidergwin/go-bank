for File in $(ls); do
  (echo "\n//Filename: "$File"" && cat "$File")
done | xclip
