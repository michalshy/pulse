use tokio::sync::mpsc;
use tokio::time::{sleep, Duration};

#[derive(Debug, Clone)]
struct FrameSample {
    frame_time: f64,
    gpu_time: f64,
    cpu_time: f64,
    timestamp_ms: u64,
}

async fn ipc_reader(tx: mpsc::Sender<FrameSample>) {
    let mut frame_number = 0u64;

    loop {
        sleep(Duration::from_millis(16)).await;

        let frame = FrameSample {
            frame_time: 16.6 + (frame_number % 10) as f64 * 0.3,
            gpu_time: 12.0,
            cpu_time: 8.0,
            timestamp_ms: frame_number * 16,
        };

        if tx.send(frame).await.is_err() {
            println!("receiver dropped");
            break;
        }

        frame_number += 1;
    }
}

async fn batch_sender(mut rx: mpsc::Receiver<FrameSample>) {
    let mut batch: Vec<FrameSample> = Vec::new();
    let batch_interval = Duration::from_millis(200);

    loop {
        let deadline = tokio::time::sleep(batch_interval);
        tokio::pin!(deadline);

        loop {
            tokio::select! {
                _ = &mut deadline => {
                    break;
                }

                frame = rx.recv() => {
                    match frame {
                        Some(f) => batch.push(f),
                        None => {
                            println!("channel closed");
                            return;
                        }
                    }
                }
            }
        }

        if !batch.is_empty() {
            println!("sending batch of {} frames", batch.len());
            // here grpc send
            batch.clear();
        }
    }
}

#[tokio::main]
async fn main() {
    let (tx, rx) = mpsc::channel::<FrameSample>(256);

    let reader = tokio::spawn(ipc_reader(tx));
    let sender = tokio::spawn(batch_sender(rx));

    sleep(Duration::from_secs(3)).await;

    reader.abort();
    sender.abort();

    println!("agent stopped!");
}