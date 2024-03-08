use log::{info, warn};

pub fn main() {
    env_logger::init();

    warn!("a warning");
    info!("an info");
    // ...
}
