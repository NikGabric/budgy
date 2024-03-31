use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct TokenClaims {
    pub sub: String,
    pub exp: usize,
    pub iat: usize,
    pub usr: UserClaims,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct UserClaims {
    pub username: String,
    pub email: String,
}
