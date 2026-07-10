from eu_licence_validator import is_valid

plates = [
    ("WPI 1234X", "PL"),
    ("B-AB 1234", "DE"),
    ("AA-123-AB", "FR"),
    ("AA-123-SS", "FR"),
    ("WPI 1234X", "XX"),
]

for plate, country in plates:
    result = is_valid(plate, country)
    print(f"is_valid({plate!r}, {country!r}) = {result}")
