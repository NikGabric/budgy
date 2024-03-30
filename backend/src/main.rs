use std::env;

use dotenvy::dotenv;
use tokio::net::TcpListener;

use crate::router::app_router;

mod schema;
mod models;
mod router;
mod handlers;
mod helpers;

#[tokio::main]
async fn main() {
    dotenv().ok();
    let db_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set!");

    let manager = deadpool_diesel::postgres::Manager::new(db_url.clone(), deadpool_diesel::Runtime::Tokio1);
    let pool = deadpool_diesel::postgres::Pool::builder(manager)
        .build()
        .unwrap();

    let app = app_router(pool);
    let listener = TcpListener::bind("localhost:3000").await.unwrap();

    axum::serve(listener, app).await.unwrap();
}
