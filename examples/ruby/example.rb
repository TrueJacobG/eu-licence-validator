require "eu-licence-validator"

plates = [
  ["WPI 1234X", "PL"],
  ["B-AB 1234", "DE"],
  ["AA-123-AB", "FR"],
  ["AA-123-SS", "FR"],
  ["WPI 1234X", "XX"],
]

plates.each do |plate, country|
  puts "isValid(#{plate.inspect}, #{country.inspect}) = #{EU::LicenceValidator.valid?(plate, country)}"
end
