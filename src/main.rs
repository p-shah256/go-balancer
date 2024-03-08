mod proxy;
mod server;

fn main() {
    // if we do this : std::thread::spawn(function())
    // you are actually calling function and passing its return value to thread::spawn.
    // This is not what you want when spawning a thread.
    // Instead, you want to pass a function or closure that can be executed by the thread.
    let server_handle = std::thread::spawn(server_thread);
    let proxy_handle = std::thread::spawn(proxy_thread);

    let _ = server_handle.join().unwrap();
    let _ = proxy_handle.join().unwrap();
}

fn server_thread() {
    let _ = server::run();
}

fn proxy_thread() {
    let _ = proxy::run();
}
