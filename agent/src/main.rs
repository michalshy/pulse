pub mod pulse {
    tonic::include_proto!("pulse");
}
pub mod service;
pub mod parser;

use tonic::transport::Server;

use crate::parser::Config;
use crate::pulse::pulse_server::PulseServer;

#[tokio::main]
async fn main() {
    // Parse configuration
    // For now we hardcode the path, later we will pass it more gracefully
    let path: String = "../pulse/pulse.toml".into();
    let config: Config = Config::new(path);
    let addr = format!(
        "{}:{}", 
        config.agent.host,
        config.agent.port
    ).parse().unwrap();

    // Prepare server
    Server::builder()
        .add_service(PulseServer::new(
            service::PulseService::new(
                reqwest::Client::new(), 
                config.ingest.endpoint
            )
        ))
        .serve(addr)
        .await
        .expect("Server failed");

    println!("Pulse Agent exits...")
}