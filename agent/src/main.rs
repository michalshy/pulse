use tokio::sync::mpsc;
use tokio::time::{interval, sleep, Duration};
use reqwest::Client;
use tokio_util::sync::CancellationToken;

async fn heartbeat_task (
    client: reqwest::Client,
    base_url: String,
    project_key: String,
    interval_secs: u64,
    token: CancellationToken,
) {
    let mut ticker = interval(Duration::from_secs(interval_secs));

    loop {
        tokio::select! {
            _ = token.cancelled() => {
                break;
            }
            _ = ticker.tick() => {
                if let Err(e) = client
                    .post(format!("{}/heartbeat/{}", base_url, project_key))
                    .send()
                    .await
                {
                    eprintln!("Heartbeat failed: {e}");
                }
            }
        }
    }
}


#[tokio::main]
async fn main() {
    dotenvy::from_path(".env").expect("Failed to load .env"); // Currently hardcoded - local settings

    let token = CancellationToken::new();
    let client = reqwest::Client::new();

    let api_url = std::env::var("API").expect("API not set");
    let project_key = std::env::var("PROJECT_KEY").expect("PROJECT_KEY not set");

    let heartbeat_time: u64 = std::env::var("HEARTBEAT_TIME")
        .expect("HEARTBEAT_TIME time not set")
        .parse()
        .expect("HEARTBEAT_TIME must be a number");

    let heartbeat = tokio::spawn(heartbeat_task(
        client.clone(), 
        api_url, 
        project_key, 
        heartbeat_time, 
        token.clone()
    ));

    tokio::signal::ctrl_c().await.expect("Failed to listen for ctrl+c");
    token.cancel();

    heartbeat.await.expect("Heartbeat task panicked!");

    println!("agent stopped!");
}