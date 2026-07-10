require "minitest/autorun"
require "json"
require_relative "../lib/eu_licence_validator"

CASES = JSON.parse(File.read(File.expand_path("../../../test_cases.json", __dir__)))

class TestLicenceValidator < Minitest::Test
  CASES.each_with_index do |tc, i|
    define_method("test_case_#{i}_#{tc['country']}_#{tc['plate'].tr(' ', '_')}") do
      got = EU::LicenceValidator.valid?(tc["plate"], tc["country"])
      assert_equal tc["expected"], got,
        "valid?(#{tc['plate'].inspect}, #{tc['country'].inspect}) = #{got}, want #{tc['expected']}"
    end
  end
end
