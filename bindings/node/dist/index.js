import { readFileSync } from "node:fs";
import { fileURLToPath } from "node:url";
import { dirname, join } from "node:path";
import { WASI } from "node:wasi";
const __dirname = dirname(fileURLToPath(import.meta.url));
let instancePromise = null;
function loadWasm() {
    if (instancePromise)
        return instancePromise;
    instancePromise = (async () => {
        const wasmPath = join(__dirname, "..", "wasm", "core.wasm");
        const bytes = readFileSync(wasmPath);
        const wasi = new WASI({ version: "preview1", args: [], env: {} });
        const mod = await WebAssembly.compile(bytes);
        const instance = (await WebAssembly.instantiate(mod, wasi.getImportObject()));
        try {
            wasi.start(instance);
        }
        catch {
            // _start may call proc_exit; safe to ignore for library use.
        }
        return instance;
    })();
    return instancePromise;
}
function writeString(memory, alloc, str) {
    const buf = Buffer.from(str, "utf8");
    const ptr = alloc(buf.length);
    const view = new Uint8Array(memory.buffer);
    view.set(buf, ptr);
    return { ptr, len: buf.length };
}
export async function isValid(plate, countryCode) {
    const instance = await loadWasm();
    const { memory, alloc, validate, dealloc } = instance.exports;
    const p = writeString(memory, alloc, plate);
    const c = writeString(memory, alloc, countryCode);
    const result = validate(p.ptr, p.len, c.ptr, c.len);
    dealloc(p.ptr);
    dealloc(c.ptr);
    return result === 1;
}
export async function isValidSync(plate, countryCode) {
    return isValid(plate, countryCode);
}
export default { isValid };
