require 'csv'
puts "Handling final text formats for: #{ARGV}"
x = CSV.read(ARGV[0])
formatted_csv = []

x.each_with_index do |line, idx|
  next if line.empty?
  next if line[1]&.include?("MTD") && line[2]&.include?("YTD")
  next if line[0]&.include?("sum")
  next if line[1]&.include?("is a citywide unit")
  next if line[0]&.empty? && line[1]&.empty? && line[2]&.empty?
  next if [
    'january',
    'february',
    'march', 
    'april', 
    'may',
    'june', 
    'july', 
    'august', 
    'september',
    'october', 
    'november',
    'december'].any? { |word| line[0].downcase.include?(word) }
  
  # Remove total movers formatting
  if line[0].include?("TOTAL Movers")
    line[0] = "TOTAL Movers"
  end


  formatted_csv << line
end

CSV.open(ARGV[0], "w") do |csv|
  formatted_csv.each { |l| csv << l }
end
