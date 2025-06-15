// convert-oklch.ts
import fs from "fs";
import path from "path";
import glob from "glob";
import { oklch, formatRgb, convert } from "culori";

// Match oklch color values
const OKLCH_REGEX = /oklch\(\s*([0-9.]+)\s+([0-9.]+)\s+([0-9.]+)\s*\)/g;

// Convert oklch to rgb string
function oklchToRgbStr(match: string, l: string, c: string, h: string): string {
  const parsed = oklch({ l: +l, c: +c, h: +h });
  const rgb = convert(formatRgb)(parsed);
  return rgb || match; // fallback to original if conversion fails
}

// List of file extensions to scan
const extensions = ["ts", "tsx", "js", "jsx", "css", "scss"];

glob("**/*.{ts,tsx,js,jsx,css,scss}", { ignore: "node_modules/**" }, (err, files) => {
  if (err) throw err;

  files.forEach((filePath) => {
    const content = fs.readFileSync(filePath, "utf-8");

    const replaced = content.replace(OKLCH_REGEX, oklchToRgbStr);

    if (replaced !== content) {
      fs.writeFileSync(filePath, replaced, "utf-8");
      console.log(`✅ Converted: ${filePath}`);
    }
  });
});
