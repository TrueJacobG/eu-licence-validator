import { test } from "node:test";
import { strictEqual } from "node:assert";
import { readFileSync } from "node:fs";
import { fileURLToPath } from "node:url";
import { dirname, join } from "node:path";
import { init, isValid } from "../dist/index.js";

const __dirname = dirname(fileURLToPath(import.meta.url));
const cases = JSON.parse(
  readFileSync(join(__dirname, "..", "..", "..", "test_cases.json"), "utf8")
);

test("isValid against test_cases.json", async () => {
  await init();
  for (const { plate, country, expected } of cases) {
    const got = isValid(plate, country);
    strictEqual(
      got,
      expected,
      `isValid(${JSON.stringify(plate)}, ${JSON.stringify(country)}) = ${got}, want ${expected}`
    );
  }
});
