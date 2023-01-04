require "open-uri"

BASE_URL = "https://raw.githubusercontent.com/shouaya/sg/main/"
files = ['lib.rb', 'main.rb', 'Gemfile']

files.each do |file_name|
    open(BASE_URL + file_name) do |file_src|
        File.open("./" + file_name, "wb") do |file|
            file.write(file_src.read)
        end
    end
end