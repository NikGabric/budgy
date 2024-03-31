use std::sync::Arc;

use deadpool_diesel::postgres::Pool;
use dotenvy::dotenv;
use tokio::net::TcpListener;

use crate::config::Config;
use crate::router::app_router;

mod schema;
mod models;
mod router;
mod handlers;
mod helpers;
mod config;
mod middleware;

#[derive(Clone)]
pub struct AppState {
    pool: Pool,
    config: Config,
}

#[tokio::main]
async fn main() {
    dotenv().ok();
    let config = Config::init();

    let manager = deadpool_diesel::postgres::Manager::new(config.db_url.clone(), deadpool_diesel::Runtime::Tokio1);
    let pool = deadpool_diesel::postgres::Pool::builder(manager)
        .build()
        .unwrap();

    let app_state = Arc::new(AppState {
        pool,
        config,
    });

    let app = app_router(app_state);
    let listener = TcpListener::bind("localhost:3000").await.unwrap();

    axum::serve(listener, app).await.unwrap();
}
