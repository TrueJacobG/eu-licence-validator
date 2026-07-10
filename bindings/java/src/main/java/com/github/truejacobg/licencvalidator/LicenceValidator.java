package com.github.truejacobg.licencvalidator;

import com.dylibso.chicory.runtime.HostFunction;
import com.dylibso.chicory.runtime.ImportValues;
import com.dylibso.chicory.runtime.Instance;
import com.dylibso.chicory.runtime.Memory;
import com.dylibso.chicory.wasi.WasiOptions;
import com.dylibso.chicory.wasi.WasiPreview1;
import com.dylibso.chicory.wasi.WasiPreview1_ModuleFactory;
import com.dylibso.chicory.wasm.Parser;

import java.io.InputStream;
import java.nio.charset.StandardCharsets;

public final class LicenceValidator {

    private static final String WASM_RESOURCE = "/core.wasm";

    private final Instance instance;
    private final Memory memory;

    private static volatile LicenceValidator instanceSingleton;

    private LicenceValidator() {
        InputStream wasmStream = LicenceValidator.class.getResourceAsStream(WASM_RESOURCE);
        if (wasmStream == null) {
            throw new IllegalStateException("core.wasm not found on classpath at " + WASM_RESOURCE);
        }
        var module = Parser.parse(wasmStream);
        WasiPreview1 wasi = WasiPreview1.builder()
                .withOptions(WasiOptions.builder().build())
                .build();
        HostFunction[] hostFunctions = WasiPreview1_ModuleFactory.toHostFunctions(wasi);
        ImportValues importValues = ImportValues.builder()
                .addFunction(hostFunctions)
                .build();
        this.instance = Instance.builder(module)
                .withImportValues(importValues)
                .withStart(false)
                .build();
        this.memory = instance.memory();

        try {
            instance.export("_start").apply();
        } catch (Exception e) {
            // _start may call proc_exit; safe to ignore for library use.
        }
    }

    private static LicenceValidator getInstance() {
        if (instanceSingleton == null) {
            synchronized (LicenceValidator.class) {
                if (instanceSingleton == null) {
                    instanceSingleton = new LicenceValidator();
                }
            }
        }
        return instanceSingleton;
    }

    public static boolean isValid(String plate, String countryCode) {
        return getInstance().doIsValid(plate, countryCode);
    }

    private boolean doIsValid(String plate, String countryCode) {
        long pPtr = writeString(plate);
        long cPtr = writeString(countryCode);

        long[] result = instance.export("validate").apply(pPtr, plate.getBytes(StandardCharsets.UTF_8).length, cPtr, countryCode.getBytes(StandardCharsets.UTF_8).length);
        instance.export("dealloc").apply(pPtr);
        instance.export("dealloc").apply(cPtr);
        return result[0] != 0;
    }

    private long writeString(String str) {
        byte[] bytes = str.getBytes(StandardCharsets.UTF_8);
        long[] ptrResult = instance.export("alloc").apply((long) bytes.length);
        long ptr = ptrResult[0];
        if (ptr == 0 && bytes.length > 0) {
            throw new IllegalStateException("wasm alloc returned 0");
        }
        memory.write((int) ptr, bytes);
        return ptr;
    }
}
