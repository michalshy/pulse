use serde::Deserialize;
use toml::Value;
use std::fs;

#[derive(Deserialize, Debug)]
pub struct ProjectConfig {
    pub key: String,
    pub name: String 
}

#[derive(Deserialize, Debug)]
pub struct AgentConfig {
    pub host: String,
    pub port: u64,
    pub binary: String,
    pub auto_start: bool,
    pub buffer_size: usize,
    pub flush_interval_secs: u64,
    pub flush_threshold: f64,
    pub heartbeat_interval_secs: u64
}

#[derive(Deserialize, Debug)]
pub struct IngestConfig {
    pub endpoint: String,
    pub timeout_secs: u64,
    pub retry_attempts: u64,
}

#[derive(Deserialize, Debug)]
pub struct BackendConfig {
    pub port: u64
}

#[derive(Deserialize, Debug)]
pub struct Config {
    pub project: ProjectConfig,
    pub agent: AgentConfig,
    pub ingest: IngestConfig,
    pub backend: BackendConfig
}

impl Config {
    pub fn new(path: String) -> Config {
        let content = fs::read_to_string(path).expect("Failed to read config!");
        let value: Value = toml::from_str(&content).expect("Failed to parse toml!");

        let config: Config = value
            .clone()
            .try_into()
            .expect("Failed to parse agent config");

        print!("
            Config discovered:
            {:?}, {:?}, {:?}, {:?}
        ", config.project, config.agent, config.ingest, config.backend);

        config
    }
}