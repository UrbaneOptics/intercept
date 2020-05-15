require 'csv'
require 'json'

puts "Retrieving violation names for: #{ARGV}"
x = CSV.read(ARGV[0])
unique_values = {}
x.each_with_index do |line, idx|
  if unique_values[:"#{line[0]}"].nil?
    unique_values[:"#{line[0]}"] = 1
  else 
    unique_values[:"#{line[0]}"] += 1
  end
end

File.open("../moving_violation_types.json","w") do |f|
  f.write(JSON.pretty_generate(unique_values))
end
