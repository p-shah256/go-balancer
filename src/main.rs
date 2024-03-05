mod server;
mod proxy;

fn main() {
   let _ = server::run();
   let _ = proxy::run();
}
