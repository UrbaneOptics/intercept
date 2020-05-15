require 'csv'
require 'json'
folders = [
  "2011_08_sum",
  "2011_09_sum",
  "2011_10_sum",
  "2011_11_sum",
  "2011_12_sum",
  "2012_01_sum",
  "2012_02_sum",
  "2012_03_sum",
  "2012_04_sum",
  "2012_05_sum",
  "2012_06_sum",
  "2012_07_sum",
  "2012_08_sum",
  "2012_09_sum",
  "2012_10_sum",
  "2012_11_sum",
  "2012_12_sum",
  "2013_01_sum",
  "2013_02_sum",
  "2013_03_sum",
  "2013_04_sum",
  "2013_05_sum",
  "2013_06_sum",
  "2013_07_sum",
  "2013_08_sum",
  "2013_09_sum",
  "2013_10_sum",
  "2013_11_sum",
  "2013_12_sum",
  "2014_01_sum",
  "2014_02_sum",
  "2014_03_sum",
  "2014_04_sum",
  "2014_05_sum",
  "2014_06_sum",
  "2014_07_sum",
  "2014_08_sum",
  "2014_09_sum",
  "2014_10_sum",
  "2014_11_sum",
  "2014_12_sum",
  "2015_01_sum",
  "2015_02_sum",
  "2015_03_sum",
  "2015_04_sum",
  "2015_05_sum",
  "2015_06_sum",
  "2015_07_sum",
  "2015_08_sum",
  "2015_09_sum",
  "2015_10_sum",
  "2015_11_sum",
  "2015_12_sum",
  "2016_01_sum",
  "2016_02_sum",
  "2016_03_sum",
  "2016_04_sum",
  "2016_05_sum",
  "2016_06_sum",
  "2016_07_sum",
  "2016_08_sum",
  "2016_09_sum",
  "2016_10_sum",
  "2016_11_sum",
  "2016_12_sum",
  "2017_01_sum",
  "2017_02_sum",
  "2017_03_sum",
  "2017_04_sum",
  "2017_05_sum",
  "2017_06_sum",
  "2017_07_sum",
  "2017_08_sum",
  "2017_09_sum",
  "2017_10_sum",
  "2017_11_sum",
  "2017_12_sum",
  "2018_01_sum",
  "2018_02_sum",
  "2018_03_sum",
  "2018_04_sum",
  "2018_05_sum",
  "2018_06_sum",
  "2018_07_sum",
  "2018_08_sum",
  "2018_09_sum",
  "2018_10_sum",
  "2018_11_sum",
  "2018_12_sum",
  "2019_01_sum",
  "2019_02_sum",
  "2019_03_sum",
  "2019_04_sum",
  "2019_05_sum",
  "2019_06_sum",
  "2019_07_sum",
  "2019_08_sum",
  "2019_09_sum",
  "2019_10_sum",
  "2019_11_sum",
]

months = {}
folders.each do |folder|
  puts "Aggregating violations occuring on #{folder} for #{ARGV}"
  # Call from the /csv directory
  # ruby ../../scripts/retrieve_violation_count.rb 013.sum.csv
  file_dir = File.dirname(__FILE__) + "/../data/csv/#{folder}/#{ARGV[0]}"
  puts "Reading file from #{file_dir}"
  months[folder] = {}
  if File.exists?(file_dir)
    file = CSV.read(file_dir)
    file.each_with_index do |line, idx|
      if !line[0].nil? && !line[1].nil?
        months[folder][line[0]] = line[1]
      end
    end  
  end
end

File.open(File.dirname(__FILE__) + "/../data/json/#{ARGV[0].sub("sum.csv","")}_aggregate.json","w") do |f|
  f.write(JSON.pretty_generate(months))
end