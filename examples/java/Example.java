package com.truejacobg.examples;

import com.github.truejacobg.licencvalidator.LicenceValidator;

public class Example {
    public static void main(String[] args) {
        String[][] plates = {
            {"WPI 1234X", "PL"},
            {"B-AB 1234", "DE"},
            {"AA-123-AB", "FR"},
            {"AA-123-SS", "FR"},
            {"WPI 1234X", "XX"}
        };

        for (String[] p : plates) {
            boolean result = LicenceValidator.isValid(p[0], p[1]);
            System.out.printf("isValid(\"%s\", \"%s\") = %s%n", p[0], p[1], result);
        }
    }
}
