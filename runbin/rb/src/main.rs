use std::fs;
use std::mem;

type AsmFuncType = extern "C" fn(*const u8, u64);

fn main() {
    let code = fs::read("hello.bin").expect("Failed to read file");
    let code_size = code.len();

    let mmap = unsafe {
        libc::mmap(
            std::ptr::null_mut(),
            code_size,
            libc::PROT_READ | libc::PROT_WRITE | libc::PROT_EXEC,
            libc::MAP_ANON | libc::MAP_PRIVATE,
            -1,
            0,
        )
    };
    if mmap == libc::MAP_FAILED {
        panic!("Failed to mmap");
    }

    unsafe {
        std::ptr::copy(code.as_ptr(), mmap as *mut u8, code_size);
    }

    let str = "hello, world!\n";
    let length = str.len();
    let cstr = str.as_bytes();

    let func_ptr = mmap as *const u8;
    unsafe {
        let exec_func: AsmFuncType = std::mem::transmute_copy(&func_ptr);
        exec_func(cstr.as_ptr(), length as u64);
    }
}
