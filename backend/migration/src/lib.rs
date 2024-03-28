pub use sea_orm_migration::prelude::*;

mod m20240328_000001_create_budgy_user_table;
mod m20240328_000002_create_transaction_type_table;
mod m20240328_000003_create_transaction_table;

pub struct Migrator;

#[async_trait::async_trait]
impl MigratorTrait for Migrator {
    fn migrations() -> Vec<Box<dyn MigrationTrait>> {
        vec![
            Box::new(m20240328_000001_create_budgy_user_table::Migration),
            Box::new(m20240328_000002_create_transaction_type_table::Migration),
            Box::new(m20240328_000003_create_transaction_table::Migration),
        ]
    }
}
