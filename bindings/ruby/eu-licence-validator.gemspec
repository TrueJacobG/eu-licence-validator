require_relative "lib/eu_licence_validator/version"

Gem::Specification.new do |spec|
  spec.name          = "eu-licence-validator"
  spec.version       = EU::LicenceValidator::VERSION
  spec.summary       = "European licence plate validation (shared Go/Wasm core)."
  spec.description   = "European licence plate validation. Single Go core compiled to WebAssembly, bundled for Ruby."
  spec.authors       = ["TrueJacobG"]
  spec.email         = ["jakubgradzewicz1309@gmail.com"]
  spec.license       = "MIT"
  spec.homepage      = "https://github.com/TrueJacobG/eu-licence-validator"
  spec.required_ruby_version = ">= 3.0"

  spec.files = Dir.chdir(__dir__) do
    Dir["lib/**/*.rb", "wasm/core.wasm", "README.md"]
  end
  spec.require_paths = ["lib"]

  spec.add_dependency "wasmtime", "~> 46.0"

  spec.metadata = {
    "homepage_uri" => spec.homepage,
    "source_code_uri" => "https://github.com/TrueJacobG/eu-licence-validator",
    "changelog_uri" => "https://github.com/TrueJacobG/eu-licence-validator/blob/main/CHANGELOG.md",
  }
end
