require 'csv'
puts "Handling final text formats for: #{ARGV}"
x = CSV.read(ARGV[0])
formatted_csv = []

x.each_with_index do |line, idx|
  next if line.empty?
  next if line[0].nil? 
  next if line[0].include?("is a citywide unit")
  if line.include?("\nBacking Unsafely")
    line[0].sub!("\nBacking", "Backing")
    line[1].sub!("\n", "")
    line[2] = x[idx+1][2]
    x[idx+1] = []
  end

  if line.include?("Other Movers\nTOTAL Movers")
    combined_movers = line[1].split("\n")
    line[0].sub!("\nTOTAL Movers", "")
    line[1] = combined_movers[0]
    x[idx+1][0] = "TOTAL Movers"
    x[idx+1][1] = combined_movers[1]
  end
  formatted_csv << line
end

CSV.open(ARGV[0], "w") do |csv|
  formatted_csv.each { |l| csv << l }
end
