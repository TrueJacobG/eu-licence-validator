require_relative "eu_licence_validator/version"
require "wasmtime"

module EU
  module LicenceValidator
    WASM_PATH = File.expand_path("../wasm/core.wasm", __dir__)

    @engine = Wasmtime::Engine.new
    @module = Wasmtime::Module.from_file(@engine, WASM_PATH)
    @linker = Wasmtime::Linker.new(@engine)

    @linker.func_new("wasi_snapshot_preview1", "proc_exit", [:i32], []) do |code|
      raise Wasmtime::WasiExit.new(code)
    end
    @linker.func_new("wasi_snapshot_preview1", "fd_write", [:i32, :i32, :i32, :i32], [:i32]) { |_fd, _ptr, _n, _ret| 0 }
    @linker.func_new("wasi_snapshot_preview1", "random_get", [:i32, :i32], [:i32]) { |_ptr, _len| 0 }

    @store = Wasmtime::Store.new(@engine)
    @instance = @linker.instantiate(@store, @module)

    begin
      @instance.invoke("_start")
    rescue Wasmtime::WasiExit, SystemExit
    end

    class << self
      def valid?(plate, country_code)
        exports = @instance.exports

        p_ptr, p_len = write_string(exports, plate)
        c_ptr, c_len = write_string(exports, country_code)

        validate = exports["validate"].to_func
        dealloc = exports["dealloc"].to_func
        result = validate.call(p_ptr, p_len, c_ptr, c_len)
        dealloc.call(p_ptr)
        dealloc.call(c_ptr)
        result == 1
      end
      alias is_valid valid?

      private

      def write_string(exports, str)
        bytes = str.encode("UTF-8").bytes
        alloc = exports["alloc"].to_func
        ptr = alloc.call(bytes.length)
        return [0, 0] if ptr == 0 && bytes.empty?

        memory = exports["memory"].to_memory
        memory.write(ptr, bytes.pack("C*"))
        [ptr, bytes.length]
      end
    end
  end
end
