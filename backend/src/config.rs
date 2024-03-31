#[derive(Debug, Clone)]
pub struct Config {
    pub db_url: String,
    pub jwt_secret: String,
}

impl Config {
    pub fn init() -> Config {
        let db_url = std::env::var("DATABASE_URL").expect("DATABASE_URL must be set");
        let jwt_secret = std::env::var("JWT_SECRET").expect("JWT_SECRET must be set");
        Config {
            db_url,
            jwt_secret,
        }
    }
}