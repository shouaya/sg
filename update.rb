require "open-uri"

BASE_URL = "https://raw.githubusercontent.com/shouaya/sg/main/"
files = ['lib.rb', 'main.rb', 'Gemfile', 'steps_demo.xls', 'README.md']

files.each do |file_name|
    open(BASE_URL + file_name) do |file_src|
        file_target = file_name
        if file_name == 'steps_demo.xls' && !File.exist?('steps.xls')
            file_target = 'steps.xls'
        end
        File.open("./" + file_target, "wb") do |file|
            file.write(file_src.read)
        end
    end
end