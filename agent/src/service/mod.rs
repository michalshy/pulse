use tonic::{Request, Response, Status};
use tonic::Streaming;
use crate::pulse::{HeartbeatRequest, HeartbeatResponse};
use crate::pulse::{Ack, LogEntry, Metric, pulse_server::{Pulse, PulseServer}};

pub struct PulseService {
    http: reqwest::Client,
    backend_url: String
}

impl PulseService {
    pub fn new(http: reqwest::Client, backend_url: String) -> PulseService {
        return PulseService { http, backend_url }
    }
}


#[tonic::async_trait]
impl Pulse for PulseService {
    async fn ingest_logs(
        &self,
        request: Request<Streaming<LogEntry>>
    ) -> Result<Response<Ack>, Status> {
        let mut stream = request.into_inner();

        while let Some(entry) = stream.message().await? {
            println!("log: [{}] {}", entry.level, entry.message);
        };

        Ok(Response::new(Ack { success: true }))
    }

    async fn ingest_metrics(
        &self,
        request: Request<Streaming<Metric>>
    ) -> Result<Response<Ack>, Status> {
        let mut stream = request.into_inner();

        while let Some(entry) = stream.message().await? {
            println!("log: [{}] {}", entry.name, entry.value);
        };

        Ok(Response::new(Ack { success: true }))
    }

    async fn heartbeat(
        &self,
        request: Request<HeartbeatRequest>
    ) -> Result<Response<HeartbeatResponse>, Status> {
        let req = request.into_inner();

        Ok(Response::new(HeartbeatResponse { 
            timestamp: req.timestamp, 
            ok: true 
        }))
    }
}