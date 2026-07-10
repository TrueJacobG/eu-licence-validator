const WASI_IMPORTS = {
    wasi_snapshot_preview1: {
        fd_write: () => 0,
        proc_exit: () => {
            throw new Error("wasi proc_exit");
        },
        random_get: () => 0,
    },
};
const encoder = new TextEncoder();
let instance = null;
let initPromise = null;
async function loadWasmBytes() {
    const wasmUrl = new URL("../wasm/core.wasm", import.meta.url);
    if (typeof process !== "undefined" &&
        typeof process.versions?.node === "string") {
        const { readFileSync } = await import(/* @vite-ignore */ "node:fs");
        return readFileSync(wasmUrl).buffer;
    }
    const response = await fetch(wasmUrl);
    return await response.arrayBuffer();
}
export async function init() {
    if (instance)
        return;
    if (initPromise)
        return initPromise;
    initPromise = (async () => {
        const bytes = await loadWasmBytes();
        const mod = await WebAssembly.compile(bytes);
        const inst = (await WebAssembly.instantiate(mod, WASI_IMPORTS));
        try {
            inst.exports._start();
        }
        catch {
            // _start calls proc_exit after initializing globals; safe to ignore.
        }
        instance = inst;
    })();
    return initPromise;
}
export function isValid(plate, countryCode) {
    if (!instance) {
        throw new Error("eu-licence-validator: call init() before isValid()");
    }
    const { memory, alloc, validate, dealloc } = instance.exports;
    const pBuf = encoder.encode(plate);
    const cBuf = encoder.encode(countryCode);
    const pPtr = alloc(pBuf.length);
    const cPtr = alloc(cBuf.length);
    new Uint8Array(memory.buffer).set(pBuf, pPtr);
    new Uint8Array(memory.buffer).set(cBuf, cPtr);
    const result = validate(pPtr, pBuf.length, cPtr, cBuf.length);
    dealloc(pPtr);
    dealloc(cPtr);
    return result === 1;
}
export default { init, isValid };
