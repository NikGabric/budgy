use sea_orm_migration::prelude::*;

use super::m20240328_000001_create_budgy_user_table::BudgyUser;

pub struct Migration;

impl MigrationName for Migration {
    fn name(&self) -> &str {
        "m20240328_000002_create_transaction_type_table"
    }
}

#[async_trait::async_trait]
impl MigrationTrait for Migration {
    async fn up(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        manager
            .create_table(
                Table::create()
                    .table(TransactionType::Table)
                    .if_not_exists()
                    .col(
                        ColumnDef::new(TransactionType::TransactionTypeId)
                            .integer()
                            .not_null()
                            .auto_increment()
                            .primary_key(),
                    )
                    .col(ColumnDef::new(TransactionType::Title).string().not_null())
                    .col(ColumnDef::new(TransactionType::Description).string().not_null())
                    .col(ColumnDef::new(TransactionType::BudgyUserId).integer().not_null())
                    .foreign_key(
                        ForeignKey::create()
                            .name("fk-transaction_type-user_id")
                            .from(TransactionType::Table, TransactionType::BudgyUserId)
                            .to(BudgyUser::Table, BudgyUser::BudgyUserId),
                    )
                    .to_owned(),
            )
            .await
    }

    async fn down(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        manager
            .drop_table(Table::drop().table(TransactionType::Table).to_owned())
            .await
    }
}

#[derive(DeriveIden)]
pub enum TransactionType {
    Table,
    TransactionTypeId,
    Title,
    Description,
    BudgyUserId,
}
