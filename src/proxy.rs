// from tokio imoort net::TcpListener - listens for incoming TCP connections +
// from tokio import net::TcpStream - represents a TCP stream between a client and a server +
// once the handshake is complete, TCP listener returns a TCP stream
use std::net::SocketAddr;
use tokio::{
    net::{TcpListener, TcpStream},
    try_join,
};

const PROXY_ADDRESS: ([u8; 4], u16) = ([127, 0, 0, 1], 3000);
const SERVER_ADDRESS: ([u8; 4], u16) = ([127, 0, 0, 1], 1973);

// Using a runtime,involves more than just importing.
// It refers to setting up and utilizing a specific runtime environment designed for asynchronous programming.
#[tokio::main]
pub async fn run() -> Result<(), Box<dyn std::error::Error>> {
    let proxy: SocketAddr = PROXY_ADDRESS.into();
    println!("proxy is running at {:?}", PROXY_ADDRESS);

    let proxy_server = TcpListener::bind(proxy).await?;

    // proxyserver.accept returns two values, the first is the client connection and
    // the second is the client's address
    // as far as it return Ok, spawn a new task to handle the client connection
    // while loop will keep accepting new connections and spawning new tasks to handle them
    while let Ok((client, _)) = proxy_server.accept().await {
        tokio::spawn(async move {
            let _ = handle_client_conn(client).await;
        });
    }

    Ok(())
}

async fn handle_client_conn(mut client_conn: TcpStream) -> Result<(), Box<dyn std::error::Error>> {
    let server: SocketAddr = SERVER_ADDRESS.into();

    let mut main_server_conn = TcpStream::connect(server).await?;
    let (mut client_recv, mut client_send) = client_conn.split();
    let (mut server_recv, mut server_send) = main_server_conn.split();

    // AWAITING for either client or server to send data, when then happens it triggers
    // copy, which will copy the data from the client to the server and from the server to the client
    let handle_one = async { tokio::io::copy(&mut server_recv, &mut client_send).await };

    let handle_two = async { tokio::io::copy(&mut client_recv, &mut server_send).await };

    try_join!(handle_one, handle_two)?;

    Ok(())
}
