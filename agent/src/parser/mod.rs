use serde::Deserialize;
use toml::Value;
use std::fs;

#[derive(Deserialize)]
pub struct Agent {
    pub host: String,
    pub port: u64,
    pub binary: String,
    pub auto_start: bool,
    pub buffer_size: usize,
    pub flush_interval_secs: u64,
    pub flush_threshold: f64,
    pub heartbeat_interval_secs: u64
}

impl Agent {
    pub fn new(path: String) -> Agent {
        let content = fs::read_to_string(path).expect("Failed to read config!");
        let value: Value = toml::from_str(&content).expect("Failed to parse toml!");

        let agent: Agent = value["agent"]
            .clone()
            .try_into()
            .expect("Failed to parse agent config");

        println!("agent discovered at {}:{}", agent.host, agent.port);
        agent
    }
}