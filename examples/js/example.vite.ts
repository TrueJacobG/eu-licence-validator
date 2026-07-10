import { init, isValid } from "@truejacobg/eu-licence-validator";

await init();

const plates: [string, string][] = [
  ["WPI 1234X", "PL"],
  ["B-AB 1234", "DE"],
  ["AA-123-AB", "FR"],
  ["AA-123-SS", "FR"],
  ["WPI 1234X", "XX"],
];

export function runExamples(): void {
  for (const [plate, country] of plates) {
    const result = isValid(plate, country);
    console.log(`isValid(${JSON.stringify(plate)}, ${JSON.stringify(country)}) = ${result}`);
  }
}
