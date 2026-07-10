package io.github.truejacobg.licencvalidator;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.InputStream;
import java.util.List;
import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class LicenceValidatorTest {

    private static final String CASES_PATH = "/test_cases.json";

    @SuppressWarnings("unchecked")
    private List<Map<String, Object>> loadCases() throws Exception {
        InputStream stream = LicenceValidatorTest.class.getResourceAsStream(CASES_PATH);
        if (stream == null) {
            throw new IllegalStateException("test_cases.json not found on classpath at " + CASES_PATH);
        }
        return new ObjectMapper().readValue(stream, List.class);
    }

    @Test
    void testAllCases() throws Exception {
        for (Map<String, Object> tc : loadCases()) {
            String plate = (String) tc.get("plate");
            String country = (String) tc.get("country");
            boolean expected = (Boolean) tc.get("expected");
            boolean got = LicenceValidator.isValid(plate, country);
            assertEquals(expected, got,
                "isValid(\"" + plate + "\", \"" + country + "\") = " + got + ", want " + expected);
        }
    }
}
